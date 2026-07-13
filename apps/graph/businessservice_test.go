package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestBusinessServiceListsEveryPageWithBackendCredentials(t *testing.T) {
	t.Helper()
	var requests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if got := r.Header.Get("Authorization"); got != "Bearer secret" {
			t.Fatalf("Authorization = %q", got)
		}
		if r.URL.Path != "/api/v1/businesses" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("cursor") == "" {
			_, _ = w.Write([]byte(`{"items":[{"id":"one","display_name":"One","slug":"one","status":"active","visibility":"internal","version":1}],"next_cursor":"opaque"}`))
			return
		}
		_, _ = w.Write([]byte(`{"items":[{"id":"two","display_name":"Two","slug":"two","status":"draft","visibility":"internal","version":1}]}`))
	}))
	defer server.Close()

	client, err := newBusinessHTTPClient(BusinessServiceConfig{BaseURL: server.URL, AccessToken: "secret"}, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	service := &BusinessService{client: client}
	items, err := service.ListBusinesses()
	if err != nil {
		t.Fatal(err)
	}
	if requests != 2 || len(items) != 2 || items[1].DisplayName != "Two" {
		t.Fatalf("requests=%d items=%+v", requests, items)
	}
}

func TestBusinessServiceCreateRecordOwnsIdempotencyKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/v1/businesses/business/records" {
			t.Fatalf("%s %s", r.Method, r.URL.Path)
		}
		if key := r.Header.Get("Idempotency-Key"); key != "create-strategy-1" {
			t.Fatalf("idempotency key %q", key)
		}
		var request CreateRecordInput
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatal(err)
		}
		if request.Key != "strategy" || request.Revision.Title != "Strategy" {
			t.Fatalf("request=%+v", request)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"record":{"id":"record","business_id":"business","key":"strategy","type_code":"STR","current_revision":1,"version":1},"revision":{"id":"revision","record_id":"record","business_id":"business","revision":1,"record_version":1,"type_code":"STR","title":"Strategy","data":{},"content_hash":"hash","graph_hash":"graph","created_by":"actor","created_at":"2026-01-01T00:00:00Z"},"relationships":[]}`))
	}))
	defer server.Close()

	client, err := newBusinessHTTPClient(BusinessServiceConfig{BaseURL: server.URL}, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	service := &BusinessService{client: client}
	view, err := service.CreateRecord("business", CreateRecordInput{
		Key:      "strategy",
		TypeCode: "STR",
		Revision: RevisionInput{Title: "Strategy", DataJSON: "{}", Relationships: []RelationshipInput{}},
	}, "create-strategy-1")
	if err != nil {
		t.Fatal(err)
	}
	if view.Record.Version != 1 || view.Revision.Title != "Strategy" {
		t.Fatalf("view=%+v", view)
	}
}

func TestBusinessServiceReviseRecordUsesVersionETagAndDecodesConflict(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("If-Match"); got != `"v7"` {
			t.Fatalf("If-Match = %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusPreconditionFailed)
		_, _ = w.Write([]byte(`{"error":{"code":"BUSINESS_PRECONDITION_FAILED","message":"record version is stale","request_id":"request-1","tenant_id":"tenant","details":{"expected":7}}}`))
	}))
	defer server.Close()

	client, err := newBusinessHTTPClient(BusinessServiceConfig{BaseURL: server.URL}, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	service := &BusinessService{client: client}
	_, err = service.ReviseRecord("business", "record", 7, ReviseRecordInput{Revision: RevisionInput{Title: "Next", DataJSON: "{}"}})
	if err == nil || !strings.HasPrefix(err.Error(), "BUSINESS_PRECONDITION_FAILED:") {
		t.Fatalf("error=%v", err)
	}
	apiError, ok := err.(*BusinessAPIError)
	if !ok || apiError.RequestID != "request-1" {
		t.Fatalf("api error=%#v", err)
	}
}

func TestNormalizeBusinessBaseURLAcceptsAPIPathAndRejectsCredentials(t *testing.T) {
	parsed, err := normalizeBusinessBaseURL("https://business.example.com/api/v1/")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.String() != "https://business.example.com" {
		t.Fatalf("url=%q", parsed.String())
	}
	if _, err := normalizeBusinessBaseURL("https://user:pass@business.example.com"); err == nil {
		t.Fatal("expected credentials to be rejected")
	}
}

func TestBusinessServiceConfigureExplicitlyPreservesOrClearsHiddenCredentials(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/live", "/ready":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client, err := newBusinessHTTPClient(BusinessServiceConfig{
		BaseURL:     server.URL,
		AccessToken: "secret",
		TenantID:    "tenant",
		ActorID:     "actor",
	}, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	service := &BusinessService{client: client}

	status, err := service.Configure(BusinessServiceConfig{
		BaseURL:             server.URL,
		PreserveCredentials: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !status.Ready || !status.HasAccessToken || !status.HasDevelopmentIdentity {
		t.Fatalf("preserved status=%+v", status)
	}
	preserved := service.snapshotClient()
	if preserved.accessToken != "secret" || preserved.tenantID != "tenant" || preserved.actorID != "actor" {
		t.Fatalf("preserved client=%+v", preserved)
	}

	status, err = service.Configure(BusinessServiceConfig{BaseURL: server.URL})
	if err != nil {
		t.Fatal(err)
	}
	if !status.Ready || status.HasAccessToken || status.HasDevelopmentIdentity {
		t.Fatalf("cleared status=%+v", status)
	}
	cleared := service.snapshotClient()
	if cleared.accessToken != "" || cleared.tenantID != "" || cleared.actorID != "" {
		t.Fatalf("cleared client=%+v", cleared)
	}
}

func TestBusinessServicePreservesOpaqueJSONAcrossReadAndRevision(t *testing.T) {
	const largeInteger = "9007199254740993"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			_, _ = w.Write([]byte(`{"record":{"id":"record","business_id":"business","key":"strategy","type_code":"STR","current_revision":1,"version":1},"revision":{"id":"revision","record_id":"record","business_id":"business","revision":1,"record_version":1,"type_code":"STR","title":"Strategy","data":{"counter":` + largeInteger + `},"content_hash":"hash","graph_hash":"graph","created_by":"actor","created_at":"2026-01-01T00:00:00Z"},"relationships":[{"id":"relationship","revision_id":"revision","source_record_id":"record","target_key":"vision","kind":"enables","resolution":"unresolved","provenance":{"line":` + largeInteger + `},"created_at":"2026-01-01T00:00:00Z"}]}`))
		case http.MethodPatch:
			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Contains(body, []byte(`"counter":`+largeInteger)) || !bytes.Contains(body, []byte(`"line":`+largeInteger)) {
				t.Fatalf("opaque JSON changed in request: %s", body)
			}
			_, _ = w.Write([]byte(`{"record":{"id":"record","business_id":"business","key":"strategy","type_code":"STR","current_revision":2,"version":2},"revision":{"id":"revision-2","record_id":"record","business_id":"business","revision":2,"record_version":2,"type_code":"STR","title":"Strategy updated","data":{"counter":` + largeInteger + `},"content_hash":"hash-2","graph_hash":"graph-2","created_by":"actor","created_at":"2026-01-01T00:01:00Z"},"relationships":[]}`))
		default:
			t.Fatalf("unexpected method %s", r.Method)
		}
	}))
	defer server.Close()

	client, err := newBusinessHTTPClient(BusinessServiceConfig{BaseURL: server.URL}, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	service := &BusinessService{client: client}
	view, err := service.GetRecord("business", "record")
	if err != nil {
		t.Fatal(err)
	}
	if view.Revision.DataJSON != `{"counter":`+largeInteger+`}` {
		t.Fatalf("data_json=%q", view.Revision.DataJSON)
	}
	if len(view.Relationships) != 1 || view.Relationships[0].ProvenanceJSON != `{"line":`+largeInteger+`}` {
		t.Fatalf("relationships=%+v", view.Relationships)
	}
	_, err = service.ReviseRecord("business", "record", 1, ReviseRecordInput{
		TypeCode: "STR",
		Revision: RevisionInput{
			Title:    "Strategy updated",
			DataJSON: view.Revision.DataJSON,
			Relationships: []RelationshipInput{{
				Kind:           "enables",
				TargetKey:      "vision",
				ProvenanceJSON: view.Relationships[0].ProvenanceJSON,
			}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBusinessServiceRejectsBearerTokenOverRemoteHTTP(t *testing.T) {
	_, err := newBusinessHTTPClient(BusinessServiceConfig{
		BaseURL:     "http://business.example.com",
		AccessToken: "secret",
	}, nil)
	if err == nil || !strings.Contains(err.Error(), "require HTTPS") {
		t.Fatalf("error=%v", err)
	}
}

func TestBusinessServiceConfigureDoesNotLeakOrDiscardWorkingConnection(t *testing.T) {
	working := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer working.Close()

	var candidateRequests int
	candidate := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		candidateRequests++
		if token := r.Header.Get("Authorization"); token != "" {
			t.Fatalf("candidate received hidden credential %q", token)
		}
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/ready" {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"error":{"code":"BUSINESS_UNAVAILABLE","message":"candidate is not ready"}}`))
			return
		}
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer candidate.Close()

	client, err := newBusinessHTTPClient(BusinessServiceConfig{
		BaseURL:     working.URL,
		AccessToken: "secret",
	}, working.Client())
	if err != nil {
		t.Fatal(err)
	}
	service := &BusinessService{client: client}

	_, err = service.Configure(BusinessServiceConfig{
		BaseURL:             candidate.URL,
		PreserveCredentials: true,
	})
	if err == nil || !strings.Contains(err.Error(), "cannot be carried") {
		t.Fatalf("cross-origin preserve error=%v", err)
	}
	if candidateRequests != 0 {
		t.Fatalf("candidate requests after blocked preserve=%d", candidateRequests)
	}

	_, err = service.Configure(BusinessServiceConfig{BaseURL: candidate.URL})
	if err == nil || !strings.Contains(err.Error(), "candidate is not ready") {
		t.Fatalf("unready candidate error=%v", err)
	}
	if service.snapshotClient() != client {
		t.Fatal("failed candidate replaced the working client")
	}
}

