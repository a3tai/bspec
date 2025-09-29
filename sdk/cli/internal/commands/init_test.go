package commands

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestInitCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		setup   func(string) // Setup function with tmpDir
		verify  func(string, *testing.T) // Verify function with project path
	}{
		{
			name:    "init help",
			args:    []string{"init", "--help"},
			wantErr: false,
			verify: func(projectPath string, t *testing.T) {
				// No verification needed for help
			},
		},
		{
			name:    "init without project name",
			args:    []string{"init"},
			wantErr: true,
			verify: func(projectPath string, t *testing.T) {
				// No verification needed for error case
			},
		},
		{
			name:    "init basic project",
			args:    []string{"init", "test-project"},
			wantErr: false,
			verify: func(projectPath string, t *testing.T) {
				// Verify directory structure
				expectedDirs := []string{
					"documents",
					"assets",
					"computed",
				}
				for _, dir := range expectedDirs {
					dirPath := filepath.Join(projectPath, dir)
					if _, err := os.Stat(dirPath); os.IsNotExist(err) {
						t.Errorf("Expected directory %s to exist", dirPath)
					}
				}

				// Verify manifest.json exists and is valid
				manifestPath := filepath.Join(projectPath, "manifest.json")
				if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
					t.Errorf("Expected manifest.json to exist")
					return
				}

				// Parse manifest.json
				manifestData, err := os.ReadFile(manifestPath)
				if err != nil {
					t.Errorf("Failed to read manifest.json: %v", err)
					return
				}

				var manifest map[string]interface{}
				if err := json.Unmarshal(manifestData, &manifest); err != nil {
					t.Errorf("Failed to parse manifest.json: %v", err)
					return
				}

				// Verify required fields
				requiredFields := []string{"name", "bspec_version", "conformance_level", "industry_profile"}
				for _, field := range requiredFields {
					if _, exists := manifest[field]; !exists {
						t.Errorf("Expected manifest.json to contain field: %s", field)
					}
				}

				// Verify README.md exists
				readmePath := filepath.Join(projectPath, "README.md")
				if _, err := os.Stat(readmePath); os.IsNotExist(err) {
					t.Errorf("Expected README.md to exist")
				}
			},
		},
		{
			name:    "init with custom author",
			args:    []string{"init", "test-project-author", "--author", "Test Author"},
			wantErr: false,
			verify: func(projectPath string, t *testing.T) {
				// Verify manifest contains author
				manifestPath := filepath.Join(projectPath, "manifest.json")
				manifestData, err := os.ReadFile(manifestPath)
				if err != nil {
					t.Errorf("Failed to read manifest.json: %v", err)
					return
				}

				var manifest map[string]interface{}
				if err := json.Unmarshal(manifestData, &manifest); err != nil {
					t.Errorf("Failed to parse manifest.json: %v", err)
					return
				}

				if author, exists := manifest["author"]; !exists || author != "Test Author" {
					t.Errorf("Expected author to be 'Test Author', got: %v", author)
				}
			},
		},
		{
			name:    "init with force flag",
			args:    []string{"init", "test-project-force", "--force"},
			wantErr: false,
			setup: func(tmpDir string) {
				// Create existing directory
				projectPath := filepath.Join(tmpDir, "test-project-force")
				os.MkdirAll(projectPath, 0755)
				// Create a dummy file
				os.WriteFile(filepath.Join(projectPath, "dummy.txt"), []byte("test"), 0644)
			},
			verify: func(projectPath string, t *testing.T) {
				// Verify old content was removed and new structure created
				dummyPath := filepath.Join(projectPath, "dummy.txt")
				if _, err := os.Stat(dummyPath); !os.IsNotExist(err) {
					t.Errorf("Expected dummy.txt to be removed")
				}

				// Verify new structure exists
				manifestPath := filepath.Join(projectPath, "manifest.json")
				if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
					t.Errorf("Expected manifest.json to exist after force init")
				}
			},
		},
		{
			name:    "init existing directory without force",
			args:    []string{"init", "test-project-existing"},
			wantErr: true,
			setup: func(tmpDir string) {
				// Create existing directory
				projectPath := filepath.Join(tmpDir, "test-project-existing")
				os.MkdirAll(projectPath, 0755)
			},
			verify: func(projectPath string, t *testing.T) {
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
			if tt.setup != nil {
				tt.setup(tmpDir)
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
			if !tt.wantErr && tt.verify != nil {
				projectName := "test-project"
				if len(tt.args) > 1 && tt.args[1] != "--help" {
					projectName = tt.args[1]
				}
				projectPath := filepath.Join(tmpDir, projectName)
				tt.verify(projectPath, t)
			}
		})
	}
}

func TestInitFlags(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)
	os.Chdir(tmpDir)

	tests := []struct {
		name         string
		args         []string
		expectedFlag string
		expectedVal  string
	}{
		{
			name:         "conformance flag",
			args:         []string{"init", "test-conformance", "--conformance", "bronze"},
			expectedFlag: "conformance_level",
			expectedVal:  "bronze",
		},
		{
			name:         "industry flag",
			args:         []string{"init", "test-industry", "--industry", "healthcare"},
			expectedFlag: "industry_profile",
			expectedVal:  "healthcare",
		},
		{
			name:         "description flag",
			args:         []string{"init", "test-desc", "--description", "Test project description"},
			expectedFlag: "description",
			expectedVal:  "Test project description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the command state
			resetRootCmd()

			// Execute the command
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Verify the flag value in manifest.json
			projectName := tt.args[1]
			manifestPath := filepath.Join(tmpDir, projectName, "manifest.json")
			manifestData, err := os.ReadFile(manifestPath)
			if err != nil {
				t.Errorf("Failed to read manifest.json: %v", err)
				return
			}

			var manifest map[string]interface{}
			if err := json.Unmarshal(manifestData, &manifest); err != nil {
				t.Errorf("Failed to parse manifest.json: %v", err)
				return
			}

			if val, exists := manifest[tt.expectedFlag]; !exists || val != tt.expectedVal {
				t.Errorf("Expected %s to be '%s', got: %v", tt.expectedFlag, tt.expectedVal, val)
			}
		})
	}
}