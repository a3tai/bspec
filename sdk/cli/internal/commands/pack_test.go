package commands

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		setup   func(string) string // Setup function returns project directory
		verify  func(string, *testing.T) // Verify function with bspec file path
	}{
		{
			name:    "pack help",
			args:    []string{"pack", "--help"},
			wantErr: false,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(bspecPath string, t *testing.T) {
				// No verification needed for help
			},
		},
		{
			name:    "pack without directory",
			args:    []string{"pack"},
			wantErr: true,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(bspecPath string, t *testing.T) {
				// No verification needed for error case
			},
		},
		{
			name:    "pack basic project",
			args:    []string{"pack", "test-project", "output.bspec"},
			wantErr: false,
			setup: func(tmpDir string) string {
				// Create a basic BSpec project structure
				projectDir := filepath.Join(tmpDir, "test-project")
				os.MkdirAll(projectDir, 0755)
				os.MkdirAll(filepath.Join(projectDir, "documents"), 0755)
				os.MkdirAll(filepath.Join(projectDir, "assets"), 0755)

				// Create manifest.json
				manifest := map[string]interface{}{
					"name":             "Test Project",
					"version":          "1.0.0",
					"conformance_level": "standard",
					"industry_profile": "generic",
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

This is a test document.
`
				os.WriteFile(filepath.Join(projectDir, "documents", "TST-test-doc.md"), []byte(doc), 0644)

				return projectDir
			},
			verify: func(bspecPath string, t *testing.T) {
				// Verify the .bspec file was created
				if _, err := os.Stat(bspecPath); os.IsNotExist(err) {
					t.Errorf("Expected .bspec file to be created at %s", bspecPath)
				}

				// Check file size is reasonable (should contain data)
				if info, err := os.Stat(bspecPath); err == nil {
					if info.Size() < 100 {
						t.Errorf("Expected .bspec file to be larger than 100 bytes, got %d", info.Size())
					}
				}
			},
		},
		{
			name:    "pack with verbose flag",
			args:    []string{"pack", "test-project-verbose", "output-verbose.bspec", "--verbose"},
			wantErr: false,
			setup: func(tmpDir string) string {
				// Create a minimal project
				projectDir := filepath.Join(tmpDir, "test-project-verbose")
				os.MkdirAll(projectDir, 0755)

				manifest := map[string]interface{}{
					"name":    "Test Project Verbose",
					"version": "1.0.0",
				}
				manifestData, _ := json.MarshalIndent(manifest, "", "  ")
				os.WriteFile(filepath.Join(projectDir, "manifest.json"), manifestData, 0644)

				return projectDir
			},
			verify: func(bspecPath string, t *testing.T) {
				// Verify the .bspec file was created
				if _, err := os.Stat(bspecPath); os.IsNotExist(err) {
					t.Errorf("Expected .bspec file to be created at %s", bspecPath)
				}
			},
		},
		{
			name:    "pack non-existent directory",
			args:    []string{"pack", "non-existent-dir", "output.bspec"},
			wantErr: true,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(bspecPath string, t *testing.T) {
				// No verification needed for error case
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
			var projectDir string
			if tt.setup != nil {
				projectDir = tt.setup(tmpDir)
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

			// Run verification if provided and no error expected
			if !tt.wantErr && tt.verify != nil && len(tt.args) > 2 && tt.args[2] != "--help" {
				bspecPath := filepath.Join(tmpDir, tt.args[2])
				tt.verify(bspecPath, t)
			}

			// Silence unused variable warning
			_ = projectDir

			// For verbose tests, check that output contains verbose information
			if !tt.wantErr && strings.Contains(strings.Join(tt.args, " "), "--verbose") {
				output := buf.String()
				if !strings.Contains(output, "Packing") && !strings.Contains(output, "Created") {
					t.Errorf("Expected verbose output to contain progress information")
				}
			}
		})
	}
}