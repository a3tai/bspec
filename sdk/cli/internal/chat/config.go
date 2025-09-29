package chat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Config holds the configuration for chat functionality
type Config struct {
	// API key from environment or settings
	APIKey string

	// LLM provider from BSPEC_LLM (supports "openrouter", "openai")
	LLMProvider string

	// Model name from BSPEC_MODEL (defaults to "x-ai/grok-4-fast:free")
	Model string

	// Base URL for API endpoint
	BaseURL string

	// Maximum tokens for responses
	MaxTokens int

	// Temperature for creativity (0.0 to 1.0)
	Temperature float32

	// Settings file path
	SettingsPath string

	// Prompts directory path
	PromptsPath string

	// MCP servers configuration
	MCPServers []MCPServer
}

// MCPServer represents an MCP server configuration
type MCPServer struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled"`
}

// Settings represents the stored configuration
type Settings struct {
	LLMProvider string      `json:"llm_provider,omitempty"`
	Model       string      `json:"model,omitempty"`
	MaxTokens   int         `json:"max_tokens,omitempty"`
	Temperature float32     `json:"temperature,omitempty"`
	MCPServers  []MCPServer `json:"mcp_servers,omitempty"`
}

// NewConfig creates a new configuration from environment variables and settings file
func NewConfig() (*Config, error) {
	// Set up paths
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	bspecDir := filepath.Join(homeDir, ".bspec")
	settingsPath := filepath.Join(bspecDir, "settings.json")
	promptsPath := filepath.Join(bspecDir, "prompts")

	config := &Config{
		// Default values - OpenRouter with Grok
		LLMProvider:  "openrouter",
		Model:        "x-ai/grok-4-fast:free",
		MaxTokens:    2000,
		Temperature:  0.7,
		SettingsPath: settingsPath,
		PromptsPath:  promptsPath,
		MCPServers: []MCPServer{
			{
				Name:        "bspec",
				URL:         "https://mcp.bspec.dev/mcp",
				Description: "BSpec Business Specification Standard MCP server",
				Enabled:     true,
			},
		},
	}

	// Load settings from file if it exists
	if err := config.loadSettings(); err != nil {
		// If settings file doesn't exist, create it with defaults
		if os.IsNotExist(err) {
			if err := config.createDefaultSettings(); err != nil {
				return nil, fmt.Errorf("failed to create default settings: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to load settings: %w", err)
		}
	}

	// Read API key (required) - prioritize OPENROUTER_API_KEY for the default provider
	config.APIKey = os.Getenv("OPENROUTER_API_KEY")
	if config.APIKey == "" {
		config.APIKey = os.Getenv("BSPEC_OPENAI_API_KEY") // Backward compatibility for OpenAI
	}
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required. Set OPENROUTER_API_KEY (for OpenRouter) or BSPEC_OPENAI_API_KEY (for OpenAI) environment variable")
	}

	// Read LLM provider (optional, defaults to "openrouter")
	if provider := os.Getenv("BSPEC_LLM"); provider != "" {
		config.LLMProvider = strings.ToLower(strings.TrimSpace(provider))
	}

	// Set base URL based on provider
	switch config.LLMProvider {
	case "openrouter":
		config.BaseURL = "https://openrouter.ai/api/v1"
	case "openai":
		config.BaseURL = "https://api.openai.com/v1"
	default:
		return nil, fmt.Errorf("unsupported LLM provider: %s (supported: openrouter, openai)", config.LLMProvider)
	}

	// Read model name and set appropriate default based on provider
	if model := os.Getenv("BSPEC_MODEL"); model != "" {
		config.Model = strings.TrimSpace(model)
	} else {
		// Set provider-specific defaults
		switch config.LLMProvider {
		case "openrouter":
			config.Model = "x-ai/grok-4-fast:free"
		case "openai":
			config.Model = "gpt-4o-mini"
		}
	}

	return config, nil
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.APIKey == "" {
		return fmt.Errorf("API key cannot be empty")
	}

	if c.LLMProvider == "" {
		return fmt.Errorf("LLM provider cannot be empty")
	}

	if c.Model == "" {
		return fmt.Errorf("model cannot be empty")
	}

	if c.BaseURL == "" {
		return fmt.Errorf("base URL cannot be empty")
	}

	if c.MaxTokens <= 0 {
		return fmt.Errorf("max tokens must be positive")
	}

	if c.Temperature < 0.0 || c.Temperature > 1.0 {
		return fmt.Errorf("temperature must be between 0.0 and 1.0")
	}

	return nil
}

// String returns a string representation of the config (without exposing the API key)
func (c *Config) String() string {
	apiKeyMask := "***"
	if len(c.APIKey) > 7 {
		apiKeyMask = c.APIKey[:4] + "..." + c.APIKey[len(c.APIKey)-4:]
	}

	return fmt.Sprintf("Config{Provider: %s, Model: %s, BaseURL: %s, APIKey: %s, MaxTokens: %d, Temperature: %.2f}",
		c.LLMProvider, c.Model, c.BaseURL, apiKeyMask, c.MaxTokens, c.Temperature)
}

