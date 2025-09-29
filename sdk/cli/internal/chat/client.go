package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/a3tai/bspec/cli/internal/chat/tools"
)

// Message represents a chat message
type Message struct {
	Role    string    `json:"role"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

// StreamResponse represents a streaming response
type StreamResponse struct {
	Content string
	Done    bool
	Error   error
}

// DebugCallback defines the interface for debug logging during tool execution
type DebugCallback interface {
	ShowDebug(message string)
}

// Client wraps the OpenAI client with additional functionality
type Client struct {
	config        *Config
	openai        *openai.Client
	messages      []Message
	tools         *tools.Registry
	debugCallback DebugCallback
}

// NewClient creates a new chat client
func NewClient(config *Config) (*Client, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Create OpenAI client with custom base URL for OpenRouter compatibility
	clientConfig := openai.DefaultConfig(config.APIKey)
	clientConfig.BaseURL = config.BaseURL

	// Get current working directory for tools
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	// Create tools registry and register all tools
	toolsRegistry := tools.NewRegistry(workingDir)
	if err := registerBSpecTools(toolsRegistry); err != nil {
		return nil, fmt.Errorf("failed to register tools: %w", err)
	}

	client := &Client{
		config:        config,
		openai:        openai.NewClientWithConfig(clientConfig),
		messages:      make([]Message, 0),
		tools:         toolsRegistry,
		debugCallback: nil, // Set later via SetDebugCallback
	}

	return client, nil
}

// AddMessage adds a message to the conversation history
func (c *Client) AddMessage(role, content string) {
	c.messages = append(c.messages, Message{
		Role:    role,
		Content: content,
		Time:    time.Now(),
	})
}

// GetMessages returns the conversation history
func (c *Client) GetMessages() []Message {
	return c.messages
}

// ClearHistory clears the conversation history
func (c *Client) ClearHistory() {
	c.messages = make([]Message, 0)
}

// SetDebugCallback sets the debug callback for tool execution
func (c *Client) SetDebugCallback(callback DebugCallback) {
	c.debugCallback = callback
}

// convertToOpenAIMessages converts our messages to OpenAI format
func (c *Client) convertToOpenAIMessages() []openai.ChatCompletionMessage {
	openaiMessages := make([]openai.ChatCompletionMessage, len(c.messages))
	for i, msg := range c.messages {
		openaiMessages[i] = openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}
	return openaiMessages
}

// convertToolsToOpenAI converts tools to OpenAI tool format
func (c *Client) convertToolsToOpenAI() []openai.Tool {
	toolDefs := c.tools.ToOpenAITools()
	openaiTools := make([]openai.Tool, len(toolDefs))

	for i, toolDef := range toolDefs {
		openaiTools[i] = openai.Tool{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        toolDef["function"].(map[string]interface{})["name"].(string),
				Description: toolDef["function"].(map[string]interface{})["description"].(string),
				Parameters:  toolDef["function"].(map[string]interface{})["parameters"],
			},
		}
	}

	return openaiTools
}

// SendMessage sends a message and returns a streaming response
func (c *Client) SendMessage(ctx context.Context, userMessage string) (<-chan StreamResponse, error) {
	// Add user message to history
	c.AddMessage(openai.ChatMessageRoleUser, userMessage)

	// Create response channel
	responseChan := make(chan StreamResponse, 10)

	go func() {
		defer close(responseChan)

		// Create chat completion request with tools
		req := openai.ChatCompletionRequest{
			Model:       c.config.Model,
			Messages:    c.convertToOpenAIMessages(),
			MaxTokens:   c.config.MaxTokens,
			Temperature: c.config.Temperature,
			Stream:      true,
			Tools:       c.convertToolsToOpenAI(),
		}

		// Create streaming completion
		stream, err := c.openai.CreateChatCompletionStream(ctx, req)
		if err != nil {
			responseChan <- StreamResponse{Error: fmt.Errorf("failed to create chat completion stream: %w", err)}
			return
		}
		defer stream.Close()

		var fullResponse strings.Builder

		// Read streaming responses
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				// Stream finished - add assistant message to history
				finalContent := fullResponse.String()
				c.AddMessage(openai.ChatMessageRoleAssistant, finalContent)

				responseChan <- StreamResponse{
					Content: "",
					Done:    true,
					Error:   nil,
				}
				return
			}

			if err != nil {
				responseChan <- StreamResponse{Error: fmt.Errorf("stream error: %w", err)}
				return
			}

			// Extract content from response
			if len(response.Choices) > 0 {
				content := response.Choices[0].Delta.Content
				if content != "" {
					fullResponse.WriteString(content)
					responseChan <- StreamResponse{
						Content: content,
						Done:    false,
						Error:   nil,
					}
				}
			}
		}
	}()

	return responseChan, nil
}

// SendMessageSync sends a message and returns the complete response (non-streaming)
func (c *Client) SendMessageSync(ctx context.Context, userMessage string) (string, error) {
	// Add user message to history
	c.AddMessage(openai.ChatMessageRoleUser, userMessage)

	// Create chat completion request with tools
	req := openai.ChatCompletionRequest{
		Model:       c.config.Model,
		Messages:    c.convertToOpenAIMessages(),
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
		Stream:      false,
		Tools:       c.convertToolsToOpenAI(),
	}

	// Create completion
	resp, err := c.openai.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned")
	}

	choice := resp.Choices[0]

	// Check if the response includes tool calls
	if len(choice.Message.ToolCalls) > 0 {
		return c.handleToolCalls(ctx, choice.Message.ToolCalls)
	}

	content := choice.Message.Content

	// Add assistant message to history
	c.AddMessage(openai.ChatMessageRoleAssistant, content)

	return content, nil
}

// handleToolCalls processes tool calls and returns the final response
func (c *Client) handleToolCalls(ctx context.Context, toolCalls []openai.ToolCall) (string, error) {
	if c.debugCallback != nil {
		c.debugCallback.ShowDebug(fmt.Sprintf("Processing %d tool call(s)", len(toolCalls)))
	}

	// Add the assistant message with tool calls to history
	c.AddMessage(openai.ChatMessageRoleAssistant, "Executing tools...")

	// Execute each tool call
	for _, toolCall := range toolCalls {
		if c.debugCallback != nil {
			c.debugCallback.ShowDebug(fmt.Sprintf("Executing tool: %s", toolCall.Function.Name))
		}

		result, err := c.executeToolCall(toolCall)
		if err != nil {
			result = fmt.Sprintf("Error executing tool: %s", err.Error())
			if c.debugCallback != nil {
				c.debugCallback.ShowDebug(fmt.Sprintf("Tool %s failed: %s", toolCall.Function.Name, err.Error()))
			}
		} else if c.debugCallback != nil {
			c.debugCallback.ShowDebug(fmt.Sprintf("Tool %s completed successfully", toolCall.Function.Name))
		}

		// Add tool result to history
		c.AddMessage("tool", fmt.Sprintf("Tool: %s\nResult: %s", toolCall.Function.Name, result))
	}

	// Send follow-up request to get final response
	req := openai.ChatCompletionRequest{
		Model:       c.config.Model,
		Messages:    c.convertToOpenAIMessages(),
		MaxTokens:   c.config.MaxTokens,
		Temperature: c.config.Temperature,
		Stream:      false,
	}

	resp, err := c.openai.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create follow-up completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned")
	}

	content := resp.Choices[0].Message.Content

	// Add final assistant message to history
	c.AddMessage(openai.ChatMessageRoleAssistant, content)

	return content, nil
}

// executeToolCall executes a single tool call
func (c *Client) executeToolCall(toolCall openai.ToolCall) (string, error) {
	// Parse the arguments
	var args map[string]interface{}
	if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
		return "", fmt.Errorf("failed to parse tool arguments: %w", err)
	}

	// Execute the tool
	return c.tools.Execute(toolCall.Function.Name, args)
}

// SetSystemPrompt sets or updates the system prompt (first message)
func (c *Client) SetSystemPrompt(systemPrompt string) {
	// Remove existing system message if present
	if len(c.messages) > 0 && c.messages[0].Role == openai.ChatMessageRoleSystem {
		c.messages = c.messages[1:]
	}

	// Add new system message at the beginning
	newMessage := Message{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemPrompt,
		Time:    time.Now(),
	}

	c.messages = append([]Message{newMessage}, c.messages...)
}

// GetLastAssistantMessage returns the last assistant message
func (c *Client) GetLastAssistantMessage() string {
	for i := len(c.messages) - 1; i >= 0; i-- {
		if c.messages[i].Role == openai.ChatMessageRoleAssistant {
			return c.messages[i].Content
		}
	}
	return ""
}

// GetConversationLength returns the number of messages in the conversation
func (c *Client) GetConversationLength() int {
	return len(c.messages)
}

// GetTools returns the tools registry
func (c *Client) GetTools() *tools.Registry {
	return c.tools
}

// registerBSpecTools registers all available BSpec tools
func registerBSpecTools(registry *tools.Registry) error {
	validator := registry.SecurityValidator()

	// Register file system tools
	if err := registry.Register(tools.NewReadBSpecFileTool(validator)); err != nil {
		return err
	}
	if err := registry.Register(tools.NewWriteBSpecFileTool(validator)); err != nil {
		return err
	}
	if err := registry.Register(tools.NewListBSpecDirectoryTool(validator)); err != nil {
		return err
	}
	if err := registry.Register(tools.NewValidateBSpecDocumentTool(validator)); err != nil {
		return err
	}

	// Register BSpec command tools
	if err := registry.Register(tools.NewBSpecQueryTool(validator)); err != nil {
		return err
	}
	if err := registry.Register(tools.NewBSpecPackTool(validator)); err != nil {
		return err
	}
	if err := registry.Register(tools.NewBSpecValidateTool(validator)); err != nil {
		return err
	}
	if err := registry.Register(tools.NewBSpecInitTool(validator)); err != nil {
		return err
	}
	if err := registry.Register(tools.NewBSpecVersionTool(validator)); err != nil {
		return err
	}

	return nil
}