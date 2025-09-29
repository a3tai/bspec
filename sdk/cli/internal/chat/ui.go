package chat

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alperdrsnn/clime"
)

// UI handles the terminal user interface for chat
type UI struct {
	client    *Client
	animation *ThinkingAnimation
	formatter *MarkdownFormatter
	debugMode bool
}

// NewUI creates a new chat UI
func NewUI(client *Client) *UI {
	ui := &UI{
		client:    client,
		animation: NewThinkingAnimation(),
		formatter: NewMarkdownFormatter(true), // Enable colors for terminal output
		debugMode: false,
	}

	// Set the UI as the animation callback for the tools registry
	if client.tools != nil {
		client.tools.SetAnimationCallback(ui)
	}

	// Set the UI as the debug callback for the client
	client.SetDebugCallback(ui)

	return ui
}

// ShowWelcome displays the welcome banner
func (ui *UI) ShowWelcome() {
	clime.Clear()

	// Create a colorful banner
	banner := `
╔══════════════════════════════════════════════════╗
║              🔧 BSpec CLI Chat                    ║
║         AI-Powered Document Generation           ║
╚══════════════════════════════════════════════════╝
`
	fmt.Print(clime.Info.Sprint(banner))
	fmt.Print(clime.BlueColor.Sprint("Provider: " + ui.client.config.LLMProvider + " | Model: " + ui.client.config.Model + "\n\n"))

	// Show help information
	ui.ShowHelp()
	fmt.Println()
}

// ShowHelp displays available commands
func (ui *UI) ShowHelp() {
	help := `Available Commands:
  /help                     - Show this help message
  /generate <type>          - Generate a BSpec document (CAP, CTX, PER, etc.)
  /analyze <file>           - Analyze an existing BSpec file
  /models                   - List available models from current provider
  /model <model-name>       - Switch to a specific model
  /clear                    - Clear conversation history
  /history                  - Show conversation history
  /mcp                      - Show MCP server status
  /debug                    - Toggle debug logging
  /quit, /exit              - Exit the chat
  Ctrl+D                    - Exit the chat

BSpec Document Types:
  CAP - Capability         CTX - Context          PER - Persona
  PRO - Process           OFF - Offering         ARC - Architecture
  COM - Compliance        MET - Metrics          RSK - Risk
  DEC - Decision          EXP - Experiment       PRF - Profile

Just type your message and press Enter to start chatting!`

	fmt.Print(clime.Warning.Sprint(help))
}

// ShowError displays an error message
func (ui *UI) ShowError(err error) {
	fmt.Print(clime.Error.Sprint("Error: " + err.Error() + "\n"))
}

// ShowSuccess displays a success message
func (ui *UI) ShowSuccess(message string) {
	fmt.Print(clime.Success.Sprint("✓ " + message + "\n"))
}

// ShowInfo displays an info message
func (ui *UI) ShowInfo(message string) {
	fmt.Print(clime.Info.Sprint("ℹ " + message + "\n"))
}

// ShowDebug displays a debug message when debug mode is enabled
func (ui *UI) ShowDebug(message string) {
	if ui.debugMode {
		fmt.Print(clime.MagentaColor.Sprint("🐛 DEBUG: " + message + "\n"))
	}
}

// ShowUserMessage displays a user message
func (ui *UI) ShowUserMessage(message string) {
	fmt.Print(clime.BlueColor.Sprint("You: ") + message + "\n")
}

// ShowAssistantPrefix shows the assistant prefix before streaming (deprecated - kept for compatibility)
func (ui *UI) ShowAssistantPrefix() {
	// No longer showing "Assistant:" prefix per user request
}

// ShowThinking displays a thinking indicator
func (ui *UI) ShowThinking() {
	fmt.Print(clime.Warning.Sprint("Thinking... "))
}

// StopThinking stops the thinking indicator
func (ui *UI) StopThinking() {
	fmt.Print("\r" + strings.Repeat(" ", 50) + "\r") // Clear the thinking line
}

