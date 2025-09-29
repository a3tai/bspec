package tools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Tool represents a callable tool that can be invoked by the LLM
type Tool interface {
	// Name returns the unique name of the tool
	Name() string

	// Description returns a description of what the tool does
	Description() string

	// Parameters returns the JSON schema for the tool's parameters
	Parameters() map[string]interface{}

	// Execute runs the tool with the given parameters and returns the result
	Execute(params map[string]interface{}) (string, error)
}

// AnimationCallback defines the interface for animation control during tool execution
type AnimationCallback interface {
	StartAnimation(toolName string)
	StopAnimation()
}

// Registry manages available tools for the chat system
type Registry struct {
	tools map[string]Tool
	validator *SecurityValidator
	animationCallback AnimationCallback
}

// SecurityValidator ensures tools operate within safe boundaries
type SecurityValidator struct {
	WorkingDir        string
	AllowedDirs       []string
	AllowedExtensions []string
}

// NewRegistry creates a new tool registry with security validation
func NewRegistry(workingDir string) *Registry {
	validator := &SecurityValidator{
		WorkingDir: workingDir,
		AllowedDirs: []string{
			"contexts", "capabilities", "processes", "personas", "offerings",
			"architecture", "compliance", "metrics", "risks", "decisions",
			"experiments", "profiles",
		},
		AllowedExtensions: []string{".md"},
	}

	return &Registry{
		tools:             make(map[string]Tool),
		validator:         validator,
		animationCallback: nil, // Set later via SetAnimationCallback
	}
}

// Register adds a tool to the registry
func (r *Registry) Register(tool Tool) error {
	if tool == nil {
		return fmt.Errorf("cannot register nil tool")
	}

	name := tool.Name()
	if name == "" {
		return fmt.Errorf("tool name cannot be empty")
	}

	if _, exists := r.tools[name]; exists {
		return fmt.Errorf("tool %s already registered", name)
	}

	r.tools[name] = tool
	return nil
}

// GetTool retrieves a tool by name
func (r *Registry) GetTool(name string) (Tool, bool) {
	tool, exists := r.tools[name]
	return tool, exists
}

// ListTools returns all registered tools
func (r *Registry) ListTools() []Tool {
	tools := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

// SecurityValidator returns the security validator
func (r *Registry) SecurityValidator() *SecurityValidator {
	return r.validator
}

// SetAnimationCallback sets the animation callback for tool execution
func (r *Registry) SetAnimationCallback(callback AnimationCallback) {
	r.animationCallback = callback
}

// Execute runs a tool with the given parameters
func (r *Registry) Execute(toolName string, params map[string]interface{}) (string, error) {
	tool, exists := r.tools[toolName]
	if !exists {
		return "", fmt.Errorf("tool %s not found", toolName)
	}

	// Start animation if callback is available
	if r.animationCallback != nil {
		r.animationCallback.StartAnimation(toolName)
		defer r.animationCallback.StopAnimation()
	}

	return tool.Execute(params)
}

// GetToolDisplayName returns a human-readable name for the tool for animation display
func (r *Registry) GetToolDisplayName(toolName string) string {
	// Convert tool names to human-readable format
	switch {
	case strings.Contains(toolName, "write"):
		return "Writing files"
	case strings.Contains(toolName, "read"):
		return "Reading files"
	case strings.Contains(toolName, "validate"):
		return "Validating documents"
	case strings.Contains(toolName, "list"):
		return "Listing files"
	case strings.Contains(toolName, "bspec"):
		return "Processing with BSpec"
	default:
		return "Executing tools"
	}
}

// ToOpenAITools converts registered tools to OpenAI tool definitions
func (r *Registry) ToOpenAITools() []map[string]interface{} {
	tools := make([]map[string]interface{}, 0, len(r.tools))

	for _, tool := range r.tools {
		toolDef := map[string]interface{}{
			"type": "function",
			"function": map[string]interface{}{
				"name":        tool.Name(),
				"description": tool.Description(),
				"parameters":  tool.Parameters(),
			},
		}
		tools = append(tools, toolDef)
	}

	return tools
}

// ValidatePath ensures a file path is within allowed boundaries
func (sv *SecurityValidator) ValidatePath(path string) error {
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}

	// Prevent directory traversal attacks
	if containsPathTraversal(path) {
		return fmt.Errorf("path traversal not allowed: %s", path)
	}

	// Check if path is within working directory boundaries
	if !isWithinWorkingDir(sv.WorkingDir, path) {
		return fmt.Errorf("path outside working directory: %s", path)
	}

	// Validate file extension if it's a file path
	if !sv.isAllowedExtension(path) {
		return fmt.Errorf("file extension not allowed: %s", path)
	}

	return nil
}

// ValidateDirectory ensures a directory path is allowed for BSpec operations
func (sv *SecurityValidator) ValidateDirectory(dirPath string) error {
	if err := sv.ValidatePath(dirPath); err != nil {
		return err
	}

	// Check if directory is in allowed list
	for _, allowedDir := range sv.AllowedDirs {
		if dirPath == allowedDir || containsSubPath(dirPath, allowedDir) {
			return nil
		}
	}

	return fmt.Errorf("directory not allowed for BSpec operations: %s", dirPath)
}

// Helper functions for security validation

func containsPathTraversal(path string) bool {
	return len(path) >= 2 && (path[:2] == ".." ||
		len(path) >= 3 && (path[:3] == "../" || path[len(path)-3:] == "/.." ||
		containsSubstring(path, "/../")))
}

func isWithinWorkingDir(workingDir, path string) bool {
	// For now, simple check - path should not start with / (absolute) or contain ../
	return !startsWithSlash(path) && !containsPathTraversal(path)
}

func (sv *SecurityValidator) isAllowedExtension(path string) bool {
	// If it's a directory (no extension), allow it
	if !containsSubstring(path, ".") {
		return true
	}

	for _, ext := range sv.AllowedExtensions {
		if endsWithSuffix(path, ext) {
			return true
		}
	}
	return false
}

func containsSubPath(path, subPath string) bool {
	return path == subPath ||
		(len(path) > len(subPath) && path[:len(subPath)+1] == subPath+"/")
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func startsWithSlash(s string) bool {
	return len(s) > 0 && s[0] == '/'
}

func endsWithSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// ToolCallRequest represents a tool call from the LLM
type ToolCallRequest struct {
	ID       string                 `json:"id"`
	Type     string                 `json:"type"`
	Function ToolFunctionCall       `json:"function"`
}

// ToolFunctionCall represents the function call details
type ToolFunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// ToolCallResponse represents the response to send back to the LLM
type ToolCallResponse struct {
	ToolCallID string `json:"tool_call_id"`
	Role       string `json:"role"`
	Content    string `json:"content"`
}

// ParseArguments parses the JSON arguments from a tool call
func (tcr *ToolCallRequest) ParseArguments() (map[string]interface{}, error) {
	var args map[string]interface{}
	if err := json.Unmarshal([]byte(tcr.Function.Arguments), &args); err != nil {
		return nil, fmt.Errorf("failed to parse tool arguments: %w", err)
	}
	return args, nil
}

// NewToolCallResponse creates a new tool call response
func NewToolCallResponse(toolCallID, content string) ToolCallResponse {
	return ToolCallResponse{
		ToolCallID: toolCallID,
		Role:       "tool",
		Content:    content,
	}
}