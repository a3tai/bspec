// Package bspec provides BUD document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.615862
// Generator: go-generator v1.0.0

package bspec

// BUDDocument represents a Budget (BUD) document
// Domain: financial
//
// Resource allocation and spending plans
type BUDDocument struct {
	BaseBSpecDocument




}

// NewBUDDocument creates a new BUD document with defaults
func NewBUDDocument(id, title, owner string) *BUDDocument {
	doc := &BUDDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeBUD,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the BUD document and returns any validation errors
func (d *BUDDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeBUD {
		errors = append(errors, "document type must be BUD")
	}

	// Document type specific validation




	return errors
}