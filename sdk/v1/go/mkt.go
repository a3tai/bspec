// Package bspec provides MKT document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.608140
// Generator: go-generator v1.0.0

package bspec

// MKTDocument represents a Market Definition (MKT) document
// Domain: market
//
// TAM/SAM/SOM, market boundaries and sizing
type MKTDocument struct {
	BaseBSpecDocument




}

// NewMKTDocument creates a new MKT document with defaults
func NewMKTDocument(id, title, owner string) *MKTDocument {
	doc := &MKTDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeMKT,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainMarket
	doc.Domain = &domain





	return doc
}

// Validate validates the MKT document and returns any validation errors
func (d *MKTDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeMKT {
		errors = append(errors, "document type must be MKT")
	}

	// Document type specific validation




	return errors
}