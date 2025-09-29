// Package bspec provides API document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.615310
// Generator: go-generator v1.0.0

package bspec

// APIDocument represents a APIs (API) document
// Domain: technology
//
// Interface specifications, protocols, and integrations
type APIDocument struct {
	BaseBSpecDocument




}

// NewAPIDocument creates a new API document with defaults
func NewAPIDocument(id, title, owner string) *APIDocument {
	doc := &APIDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeAPI,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainTechnology
	doc.Domain = &domain





	return doc
}

// Validate validates the API document and returns any validation errors
func (d *APIDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeAPI {
		errors = append(errors, "document type must be API")
	}

	// Document type specific validation




	return errors
}