// Package bspec provides WFL document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.613793
// Generator: go-generator v1.0.0

package bspec

// WFLDocument represents a Workflows (WFL) document
// Domain: operations
//
// Detailed task sequences within processes
type WFLDocument struct {
	BaseBSpecDocument




}

// NewWFLDocument creates a new WFL document with defaults
func NewWFLDocument(id, title, owner string) *WFLDocument {
	doc := &WFLDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeWFL,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainOperations
	doc.Domain = &domain





	return doc
}

// Validate validates the WFL document and returns any validation errors
func (d *WFLDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeWFL {
		errors = append(errors, "document type must be WFL")
	}

	// Document type specific validation




	return errors
}