package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultBusinessServiceURL = "http://localhost:7290"
	maxBusinessResponseBytes  = 36 << 20
	maxBusinessImportBytes    = 32 << 20
	maxBusinessDocumentBytes  = 4 << 20
	maxBusinessDocuments      = 1000
)

// BusinessService is the Wails boundary for the canonical A3T business API.
// Credentials and HTTP behavior stay in Go rather than the webview.
type BusinessService struct {
	mu            sync.RWMutex
	client        *businessHTTPClient
	importMu      sync.Mutex
	importReviews map[string]stagedBusinessImport
}

type BusinessServiceConfig struct {
	BaseURL             string `json:"base_url"`
	AccessToken         string `json:"access_token"`
	TenantID            string `json:"tenant_id"`
	ActorID             string `json:"actor_id"`
	PreserveCredentials bool   `json:"preserve_credentials"`
}

type BusinessServiceStatus struct {
	BaseURL                string `json:"base_url"`
	Connected              bool   `json:"connected"`
	Ready                  bool   `json:"ready"`
	AuthMode               string `json:"auth_mode"`
	HasAccessToken         bool   `json:"has_access_token"`
	HasDevelopmentIdentity bool   `json:"has_development_identity"`
	Message                string `json:"message"`
}

type Business struct {
	ID              string  `json:"id"`
	TenantID        string  `json:"tenant_id"`
	OrganizationID  *string `json:"organization_id,omitempty"`
	Slug            string  `json:"slug"`
	DisplayName     string  `json:"display_name"`
	LegalName       string  `json:"legal_name,omitempty"`
	Description     string  `json:"description,omitempty"`
	Status          string  `json:"status"`
	Visibility      string  `json:"visibility"`
	OntologyVersion string  `json:"ontology_version,omitempty"`
	Version         int64   `json:"version"`
	CreatedBy       string  `json:"created_by"`
	UpdatedBy       string  `json:"updated_by"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type CreateBusinessInput struct {
	Slug            string  `json:"slug"`
	DisplayName     string  `json:"display_name"`
	LegalName       string  `json:"legal_name,omitempty"`
	Description     string  `json:"description,omitempty"`
	Status          string  `json:"status,omitempty"`
	Visibility      string  `json:"visibility,omitempty"`
	OntologyVersion string  `json:"ontology_version,omitempty"`
	OrganizationID  *string `json:"organization_id,omitempty"`
}

type UpdateBusinessInput struct {
	DisplayName     *string `json:"display_name,omitempty"`
	LegalName       *string `json:"legal_name,omitempty"`
	Description     *string `json:"description,omitempty"`
	Status          *string `json:"status,omitempty"`
	Visibility      *string `json:"visibility,omitempty"`
	OntologyVersion *string `json:"ontology_version,omitempty"`
}

type BusinessIdentifier struct {
	ID              string  `json:"id"`
	BusinessID      string  `json:"business_id"`
	Scheme          string  `json:"scheme"`
	Value           string  `json:"value"`
	NormalizedValue string  `json:"normalized_value"`
	Primary         bool    `json:"primary"`
	Version         int64   `json:"version"`
	VerifiedAt      *string `json:"verified_at,omitempty"`
	MetadataJSON    string  `json:"metadata_json,omitempty"`
	CreatedBy       string  `json:"created_by"`
	CreatedAt       string  `json:"created_at"`
}

type BusinessRecord struct {
	ID                string  `json:"id"`
	BusinessID        string  `json:"business_id"`
	Key               string  `json:"key"`
	TypeCode          string  `json:"type_code"`
	CurrentRevisionID *string `json:"current_revision_id,omitempty"`
	CurrentRevision   int64   `json:"current_revision"`
	Version           int64   `json:"version"`
	CreatedBy         string  `json:"created_by"`
	UpdatedBy         string  `json:"updated_by"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type RecordRevision struct {
	ID                string  `json:"id"`
	RecordID          string  `json:"record_id"`
	BusinessID        string  `json:"business_id"`
	Revision          int64   `json:"revision"`
	RecordVersion     int64   `json:"record_version"`
	TypeCode          string  `json:"type_code"`
	Title             string  `json:"title"`
	Status            string  `json:"status,omitempty"`
	Domain            string  `json:"domain,omitempty"`
	Visibility        string  `json:"visibility,omitempty"`
	OwnerRef          string  `json:"owner_ref,omitempty"`
	SchemaVersion     string  `json:"schema_version,omitempty"`
	SourceVersion     string  `json:"source_version,omitempty"`
	DataJSON          string  `json:"data_json"`
	NarrativeMarkdown string  `json:"narrative_markdown,omitempty"`
	RawMarkdown       string  `json:"raw_markdown,omitempty"`
	SourceHash        string  `json:"source_hash,omitempty"`
	ContentHash       string  `json:"content_hash"`
	GraphHash         string  `json:"graph_hash"`
	SourceURI         string  `json:"source_uri,omitempty"`
	SourceMediaType   string  `json:"source_media_type,omitempty"`
	EffectiveAt       *string `json:"effective_at,omitempty"`
	CreatedBy         string  `json:"created_by"`
	CreatedAt         string  `json:"created_at"`
}

type RecordRelationship struct {
	ID             string  `json:"id"`
	RevisionID     string  `json:"revision_id"`
	SourceRecordID string  `json:"source_record_id"`
	TargetRecordID *string `json:"target_record_id,omitempty"`
	TargetKey      string  `json:"target_key"`
	Kind           string  `json:"kind"`
	Resolution     string  `json:"resolution"`
	Strength       string  `json:"strength,omitempty"`
	ProvenanceJSON string  `json:"provenance_json,omitempty"`
	CreatedAt      string  `json:"created_at"`
}

type RecordView struct {
	Record        BusinessRecord       `json:"record"`
	Revision      RecordRevision       `json:"revision"`
	Relationships []RecordRelationship `json:"relationships"`
}

type RelationshipInput struct {
	Kind           string `json:"kind"`
	TargetKey      string `json:"target_key"`
	Strength       string `json:"strength,omitempty"`
	ProvenanceJSON string `json:"provenance_json,omitempty"`
}

type RevisionInput struct {
	Title             string              `json:"title"`
	Status            string              `json:"status,omitempty"`
	Domain            string              `json:"domain,omitempty"`
	Visibility        string              `json:"visibility,omitempty"`
	OwnerRef          string              `json:"owner_ref,omitempty"`
	SchemaVersion     string              `json:"schema_version,omitempty"`
	SourceVersion     string              `json:"source_version,omitempty"`
	DataJSON          string              `json:"data_json"`
	NarrativeMarkdown string              `json:"narrative_markdown,omitempty"`
	RawMarkdown       string              `json:"raw_markdown,omitempty"`
	SourceURI         string              `json:"source_uri,omitempty"`
	SourceMediaType   string              `json:"source_media_type,omitempty"`
	EffectiveAt       *string             `json:"effective_at,omitempty"`
	Relationships     []RelationshipInput `json:"relationships"`
}

type CreateRecordInput struct {
	Key      string        `json:"key"`
	TypeCode string        `json:"type_code"`
	Revision RevisionInput `json:"revision"`
}

type ReviseRecordInput struct {
	TypeCode string        `json:"type_code,omitempty"`
	Revision RevisionInput `json:"revision"`
}

func (identifier *BusinessIdentifier) UnmarshalJSON(data []byte) error {
	type identifierAlias BusinessIdentifier
	var decoded identifierAlias
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var raw struct {
		Metadata json.RawMessage `json:"metadata"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*identifier = BusinessIdentifier(decoded)
	if len(raw.Metadata) > 0 {
		identifier.MetadataJSON = string(raw.Metadata)
	}
	return nil
}

func (revision *RecordRevision) UnmarshalJSON(data []byte) error {
	type revisionAlias RecordRevision
	var decoded revisionAlias
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var raw struct {
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*revision = RecordRevision(decoded)
	if len(raw.Data) > 0 {
		revision.DataJSON = string(raw.Data)
	}
	if revision.DataJSON == "" {
		revision.DataJSON = "{}"
	}
	return nil
}

func (relationship *RecordRelationship) UnmarshalJSON(data []byte) error {
	type relationshipAlias RecordRelationship
	var decoded relationshipAlias
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var raw struct {
		Provenance json.RawMessage `json:"provenance"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*relationship = RecordRelationship(decoded)
	if len(raw.Provenance) > 0 {
		relationship.ProvenanceJSON = string(raw.Provenance)
	}
	return nil
}

func (input RelationshipInput) MarshalJSON() ([]byte, error) {
	var provenance json.RawMessage
	if strings.TrimSpace(input.ProvenanceJSON) != "" {
		var err error
		provenance, err = validatedRawJSON("relationship provenance", input.ProvenanceJSON, "")
		if err != nil {
			return nil, err
		}
	}
	return json.Marshal(struct {
		Kind       string          `json:"kind"`
		TargetKey  string          `json:"target_key"`
		Strength   string          `json:"strength,omitempty"`
		Provenance json.RawMessage `json:"provenance,omitempty"`
	}{
		Kind:       input.Kind,
		TargetKey:  input.TargetKey,
		Strength:   input.Strength,
		Provenance: provenance,
	})
}

func (input RevisionInput) MarshalJSON() ([]byte, error) {
	data, err := validatedRawJSON("record data", input.DataJSON, "{}")
	if err != nil {
		return nil, err
	}
	return json.Marshal(struct {
		Title             string              `json:"title"`
		Status            string              `json:"status,omitempty"`
		Domain            string              `json:"domain,omitempty"`
		Visibility        string              `json:"visibility,omitempty"`
		OwnerRef          string              `json:"owner_ref,omitempty"`
		SchemaVersion     string              `json:"schema_version,omitempty"`
		SourceVersion     string              `json:"source_version,omitempty"`
		Data              json.RawMessage     `json:"data"`
		NarrativeMarkdown string              `json:"narrative_markdown,omitempty"`
		RawMarkdown       string              `json:"raw_markdown,omitempty"`
		SourceURI         string              `json:"source_uri,omitempty"`
		SourceMediaType   string              `json:"source_media_type,omitempty"`
		EffectiveAt       *string             `json:"effective_at,omitempty"`
		Relationships     []RelationshipInput `json:"relationships"`
	}{
		Title:             input.Title,
		Status:            input.Status,
		Domain:            input.Domain,
		Visibility:        input.Visibility,
		OwnerRef:          input.OwnerRef,
		SchemaVersion:     input.SchemaVersion,
		SourceVersion:     input.SourceVersion,
		Data:              data,
		NarrativeMarkdown: input.NarrativeMarkdown,
		RawMarkdown:       input.RawMarkdown,
		SourceURI:         input.SourceURI,
		SourceMediaType:   input.SourceMediaType,
		EffectiveAt:       input.EffectiveAt,
		Relationships:     input.Relationships,
	})
}

func validatedRawJSON(field, value, fallback string) (json.RawMessage, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		value = fallback
	}
	if value == "" {
		return nil, nil
	}
	raw := json.RawMessage(value)
	if !json.Valid(raw) {
		return nil, fmt.Errorf("%s must be valid JSON", field)
	}
	return raw, nil
}

type GraphNode struct {
	ID         string `json:"id"`
	Key        string `json:"key"`
	TypeCode   string `json:"type_code"`
	Title      string `json:"title"`
	Status     string `json:"status,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Version    int64  `json:"version"`
}

type GraphEdge struct {
	ID        string  `json:"id"`
	SourceID  string  `json:"source_id"`
	TargetID  *string `json:"target_id,omitempty"`
	TargetRef string  `json:"target_ref,omitempty"`
	Kind      string  `json:"kind"`
	Strength  string  `json:"strength,omitempty"`
}

type Neighborhood struct {
	BusinessID       string      `json:"business_id"`
	RootRecordID     string      `json:"root_record_id"`
	Depth            int         `json:"depth"`
	Nodes            []GraphNode `json:"nodes"`
	Edges            []GraphEdge `json:"edges"`
	ProjectedThrough string      `json:"projected_through,omitempty"`
	Truncated        bool        `json:"truncated"`
}

type PathResult struct {
	BusinessID string      `json:"business_id"`
	From       string      `json:"from"`
	To         string      `json:"to"`
	Nodes      []GraphNode `json:"nodes"`
	Edges      []GraphEdge `json:"edges"`
	Found      bool        `json:"found"`
}

type ProjectionStatus struct {
	BusinessID    string `json:"business_id"`
	PendingEvents int64  `json:"pending_events"`
	Status        string `json:"status"`
}

type ImportDiagnostic struct {
	Path     string `json:"path,omitempty"`
	Severity string `json:"severity"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

type ImportReport struct {
	ID             string             `json:"id"`
	BusinessID     string             `json:"business_id"`
	SourceName     string             `json:"source_name"`
	SourceHash     string             `json:"source_hash"`
	CreatedRecords int                `json:"created_records"`
	RevisedRecords int                `json:"revised_records"`
	Unchanged      int                `json:"unchanged_records"`
	ResolvedEdges  int                `json:"resolved_relationships"`
	UnresolvedRefs int                `json:"unresolved_references"`
	Diagnostics    []ImportDiagnostic `json:"diagnostics,omitempty"`
	DryRun         bool               `json:"dry_run"`
	CreatedBy      string             `json:"created_by"`
	CreatedAt      string             `json:"created_at"`
}

type ImportReviewResult struct {
	ReviewID string       `json:"review_id"`
	Report   ImportReport `json:"report"`
}

type stagedBusinessImport struct {
	BusinessID     string
	ServiceURL     string
	SourceName     string
	Documents      []documentImport
	Archive        []byte
	ApplyOperation string
	ReviewedAt     time.Time
}

type businessPage struct {
	Items      []Business `json:"items"`
	NextCursor string     `json:"next_cursor,omitempty"`
}

type RecordPage struct {
	Items      []RecordView `json:"items"`
	NextCursor string       `json:"next_cursor,omitempty"`
}

type RevisionPage struct {
	Items      []RecordRevision `json:"items"`
	NextCursor string           `json:"next_cursor,omitempty"`
}

type identifierList struct {
	Items []BusinessIdentifier `json:"items"`
}

type documentImport struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

type documentImportRequest struct {
	SourceName      string           `json:"source_name"`
	OntologyVersion string           `json:"ontology_version,omitempty"`
	Documents       []documentImport `json:"documents"`
	DryRun          bool             `json:"dry_run"`
}

type businessErrorEnvelope struct {
	Error struct {
		Code      string         `json:"code"`
		Message   string         `json:"message"`
		RequestID string         `json:"request_id"`
		TenantID  string         `json:"tenant_id"`
		Details   map[string]any `json:"details"`
	} `json:"error"`
}

type BusinessAPIError struct {
	StatusCode int
	Code       string
	Message    string
	RequestID  string
	Details    map[string]any
}

func (e *BusinessAPIError) Error() string {
	message := strings.TrimSpace(e.Message)
	if message == "" {
		message = http.StatusText(e.StatusCode)
	}
	if e.Code != "" {
		return e.Code + ": " + message
	}
	return fmt.Sprintf("business service returned HTTP %d: %s", e.StatusCode, message)
}

type businessHTTPClient struct {
	baseURL     *url.URL
	accessToken string
	tenantID    string
	actorID     string
	httpClient  *http.Client
}

func NewBusinessService() *BusinessService {
	client, err := newBusinessHTTPClient(defaultBusinessConfig(), nil)
	if err != nil {
		client, _ = newBusinessHTTPClient(BusinessServiceConfig{BaseURL: defaultBusinessServiceURL}, nil)
	}
	return &BusinessService{client: client, importReviews: make(map[string]stagedBusinessImport)}
}

func defaultBusinessConfig() BusinessServiceConfig {
	baseURL := firstEnvironmentValue("A3T_BUSINESS_API_URL", "BUSINESS_API_URL")
	if baseURL == "" {
		baseURL = defaultBusinessServiceURL
	}
	return BusinessServiceConfig{
		BaseURL:     baseURL,
		AccessToken: firstEnvironmentValue("A3T_BUSINESS_ACCESS_TOKEN", "A3T_ACCESS_TOKEN"),
		TenantID:    firstEnvironmentValue("A3T_BUSINESS_TENANT_ID", "A3T_TENANT_ID"),
		ActorID:     firstEnvironmentValue("A3T_BUSINESS_ACTOR_ID", "A3T_ACTOR_ID"),
	}
}

func firstEnvironmentValue(names ...string) string {
	for _, name := range names {
		if value := strings.TrimSpace(os.Getenv(name)); value != "" {
			return value
		}
	}
	return ""
}

func newBusinessHTTPClient(config BusinessServiceConfig, httpClient *http.Client) (*businessHTTPClient, error) {
	baseURL, err := normalizeBusinessBaseURL(config.BaseURL)
	if err != nil {
		return nil, err
	}
	accessToken := strings.TrimSpace(config.AccessToken)
	if accessToken != "" && baseURL.Scheme == "http" && !isLoopbackHostname(baseURL.Hostname()) {
		return nil, errors.New("bearer credentials require HTTPS unless the service is on the local machine")
	}
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 2 * time.Minute}
	}
	return &businessHTTPClient{
		baseURL:     baseURL,
		accessToken: accessToken,
		tenantID:    strings.TrimSpace(config.TenantID),
		actorID:     strings.TrimSpace(config.ActorID),
		httpClient:  httpClient,
	}, nil
}

