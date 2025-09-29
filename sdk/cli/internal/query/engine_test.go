package query

import (
	"testing"

	"github.com/a3tai/bspec/cli/internal/archive"
)

func TestNewQueryEngine(t *testing.T) {
	// Create a test archive
	testArchive := &archive.BSpecArchive{
		Manifest: archive.Manifest{
			Name: "Test Project",
		},
		Documents: map[string]archive.BSpecDocument{
			"test.md": {
				ID:    "TST-test",
				Title: "Test Document",
				Type:  "TST",
			},
		},
	}

	engine := NewQueryEngine(testArchive)
	if engine == nil {
		t.Error("Expected query engine to be non-nil")
	}
	if engine.archive != testArchive {
		t.Error("Expected query engine to contain the test archive")
	}
}

func TestQueryStruct(t *testing.T) {
	// Test that the Query struct can be created with various fields
	query := Query{
		Type:   "TST",
		Domain: "test",
		Status: "Draft",
		Owner:  "Test Owner",
		Tags:   []string{"test", "example"},
		Search: "test content",
	}

	if query.Type != "TST" {
		t.Errorf("Expected type 'TST', got '%s'", query.Type)
	}
	if query.Domain != "test" {
		t.Errorf("Expected domain 'test', got '%s'", query.Domain)
	}
	if query.Status != "Draft" {
		t.Errorf("Expected status 'Draft', got '%s'", query.Status)
	}
	if query.Owner != "Test Owner" {
		t.Errorf("Expected owner 'Test Owner', got '%s'", query.Owner)
	}
	if len(query.Tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(query.Tags))
	}
	if query.Search != "test content" {
		t.Errorf("Expected search 'test content', got '%s'", query.Search)
	}
}

func TestQueryEngineStruct(t *testing.T) {
	// Test the QueryEngine struct
	testArchive := &archive.BSpecArchive{
		Manifest: archive.Manifest{
			Name: "Test Project",
		},
	}

	engine := &QueryEngine{
		archive: testArchive,
	}

	if engine.archive != testArchive {
		t.Error("Expected engine to contain the test archive")
	}
}

func TestEmptyQuery(t *testing.T) {
	// Test that an empty query can be created
	var query Query

	if query.Type != "" {
		t.Errorf("Expected empty type, got '%s'", query.Type)
	}
	if query.Domain != "" {
		t.Errorf("Expected empty domain, got '%s'", query.Domain)
	}
	if len(query.Tags) != 0 {
		t.Errorf("Expected no tags, got %d", len(query.Tags))
	}
}