package chat

import (
	"fmt"
	"regexp"
	"strings"
)

// FileOperation represents a file operation parsed from LLM output
type FileOperation struct {
	FilePath string
	Content  string
}

// ParsedOutput represents the structured output from an LLM
type ParsedOutput struct {
	TextResponse   string
	FileOperations []FileOperation
}

// StructuredOutputParser parses LLM responses for file operations
type StructuredOutputParser struct {
	filePattern *regexp.Regexp
}

// NewStructuredOutputParser creates a new output parser
func NewStructuredOutputParser() *StructuredOutputParser {
	// Pattern to match FILE: path CONTENT: content END_FILE blocks
	filePattern := regexp.MustCompile(`(?s)FILE:\s*(.+?)\s*CONTENT:\s*(.*?)\s*END_FILE`)

	return &StructuredOutputParser{
		filePattern: filePattern,
	}
}

// Parse extracts file operations and regular text from LLM output
func (p *StructuredOutputParser) Parse(output string) ParsedOutput {
	result := ParsedOutput{
		FileOperations: make([]FileOperation, 0),
	}

	// Find all file operations
	matches := p.filePattern.FindAllStringSubmatch(output, -1)

	if len(matches) == 0 {
		// No file operations found, return entire output as text
		result.TextResponse = output
		return result
	}

	// Extract file operations
	for _, match := range matches {
		if len(match) >= 3 {
			filePath := strings.TrimSpace(match[1])
			content := strings.TrimSpace(match[2])

			result.FileOperations = append(result.FileOperations, FileOperation{
				FilePath: filePath,
				Content:  content,
			})
		}
	}

	// Remove file operation blocks from the output to get remaining text
	textResponse := p.filePattern.ReplaceAllString(output, "")
	result.TextResponse = strings.TrimSpace(textResponse)

	return result
}

// ValidateFileOperations validates that file operations are safe
func (p *StructuredOutputParser) ValidateFileOperations(operations []FileOperation, workingDir string) error {
	for i, op := range operations {
		// Validate file path
		if op.FilePath == "" {
			return fmt.Errorf("file operation %d has empty file path", i+1)
		}

		// Check for path traversal attempts
		if strings.Contains(op.FilePath, "..") {
			return fmt.Errorf("file operation %d contains path traversal: %s", i+1, op.FilePath)
		}

		// Ensure it's a markdown file in a valid BSpec directory
		if !strings.HasSuffix(op.FilePath, ".md") {
			return fmt.Errorf("file operation %d is not a markdown file: %s", i+1, op.FilePath)
		}

		// Check if it's in a valid BSpec directory
		validDirs := []string{
			"contexts/", "capabilities/", "processes/", "personas/", "offerings/",
			"architecture/", "compliance/", "metrics/", "risks/", "decisions/",
			"experiments/", "profiles/",
		}

		validDir := false
		for _, dir := range validDirs {
			if strings.HasPrefix(op.FilePath, dir) {
				validDir = true
				break
			}
		}

		if !validDir {
			return fmt.Errorf("file operation %d is not in a valid BSpec directory: %s", i+1, op.FilePath)
		}

		// Validate content has basic BSpec structure
		if err := p.validateBSpecContent(op.Content); err != nil {
			return fmt.Errorf("file operation %d has invalid BSpec content: %w", i+1, err)
		}
	}

	return nil
}

// validateBSpecContent performs basic validation of BSpec document content
func (p *StructuredOutputParser) validateBSpecContent(content string) error {
	if content == "" {
		return fmt.Errorf("content cannot be empty")
	}

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

	// Check for required fields
	requiredFields := []string{"id:", "title:", "status:", "version:", "owner:"}
	for _, field := range requiredFields {
		if !strings.Contains(frontmatter, field) {
			return fmt.Errorf("missing required field: %s", strings.TrimSuffix(field, ":"))
		}
	}

	return nil
}

// FormatFileOperationSummary creates a summary of file operations for display
func (p *StructuredOutputParser) FormatFileOperationSummary(operations []FileOperation) string {
	if len(operations) == 0 {
		return "No file operations detected."
	}

	var summary strings.Builder
	summary.WriteString(fmt.Sprintf("Detected %d file operation(s):\n", len(operations)))

	for i, op := range operations {
		summary.WriteString(fmt.Sprintf("  %d. %s (%d bytes)\n", i+1, op.FilePath, len(op.Content)))
	}

	return summary.String()
}

// ExecuteFileOperations executes validated file operations using tools
func (client *Client) ExecuteFileOperations(operations []FileOperation) (string, error) {
	if len(operations) == 0 {
		return "No file operations to execute.", nil
	}

	var results strings.Builder
	results.WriteString(fmt.Sprintf("Executing %d file operation(s):\n\n", len(operations)))

	for i, op := range operations {
		// Use the write tool to create the file
		params := map[string]interface{}{
			"file_path": op.FilePath,
			"content":   op.Content,
		}

		result, err := client.tools.Execute("write_bspec_file", params)
		if err != nil {
			errorMsg := fmt.Sprintf("  %d. FAILED: %s - %s\n", i+1, op.FilePath, err.Error())
			results.WriteString(errorMsg)
			continue
		}

		results.WriteString(fmt.Sprintf("  %d. SUCCESS: %s - %s\n", i+1, op.FilePath, result))
	}

	return results.String(), nil
}