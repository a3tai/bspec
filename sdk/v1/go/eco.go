// Package bspec provides ECO document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.608735
// Generator: go-generator v1.0.0

package bspec

// ECODocument represents a Ecosystem (ECO) document
// Domain: market
//
// Partners, suppliers, distributors, and ecosystem players
type ECODocument struct {
	BaseBSpecDocument




}

// NewECODocument creates a new ECO document with defaults
func NewECODocument(id, title, owner string) *ECODocument {
	doc := &ECODocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeECO,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainMarket
	doc.Domain = &domain





	return doc
}

// Validate validates the ECO document and returns any validation errors
func (d *ECODocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeECO {
		errors = append(errors, "document type must be ECO")
	}

	// Document type specific validation




	return errors
}