func isLoopbackHostname(hostname string) bool {
	hostname = strings.TrimSuffix(strings.ToLower(strings.TrimSpace(hostname)), ".")
	if hostname == "localhost" || strings.HasSuffix(hostname, ".localhost") {
		return true
	}
	address := net.ParseIP(hostname)
	return address != nil && address.IsLoopback()
}

func normalizeBusinessBaseURL(raw string) (*url.URL, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		raw = defaultBusinessServiceURL
	}
	parsed, err := url.Parse(raw)
	if err != nil {
		return nil, fmt.Errorf("invalid business service URL: %w", err)
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return nil, errors.New("business service URL must use http or https")
	}
	if parsed.Host == "" || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" {
		return nil, errors.New("business service URL must contain only a scheme, host, and optional base path")
	}
	parsed.Path = strings.TrimSuffix(parsed.Path, "/")
	parsed.Path = strings.TrimSuffix(parsed.Path, "/api/v1")
	return parsed, nil
}

func (s *BusinessService) snapshotClient() *businessHTTPClient {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.client
}

func (s *BusinessService) Configure(config BusinessServiceConfig) (*BusinessServiceStatus, error) {
	current := s.snapshotClient()
	if config.PreserveCredentials {
		candidateURL, err := normalizeBusinessBaseURL(config.BaseURL)
		if err != nil {
			return nil, err
		}
		if candidateURL.String() != current.baseURL.String() {
			return nil, errors.New("hidden credentials cannot be carried to a different service URL; enter the new credentials or turn off credential preservation")
		}
		if strings.TrimSpace(config.AccessToken) == "" {
			config.AccessToken = current.accessToken
		}
		if strings.TrimSpace(config.TenantID) == "" {
			config.TenantID = current.tenantID
		}
		if strings.TrimSpace(config.ActorID) == "" {
			config.ActorID = current.actorID
		}
	}
	client, err := newBusinessHTTPClient(config, nil)
	if err != nil {
		return nil, err
	}
	status := businessServiceStatus(client)
	if !status.Ready {
		return nil, errors.New(status.Message)
	}
	s.mu.Lock()
	s.client = client
	s.mu.Unlock()
	s.importMu.Lock()
	s.importReviews = make(map[string]stagedBusinessImport)
	s.importMu.Unlock()
	return status, nil
}

