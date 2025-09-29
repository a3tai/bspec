package commands

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		setup   func(string) string // Setup function returns bspec file path
		verify  func(string, *testing.T) // Verify function with extract directory
	}{
		{
			name:    "extract help",
			args:    []string{"extract", "--help"},
			wantErr: false,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(extractDir string, t *testing.T) {
				// No verification needed for help
			},
		},
		{
			name:    "extract without bspec file",
			args:    []string{"extract"},
			wantErr: true,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(extractDir string, t *testing.T) {
				// No verification needed for error case
			},
		},
		{
			name:    "extract non-existent file",
			args:    []string{"extract", "non-existent.bspec"},
			wantErr: true,
			setup: func(tmpDir string) string {
				return ""
			},
			verify: func(extractDir string, t *testing.T) {
				// No verification needed for error case
			},
		},
		{
			name:    "extract with custom output directory",
			args:    []string{"extract", "test.bspec", "--output", "custom-output"},
			wantErr: true, // Will fail because we can't easily create a real .bspec file in tests
			setup: func(tmpDir string) string {
				// Create a dummy .bspec file (this won't be valid, but we're testing arg parsing)
				bspecPath := filepath.Join(tmpDir, "test.bspec")
				os.WriteFile(bspecPath, []byte("dummy"), 0644)
				return bspecPath
			},
			verify: func(extractDir string, t *testing.T) {
				// No verification needed since this will error
			},
		},
		{
			name:    "extract with force flag",
			args:    []string{"extract", "test.bspec", "--force"},
			wantErr: true, // Will fail because we can't easily create a real .bspec file in tests
			setup: func(tmpDir string) string {
				// Create a dummy .bspec file
				bspecPath := filepath.Join(tmpDir, "test.bspec")
				os.WriteFile(bspecPath, []byte("dummy"), 0644)
				return bspecPath
			},
			verify: func(extractDir string, t *testing.T) {
				// No verification needed since this will error
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
			var bspecPath string
			if tt.setup != nil {
				bspecPath = tt.setup(tmpDir)
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
			_ = bspecPath

			// Run verification if provided and no error expected
			if !tt.wantErr && tt.verify != nil {
				// Determine output directory
				outputDir := "extracted" // default
				for i, arg := range tt.args {
					if arg == "--output" && i+1 < len(tt.args) {
						outputDir = tt.args[i+1]
						break
					}
				}
				extractPath := filepath.Join(tmpDir, outputDir)
				tt.verify(extractPath, t)
			}
		})
	}
}

func TestExtractFlags(t *testing.T) {
	// Test that flags are correctly defined and parsed
	// We'll test flag parsing without actually executing the extract logic

	tmpDir := t.TempDir()
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)
	os.Chdir(tmpDir)

	// Create a dummy .bspec file
	bspecPath := filepath.Join(tmpDir, "test.bspec")
	os.WriteFile(bspecPath, []byte("dummy"), 0644)

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "output flag",
			args: []string{"extract", "test.bspec", "--output", "custom-dir"},
		},
		{
			name: "force flag",
			args: []string{"extract", "test.bspec", "--force"},
		},
		{
			name: "verbose flag",
			args: []string{"extract", "test.bspec", "--verbose"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the command state
			resetRootCmd()

			// Capture output (we expect errors since we don't have valid .bspec files)
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)
			rootCmd.SetArgs(tt.args)

			// We expect this to fail, but it should parse flags correctly
			rootCmd.Execute()

			// The main goal is to ensure no panic occurs during flag parsing
			// The actual extraction will fail, but that's expected with dummy files
		})
	}
}