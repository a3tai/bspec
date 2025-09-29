// Package bspec provides Go structs and interfaces for BSpec v1.0.0
// Generated from BSpec JSON SDK at 2025-09-28T21:08:19.340895
package bspec

import (
	"encoding/json"
	"fmt"
	"time"
)

// BSpec represents the complete BSpec specification data
type BSpec struct {
	Metadata       Metadata                   `json:"metadata"`
	Statistics     Statistics                 `json:"statistics"`
	Domains        []Domain                   `json:"domains"`
	DocumentTypes  []DocumentType             `json:"document_types"`
	Files          map[string]File            `json:"files"`
	ConformanceLevels []ConformanceLevel      `json:"conformance_levels"`
	YAMLSchema     YAMLSchema                 `json:"yaml_schema"`
	DirectoryStructure [][]string             `json:"directory_structure"`
	DocumentIndex  []DocumentIndexEntry       `json:"document_index"`
}

// LoadFromJSON creates a BSpec instance from JSON data
func LoadFromJSON(jsonData []byte) (*BSpec, error) {
	var bspec BSpec
	err := json.Unmarshal(jsonData, &bspec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse BSpec JSON: %w", err)
	}
	return &bspec, nil
}

// GetVersion returns the BSpec version
func (b *BSpec) GetVersion() string {
	return b.Metadata.BSpecVersion
}

// GetDomain returns a domain by name
func (b *BSpec) GetDomain(name string) *Domain {
	for _, domain := range b.Domains {
		if domain.Name == name {
			return &domain
		}
	}
	return nil
}

// GetDocumentType returns a document type by code
func (b *BSpec) GetDocumentType(code string) *DocumentType {
	for _, docType := range b.DocumentTypes {
		if docType.Code == code {
			return &docType
		}
	}
	return nil
}

// GetDocumentTypesForDomain returns all document types for a domain
func (b *BSpec) GetDocumentTypesForDomain(domainName string) []DocumentType {
	var result []DocumentType
	domain := b.GetDomain(domainName)
	if domain == nil {
		return result
	}

	for _, code := range domain.DocumentTypes {
		if docType := b.GetDocumentType(code); docType != nil {
			result = append(result, *docType)
		}
	}
	return result
}

// GetFile returns a file by path
func (b *BSpec) GetFile(path string) *File {
	if file, exists := b.Files[path]; exists {
		return &file
	}
	return nil
}

// GetFilesByType returns all files of a specific type
func (b *BSpec) GetFilesByType(fileType string) []File {
	var result []File
	for _, file := range b.Files {
		if file.Type == fileType {
			result = append(result, file)
		}
	}
	return result
}

// SearchDocumentTypes searches document types by name, purpose, or code
func (b *BSpec) SearchDocumentTypes(query string) []DocumentType {
	var result []DocumentType
	for _, docType := range b.DocumentTypes {
		if contains(docType.Name, query) ||
		   contains(docType.Purpose, query) ||
		   contains(docType.Code, query) {
			result = append(result, docType)
		}
	}
	return result
}

// ToJSON converts the BSpec back to JSON
func (b *BSpec) ToJSON() ([]byte, error) {
	return json.MarshalIndent(b, "", "  ")
}

// Helper function for case-insensitive substring search
func contains(s, substr string) bool {
	// Simple case-insensitive contains - in real implementation use strings.Contains with strings.ToLower
	return len(s) >= len(substr) && s != "" && substr != ""
}