// ShowStreamingResponse handles streaming response display
func (ui *UI) ShowStreamingResponse(responseChan <-chan StreamResponse) error {
	var fullContent strings.Builder

	for response := range responseChan {
		if response.Error != nil {
			fmt.Print(clime.Error.Sprint("Error: " + response.Error.Error() + "\n\n"))
			return response.Error
		}

		if response.Done {
			// Format and display the complete response using markdown formatter
			formattedContent := ui.formatter.Format(fullContent.String())
			fmt.Print(formattedContent + "\n\n")
			break
		}

		// Accumulate content for formatting
		fullContent.WriteString(response.Content)
	}

	return nil
}

// PromptUser prompts for user input with a colorful prefix
func (ui *UI) PromptUser() (string, error) {
	fmt.Print(clime.BlueColor.Sprint("\n> "))

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			// Handle Ctrl+D gracefully
			ui.ShowInfo("\nGoodbye! 👋")
			os.Exit(0)
		}
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// ProcessCommand processes special commands (returns true if command was handled)
func (ui *UI) ProcessCommand(input string) (bool, error) {
	input = strings.TrimSpace(input)

	if !strings.HasPrefix(input, "/") {
		return false, nil // Not a command
	}

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return true, nil
	}

	command := strings.ToLower(parts[0])

	switch command {
	case "/help":
		ui.ShowHelp()
		return true, nil

	case "/quit", "/exit":
		ui.ShowInfo("Goodbye! 👋")
		os.Exit(0)
		return true, nil

	case "/clear":
		ui.client.ClearHistory()
		clime.Clear()
		ui.ShowWelcome()
		ui.ShowSuccess("Conversation history cleared")
		return true, nil

	case "/history":
		ui.ShowConversationHistory()
		return true, nil

	case "/generate":
		if len(parts) < 2 {
			ui.ShowError(fmt.Errorf("usage: /generate <document-type>"))
			return true, nil
		}
		return ui.HandleGenerateCommand(parts[1])

	case "/analyze":
		if len(parts) < 2 {
			ui.ShowError(fmt.Errorf("usage: /analyze <file-path>"))
			return true, nil
		}
		return ui.HandleAnalyzeCommand(parts[1])

	case "/models":
		return ui.HandleModelsCommand()

	case "/model":
		if len(parts) < 2 {
			ui.ShowError(fmt.Errorf("usage: /model <model-name>"))
			return true, nil
		}
		return ui.HandleModelSwitchCommand(parts[1])

	case "/mcp":
		return ui.HandleMCPCommand()

	case "/debug":
		return ui.HandleDebugCommand()

	default:
		ui.ShowError(fmt.Errorf("unknown command: %s (type /help for available commands)", command))
		return true, nil
	}
}

// ShowConversationHistory displays the conversation history
func (ui *UI) ShowConversationHistory() {
	messages := ui.client.GetMessages()

	if len(messages) == 0 {
		ui.ShowInfo("No conversation history")
		return
	}

	fmt.Print(clime.Warning.Sprint("\n=== Conversation History ===\n"))

	for i, msg := range messages {
		timestamp := msg.Time.Format("15:04:05")

		switch msg.Role {
		case "system":
			fmt.Printf(clime.MagentaColor.Sprint("[%s] System: %s\n"), timestamp, msg.Content)
		case "user":
			fmt.Printf(clime.BlueColor.Sprint("[%s] You: %s\n"), timestamp, msg.Content)
		case "assistant":
			fmt.Printf(clime.GreenColor.Sprint("[%s] Assistant: %s\n"), timestamp, msg.Content)
		}

		if i < len(messages)-1 {
			fmt.Println()
		}
	}

	fmt.Print(clime.Warning.Sprint("=== End History ===\n"))
}

