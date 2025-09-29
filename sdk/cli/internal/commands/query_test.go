package commands

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestQueryCommand(t *testing.T) {
	// Create a temporary test directory
	tmpDir := t.TempDir()

	// Create a test bspec file structure
	testProjectDir := filepath.Join(tmpDir, "test-project")
	err := os.MkdirAll(testProjectDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create manifest.json
	manifest := map[string]interface{}{
		"name":             "Test Project",
		"version":          "1.0.0",
		"conformance_level": "standard",
		"industry_profile": "generic",
	}
	manifestData, _ := json.MarshalIndent(manifest, "", "  ")
	err = os.WriteFile(filepath.Join(testProjectDir, "manifest.json"), manifestData, 0644)
	if err != nil {
		t.Fatalf("Failed to create manifest.json: %v", err)
	}

	// Create test documents
	docsDir := filepath.Join(testProjectDir, "documents")
	err = os.MkdirAll(docsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create documents directory: %v", err)
	}

	// Create a test MSN document
	msnDoc := `---
id: MSN-test-mission
title: Test Mission Statement
type: MSN
status: Draft
owner: Test Owner
domain: strategic
version: 1.0.0
---

# Test Mission Statement

This is a test mission statement for our organization.
`
	err = os.WriteFile(filepath.Join(docsDir, "MSN-test-mission.md"), []byte(msnDoc), 0644)
	if err != nil {
		t.Fatalf("Failed to create test document: %v", err)
	}

	// Create a test CAP document
	capDoc := `---
id: CAP-test-capability
title: Test Capability
type: CAP
status: Accepted
owner: Test Owner
domain: product
version: 1.0.0
---

# Test Capability

This is a test capability document.
`
	err = os.WriteFile(filepath.Join(docsDir, "CAP-test-capability.md"), []byte(capDoc), 0644)
	if err != nil {
		t.Fatalf("Failed to create test capability document: %v", err)
	}

	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		contains string
	}{
		{
			name:     "query help",
			args:     []string{"query", "--help"},
			wantErr:  false,
			contains: "Query BSpec documents with structured queries",
		},
		{
			name:     "query without file",
			args:     []string{"query"},
			wantErr:  true,
			contains: "requires exactly 1 arg",
		},
		{
			name:     "query with non-existent file",
			args:     []string{"query", "non-existent.bspec"},
			wantErr:  true,
			contains: "",
		},
		{
			name:     "query all documents",
			args:     []string{"query", testProjectDir},
			wantErr:  false,
			contains: "MSN-test-mission",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the command state
			resetRootCmd()

			// Capture output
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			output := buf.String()
			if tt.contains != "" && !strings.Contains(output, tt.contains) {
				t.Errorf("Expected output to contain '%s', got: %s", tt.contains, output)
			}
		})
	}
}

func TestQueryFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		flagName string
		expected interface{}
	}{
		{
			name:     "type flag",
			args:     []string{"query", "test.bspec", "--type", "MSN"},
			flagName: "type",
			expected: "MSN",
		},
		{
			name:     "domain flag",
			args:     []string{"query", "test.bspec", "--domain", "strategic"},
			flagName: "domain",
			expected: "strategic",
		},
		{
			name:     "status flag",
			args:     []string{"query", "test.bspec", "--status", "Draft"},
			flagName: "status",
			expected: "Draft",
		},
		{
			name:     "owner flag",
			args:     []string{"query", "test.bspec", "--owner", "John Doe"},
			flagName: "owner",
			expected: "John Doe",
		},
		{
			name:     "limit flag",
			args:     []string{"query", "test.bspec", "--limit", "5"},
			flagName: "limit",
			expected: 5,
		},
		{
			name:     "fields flag",
			args:     []string{"query", "test.bspec", "--fields", "id,title,type"},
			flagName: "fields",
			expected: "id,title,type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test just verifies that flags are parsed correctly
			// We'll use a mock command to avoid file I/O
			testCmd := &cobra.Command{
				Use: "query",
				Run: func(cmd *cobra.Command, args []string) {
					// Check flag values
					switch tt.flagName {
					case "limit":
						if val, _ := cmd.Flags().GetInt(tt.flagName); val != tt.expected {
							t.Errorf("Expected %s to be %v, got %v", tt.flagName, tt.expected, val)
						}
					default:
						if val, _ := cmd.Flags().GetString(tt.flagName); val != tt.expected {
							t.Errorf("Expected %s to be %v, got %v", tt.flagName, tt.expected, val)
						}
					}
				},
			}

			// Add the same flags as the real query command
			testCmd.Flags().StringP("type", "t", "", "Filter by document type")
			testCmd.Flags().StringP("domain", "d", "", "Filter by domain")
			testCmd.Flags().StringP("status", "s", "", "Filter by status")
			testCmd.Flags().StringP("owner", "o", "", "Filter by owner")
			testCmd.Flags().String("search", "", "Text search in document content")
			testCmd.Flags().String("fields", "", "Comma-separated list of fields to include")
			testCmd.Flags().IntP("limit", "l", 0, "Limit number of results")
			testCmd.Flags().String("json", "", "JSON query string")

			// Set args and execute
			testCmd.SetArgs(tt.args[1:]) // Remove "query" from args
			err := testCmd.Execute()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}