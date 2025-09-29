// Package bspec provides EXP document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.618139
// Generator: go-generator v1.0.0

package bspec

// EXPDocument represents a Expansion (EXP) document
// Domain: growth
//
// Geographic, market, or product expansion plans
type EXPDocument struct {
	BaseBSpecDocument




}

// NewEXPDocument creates a new EXP document with defaults
func NewEXPDocument(id, title, owner string) *EXPDocument {
	doc := &EXPDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeEXP,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainGrowth
	doc.Domain = &domain





	return doc
}

// Validate validates the EXP document and returns any validation errors
func (d *EXPDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeEXP {
		errors = append(errors, "document type must be EXP")
	}

	// Document type specific validation




	return errors
}