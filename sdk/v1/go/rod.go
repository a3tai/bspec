// Package bspec provides ROD document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.611361
// Generator: go-generator v1.0.0

package bspec

// RODDocument represents a Roadmap (ROD) document
// Domain: product
//
// Planned development and enhancement timeline
type RODDocument struct {
	BaseBSpecDocument




}

// NewRODDocument creates a new ROD document with defaults
func NewRODDocument(id, title, owner string) *RODDocument {
	doc := &RODDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeROD,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainProduct
	doc.Domain = &domain





	return doc
}

// Validate validates the ROD document and returns any validation errors
func (d *RODDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeROD {
		errors = append(errors, "document type must be ROD")
	}

	// Document type specific validation




	return errors
}