package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"github.com/a3tai/bspec/cli/internal/sdk"
)

// ReadBSpecFileTool reads BSpec markdown files within the project
type ReadBSpecFileTool struct {
	validator *SecurityValidator
}

// NewReadBSpecFileTool creates a new file reading tool
func NewReadBSpecFileTool(validator *SecurityValidator) *ReadBSpecFileTool {
	return &ReadBSpecFileTool{validator: validator}
}

func (t *ReadBSpecFileTool) Name() string {
	return "read_bspec_file"
}

func (t *ReadBSpecFileTool) Description() string {
	return "Read the contents of a BSpec markdown file within the current project. Only works with .md files in standard BSpec directories (contexts/, capabilities/, etc.)"
}

func (t *ReadBSpecFileTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"file_path": map[string]interface{}{
				"type":        "string",
				"description": "Path to the BSpec markdown file relative to project root (e.g., 'contexts/CTX-example-v1.0.0.md')",
			},
		},
		"required": []string{"file_path"},
	}
}

func (t *ReadBSpecFileTool) Execute(params map[string]interface{}) (string, error) {
	filePath, ok := params["file_path"].(string)
	if !ok {
		return "", fmt.Errorf("file_path parameter is required and must be a string")
	}

	// Validate the file path
	if err := t.validator.ValidatePath(filePath); err != nil {
		return "", fmt.Errorf("invalid file path: %w", err)
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	return string(content), nil
}

// WriteBSpecFileTool writes BSpec markdown files with validation
type WriteBSpecFileTool struct {
	validator *SecurityValidator
}

// NewWriteBSpecFileTool creates a new file writing tool
func NewWriteBSpecFileTool(validator *SecurityValidator) *WriteBSpecFileTool {
	return &WriteBSpecFileTool{validator: validator}
}

func (t *WriteBSpecFileTool) Name() string {
	return "write_bspec_file"
}

func (t *WriteBSpecFileTool) Description() string {
	return "Write content to a BSpec markdown file. Creates directories if needed. Content must include proper YAML frontmatter with required fields (id, title, status, version, owner)."
}

func (t *WriteBSpecFileTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"file_path": map[string]interface{}{
				"type":        "string",
				"description": "Path to the BSpec markdown file relative to project root (e.g., 'contexts/CTX-example-v1.0.0.md')",
			},
			"content": map[string]interface{}{
				"type":        "string",
				"description": "Full content of the file including YAML frontmatter and markdown body",
			},
		},
		"required": []string{"file_path", "content"},
	}
}

func (t *WriteBSpecFileTool) Execute(params map[string]interface{}) (string, error) {
	filePath, ok := params["file_path"].(string)
	if !ok {
		return "", fmt.Errorf("file_path parameter is required and must be a string")
	}

	content, ok := params["content"].(string)
	if !ok {
		return "", fmt.Errorf("content parameter is required and must be a string")
	}

	// Validate the file path
	if err := t.validator.ValidatePath(filePath); err != nil {
		return "", fmt.Errorf("invalid file path: %w", err)
	}

	// Validate that content has proper BSpec structure
	if err := validateBSpecContent(content); err != nil {
		return "", fmt.Errorf("invalid BSpec content: %w", err)
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write the file
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return fmt.Sprintf("Successfully wrote %d bytes to %s", len(content), filePath), nil
}

// ListBSpecDirectoryTool lists contents of BSpec project directories
type ListBSpecDirectoryTool struct {
	validator *SecurityValidator
}

// NewListBSpecDirectoryTool creates a new directory listing tool
func NewListBSpecDirectoryTool(validator *SecurityValidator) *ListBSpecDirectoryTool {
	return &ListBSpecDirectoryTool{validator: validator}
}

func (t *ListBSpecDirectoryTool) Name() string {
	return "list_bspec_directory"
}

func (t *ListBSpecDirectoryTool) Description() string {
	return "List contents of a BSpec project directory. Only works with standard BSpec directories (contexts/, capabilities/, processes/, etc.)"
}

func (t *ListBSpecDirectoryTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"directory_path": map[string]interface{}{
				"type":        "string",
				"description": "Path to the BSpec directory relative to project root (e.g., 'contexts', 'capabilities')",
			},
		},
		"required": []string{"directory_path"},
	}
}

