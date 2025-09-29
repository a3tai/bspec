package output

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/a3tai/bspec/cli/internal/archive"
	"github.com/a3tai/bspec/cli/internal/query"
)

// OutputFormat represents the supported output formats
type OutputFormat string

const (
	FormatJSON     OutputFormat = "json"
	FormatYAML     OutputFormat = "yaml"
	FormatMarkdown OutputFormat = "markdown"
)

// Formatter handles formatting output in different formats
type Formatter struct {
	format OutputFormat
	pretty bool
}

// NewFormatter creates a new formatter
func NewFormatter(format string, pretty bool) (*Formatter, error) {
	var outputFormat OutputFormat

	switch strings.ToLower(format) {
	case "json":
		outputFormat = FormatJSON
	case "yaml", "yml":
		outputFormat = FormatYAML
	case "markdown", "md":
		outputFormat = FormatMarkdown
	default:
		return nil, fmt.Errorf("unsupported output format: %s", format)
	}

	return &Formatter{
		format: outputFormat,
		pretty: pretty,
	}, nil
}

// FormatQueryResult formats a query result
func (f *Formatter) FormatQueryResult(result *query.QueryResult) (string, error) {
	switch f.format {
	case FormatJSON:
		return f.formatQueryResultJSON(result)
	case FormatYAML:
		return f.formatQueryResultYAML(result)
	case FormatMarkdown:
		return f.formatQueryResultMarkdown(result)
	default:
		return "", fmt.Errorf("unsupported format: %s", f.format)
	}
}

// FormatDocument formats a single document
func (f *Formatter) FormatDocument(doc archive.BSpecDocument) (string, error) {
	switch f.format {
	case FormatJSON:
		return f.formatDocumentJSON(doc)
	case FormatYAML:
		return f.formatDocumentYAML(doc)
	case FormatMarkdown:
		return f.formatDocumentMarkdown(doc)
	default:
		return "", fmt.Errorf("unsupported format: %s", f.format)
	}
}

// FormatArchiveInfo formats archive information
func (f *Formatter) FormatArchiveInfo(arch *archive.BSpecArchive) (string, error) {
	switch f.format {
	case FormatJSON:
		return f.formatArchiveInfoJSON(arch)
	case FormatYAML:
		return f.formatArchiveInfoYAML(arch)
	case FormatMarkdown:
		return f.formatArchiveInfoMarkdown(arch)
	default:
		return "", fmt.Errorf("unsupported format: %s", f.format)
	}
}

// FormatStats formats statistics
func (f *Formatter) FormatStats(stats map[string]interface{}) (string, error) {
	switch f.format {
	case FormatJSON:
		return f.formatStatsJSON(stats)
	case FormatYAML:
		return f.formatStatsYAML(stats)
	case FormatMarkdown:
		return f.formatStatsMarkdown(stats)
	default:
		return "", fmt.Errorf("unsupported format: %s", f.format)
	}
}

// JSON formatters
func (f *Formatter) formatQueryResultJSON(result *query.QueryResult) (string, error) {
	if f.pretty {
		data, err := json.MarshalIndent(result, "", "  ")
		return string(data), err
	}
	data, err := json.Marshal(result)
	return string(data), err
}

func (f *Formatter) formatDocumentJSON(doc archive.BSpecDocument) (string, error) {
	if f.pretty {
		data, err := json.MarshalIndent(doc, "", "  ")
		return string(data), err
	}
	data, err := json.Marshal(doc)
	return string(data), err
}

func (f *Formatter) formatArchiveInfoJSON(arch *archive.BSpecArchive) (string, error) {
	info := map[string]interface{}{
		"manifest":       arch.Manifest,
		"document_count": len(arch.Documents),
		"asset_count":    len(arch.Assets),
		"computed_count": len(arch.Computed),
	}

	if f.pretty {
		data, err := json.MarshalIndent(info, "", "  ")
		return string(data), err
	}
	data, err := json.Marshal(info)
	return string(data), err
}

func (f *Formatter) formatStatsJSON(stats map[string]interface{}) (string, error) {
	if f.pretty {
		data, err := json.MarshalIndent(stats, "", "  ")
		return string(data), err
	}
	data, err := json.Marshal(stats)
	return string(data), err
}

