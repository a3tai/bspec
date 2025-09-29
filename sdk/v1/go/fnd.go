// Package bspec provides FND document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616045
// Generator: go-generator v1.0.0

package bspec

// FNDDocument represents a Funding (FND) document
// Domain: financial
//
// Capital requirements and financing strategy
type FNDDocument struct {
	BaseBSpecDocument




}

// NewFNDDocument creates a new FND document with defaults
func NewFNDDocument(id, title, owner string) *FNDDocument {
	doc := &FNDDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeFND,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the FND document and returns any validation errors
func (d *FNDDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeFND {
		errors = append(errors, "document type must be FND")
	}

	// Document type specific validation




	return errors
}