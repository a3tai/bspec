// Package bspec provides AUD document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616534
// Generator: go-generator v1.0.0

package bspec

// AUDDocument represents a Audit (AUD) document
// Domain: financial
//
// Financial controls, compliance, and audit processes
type AUDDocument struct {
	BaseBSpecDocument




}

// NewAUDDocument creates a new AUD document with defaults
func NewAUDDocument(id, title, owner string) *AUDDocument {
	doc := &AUDDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeAUD,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the AUD document and returns any validation errors
func (d *AUDDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeAUD {
		errors = append(errors, "document type must be AUD")
	}

	// Document type specific validation




	return errors
}