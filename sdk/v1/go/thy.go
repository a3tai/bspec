// Package bspec provides THY document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.608023
// Generator: go-generator v1.0.0

package bspec

// THYDocument represents a Theory of Change (THY) document
// Domain: strategic
//
// Logic model connecting activities to outcomes
type THYDocument struct {
	BaseBSpecDocument




}

// NewTHYDocument creates a new THY document with defaults
func NewTHYDocument(id, title, owner string) *THYDocument {
	doc := &THYDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeTHY,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainStrategic
	doc.Domain = &domain





	return doc
}

// Validate validates the THY document and returns any validation errors
func (d *THYDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeTHY {
		errors = append(errors, "document type must be THY")
	}

	// Document type specific validation




	return errors
}