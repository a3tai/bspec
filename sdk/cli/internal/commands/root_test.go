package commands

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestRootCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		contains string
	}{
		{
			name:     "help flag",
			args:     []string{"--help"},
			wantErr:  false,
			contains: "BSpec CLI is a comprehensive toolkit",
		},
		{
			name:     "version flag",
			args:     []string{"--version"},
			wantErr:  false,
			contains: "bspec version 1.0.0",
		},
		{
			name:     "invalid flag",
			args:     []string{"--invalid"},
			wantErr:  true,
			contains: "unknown flag",
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

func TestGlobalFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		flagName string
		expected interface{}
	}{
		{
			name:     "output flag json",
			args:     []string{"--output", "json", "--help"},
			flagName: "output",
			expected: "json",
		},
		{
			name:     "output flag yaml",
			args:     []string{"-o", "yaml", "--help"},
			flagName: "output",
			expected: "yaml",
		},
		{
			name:     "verbose flag",
			args:     []string{"--verbose", "--help"},
			flagName: "verbose",
			expected: true,
		},
		{
			name:     "quiet flag",
			args:     []string{"-q", "--help"},
			flagName: "quiet",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the command state
			resetRootCmd()
			viper.Reset()

			// Set up command
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check viper value
			actual := viper.Get(tt.flagName)
			if actual != tt.expected {
				t.Errorf("Expected %s to be %v, got %v", tt.flagName, tt.expected, actual)
			}
		})
	}
}

func TestInitConfig(t *testing.T) {
	// Save original HOME
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)

	// Reset viper
	viper.Reset()

	// Test initConfig without config file
	initConfig()

	// Verify config setup
	configPaths := viper.ConfigFileUsed()
	if configPaths != "" {
		t.Errorf("Expected no config file to be used, but got: %s", configPaths)
	}
}

func TestInitConfigWithFile(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configFile := tmpDir + "/.bspec.yaml"

	configContent := `
output: json
verbose: true
`
	err := os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Save original HOME
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tmpDir)

	// Reset viper
	viper.Reset()

	// Test initConfig with config file
	initConfig()

	// Verify config values
	if viper.GetString("output") != "json" {
		t.Errorf("Expected output to be 'json', got '%s'", viper.GetString("output"))
	}
	if !viper.GetBool("verbose") {
		t.Errorf("Expected verbose to be true")
	}
}

// resetRootCmd resets the root command for testing
func resetRootCmd() {
	rootCmd.ResetFlags()
	rootCmd.ResetCommands()

	// Re-add persistent flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bspec.yaml)")
	rootCmd.PersistentFlags().StringP("output", "o", "yaml", "output format (json|yaml|markdown)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet output")

	// Re-bind flags to viper
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))

	// Re-add subcommands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(extractCmd)
	rootCmd.AddCommand(packCmd)
}