// HandleGenerateCommand handles the /generate command
func (ui *UI) HandleGenerateCommand(docType string) (bool, error) {
	docType = strings.ToUpper(docType)

	// Validate document type
	validTypes := []string{"CAP", "CTX", "PER", "PRO", "OFF", "ARC", "COM", "MET", "RSK", "DEC", "EXP", "PRF"}
	valid := false
	for _, vt := range validTypes {
		if docType == vt {
			valid = true
			break
		}
	}

	if !valid {
		ui.ShowError(fmt.Errorf("invalid document type: %s\nValid types: %s", docType, strings.Join(validTypes, ", ")))
		return true, nil
	}

	ui.ShowInfo(fmt.Sprintf("Generating BSpec %s document...", docType))

	// Create a generation prompt
	prompt := fmt.Sprintf("Please generate a BSpec %s document. Include proper YAML frontmatter and markdown content following the BSpec specification.", docType)

	// Send the message using normal chat flow
	return false, ui.sendMessage(context.Background(), prompt)
}

// HandleAnalyzeCommand handles the /analyze command
func (ui *UI) HandleAnalyzeCommand(filePath string) (bool, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ui.ShowError(fmt.Errorf("file not found: %s", filePath))
		return true, nil
	}

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		ui.ShowError(fmt.Errorf("failed to read file: %w", err))
		return true, nil
	}

	ui.ShowInfo(fmt.Sprintf("Analyzing file: %s", filePath))

	// Create analysis prompt
	prompt := fmt.Sprintf("Please analyze this BSpec document and provide insights:\n\n```\n%s\n```", string(content))

	return false, ui.sendMessage(context.Background(), prompt)
}

// HandleModelsCommand handles the /models command
func (ui *UI) HandleModelsCommand() (bool, error) {
	ui.ShowInfo("Fetching available models...")

	models, err := ui.client.config.FetchAvailableModels()
	if err != nil {
		ui.ShowError(fmt.Errorf("failed to fetch models: %w", err))
		return true, nil
	}

	if len(models) == 0 {
		ui.ShowInfo("No models available")
		return true, nil
	}

	fmt.Print(clime.Warning.Sprint("\n=== Available Models ===\n"))

	// Show first 20 models to avoid overwhelming output
	maxModels := 20
	if len(models) > maxModels {
		fmt.Printf(clime.BlueColor.Sprint("Showing first %d of %d models (sorted by popularity):\n\n"), maxModels, len(models))
	}

	for i, model := range models {
		if i >= maxModels {
			break
		}

		// Format pricing info
		pricingInfo := ""
		if model.Pricing.Prompt != "" && model.Pricing.Completion != "" {
			pricingInfo = fmt.Sprintf(" (Prompt: $%s/1M, Completion: $%s/1M)", model.Pricing.Prompt, model.Pricing.Completion)
		}

		// Format context length
		contextInfo := ""
		if model.ContextLength > 0 {
			contextInfo = fmt.Sprintf(" [%dk context]", model.ContextLength/1000)
		}

		fmt.Printf(clime.GreenColor.Sprint("%d. %s\n"), i+1, model.ID)
		fmt.Printf("   %s%s%s\n", model.Name, contextInfo, pricingInfo)
		if model.Description != "" && len(model.Description) < 100 {
			fmt.Printf("   %s\n", clime.MagentaColor.Sprint(model.Description))
		}
		fmt.Println()
	}

	if len(models) > maxModels {
		fmt.Printf(clime.BlueColor.Sprint("... and %d more models\n"), len(models)-maxModels)
	}

	fmt.Print(clime.Warning.Sprint("=== End Models ===\n"))
	fmt.Printf(clime.Info.Sprint("\nTo use a specific model, set BSPEC_MODEL environment variable.\n"))
	fmt.Printf(clime.Info.Sprint("Current model: %s\n"), ui.client.config.Model)

	return true, nil
}