// OpenRouterModel represents a model from the OpenRouter API
type OpenRouterModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Pricing     struct {
		Prompt     string `json:"prompt"`
		Completion string `json:"completion"`
	} `json:"pricing"`
	ContextLength int `json:"context_length"`
	TopProvider   struct {
		MaxCompletionTokens int `json:"max_completion_tokens"`
	} `json:"top_provider"`
}

// OpenRouterResponse represents the API response structure
type OpenRouterResponse struct {
	Data []OpenRouterModel `json:"data"`
}

// FetchAvailableModels fetches models from OpenRouter API
func (c *Config) FetchAvailableModels() ([]OpenRouterModel, error) {
	if c.LLMProvider != "openrouter" {
		return nil, fmt.Errorf("model fetching only supported for OpenRouter provider")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://openrouter.ai/api/v1/models", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add authorization header if available
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch models: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var response OpenRouterResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Data, nil
}

// loadSettings loads configuration from the settings file
func (c *Config) loadSettings() error {
	data, err := os.ReadFile(c.SettingsPath)
	if err != nil {
		return err
	}

	var settings Settings
	if err := json.Unmarshal(data, &settings); err != nil {
		return fmt.Errorf("failed to parse settings file: %w", err)
	}

	// Apply settings (environment variables take precedence)
	if settings.LLMProvider != "" && os.Getenv("BSPEC_LLM") == "" {
		c.LLMProvider = settings.LLMProvider
	}
	if settings.Model != "" && os.Getenv("BSPEC_MODEL") == "" {
		c.Model = settings.Model
	}
	if settings.MaxTokens > 0 {
		c.MaxTokens = settings.MaxTokens
	}
	if settings.Temperature > 0 {
		c.Temperature = settings.Temperature
	}
	if len(settings.MCPServers) > 0 {
		c.MCPServers = settings.MCPServers
	}

	return nil
}

// SaveSettings saves current configuration to the settings file
func (c *Config) SaveSettings() error {
	settings := Settings{
		LLMProvider: c.LLMProvider,
		Model:       c.Model,
		MaxTokens:   c.MaxTokens,
		Temperature: c.Temperature,
		MCPServers:  c.MCPServers,
	}

	// Ensure the .bspec directory exists
	if err := os.MkdirAll(filepath.Dir(c.SettingsPath), 0755); err != nil {
		return fmt.Errorf("failed to create .bspec directory: %w", err)
	}

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %w", err)
	}

	if err := os.WriteFile(c.SettingsPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write settings file: %w", err)
	}

	return nil
}

// EnsurePromptsDirectory ensures the prompts directory exists
func (c *Config) EnsurePromptsDirectory() error {
	return os.MkdirAll(c.PromptsPath, 0755)
}

// LoadPrompt loads a prompt from the prompts directory
func (c *Config) LoadPrompt(name string) (string, error) {
	promptPath := filepath.Join(c.PromptsPath, name+".md")
	data, err := os.ReadFile(promptPath)
	if err != nil {
		return "", fmt.Errorf("failed to load prompt %s: %w", name, err)
	}
	return string(data), nil
}

// LoadProjectContext loads the project context prompt with current working directory info
func (c *Config) LoadProjectContext() (string, error) {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Load context template
	contextTemplate, err := c.LoadPrompt("context")
	if err != nil {
		return "", err
	}

	// Scan for BSpec directories and documents
	documentList := c.scanBSpecDocuments(cwd)

	// Replace placeholders
	contextPrompt := strings.ReplaceAll(contextTemplate, "{PROJECT_PATH}", cwd)
	contextPrompt = strings.ReplaceAll(contextPrompt, "{DOCUMENT_LIST}", documentList)

	return contextPrompt, nil
}

// scanBSpecDocuments scans the current directory for BSpec documents and returns a summary
func (c *Config) scanBSpecDocuments(projectPath string) string {
	var result strings.Builder

	bspecDirs := []string{
		"contexts", "capabilities", "processes", "personas", "offerings",
		"architecture", "compliance", "metrics", "risks", "decisions",
		"experiments", "profiles",
	}

	for _, dir := range bspecDirs {
		dirPath := filepath.Join(projectPath, dir)
		if entries, err := os.ReadDir(dirPath); err == nil && len(entries) > 0 {
			result.WriteString(fmt.Sprintf("\n**%s/**: ", dir))
			count := 0
			for _, entry := range entries {
				if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
					if count > 0 {
						result.WriteString(", ")
					}
					// Remove .md extension for display
					name := strings.TrimSuffix(entry.Name(), ".md")
					result.WriteString(name)
					count++
					if count >= 5 { // Limit to 5 files per directory
						if len(entries) > 5 {
							result.WriteString(fmt.Sprintf(", and %d more", len(entries)-5))
						}
						break
					}
				}
			}
		}
	}

	if result.Len() == 0 {
		return "No BSpec documents found in current directory."
	}

	return result.String()
}

