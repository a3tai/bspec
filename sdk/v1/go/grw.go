// Package bspec provides GRW document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.617551
// Generator: go-generator v1.0.0

package bspec

// GRWDocument represents a Growth Model (GRW) document
// Domain: growth
//
// How the business scales customers and revenue
type GRWDocument struct {
	BaseBSpecDocument




}

// NewGRWDocument creates a new GRW document with defaults
func NewGRWDocument(id, title, owner string) *GRWDocument {
	doc := &GRWDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeGRW,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainGrowth
	doc.Domain = &domain





	return doc
}

// Validate validates the GRW document and returns any validation errors
func (d *GRWDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeGRW {
		errors = append(errors, "document type must be GRW")
	}

	// Document type specific validation




	return errors
}