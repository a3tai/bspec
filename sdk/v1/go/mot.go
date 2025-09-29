// Package bspec provides MOT document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.607638
// Generator: go-generator v1.0.0

package bspec

// MOTDocument represents a Moats (MOT) document
// Domain: strategic
//
// Competitive advantages that protect market position
type MOTDocument struct {
	BaseBSpecDocument




}

// NewMOTDocument creates a new MOT document with defaults
func NewMOTDocument(id, title, owner string) *MOTDocument {
	doc := &MOTDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeMOT,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainStrategic
	doc.Domain = &domain





	return doc
}

// Validate validates the MOT document and returns any validation errors
func (d *MOTDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeMOT {
		errors = append(errors, "document type must be MOT")
	}

	// Document type specific validation




	return errors
}