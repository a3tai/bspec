package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOpenCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		setup   func(string) string // Setup function returns bspec file/directory path
		verify  func(string, *testing.T) // Verify function with output
	}{
		{
			name:    "open help",
			args:    []string{"open", "--help"},
			wantErr: false,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(output string, t *testing.T) {
				if !strings.Contains(output, "Open and display information about a .bspec file") {
					t.Errorf("Expected help text in output")
				}
			},
		},
		{
			name:    "open without file",
			args:    []string{"open"},
			wantErr: true,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(output string, t *testing.T) {
				// No verification needed for error case
			},
		},
		{
			name:    "open non-existent file",
			args:    []string{"open", "non-existent.bspec"},
			wantErr: true,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(output string, t *testing.T) {
				// No verification needed for error case
			},
		},
		{
			name:    "open directory structure",
			args:    []string{"open", "test-project"},
			wantErr: false,
			setup: func(tmpDir string) string {
				// Create a basic BSpec project structure
				projectDir := filepath.Join(tmpDir, "test-project")
				os.MkdirAll(projectDir, 0755)
				os.MkdirAll(filepath.Join(projectDir, "documents"), 0755)

				// Create manifest.json
				manifest := map[string]interface{}{
					"name":             "Test Project",
					"version":          "1.0.0",
					"conformance_level": "standard",
					"industry_profile": "generic",
					"author":           "Test Author",
					"description":      "A test project for unit testing",
				}
				manifestData, _ := json.MarshalIndent(manifest, "", "  ")
				os.WriteFile(filepath.Join(projectDir, "manifest.json"), manifestData, 0644)

				// Create a test document
				doc := `---
id: TST-test-doc
title: Test Document
type: TST
status: Draft
owner: Test Owner
version: 1.0.0
---

# Test Document

This is a test document for verification.
`
				os.WriteFile(filepath.Join(projectDir, "documents", "TST-test-doc.md"), []byte(doc), 0644)

				return projectDir
			},
			verify: func(output string, t *testing.T) {
				// Verify that output contains project information
				expectedStrings := []string{
					"Test Project",
					"1.0.0",
					"Test Author",
					"TST-test-doc",
				}
				for _, expected := range expectedStrings {
					if !strings.Contains(output, expected) {
						t.Errorf("Expected output to contain '%s', got: %s", expected, output)
					}
				}
			},
		},
		{
			name:    "open with stats flag",
			args:    []string{"open", "test-project-stats", "--stats"},
			wantErr: false,
			setup: func(tmpDir string) string {
				// Create a minimal project for stats testing
				projectDir := filepath.Join(tmpDir, "test-project-stats")
				os.MkdirAll(projectDir, 0755)
				os.MkdirAll(filepath.Join(projectDir, "documents"), 0755)

				// Create manifest.json
				manifest := map[string]interface{}{
					"name":    "Stats Test Project",
					"version": "1.0.0",
				}
				manifestData, _ := json.MarshalIndent(manifest, "", "  ")
				os.WriteFile(filepath.Join(projectDir, "manifest.json"), manifestData, 0644)

				// Create multiple test documents
				for i := 0; i < 3; i++ {
					doc := fmt.Sprintf(`---
id: DOC-%d
title: Document %d
type: DOC
status: Draft
version: 1.0.0
---

# Document %d
`, i, i, i)
					filename := fmt.Sprintf("DOC-%d.md", i)
					os.WriteFile(filepath.Join(projectDir, "documents", filename), []byte(doc), 0644)
				}

				return projectDir
			},
			verify: func(output string, t *testing.T) {
				// Verify stats output contains statistics
				if !strings.Contains(output, "Statistics") && !strings.Contains(output, "Documents:") {
					t.Errorf("Expected output to contain statistics information")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary directory
			tmpDir := t.TempDir()
			originalDir, _ := os.Getwd()
			defer os.Chdir(originalDir)
			os.Chdir(tmpDir)

			// Run setup if provided
			var targetPath string
			if tt.setup != nil {
				targetPath = tt.setup(tmpDir)
			}

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

			// Silence unused variable warning
			_ = targetPath

			// Run verification if provided
			if tt.verify != nil {
				output := buf.String()
				tt.verify(output, t)
			}
		})
	}
}

func TestOpenFlags(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)
	os.Chdir(tmpDir)

	// Create a minimal test project
	projectDir := filepath.Join(tmpDir, "test-project")
	os.MkdirAll(projectDir, 0755)

	manifest := map[string]interface{}{
		"name":    "Flag Test Project",
		"version": "1.0.0",
	}
	manifestData, _ := json.MarshalIndent(manifest, "", "  ")
	os.WriteFile(filepath.Join(projectDir, "manifest.json"), manifestData, 0644)

	tests := []struct {
		name     string
		args     []string
		contains string
	}{
		{
			name:     "output json flag",
			args:     []string{"open", "test-project", "--output", "json"},
			contains: "Flag Test Project",
		},
		{
			name:     "output yaml flag",
			args:     []string{"open", "test-project", "-o", "yaml"},
			contains: "Flag Test Project",
		},
		{
			name:     "stats flag",
			args:     []string{"open", "test-project", "--stats"},
			contains: "Flag Test Project",
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
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			output := buf.String()
			if tt.contains != "" && !strings.Contains(output, tt.contains) {
				t.Errorf("Expected output to contain '%s', got: %s", tt.contains, output)
			}
		})
	}
}