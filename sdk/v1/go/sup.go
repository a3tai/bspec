// Package bspec provides SUP document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.612024
// Generator: go-generator v1.0.0

package bspec

// SUPDocument represents a Support (SUP) document
// Domain: product
//
// Customer support and success processes
type SUPDocument struct {
	BaseBSpecDocument




}

// NewSUPDocument creates a new SUP document with defaults
func NewSUPDocument(id, title, owner string) *SUPDocument {
	doc := &SUPDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeSUP,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainProduct
	doc.Domain = &domain





	return doc
}

// Validate validates the SUP document and returns any validation errors
func (d *SUPDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeSUP {
		errors = append(errors, "document type must be SUP")
	}

	// Document type specific validation




	return errors
}