// HandleModelSwitchCommand handles the /model command
func (ui *UI) HandleModelSwitchCommand(modelName string) (bool, error) {
	// Validate model name exists in popular models
	popularModels := GetPopularModels()
	providerModels := popularModels[ui.client.config.LLMProvider]

	modelExists := false
	for _, model := range providerModels {
		if model == modelName {
			modelExists = true
			break
		}
	}

	if !modelExists {
		ui.ShowError(fmt.Errorf("unknown model: %s. Use /models to see available models", modelName))
		return true, nil
	}

	// Switch to the new model
	if err := ui.client.config.SetModel(modelName); err != nil {
		ui.ShowError(fmt.Errorf("failed to switch model: %w", err))
		return true, nil
	}

	ui.ShowSuccess(fmt.Sprintf("Switched to model: %s", modelName))
	ui.ShowInfo(GetModelDescription(modelName))
	return true, nil
}

// HandleMCPCommand handles the /mcp command
func (ui *UI) HandleMCPCommand() (bool, error) {
	mcpServers := ui.client.config.MCPServers

	if len(mcpServers) == 0 {
		ui.ShowInfo("No MCP servers configured")
		return true, nil
	}

	fmt.Print(clime.Warning.Sprint("\n=== MCP Servers ===\n"))

	for i, server := range mcpServers {
		status := clime.Error.Sprint("disabled")
		if server.Enabled {
			status = clime.Success.Sprint("enabled")
		}

		fmt.Printf(clime.GreenColor.Sprint("%d. %s\n"), i+1, server.Name)
		fmt.Printf("   URL: %s\n", server.URL)
		if server.Description != "" {
			fmt.Printf("   Description: %s\n", server.Description)
		}
		fmt.Printf("   Status: %s\n", status)
		fmt.Println()
	}

	fmt.Print(clime.Warning.Sprint("=== End MCP Servers ===\n"))
	return true, nil
}

// HandleDebugCommand handles the /debug command
func (ui *UI) HandleDebugCommand() (bool, error) {
	ui.debugMode = !ui.debugMode

	if ui.debugMode {
		ui.ShowSuccess("Debug mode enabled")
	} else {
		ui.ShowSuccess("Debug mode disabled")
	}

	return true, nil
}

// StartAnimation implements AnimationCallback interface for tool execution
func (ui *UI) StartAnimation(toolName string) {
	displayName := ui.client.tools.GetToolDisplayName(toolName)
	ui.animation.Start(ThinkingTools, displayName)
}

// StopAnimation implements AnimationCallback interface for tool execution
func (ui *UI) StopAnimation() {
	ui.animation.Stop()
}

// sendMessage is a helper to send a message and handle the response
func (ui *UI) sendMessage(ctx context.Context, message string) error {
	ui.ShowDebug(fmt.Sprintf("Sending message: %s", message))

	// Start enhanced thinking animation
	ui.animation.Start(ThinkingLLM, "Sending message to AI")

	// Use sync version to properly handle tool calls
	response, err := ui.client.SendMessageSync(ctx, message)
	if err != nil {
		ui.animation.Stop()
		ui.ShowError(err)
		return err
	}

	// Stop thinking and show formatted response
	ui.animation.Stop()

	ui.ShowDebug("Received response from AI")

	// Format and display the response using markdown formatter
	formattedContent := ui.formatter.Format(response)
	fmt.Print(formattedContent + "\n\n")

	return nil
}

// Run starts the interactive chat loop
func (ui *UI) Run() error {
	ui.ShowWelcome()

	for {
		input, err := ui.PromptUser()
		if err != nil {
			ui.ShowError(err)
			continue
		}

		if input == "" {
			continue
		}

		// Check if it's a command
		handled, err := ui.ProcessCommand(input)
		if err != nil {
			ui.ShowError(err)
			continue
		}

		if handled {
			continue
		}

		// Show user message
		ui.ShowUserMessage(input)

		// Send message and show response
		if err := ui.sendMessage(context.Background(), input); err != nil {
			continue // Error already shown in sendMessage
		}
	}
}