func (t *ListBSpecDirectoryTool) Execute(params map[string]interface{}) (string, error) {
	dirPath, ok := params["directory_path"].(string)
	if !ok {
		return "", fmt.Errorf("directory_path parameter is required and must be a string")
	}

	// Validate the directory path
	if err := t.validator.ValidateDirectory(dirPath); err != nil {
		return "", fmt.Errorf("invalid directory path: %w", err)
	}

	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Sprintf("Directory does not exist: %s", dirPath), nil
	}

	// Read directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return "", fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	// Format the listing
	var result strings.Builder
	result.WriteString(fmt.Sprintf("Contents of %s:\n", dirPath))

	if len(entries) == 0 {
		result.WriteString("  (empty directory)\n")
	} else {
		for _, entry := range entries {
			name := entry.Name()
			if entry.IsDir() {
				result.WriteString(fmt.Sprintf("  %s/ (directory)\n", name))
			} else {
				info, _ := entry.Info()
				result.WriteString(fmt.Sprintf("  %s (%d bytes)\n", name, info.Size()))
			}
		}
	}

	return result.String(), nil
}

// ValidateBSpecDocumentTool validates BSpec document structure and cross-references
type ValidateBSpecDocumentTool struct {
	validator *SecurityValidator
}

// NewValidateBSpecDocumentTool creates a new document validation tool
func NewValidateBSpecDocumentTool(validator *SecurityValidator) *ValidateBSpecDocumentTool {
	return &ValidateBSpecDocumentTool{validator: validator}
}

func (t *ValidateBSpecDocumentTool) Name() string {
	return "validate_bspec_document"
}

func (t *ValidateBSpecDocumentTool) Description() string {
	return "Validate a BSpec document for proper YAML frontmatter, required fields, and structure compliance"
}

func (t *ValidateBSpecDocumentTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"file_path": map[string]interface{}{
				"type":        "string",
				"description": "Path to the BSpec markdown file to validate",
			},
		},
		"required": []string{"file_path"},
	}
}

func (t *ValidateBSpecDocumentTool) Execute(params map[string]interface{}) (string, error) {
	filePath, ok := params["file_path"].(string)
	if !ok {
		return "", fmt.Errorf("file_path parameter is required and must be a string")
	}

	// Validate the file path
	if err := t.validator.ValidatePath(filePath); err != nil {
		return "", fmt.Errorf("invalid file path: %w", err)
	}

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Validate BSpec structure
	if err := validateBSpecContent(string(content)); err != nil {
		return fmt.Sprintf("Validation failed for %s: %s", filePath, err.Error()), nil
	}

	return fmt.Sprintf("Document %s is valid BSpec format", filePath), nil
}

// Helper function to validate BSpec content structure using SDK
func validateBSpecContent(content string) error {
	// Check for YAML frontmatter
	if !strings.HasPrefix(content, "---") {
		return fmt.Errorf("document must start with YAML frontmatter (---)")
	}

	lines := strings.Split(content, "\n")
	if len(lines) < 3 {
		return fmt.Errorf("document too short to contain valid YAML frontmatter")
	}

	// Find the end of frontmatter
	frontmatterEnd := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			frontmatterEnd = i
			break
		}
	}

	if frontmatterEnd == -1 {
		return fmt.Errorf("YAML frontmatter not properly closed with ---")
	}

	// Extract frontmatter content
	frontmatter := strings.Join(lines[1:frontmatterEnd], "\n")

	// Parse YAML frontmatter
	var metadata map[string]interface{}
	if err := yaml.Unmarshal([]byte(frontmatter), &metadata); err != nil {
		return fmt.Errorf("invalid YAML frontmatter: %w", err)
	}

	// Create SDK converter instance to access schemas
	converter := sdk.NewSDKConverter()

	// Validate using SDK schemas
	if err := converter.ValidateDocumentMetadata(metadata); err != nil {
		return fmt.Errorf("document validation failed: %w", err)
	}

	// Additional basic validation for required fields
	requiredFields := []string{"id", "title", "status", "version", "owner"}
	for _, field := range requiredFields {
		if _, exists := metadata[field]; !exists {
			return fmt.Errorf("missing required field: %s", field)
		}
	}

	// Validate status values
	status, ok := metadata["status"].(string)
	if !ok {
		return fmt.Errorf("status field must be a string")
	}

	validStatuses := []string{"Draft", "Accepted", "Deprecated"}
	statusValid := false
	for _, validStatus := range validStatuses {
		if status == validStatus {
			statusValid = true
			break
		}
	}

	if !statusValid {
		return fmt.Errorf("status must be one of: %s", strings.Join(validStatuses, ", "))
	}

	return nil
}