func (s *BusinessService) Status() *BusinessServiceStatus {
	return businessServiceStatus(s.snapshotClient())
}

func businessServiceStatus(client *businessHTTPClient) *BusinessServiceStatus {
	status := &BusinessServiceStatus{
		BaseURL:                client.baseURL.String(),
		AuthMode:               client.authMode(),
		HasAccessToken:         client.accessToken != "",
		HasDevelopmentIdentity: client.tenantID != "" || client.actorID != "",
		Message:                "Business service is not reachable",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	var liveness map[string]any
	if err := client.doJSON(ctx, http.MethodGet, "/live", nil, nil, nil, &liveness); err != nil {
		status.Message = err.Error()
		return status
	}
	status.Connected = true

	var readiness map[string]any
	if err := client.doJSON(ctx, http.MethodGet, "/ready", nil, nil, nil, &readiness); err != nil {
		status.Message = err.Error()
		return status
	}
	status.Ready = true
	status.Message = "Canonical store ready"
	return status
}

func (c *businessHTTPClient) authMode() string {
	if c.accessToken != "" {
		return "bearer"
	}
	if c.tenantID != "" || c.actorID != "" {
		return "development headers"
	}
	return "service development identity"
}

func (s *BusinessService) ListBusinesses() ([]Business, error) {
	client := s.snapshotClient()
	var all []Business
	cursor := ""
	seen := map[string]bool{}
	for pageNumber := 0; pageNumber < 1000; pageNumber++ {
		query := url.Values{"limit": {"200"}}
		if cursor != "" {
			query.Set("cursor", cursor)
		}
		var page businessPage
		if err := client.doJSON(context.Background(), http.MethodGet, "/api/v1/businesses", query, nil, nil, &page); err != nil {
			return nil, err
		}
		all = append(all, page.Items...)
		if page.NextCursor == "" {
			return all, nil
		}
		if seen[page.NextCursor] {
			return nil, errors.New("business service returned a repeated pagination cursor")
		}
		seen[page.NextCursor] = true
		cursor = page.NextCursor
	}
	return nil, errors.New("business pagination exceeded the safety limit")
}

func (s *BusinessService) CreateBusiness(input CreateBusinessInput, operationID string) (*Business, error) {
	client := s.snapshotClient()
	var result Business
	idempotencyKey, err := businessIdempotencyKey(operationID)
	if err != nil {
		return nil, err
	}
	headers := http.Header{"Idempotency-Key": {idempotencyKey}}
	if err := client.doJSON(context.Background(), http.MethodPost, "/api/v1/businesses", nil, headers, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) GetBusiness(businessID string) (*Business, error) {
	client := s.snapshotClient()
	var result Business
	path := "/api/v1/businesses/" + url.PathEscape(strings.TrimSpace(businessID))
	if err := client.doJSON(context.Background(), http.MethodGet, path, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) UpdateBusiness(businessID string, expectedVersion int64, input UpdateBusinessInput) (*Business, error) {
	client := s.snapshotClient()
	headers := http.Header{"If-Match": {versionETag(expectedVersion)}}
	var result Business
	path := "/api/v1/businesses/" + url.PathEscape(strings.TrimSpace(businessID))
	if err := client.doJSON(context.Background(), http.MethodPatch, path, nil, headers, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) ListIdentifiers(businessID string) ([]BusinessIdentifier, error) {
	client := s.snapshotClient()
	var result identifierList
	path := businessPath(businessID) + "/identifiers"
	if err := client.doJSON(context.Background(), http.MethodGet, path, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return result.Items, nil
}

func (s *BusinessService) ListRecords(businessID, search, typeCode, status string) ([]RecordView, error) {
	client := s.snapshotClient()
	var all []RecordView
	cursor := ""
	seen := map[string]bool{}
	for pageNumber := 0; pageNumber < 1000; pageNumber++ {
		query := url.Values{"limit": {"200"}}
		if cursor != "" {
			query.Set("cursor", cursor)
		}
		if value := strings.TrimSpace(search); value != "" {
			query.Set("q", value)
		}
		if value := strings.TrimSpace(typeCode); value != "" {
			query.Set("type_code", value)
		}
		if value := strings.TrimSpace(status); value != "" {
			query.Set("status", value)
		}
		var page RecordPage
		if err := client.doJSON(context.Background(), http.MethodGet, businessPath(businessID)+"/records", query, nil, nil, &page); err != nil {
			return nil, err
		}
		all = append(all, page.Items...)
		if page.NextCursor == "" {
			return all, nil
		}
		if seen[page.NextCursor] {
			return nil, errors.New("business service returned a repeated record cursor")
		}
		seen[page.NextCursor] = true
		cursor = page.NextCursor
	}
	return nil, errors.New("record pagination exceeded the safety limit")
}

func (s *BusinessService) ListRecordPage(businessID, search, typeCode, status, cursor string, limit int) (*RecordPage, error) {
	if limit <= 0 {
		limit = 100
	}
	if limit > 200 {
		limit = 200
	}
	query := url.Values{"limit": {strconv.Itoa(limit)}}
	if value := strings.TrimSpace(cursor); value != "" {
		query.Set("cursor", value)
	}
	if value := strings.TrimSpace(search); value != "" {
		query.Set("q", value)
	}
	if value := strings.TrimSpace(typeCode); value != "" {
		query.Set("type_code", value)
	}
	if value := strings.TrimSpace(status); value != "" {
		query.Set("status", value)
	}
	var page RecordPage
	if err := s.snapshotClient().doJSON(context.Background(), http.MethodGet, businessPath(businessID)+"/records", query, nil, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (s *BusinessService) GetRecord(businessID, recordID string) (*RecordView, error) {
	client := s.snapshotClient()
	var result RecordView
	path := businessPath(businessID) + "/records/" + url.PathEscape(strings.TrimSpace(recordID))
	if err := client.doJSON(context.Background(), http.MethodGet, path, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) GetRecordsByIDs(businessID string, recordIDs []string) ([]RecordView, error) {
	unique := make([]string, 0, len(recordIDs))
	seen := make(map[string]bool, len(recordIDs))
	for _, recordID := range recordIDs {
		recordID = strings.TrimSpace(recordID)
		if recordID != "" && !seen[recordID] {
			seen[recordID] = true
			unique = append(unique, recordID)
		}
	}
	if len(unique) > 100 {
		return nil, errors.New("record metadata batch is limited to 100 IDs")
	}
	if len(unique) == 0 {
		return []RecordView{}, nil
	}

	client := s.snapshotClient()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	results := make([]RecordView, len(unique))
	semaphore := make(chan struct{}, 8)
	var wait sync.WaitGroup
	var resultMu sync.Mutex
	var firstErr error
	for index, recordID := range unique {
		wait.Add(1)
		go func() {
			defer wait.Done()
			select {
			case semaphore <- struct{}{}:
				defer func() { <-semaphore }()
			case <-ctx.Done():
				return
			}
			path := businessPath(businessID) + "/records/" + url.PathEscape(recordID)
			if err := client.doJSON(ctx, http.MethodGet, path, nil, nil, nil, &results[index]); err != nil {
				resultMu.Lock()
				if firstErr == nil {
					firstErr = err
					cancel()
				}
				resultMu.Unlock()
			}
		}()
	}
	wait.Wait()
	if firstErr != nil {
		return nil, firstErr
	}
	return results, nil
}

func (s *BusinessService) CreateRecord(businessID string, input CreateRecordInput, operationID string) (*RecordView, error) {
	client := s.snapshotClient()
	var result RecordView
	idempotencyKey, err := businessIdempotencyKey(operationID)
	if err != nil {
		return nil, err
	}
	headers := http.Header{"Idempotency-Key": {idempotencyKey}}
	if err := client.doJSON(context.Background(), http.MethodPost, businessPath(businessID)+"/records", nil, headers, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) ReviseRecord(businessID, recordID string, expectedVersion int64, input ReviseRecordInput) (*RecordView, error) {
	client := s.snapshotClient()
	var result RecordView
	headers := http.Header{"If-Match": {versionETag(expectedVersion)}}
	path := businessPath(businessID) + "/records/" + url.PathEscape(strings.TrimSpace(recordID))
	if err := client.doJSON(context.Background(), http.MethodPatch, path, nil, headers, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) ListRevisions(businessID, recordID string) ([]RecordRevision, error) {
	client := s.snapshotClient()
	var all []RecordRevision
	cursor := ""
	seen := map[string]bool{}
	path := businessPath(businessID) + "/records/" + url.PathEscape(strings.TrimSpace(recordID)) + "/revisions"
	for pageNumber := 0; pageNumber < 1000; pageNumber++ {
		query := url.Values{"limit": {"200"}}
		if cursor != "" {
			query.Set("cursor", cursor)
		}
		var page RevisionPage
		if err := client.doJSON(context.Background(), http.MethodGet, path, query, nil, nil, &page); err != nil {
			return nil, err
		}
		all = append(all, page.Items...)
		if page.NextCursor == "" {
			return all, nil
		}
		if seen[page.NextCursor] {
			return nil, errors.New("business service returned a repeated revision cursor")
		}
		seen[page.NextCursor] = true
		cursor = page.NextCursor
	}
	return nil, errors.New("revision pagination exceeded the safety limit")
}

func (s *BusinessService) ListRevisionPage(businessID, recordID, cursor string, limit int) (*RevisionPage, error) {
	if limit <= 0 {
		limit = 100
	}
	if limit > 200 {
		limit = 200
	}
	query := url.Values{"limit": {strconv.Itoa(limit)}}
	if value := strings.TrimSpace(cursor); value != "" {
		query.Set("cursor", value)
	}
	path := businessPath(businessID) + "/records/" + url.PathEscape(strings.TrimSpace(recordID)) + "/revisions"
	var page RevisionPage
	if err := s.snapshotClient().doJSON(context.Background(), http.MethodGet, path, query, nil, nil, &page); err != nil {
		return nil, err
	}
	return &page, nil
}

func (s *BusinessService) GetRevision(businessID, recordID, revisionID string) (*RecordView, error) {
	client := s.snapshotClient()
	var result RecordView
	path := businessPath(businessID) + "/records/" + url.PathEscape(strings.TrimSpace(recordID)) + "/revisions/" + url.PathEscape(strings.TrimSpace(revisionID))
	if err := client.doJSON(context.Background(), http.MethodGet, path, nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) Neighborhood(businessID, recordID string, depth, limit int, kinds []string) (*Neighborhood, error) {
	client := s.snapshotClient()
	query := url.Values{
		"record_id": {strings.TrimSpace(recordID)},
		"depth":     {strconv.Itoa(depth)},
		"limit":     {strconv.Itoa(limit)},
	}
	for _, kind := range kinds {
		if kind = strings.TrimSpace(kind); kind != "" {
			query.Add("kind", kind)
		}
	}
	var result Neighborhood
	if err := client.doJSON(context.Background(), http.MethodGet, businessPath(businessID)+"/graph/neighborhood", query, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) FindPath(businessID, fromRecordID, toRecordID string, maxDepth int, kinds []string) (*PathResult, error) {
	client := s.snapshotClient()
	query := url.Values{
		"from_record_id": {strings.TrimSpace(fromRecordID)},
		"to_record_id":   {strings.TrimSpace(toRecordID)},
		"max_depth":      {strconv.Itoa(maxDepth)},
	}
	for _, kind := range kinds {
		if kind = strings.TrimSpace(kind); kind != "" {
			query.Add("kind", kind)
		}
	}
	var result PathResult
	if err := client.doJSON(context.Background(), http.MethodGet, businessPath(businessID)+"/graph/paths", query, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) ProjectionStatus(businessID string) (*ProjectionStatus, error) {
	client := s.snapshotClient()
	var result ProjectionStatus
	if err := client.doJSON(context.Background(), http.MethodGet, businessPath(businessID)+"/projection/status", nil, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) ReviewImportSource(businessID, inputPath string) (*ImportReviewResult, error) {
	absPath, info, err := resolveBusinessImportPath(inputPath)
	if err != nil {
		return nil, err
	}
	client := s.snapshotClient()
	staged := stagedBusinessImport{
		BusinessID:     strings.TrimSpace(businessID),
		ServiceURL:     client.baseURL.String(),
		SourceName:     filepath.Base(absPath),
		ApplyOperation: newBusinessIdempotencyKey(),
		ReviewedAt:     time.Now(),
	}
	if info.IsDir() {
		staged.SourceName = filepath.Base(absPath)
		staged.Documents, err = readBusinessImportDirectory(absPath)
	} else if strings.EqualFold(filepath.Ext(absPath), ".bspec") {
		staged.Archive, err = readBusinessImportArchive(absPath, info)
	} else {
		return nil, errors.New("choose a .bspec archive or a directory containing BSpec Markdown")
	}
	if err != nil {
		return nil, err
	}

	report, err := s.sendStagedImport(client, staged, true, newBusinessIdempotencyKey())
	if err != nil {
		return nil, err
	}
	reviewID := newBusinessIdempotencyKey()
	s.storeImportReview(reviewID, staged)
	return &ImportReviewResult{ReviewID: reviewID, Report: *report}, nil
}

func (s *BusinessService) ApplyReviewedImport(businessID, reviewID string) (*ImportReport, error) {
	staged, err := s.importReview(strings.TrimSpace(reviewID))
	if err != nil {
		return nil, err
	}
	if staged.BusinessID != strings.TrimSpace(businessID) {
		return nil, errors.New("the reviewed import belongs to a different business")
	}
	client := s.snapshotClient()
	if staged.ServiceURL != client.baseURL.String() {
		return nil, errors.New("the business service connection changed after review; run the dry run again")
	}
	return s.sendStagedImport(client, staged, false, staged.ApplyOperation)
}

func resolveBusinessImportPath(inputPath string) (string, fs.FileInfo, error) {
	inputPath = strings.TrimSpace(inputPath)
	if inputPath == "" {
		return "", nil, errors.New("import path is required")
	}
	if strings.HasPrefix(inputPath, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", nil, err
		}
		inputPath = filepath.Join(home, strings.TrimPrefix(inputPath, "~/"))
	}
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		return "", nil, err
	}
	info, err := os.Stat(absPath)
	if err != nil {
		return "", nil, err
	}
	return absPath, info, nil
}

func readBusinessImportDirectory(root string) ([]documentImport, error) {
	documents := make([]documentImport, 0)
	totalBytes := int64(0)
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		if strings.HasPrefix(entry.Name(), "._") {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".md" && ext != ".markdown" {
			return nil
		}
		if len(documents) >= maxBusinessDocuments {
			return fmt.Errorf("import contains more than %d Markdown documents", maxBusinessDocuments)
		}
		if info.Size() > maxBusinessDocumentBytes {
			return fmt.Errorf("%s exceeds the 4 MiB document limit", path)
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		totalBytes += int64(len(content))
		if totalBytes > maxBusinessImportBytes-(1<<20) {
			return errors.New("Markdown import is too large for the 32 MiB service request limit")
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		documents = append(documents, documentImport{Path: filepath.ToSlash(relative), Content: string(content)})
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(documents) == 0 {
		return nil, errors.New("selected directory contains no Markdown documents")
	}
	return documents, nil
}

func readBusinessImportArchive(archivePath string, info fs.FileInfo) ([]byte, error) {
	if info.Size() > maxBusinessImportBytes-(1<<20) {
		return nil, errors.New("BSpec archive is too large for the 32 MiB service request limit")
	}
	archive, err := os.ReadFile(archivePath)
	if err != nil {
		return nil, err
	}
	if len(archive) > maxBusinessImportBytes-(1<<20) {
		return nil, errors.New("BSpec archive is too large for the 32 MiB service request limit")
	}
	return archive, nil
}

func (s *BusinessService) sendStagedImport(client *businessHTTPClient, staged stagedBusinessImport, dryRun bool, operationID string) (*ImportReport, error) {
	if len(staged.Documents) > 0 {
		request := documentImportRequest{
			SourceName: staged.SourceName,
			Documents:  staged.Documents,
			DryRun:     dryRun,
		}
		var result ImportReport
		headers := http.Header{"Idempotency-Key": {operationID}}
		if err := client.doJSON(context.Background(), http.MethodPost, businessPath(staged.BusinessID)+"/imports", nil, headers, request, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	if len(staged.Archive) == 0 {
		return nil, errors.New("reviewed import payload is empty")
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	partHeader := make(textproto.MIMEHeader)
	partHeader.Set("Content-Disposition", mime.FormatMediaType("form-data", map[string]string{
		"name":     "archive",
		"filename": staged.SourceName,
	}))
	partHeader.Set("Content-Type", "application/gzip")
	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return nil, err
	}
	if _, err := part.Write(staged.Archive); err != nil {
		return nil, err
	}
	if err := writer.WriteField("source_name", staged.SourceName); err != nil {
		return nil, err
	}
	if err := writer.WriteField("dry_run", strconv.FormatBool(dryRun)); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	var result ImportReport
	headers := http.Header{
		"Content-Type":    {writer.FormDataContentType()},
		"Idempotency-Key": {operationID},
	}
	if err := client.doRaw(context.Background(), http.MethodPost, businessPath(staged.BusinessID)+"/imports", nil, headers, &body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BusinessService) storeImportReview(reviewID string, staged stagedBusinessImport) {
	s.importMu.Lock()
	defer s.importMu.Unlock()
	if s.importReviews == nil {
		s.importReviews = make(map[string]stagedBusinessImport)
	}
	cutoff := time.Now().Add(-30 * time.Minute)
	for id, candidate := range s.importReviews {
		if candidate.ReviewedAt.Before(cutoff) {
			delete(s.importReviews, id)
		}
	}
	for len(s.importReviews) >= 4 {
		var oldestID string
		var oldestTime time.Time
		for id, candidate := range s.importReviews {
			if oldestID == "" || candidate.ReviewedAt.Before(oldestTime) {
				oldestID = id
				oldestTime = candidate.ReviewedAt
			}
		}
		delete(s.importReviews, oldestID)
	}
	s.importReviews[reviewID] = staged
}

func (s *BusinessService) importReview(reviewID string) (stagedBusinessImport, error) {
	if reviewID == "" {
		return stagedBusinessImport{}, errors.New("import review ID is required")
	}
	s.importMu.Lock()
	defer s.importMu.Unlock()
	staged, ok := s.importReviews[reviewID]
	if !ok {
		return stagedBusinessImport{}, errors.New("import review is unavailable; run the dry run again")
	}
	if staged.ReviewedAt.Before(time.Now().Add(-30 * time.Minute)) {
		delete(s.importReviews, reviewID)
		return stagedBusinessImport{}, errors.New("import review expired; run the dry run again")
	}
	return staged, nil
}

func businessPath(businessID string) string {
	return "/api/v1/businesses/" + url.PathEscape(strings.TrimSpace(businessID))
}

func versionETag(version int64) string {
	return fmt.Sprintf("\"v%d\"", version)
}

func newBusinessIdempotencyKey() string {
	buffer := make([]byte, 16)
	if _, err := rand.Read(buffer); err != nil {
		return "bspec-" + strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return "bspec-" + hex.EncodeToString(buffer)
}

func businessIdempotencyKey(operationID string) (string, error) {
	operationID = strings.TrimSpace(operationID)
	if operationID == "" {
		return newBusinessIdempotencyKey(), nil
	}
	if len(operationID) < 8 || len(operationID) > 200 {
		return "", errors.New("operation ID must be between 8 and 200 characters")
	}
	return operationID, nil
}

func (c *businessHTTPClient) doJSON(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	headers http.Header,
	input any,
	output any,
) error {
	var body io.Reader
	requestHeaders := cloneHTTPHeader(headers)
	if input != nil {
		payload, err := json.Marshal(input)
		if err != nil {
			return fmt.Errorf("encode business request: %w", err)
		}
		if len(payload) > maxBusinessImportBytes {
			return errors.New("business request exceeds the 32 MiB client limit")
		}
		body = bytes.NewReader(payload)
		requestHeaders.Set("Content-Type", "application/json")
	}
	return c.doRaw(ctx, method, path, query, requestHeaders, body, output)
}

func (c *businessHTTPClient) doRaw(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	headers http.Header,
	body io.Reader,
	output any,
) error {
	endpoint := *c.baseURL
	endpoint.Path = strings.TrimSuffix(c.baseURL.Path, "/") + "/" + strings.TrimPrefix(path, "/")
	endpoint.RawQuery = query.Encode()

	request, err := http.NewRequestWithContext(ctx, method, endpoint.String(), body)
	if err != nil {
		return fmt.Errorf("create business request: %w", err)
	}
	for name, values := range headers {
		for _, value := range values {
			request.Header.Add(name, value)
		}
	}
	request.Header.Set("Accept", "application/json")
	if c.accessToken != "" {
		request.Header.Set("Authorization", "Bearer "+c.accessToken)
	} else {
		if c.tenantID != "" {
			request.Header.Set("X-A3T-Tenant-ID", c.tenantID)
		}
		if c.actorID != "" {
			request.Header.Set("X-A3T-Actor-ID", c.actorID)
		}
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("business service request failed: %w", err)
	}
	defer response.Body.Close()

	payload, err := io.ReadAll(io.LimitReader(response.Body, maxBusinessResponseBytes+1))
	if err != nil {
		return fmt.Errorf("read business response: %w", err)
	}
	if len(payload) > maxBusinessResponseBytes {
		return errors.New("business service response exceeded the 36 MiB client limit")
	}
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return decodeBusinessAPIError(response, payload)
	}
	if output == nil || len(bytes.TrimSpace(payload)) == 0 {
		return nil
	}
	if err := json.Unmarshal(payload, output); err != nil {
		return fmt.Errorf("decode business response: %w", err)
	}
	return nil
}

func decodeBusinessAPIError(response *http.Response, payload []byte) error {
	var envelope businessErrorEnvelope
	if err := json.Unmarshal(payload, &envelope); err == nil && (envelope.Error.Code != "" || envelope.Error.Message != "") {
		return &BusinessAPIError{
			StatusCode: response.StatusCode,
			Code:       envelope.Error.Code,
			Message:    envelope.Error.Message,
			RequestID:  envelope.Error.RequestID,
			Details:    envelope.Error.Details,
		}
	}
	message := strings.TrimSpace(string(payload))
	if len(message) > 512 {
		message = message[:512]
	}
	return &BusinessAPIError{
		StatusCode: response.StatusCode,
		Message:    message,
		RequestID:  response.Header.Get("X-Request-ID"),
	}
}

func cloneHTTPHeader(source http.Header) http.Header {
	if source == nil {
		return make(http.Header)
	}
	return source.Clone()
}
