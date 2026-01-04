package commands

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		contains []string
	}{
		{
			name: "default version output",
			args: []string{"version"},
			contains: []string{
				"BSpec CLI version",
				"Build date:",
				"Git commit:",
			},
		},
		{
			name: "short version output",
			args: []string{"version", "--short"},
			contains: []string{
				Version, // Use actual version
			},
		},
		{
			name: "short version output with flag shorthand",
			args: []string{"version", "-s"},
			contains: []string{
				Version, // Use actual version
			},
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
			}

			output := buf.String()
			for _, expected := range tt.contains {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', got: %s", expected, output)
				}
			}
		})
	}
}

func TestVersionVariables(t *testing.T) {
	// Test that version variables are set correctly (not empty)
	if Version == "" {
		t.Error("Version should not be empty")
	}
	if BuildDate == "" {
		t.Error("BuildDate should not be empty")
	}
	if GitCommit == "" {
		t.Error("GitCommit should not be empty")
	}

	// Version should be a valid semver format (X.Y.Z)
	if len(Version) < 5 { // Minimum "0.0.0"
		t.Errorf("Version '%s' does not appear to be a valid semver format", Version)
	}
}