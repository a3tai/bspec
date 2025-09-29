// Package bspec provides INV document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616150
// Generator: go-generator v1.0.0

package bspec

// INVDocument represents a Investment (INV) document
// Domain: financial
//
// Capital allocation and investment decisions
type INVDocument struct {
	BaseBSpecDocument




}

// NewINVDocument creates a new INV document with defaults
func NewINVDocument(id, title, owner string) *INVDocument {
	doc := &INVDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeINV,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the INV document and returns any validation errors
func (d *INVDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeINV {
		errors = append(errors, "document type must be INV")
	}

	// Document type specific validation




	return errors
}