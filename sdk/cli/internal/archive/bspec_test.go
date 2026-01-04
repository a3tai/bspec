package archive

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestExtract(t *testing.T) {
	// We can't easily test Extract without creating real .bspec files
	// So we'll test the API exists and handles basic error cases
	tmpDir := t.TempDir()

	// Test with non-existent file
	err := Extract("non-existent.bspec", tmpDir)
	if err == nil {
		t.Error("Expected error when extracting non-existent file")
	}
}

func TestPack(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a basic project structure
	projectDir := filepath.Join(tmpDir, "test-project")
	os.MkdirAll(projectDir, 0755)
	os.MkdirAll(filepath.Join(projectDir, "documents"), 0755)

	// Create manifest.json
	manifest := map[string]interface{}{
		"format_version":   "1.0",
		"bspec_version":    "1.0.0",
		"name":             "Test Project",
		"conformance_level": "standard",
	}
	manifestData, _ := json.MarshalIndent(manifest, "", "  ")
	os.WriteFile(filepath.Join(projectDir, "manifest.json"), manifestData, 0644)

	// Create a test document
	doc := `---
id: TST-test
title: Test Document
type: TST
status: Draft
owner: Test Owner
version: 1.0.0
---

# Test Document
`
	os.WriteFile(filepath.Join(projectDir, "documents", "TST-test.md"), []byte(doc), 0644)

	// Test packing
	bspecPath := filepath.Join(tmpDir, "test.bspec")
	err := Pack(projectDir, bspecPath)
	if err != nil {
		t.Errorf("Unexpected error packing directory: %v", err)
	}

	// Verify the .bspec file was created
	if _, err := os.Stat(bspecPath); os.IsNotExist(err) {
		t.Error("Expected .bspec file to be created")
	}
}

func TestRead(t *testing.T) {
	// Test reading non-existent file
	_, err := Read("non-existent.bspec")
	if err == nil {
		t.Error("Expected error when reading non-existent file")
	}
}

func TestParseDocument(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		wantErr  bool
		verify   func(*BSpecDocument, *testing.T)
	}{
		{
			name: "valid document with frontmatter",
			content: `---
id: TST-valid
title: Valid Test Document
type: TST
status: Draft
owner: Test Owner
version: 1.0.0
domain: test
---

# Valid Test Document

This is a valid test document with proper frontmatter.
`,
			wantErr: false,
			verify: func(doc *BSpecDocument, t *testing.T) {
				if doc.ID != "TST-valid" {
					t.Errorf("Expected ID 'TST-valid', got '%s'", doc.ID)
				}
				if doc.Title != "Valid Test Document" {
					t.Errorf("Expected title 'Valid Test Document', got '%s'", doc.Title)
				}
				if doc.Type != "TST" {
					t.Errorf("Expected type 'TST', got '%s'", doc.Type)
				}
				if doc.Status != "Draft" {
					t.Errorf("Expected status 'Draft', got '%s'", doc.Status)
				}
				if doc.Content == "" {
					t.Error("Expected content to have value")
				}
			},
		},
		{
			name: "document without frontmatter",
			content: `# Document Without Frontmatter

This document doesn't have YAML frontmatter.
`,
			wantErr: true,
			verify:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := parseDocument([]byte(tt.content))

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Run verification if provided
			if !tt.wantErr && tt.verify != nil && doc != nil {
				tt.verify(doc, t)
			}
		})
	}
}

func TestManifestStruct(t *testing.T) {
	// Test that the Manifest struct can be properly marshaled/unmarshaled
	manifest := Manifest{
		FormatVersion:   "1.0",
		BSpecVersion:    "1.0.0",
		Name:            "Test Project",
		Description:     "A test project",
		ConformanceLevel: "standard",
		IndustryProfile: "generic",
	}

	// Marshal to JSON
	data, err := json.Marshal(manifest)
	if err != nil {
		t.Errorf("Failed to marshal manifest: %v", err)
	}

	// Unmarshal back
	var unmarshaled Manifest
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal manifest: %v", err)
	}

	// Verify fields
	if unmarshaled.Name != manifest.Name {
		t.Errorf("Expected name '%s', got '%s'", manifest.Name, unmarshaled.Name)
	}
}

func TestBSpecDocumentStruct(t *testing.T) {
	// Test that the BSpecDocument struct works correctly
	doc := BSpecDocument{
		ID:      "TST-test",
		Title:   "Test Document",
		Type:    "TST",
		Status:  "Draft",
		Version: "1.0.0",
		Owner:   "Test Owner",
		Content: "# Test Document\n\nThis is a test.",
	}

	// Marshal to JSON
	data, err := json.Marshal(doc)
	if err != nil {
		t.Errorf("Failed to marshal document: %v", err)
	}

	// Unmarshal back
	var unmarshaled BSpecDocument
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal document: %v", err)
	}

	// Verify fields
	if unmarshaled.ID != doc.ID {
		t.Errorf("Expected ID '%s', got '%s'", doc.ID, unmarshaled.ID)
	}
	if unmarshaled.Title != doc.Title {
		t.Errorf("Expected title '%s', got '%s'", doc.Title, unmarshaled.Title)
	}
}