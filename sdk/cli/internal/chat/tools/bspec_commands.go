package tools

import (
	"fmt"
	"os/exec"
	"strings"
)

// BSpecQueryTool executes bspec query commands
type BSpecQueryTool struct {
	validator *SecurityValidator
}

// NewBSpecQueryTool creates a new BSpec query tool
func NewBSpecQueryTool(validator *SecurityValidator) *BSpecQueryTool {
	return &BSpecQueryTool{validator: validator}
}

func (t *BSpecQueryTool) Name() string {
	return "bspec_query"
}

func (t *BSpecQueryTool) Description() string {
	return "Execute bspec query command to analyze project documents with optional filtering by type, domain, and status"
}

func (t *BSpecQueryTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"output_format": map[string]interface{}{
				"type":        "string",
				"description": "Output format: markdown, yaml, or json",
				"enum":        []string{"markdown", "yaml", "json"},
				"default":     "markdown",
			},
			"document_type": map[string]interface{}{
				"type":        "string",
				"description": "Filter by document type (CAP, CTX, PER, etc.)",
			},
			"domain": map[string]interface{}{
				"type":        "string",
				"description": "Filter by domain (e.g., strategic, operational)",
			},
			"status": map[string]interface{}{
				"type":        "string",
				"description": "Filter by status (Draft, Accepted, Deprecated)",
				"enum":        []string{"Draft", "Accepted", "Deprecated"},
			},
		},
	}
}

func (t *BSpecQueryTool) Execute(params map[string]interface{}) (string, error) {
	// Build the command arguments
	args := []string{"query", "."}

	// Add output format
	if format, ok := params["output_format"].(string); ok && format != "" {
		args = append(args, "--output="+format)
	} else {
		args = append(args, "--output=markdown")
	}

	// Add type filter
	if docType, ok := params["document_type"].(string); ok && docType != "" {
		args = append(args, "--type="+docType)
	}

	// Add domain filter
	if domain, ok := params["domain"].(string); ok && domain != "" {
		args = append(args, "--domain="+domain)
	}

	// Add status filter
	if status, ok := params["status"].(string); ok && status != "" {
		args = append(args, "--status="+status)
	}

	// Execute the command
	cmd := exec.Command("bspec", args...)
	cmd.Dir = t.validator.WorkingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("bspec query failed: %w\nOutput: %s", err, string(output))
	}

	return string(output), nil
}

// BSpecPackTool packages the current BSpec project
type BSpecPackTool struct {
	validator *SecurityValidator
}

// NewBSpecPackTool creates a new BSpec pack tool
func NewBSpecPackTool(validator *SecurityValidator) *BSpecPackTool {
	return &BSpecPackTool{validator: validator}
}

func (t *BSpecPackTool) Name() string {
	return "bspec_pack"
}

func (t *BSpecPackTool) Description() string {
	return "Package the current BSpec project into a distributable .bspec file"
}

func (t *BSpecPackTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"output_file": map[string]interface{}{
				"type":        "string",
				"description": "Name of the output .bspec file (e.g., 'my-project.bspec')",
			},
			"verbose": map[string]interface{}{
				"type":        "boolean",
				"description": "Enable verbose output",
				"default":     false,
			},
		},
		"required": []string{"output_file"},
	}
}

func (t *BSpecPackTool) Execute(params map[string]interface{}) (string, error) {
	outputFile, ok := params["output_file"].(string)
	if !ok {
		return "", fmt.Errorf("output_file parameter is required and must be a string")
	}

	// Validate output file name
	if !strings.HasSuffix(outputFile, ".bspec") {
		outputFile += ".bspec"
	}

	// Build the command arguments
	args := []string{"pack", ".", outputFile}

	// Add verbose flag if requested
	if verbose, ok := params["verbose"].(bool); ok && verbose {
		args = append(args, "--verbose")
	}

	// Execute the command
	cmd := exec.Command("bspec", args...)
	cmd.Dir = t.validator.WorkingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("bspec pack failed: %w\nOutput: %s", err, string(output))
	}

	return string(output), nil
}

// BSpecValidateTool validates the current BSpec project
type BSpecValidateTool struct {
	validator *SecurityValidator
}

// NewBSpecValidateTool creates a new BSpec validation tool
func NewBSpecValidateTool(validator *SecurityValidator) *BSpecValidateTool {
	return &BSpecValidateTool{validator: validator}
}

func (t *BSpecValidateTool) Name() string {
	return "bspec_validate"
}

func (t *BSpecValidateTool) Description() string {
	return "Validate the current BSpec project structure and document integrity"
}

func (t *BSpecValidateTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type":       "object",
		"properties": map[string]interface{}{},
	}
}

func (t *BSpecValidateTool) Execute(params map[string]interface{}) (string, error) {
	// Execute validation by running a query and checking for errors
	cmd := exec.Command("bspec", "query", ".", "--output=json")
	cmd.Dir = t.validator.WorkingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Validation failed: %s\nOutput: %s", err.Error(), string(output)), nil
	}

	return "BSpec project validation passed successfully", nil
}

// BSpecInitTool initializes new BSpec directories
type BSpecInitTool struct {
	validator *SecurityValidator
}

// NewBSpecInitTool creates a new BSpec init tool
func NewBSpecInitTool(validator *SecurityValidator) *BSpecInitTool {
	return &BSpecInitTool{validator: validator}
}

func (t *BSpecInitTool) Name() string {
	return "bspec_init"
}

func (t *BSpecInitTool) Description() string {
	return "Initialize a new BSpec project structure in the current directory"
}

func (t *BSpecInitTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"project_name": map[string]interface{}{
				"type":        "string",
				"description": "Name of the BSpec project",
			},
		},
		"required": []string{"project_name"},
	}
}

func (t *BSpecInitTool) Execute(params map[string]interface{}) (string, error) {
	projectName, ok := params["project_name"].(string)
	if !ok {
		return "", fmt.Errorf("project_name parameter is required and must be a string")
	}

	// Execute the command
	cmd := exec.Command("bspec", "init", projectName)
	cmd.Dir = t.validator.WorkingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("bspec init failed: %w\nOutput: %s", err, string(output))
	}

	return string(output), nil
}

// BSpecVersionTool shows BSpec CLI version information
type BSpecVersionTool struct {
	validator *SecurityValidator
}

// NewBSpecVersionTool creates a new BSpec version tool
func NewBSpecVersionTool(validator *SecurityValidator) *BSpecVersionTool {
	return &BSpecVersionTool{validator: validator}
}

func (t *BSpecVersionTool) Name() string {
	return "bspec_version"
}

func (t *BSpecVersionTool) Description() string {
	return "Show BSpec CLI version information"
}

func (t *BSpecVersionTool) Parameters() map[string]interface{} {
	return map[string]interface{}{
		"type":       "object",
		"properties": map[string]interface{}{},
	}
}

func (t *BSpecVersionTool) Execute(params map[string]interface{}) (string, error) {
	// Execute the command
	cmd := exec.Command("bspec", "version")
	cmd.Dir = t.validator.WorkingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("bspec version failed: %w\nOutput: %s", err, string(output))
	}

	return string(output), nil
}