// YAML formatters
func (f *Formatter) formatQueryResultYAML(result *query.QueryResult) (string, error) {
	data, err := yaml.Marshal(result)
	return string(data), err
}

func (f *Formatter) formatDocumentYAML(doc archive.BSpecDocument) (string, error) {
	data, err := yaml.Marshal(doc)
	return string(data), err
}

func (f *Formatter) formatArchiveInfoYAML(arch *archive.BSpecArchive) (string, error) {
	info := map[string]interface{}{
		"manifest":       arch.Manifest,
		"document_count": len(arch.Documents),
		"asset_count":    len(arch.Assets),
		"computed_count": len(arch.Computed),
	}

	data, err := yaml.Marshal(info)
	return string(data), err
}

func (f *Formatter) formatStatsYAML(stats map[string]interface{}) (string, error) {
	data, err := yaml.Marshal(stats)
	return string(data), err
}

// Markdown formatters
func (f *Formatter) formatQueryResultMarkdown(result *query.QueryResult) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Query Results\n\n"))
	sb.WriteString(fmt.Sprintf("**Total Documents:** %d\n\n", result.Total))

	if len(result.Documents) == 0 {
		sb.WriteString("No documents found matching the query criteria.\n")
		return sb.String(), nil
	}

	sb.WriteString("## Documents\n\n")

	for i, doc := range result.Documents {
		sb.WriteString(fmt.Sprintf("### %d. %s\n\n", i+1, doc.Title))
		sb.WriteString(fmt.Sprintf("- **ID:** %s\n", doc.ID))
		sb.WriteString(fmt.Sprintf("- **Type:** %s\n", doc.Type))
		sb.WriteString(fmt.Sprintf("- **Status:** %s\n", doc.Status))
		sb.WriteString(fmt.Sprintf("- **Version:** %s\n", doc.Version))
		sb.WriteString(fmt.Sprintf("- **Owner:** %s\n", doc.Owner))
		if doc.Domain != "" {
			sb.WriteString(fmt.Sprintf("- **Domain:** %s\n", doc.Domain))
		}
		sb.WriteString(fmt.Sprintf("- **Created:** %s\n", doc.Created))
		sb.WriteString(fmt.Sprintf("- **Updated:** %s\n", doc.Updated))

		if doc.Content != "" {
			sb.WriteString("\n**Content:**\n\n")
			// Truncate content for summary view
			content := doc.Content
			if len(content) > 200 {
				content = content[:200] + "..."
			}
			sb.WriteString(content)
		}

		sb.WriteString("\n\n---\n\n")
	}

	return sb.String(), nil
}

func (f *Formatter) formatDocumentMarkdown(doc archive.BSpecDocument) (string, error) {
	var sb strings.Builder

	// Write frontmatter
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("id: %s\n", doc.ID))
	sb.WriteString(fmt.Sprintf("title: %s\n", doc.Title))
	sb.WriteString(fmt.Sprintf("type: %s\n", doc.Type))
	sb.WriteString(fmt.Sprintf("status: %s\n", doc.Status))
	sb.WriteString(fmt.Sprintf("version: %s\n", doc.Version))
	sb.WriteString(fmt.Sprintf("owner: %s\n", doc.Owner))
	if doc.Domain != "" {
		sb.WriteString(fmt.Sprintf("domain: %s\n", doc.Domain))
	}
	sb.WriteString(fmt.Sprintf("created: %s\n", doc.Created))
	sb.WriteString(fmt.Sprintf("updated: %s\n", doc.Updated))

	// Add metadata if present
	if len(doc.Metadata) > 0 {
		for key, value := range doc.Metadata {
			sb.WriteString(fmt.Sprintf("%s: %v\n", key, value))
		}
	}

	sb.WriteString("---\n\n")

	// Write content
	if doc.Content != "" {
		sb.WriteString(doc.Content)
	}

	return sb.String(), nil
}