func TestBusinessServiceAppliesExactReviewedImportWithStableOperation(t *testing.T) {
	var applyKey string
	var applyRequests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request documentImportRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatal(err)
		}
		if len(request.Documents) != 1 {
			t.Fatalf("documents=%+v", request.Documents)
		}
		if !request.DryRun {
			applyRequests++
			if request.Documents[0].Content != "reviewed bytes" {
				t.Fatalf("applied unreviewed content %q", request.Documents[0].Content)
			}
			if applyKey == "" {
				applyKey = r.Header.Get("Idempotency-Key")
			} else if got := r.Header.Get("Idempotency-Key"); got != applyKey {
				t.Fatalf("apply key changed from %q to %q", applyKey, got)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"import","business_id":"business","source_name":"source","source_hash":"hash","dry_run":` + strconv.FormatBool(request.DryRun) + `,"created_by":"actor","created_at":"2026-01-01T00:00:00Z"}`))
	}))
	defer server.Close()

	client, err := newBusinessHTTPClient(BusinessServiceConfig{BaseURL: server.URL}, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	service := &BusinessService{client: client}
	directory := t.TempDir()
	path := filepath.Join(directory, "record.md")
	if err := os.WriteFile(path, []byte("reviewed bytes"), 0o600); err != nil {
		t.Fatal(err)
	}
	review, err := service.ReviewImportSource("business", directory)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("changed after review"), 0o600); err != nil {
		t.Fatal(err)
	}
	for range 2 {
		report, err := service.ApplyReviewedImport("business", review.ReviewID)
		if err != nil {
			t.Fatal(err)
		}
		if report.DryRun {
			t.Fatalf("apply report=%+v", report)
		}
	}
	if applyRequests != 2 || len(applyKey) < 8 {
		t.Fatalf("apply requests=%d key=%q", applyRequests, applyKey)
	}
}

func TestBusinessServiceLiveReadAndDryRun(t *testing.T) {
	if os.Getenv("BUSINESS_E2E") != "1" {
		t.Skip("set BUSINESS_E2E=1 to exercise the local business service")
	}

	service := NewBusinessService()
	status := service.Status()
	if !status.Ready {
		t.Fatalf("business service is not ready: %s", status.Message)
	}
	businesses, err := service.ListBusinesses()
	if err != nil {
		t.Fatal(err)
	}
	if len(businesses) == 0 {
		t.Fatal("live business service has no business for a dry-run target")
	}
	if _, err := service.ListRecordPage(businesses[0].ID, "", "", "", "", 25); err != nil {
		t.Fatal(err)
	}

	directory := t.TempDir()
	markdown := "---\nid: workbench-e2e-preview\ntitle: Workbench E2E Preview\ntype: STR\n---\n\n# Workbench E2E Preview\n\nThis dry run must not write canonical state.\n"
	if err := os.WriteFile(filepath.Join(directory, "preview.md"), []byte(markdown), 0o600); err != nil {
		t.Fatal(err)
	}
	review, err := service.ReviewImportSource(businesses[0].ID, directory)
	if err != nil {
		t.Fatal(err)
	}
	if !review.Report.DryRun {
		t.Fatalf("expected dry-run report: %+v", review.Report)
	}
}