// GetPopularModels returns a list of popular models for each provider
func GetPopularModels() map[string][]string {
	return map[string][]string{
		"openrouter": {
			"x-ai/grok-4-fast:free",
			"x-ai/grok-2-1212",
			"x-ai/grok-vision-beta",
			"anthropic/claude-3.5-sonnet",
			"anthropic/claude-3.5-haiku",
			"openai/gpt-4o",
			"openai/gpt-4o-mini",
			"openai/o1-preview",
			"openai/o1-mini",
			"google/gemini-pro-1.5",
			"google/gemini-flash-1.5",
			"meta-llama/llama-3.1-405b-instruct",
			"mistralai/mistral-large",
			"deepseek/deepseek-chat",
		},
		"openai": {
			"gpt-4o",
			"gpt-4o-mini",
			"o1-preview",
			"o1-mini",
			"gpt-4-turbo",
		},
	}
}

// GetModelDescription returns a description for common models
func GetModelDescription(model string) string {
	descriptions := map[string]string{
		"x-ai/grok-4-fast:free":        "Grok 4 Fast Free - xAI's fast model optimized for quick responses (free tier)",
		"x-ai/grok-2-1212":             "Grok 2 - xAI's latest model with strong reasoning and real-time data",
		"x-ai/grok-vision-beta":        "Grok Vision Beta - Multimodal model with image understanding",
		"anthropic/claude-3.5-sonnet":  "Claude 3.5 Sonnet - Excellent for analysis, coding, and writing",
		"anthropic/claude-3.5-haiku":   "Claude 3.5 Haiku - Fast and efficient for simple tasks",
		"openai/gpt-4o":                "GPT-4o - Advanced reasoning and multimodal capabilities",
		"openai/gpt-4o-mini":           "GPT-4o Mini - Faster and more affordable version of GPT-4o",
		"openai/o1-preview":            "O1 Preview - OpenAI's reasoning model for complex problems",
		"openai/o1-mini":               "O1 Mini - Lightweight reasoning model",
		"google/gemini-pro-1.5":        "Gemini Pro 1.5 - Google's advanced model with long context",
		"google/gemini-flash-1.5":      "Gemini Flash 1.5 - Fast and efficient version",
		"meta-llama/llama-3.1-405b-instruct": "Llama 3.1 405B - Meta's largest open model",
		"mistralai/mistral-large":      "Mistral Large - European AI with strong reasoning",
		"deepseek/deepseek-chat":       "DeepSeek Chat - Chinese model with strong coding abilities",
		"gpt-4o":                       "GPT-4o - Advanced reasoning and multimodal capabilities",
		"gpt-4o-mini":                  "GPT-4o Mini - Faster and more affordable version of GPT-4o",
		"o1-preview":                   "O1 Preview - OpenAI's reasoning model for complex problems",
		"o1-mini":                      "O1 Mini - Lightweight reasoning model",
		"gpt-4-turbo":                  "GPT-4 Turbo - Enhanced version of GPT-4",
	}

	if desc, exists := descriptions[model]; exists {
		return desc
	}
	return "No description available"
}

// SetModel updates the current model and saves to settings
func (c *Config) SetModel(model string) error {
	c.Model = model
	return c.SaveSettings()
}

// GetEnabledMCPServers returns all enabled MCP servers
func (c *Config) GetEnabledMCPServers() []MCPServer {
	var enabled []MCPServer
	for _, server := range c.MCPServers {
		if server.Enabled {
			enabled = append(enabled, server)
		}
	}
	return enabled
}

// AddMCPServer adds a new MCP server to the configuration
func (c *Config) AddMCPServer(server MCPServer) error {
	// Check if server with same name already exists
	for i, existing := range c.MCPServers {
		if existing.Name == server.Name {
			c.MCPServers[i] = server // Update existing
			return c.SaveSettings()
		}
	}

	c.MCPServers = append(c.MCPServers, server)
	return c.SaveSettings()
}

// EnableMCPServer enables or disables an MCP server by name
func (c *Config) EnableMCPServer(name string, enabled bool) error {
	for i, server := range c.MCPServers {
		if server.Name == name {
			c.MCPServers[i].Enabled = enabled
			return c.SaveSettings()
		}
	}
	return fmt.Errorf("MCP server '%s' not found", name)
}

// RemoveMCPServer removes an MCP server by name
func (c *Config) RemoveMCPServer(name string) error {
	for i, server := range c.MCPServers {
		if server.Name == name {
			c.MCPServers = append(c.MCPServers[:i], c.MCPServers[i+1:]...)
			return c.SaveSettings()
		}
	}
	return fmt.Errorf("MCP server '%s' not found", name)
}

// createDefaultSettings creates the default settings file
func (c *Config) createDefaultSettings() error {
	// Create the default settings with current configuration
	return c.SaveSettings()
}