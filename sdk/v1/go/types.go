// Types for BSpec v1.0.0
// Generated from BSpec JSON SDK at 2025-09-28T21:08:19.341307
package bspec

import "time"

// Metadata contains generation metadata
type Metadata struct {
	BSpecVersion string    `json:"bspec_version"`
	GeneratedAt  time.Time `json:"generated_at"`
	Generator    string    `json:"generator"`
	SourceSpec   string    `json:"source_spec"`
}

// Statistics contains specification statistics
type Statistics struct {
	TotalFiles         int `json:"total_files"`
	TotalDocumentTypes int `json:"total_document_types"`
	TotalDomains       int `json:"total_domains"`
}

// Domain represents a business domain
type Domain struct {
	Name          string   `json:"name"`
	DisplayName   string   `json:"display_name"`
	Description   string   `json:"description"`
	Emoji         string   `json:"emoji"`
	DocumentTypes []string `json:"document_types"`
	DocumentCount int      `json:"document_count"`
}

// DocumentType represents a BSpec document type
type DocumentType struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Purpose  string `json:"purpose"`
	Domain   string `json:"domain"`
	Examples []DocumentExample `json:"examples"`
}

// DocumentExample represents an example document
type DocumentExample struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// File represents a file in the specification
type File struct {
	Path           string                 `json:"path"`
	Type           string                 `json:"type"`
	Size           int                    `json:"size"`
	Content        string                 `json:"content"`
	HasFrontmatter bool                   `json:"has_frontmatter"`
	Frontmatter    map[string]interface{} `json:"frontmatter,omitempty"`
	ParsedContent  string                 `json:"parsed_content,omitempty"`
}

// ConformanceLevel represents a conformance level
type ConformanceLevel struct {
	Name         string `json:"name"`
	DisplayName  string `json:"display_name"`
	Description  string `json:"description"`
	Emoji        string `json:"emoji"`
	MinDocuments int    `json:"min_documents"`
}

// YAMLSchema represents the YAML schema definition
type YAMLSchema struct {
	RequiredFields []string               `json:"required_fields"`
	OptionalFields []string               `json:"optional_fields"`
	FieldTypes     map[string]string      `json:"field_types"`
	Enums          map[string][]string    `json:"enums"`
	Defaults       map[string]interface{} `json:"defaults"`
}

// DocumentIndexEntry represents an entry in the document index
type DocumentIndexEntry struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	Owner    string `json:"owner"`
	Status   string `json:"status"`
	Version  string `json:"version"`
}
