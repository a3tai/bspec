// Package bspec provides STR document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.607430
// Generator: go-generator v1.0.0

package bspec

// STRDocument represents a Strategy (STR) document
// Domain: strategic
//
// How the organization will achieve its vision and compete
type STRDocument struct {
	BaseBSpecDocument




}

// NewSTRDocument creates a new STR document with defaults
func NewSTRDocument(id, title, owner string) *STRDocument {
	doc := &STRDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeSTR,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainStrategic
	doc.Domain = &domain





	return doc
}

// Validate validates the STR document and returns any validation errors
func (d *STRDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeSTR {
		errors = append(errors, "document type must be STR")
	}

	// Document type specific validation




	return errors
}