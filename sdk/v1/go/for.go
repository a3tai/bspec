// Package bspec provides FOR document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.615949
// Generator: go-generator v1.0.0

package bspec

// FORDocument represents a Forecasts (FOR) document
// Domain: financial
//
// Forward-looking financial predictions and scenarios
type FORDocument struct {
	BaseBSpecDocument




}

// NewFORDocument creates a new FOR document with defaults
func NewFORDocument(id, title, owner string) *FORDocument {
	doc := &FORDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeFOR,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the FOR document and returns any validation errors
func (d *FORDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeFOR {
		errors = append(errors, "document type must be FOR")
	}

	// Document type specific validation




	return errors
}