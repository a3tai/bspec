package cmd

import "github.com/a3tai/bspec/cli/internal/commands"

// Execute runs the root command and handles errors
func Execute() error {
	return commands.Execute()
}