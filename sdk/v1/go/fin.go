// Package bspec provides FIN document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.615766
// Generator: go-generator v1.0.0

package bspec

// FINDocument represents a Financial Model (FIN) document
// Domain: financial
//
// Comprehensive P&L, balance sheet, cash flow projections
type FINDocument struct {
	BaseBSpecDocument




}

// NewFINDocument creates a new FIN document with defaults
func NewFINDocument(id, title, owner string) *FINDocument {
	doc := &FINDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeFIN,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the FIN document and returns any validation errors
func (d *FINDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeFIN {
		errors = append(errors, "document type must be FIN")
	}

	// Document type specific validation




	return errors
}