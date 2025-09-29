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
				"BSpec CLI version 1.0.0",
				"Build date: unknown",
				"Git commit: unknown",
			},
		},
		{
			name: "short version output",
			args: []string{"version", "--short"},
			contains: []string{
				"1.0.0",
			},
		},
		{
			name: "short version output with flag shorthand",
			args: []string{"version", "-s"},
			contains: []string{
				"1.0.0",
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
	// Test that version variables are set correctly
	if Version == "" {
		t.Error("Version should not be empty")
	}
	if BuildDate == "" {
		t.Error("BuildDate should not be empty")
	}
	if GitCommit == "" {
		t.Error("GitCommit should not be empty")
	}

	// Test default values
	expectedVersion := "1.0.0"
	if Version != expectedVersion {
		t.Errorf("Expected Version to be '%s', got '%s'", expectedVersion, Version)
	}

	expectedBuildDate := "unknown"
	if BuildDate != expectedBuildDate {
		t.Errorf("Expected BuildDate to be '%s', got '%s'", expectedBuildDate, BuildDate)
	}

	expectedGitCommit := "unknown"
	if GitCommit != expectedGitCommit {
		t.Errorf("Expected GitCommit to be '%s', got '%s'", expectedGitCommit, GitCommit)
	}
}