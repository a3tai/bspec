package commands

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/a3tai/bspec/cli/internal/chat"
)

//go:embed prompts
var promptsFS embed.FS

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Interactive AI chat for BSpec document generation and analysis",
	Long: `Start an interactive AI-powered chat session for generating and analyzing BSpec documents.

The chat command provides:
  • Real-time streaming responses from AI models
  • BSpec-aware document generation (CAP, CTX, PER, etc.)
  • Document analysis and validation
  • Interactive command system with help and history

Environment Variables:
  OPENROUTER_API_KEY      OpenRouter API key (required for default provider)
  BSPEC_LLM               LLM provider: "openrouter" or "openai" (default: "openrouter")
  BSPEC_MODEL             Model to use (default: "openai/gpt-5" for OpenRouter)

  For OpenAI (backward compatibility):
  BSPEC_OPENAI_API_KEY    OpenAI API key

Available Chat Commands:
  /help                   Show help and available commands
  /generate <type>        Generate a BSpec document (CAP, CTX, PER, etc.)
  /analyze <file>         Analyze an existing BSpec file
  /models                 List available models from current provider
  /clear                  Clear conversation history
  /history                Show conversation history
  /quit, /exit            Exit the chat

Examples:
  OPENROUTER_API_KEY=sk-... bspec chat                      # Start with OpenRouter (default)
  OPENROUTER_API_KEY=sk-... BSPEC_MODEL=openai/gpt-4o bspec chat  # Use GPT-4o via OpenRouter
  BSPEC_LLM=openai BSPEC_OPENAI_API_KEY=sk-... bspec chat   # Use OpenAI directly

Then in chat:
  > /models                           # List available models
  > /generate CAP                     # Generate a Capability document
  > /analyze myfile.md                # Analyze a BSpec file
  > How do I create a Context document?  # Ask questions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load configuration from environment variables (this will create directories)
		config, err := chat.NewConfig()
		if err != nil {
			return fmt.Errorf("configuration error: %w", err)
		}

		// Always ensure prompts directory exists and install default prompts
		// This happens even if API key validation fails later
		if err := setupPrompts(config); err != nil {
			// Don't fail on prompt setup errors, just warn
			fmt.Printf("Warning: failed to setup prompts: %v\n", err)
		}

		// Create chat client
		client, err := chat.NewClient(config)
		if err != nil {
			return fmt.Errorf("failed to create chat client: %w", err)
		}

		// Load default system prompt
		systemPrompt, err := config.LoadPrompt("default")
		if err != nil {
			// Fallback to built-in prompt if file loading fails
			systemPrompt = chat.BSpecSystemPrompt()
		}

		// Try to add project context if we're in a BSpec project
		if projectContext, err := config.LoadProjectContext(); err == nil {
			systemPrompt += "\n\n" + projectContext
		}

		client.SetSystemPrompt(systemPrompt)

		// Create UI
		ui := chat.NewUI(client)

		// Check if we're in a terminal
		if !isTerminal() {
			return fmt.Errorf("chat command requires an interactive terminal")
		}

		// Start the interactive chat
		return ui.Run()
	},
}

// isTerminal checks if we're running in an interactive terminal
func isTerminal() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// setupPrompts ensures the prompts directory exists and copies default prompts
func setupPrompts(config *chat.Config) error {
	// Ensure prompts directory exists
	if err := config.EnsurePromptsDirectory(); err != nil {
		return err
	}

	prompts := []string{"default.md", "context.md"}

	for _, promptFile := range prompts {
		promptPath := filepath.Join(config.PromptsPath, promptFile)

		// Skip if file already exists
		if _, err := os.Stat(promptPath); err == nil {
			continue
		}

		// Copy prompt from embedded filesystem
		data, err := promptsFS.ReadFile("prompts/" + promptFile)
		if err != nil {
			return fmt.Errorf("failed to read embedded prompt %s: %w", promptFile, err)
		}

		if err := os.WriteFile(promptPath, data, 0644); err != nil {
			return fmt.Errorf("failed to write prompt %s: %w", promptFile, err)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Add flags for chat configuration
	chatCmd.Flags().StringP("model", "m", "", "Override the model to use (overrides BSPEC_MODEL)")
	chatCmd.Flags().StringP("provider", "p", "", "Override the LLM provider (overrides BSPEC_LLM)")
	chatCmd.Flags().BoolP("no-history", "", false, "Start with empty conversation history")

	// Mark flags as hidden for now since we're using environment variables
	chatCmd.Flags().MarkHidden("model")
	chatCmd.Flags().MarkHidden("provider")
	chatCmd.Flags().MarkHidden("no-history")
}