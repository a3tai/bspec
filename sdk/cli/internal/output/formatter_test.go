package output

import (
	"testing"
)

func TestFormatterConstants(t *testing.T) {
	// Test that output format constants are defined correctly
	if FormatJSON != "json" {
		t.Errorf("Expected FormatJSON to be 'json', got '%s'", FormatJSON)
	}
	if FormatYAML != "yaml" {
		t.Errorf("Expected FormatYAML to be 'yaml', got '%s'", FormatYAML)
	}
	if FormatMarkdown != "markdown" {
		t.Errorf("Expected FormatMarkdown to be 'markdown', got '%s'", FormatMarkdown)
	}
}

func TestFormatterStruct(t *testing.T) {
	// Test the Formatter struct fields
	formatter := &Formatter{
		format: FormatJSON,
		pretty: true,
	}

	if formatter.format != FormatJSON {
		t.Errorf("Expected format %s, got %s", FormatJSON, formatter.format)
	}
	if !formatter.pretty {
		t.Error("Expected pretty to be true")
	}
}

func TestOutputFormatString(t *testing.T) {
	// Test that OutputFormat type works as expected
	var format OutputFormat = "json"
	if format != FormatJSON {
		t.Errorf("Expected format to equal FormatJSON")
	}

	format = "yaml"
	if format != FormatYAML {
		t.Errorf("Expected format to equal FormatYAML")
	}

	format = "markdown"
	if format != FormatMarkdown {
		t.Errorf("Expected format to equal FormatMarkdown")
	}
}