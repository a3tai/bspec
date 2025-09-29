// Package bspec provides FEA document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.611246
// Generator: go-generator v1.0.0

package bspec

// FEADocument represents a Features (FEA) document
// Domain: product
//
// Specific capabilities of products and services
type FEADocument struct {
	BaseBSpecDocument




}

// NewFEADocument creates a new FEA document with defaults
func NewFEADocument(id, title, owner string) *FEADocument {
	doc := &FEADocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeFEA,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainProduct
	doc.Domain = &domain





	return doc
}

// Validate validates the FEA document and returns any validation errors
func (d *FEADocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeFEA {
		errors = append(errors, "document type must be FEA")
	}

	// Document type specific validation




	return errors
}