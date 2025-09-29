// Package bspec provides UNT document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.613319
// Generator: go-generator v1.0.0

package bspec

// UNTDocument represents a Unit Economics (UNT) document
// Domain: business-model
//
// Per-customer or per-unit financial metrics
type UNTDocument struct {
	BaseBSpecDocument




}

// NewUNTDocument creates a new UNT document with defaults
func NewUNTDocument(id, title, owner string) *UNTDocument {
	doc := &UNTDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeUNT,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainBusinessModel
	doc.Domain = &domain





	return doc
}

// Validate validates the UNT document and returns any validation errors
func (d *UNTDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeUNT {
		errors = append(errors, "document type must be UNT")
	}

	// Document type specific validation




	return errors
}