func (f *Formatter) formatArchiveInfoMarkdown(arch *archive.BSpecArchive) (string, error) {
	var sb strings.Builder

	sb.WriteString("# BSpec Archive Information\n\n")

	// Manifest information
	sb.WriteString("## Manifest\n\n")
	sb.WriteString(fmt.Sprintf("- **Name:** %s\n", arch.Manifest.Name))
	sb.WriteString(fmt.Sprintf("- **Description:** %s\n", arch.Manifest.Description))
	sb.WriteString(fmt.Sprintf("- **Author:** %s\n", arch.Manifest.Author))
	sb.WriteString(fmt.Sprintf("- **BSpec Version:** %s\n", arch.Manifest.BSpecVersion))
	sb.WriteString(fmt.Sprintf("- **Format Version:** %s\n", arch.Manifest.FormatVersion))
	sb.WriteString(fmt.Sprintf("- **Created:** %s\n", arch.Manifest.CreatedAt.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("- **Updated:** %s\n", arch.Manifest.UpdatedAt.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("- **Conformance Level:** %s\n", arch.Manifest.ConformanceLevel))
	sb.WriteString(fmt.Sprintf("- **Industry Profile:** %s\n", arch.Manifest.IndustryProfile))

	// Content statistics
	sb.WriteString("\n## Content Statistics\n\n")
	sb.WriteString(fmt.Sprintf("- **Documents:** %d\n", len(arch.Documents)))
	sb.WriteString(fmt.Sprintf("- **Assets:** %d\n", len(arch.Assets)))
	sb.WriteString(fmt.Sprintf("- **Computed Files:** %d\n", len(arch.Computed)))

	// Document types
	if len(arch.Manifest.DocumentTypes) > 0 {
		sb.WriteString("\n## Document Types\n\n")
		sort.Strings(arch.Manifest.DocumentTypes)
		for _, docType := range arch.Manifest.DocumentTypes {
			sb.WriteString(fmt.Sprintf("- %s\n", docType))
		}
	}

	// Domains
	if len(arch.Manifest.Domains) > 0 {
		sb.WriteString("\n## Domains\n\n")
		sort.Strings(arch.Manifest.Domains)
		for _, domain := range arch.Manifest.Domains {
			sb.WriteString(fmt.Sprintf("- %s\n", domain))
		}
	}

	return sb.String(), nil
}

func (f *Formatter) formatStatsMarkdown(stats map[string]interface{}) (string, error) {
	var sb strings.Builder

	sb.WriteString("# Archive Statistics\n\n")

	// Basic counts
	if totalDocs, ok := stats["total_documents"].(int); ok {
		sb.WriteString(fmt.Sprintf("**Total Documents:** %d\n", totalDocs))
	}
	if totalAssets, ok := stats["total_assets"].(int); ok {
		sb.WriteString(fmt.Sprintf("**Total Assets:** %d\n\n", totalAssets))
	}

	// Status distribution
	if statusDist, ok := stats["status_distribution"].(map[string]int); ok && len(statusDist) > 0 {
		sb.WriteString("## Status Distribution\n\n")
		for status, count := range statusDist {
			sb.WriteString(fmt.Sprintf("- **%s:** %d\n", status, count))
		}
		sb.WriteString("\n")
	}

	// Type distribution
	if typeDist, ok := stats["type_distribution"].(map[string]int); ok && len(typeDist) > 0 {
		sb.WriteString("## Document Type Distribution\n\n")
		for docType, count := range typeDist {
			sb.WriteString(fmt.Sprintf("- **%s:** %d\n", docType, count))
		}
		sb.WriteString("\n")
	}

	// Document types
	if docTypes, ok := stats["document_types"].([]string); ok && len(docTypes) > 0 {
		sb.WriteString("## Available Document Types\n\n")
		sort.Strings(docTypes)
		for _, docType := range docTypes {
			sb.WriteString(fmt.Sprintf("- %s\n", docType))
		}
		sb.WriteString("\n")
	}

	// Domains
	if domains, ok := stats["domains"].([]string); ok && len(domains) > 0 {
		sb.WriteString("## Available Domains\n\n")
		sort.Strings(domains)
		for _, domain := range domains {
			sb.WriteString(fmt.Sprintf("- %s\n", domain))
		}
		sb.WriteString("\n")
	}

	// Owners
	if owners, ok := stats["owners"].([]string); ok && len(owners) > 0 {
		sb.WriteString("## Document Owners\n\n")
		sort.Strings(owners)
		for _, owner := range owners {
			sb.WriteString(fmt.Sprintf("- %s\n", owner))
		}
	}

	return sb.String(), nil
}