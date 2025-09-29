// Package bspec provides PUR document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.607860
// Generator: go-generator v1.0.0

package bspec

// PURDocument represents a Purpose (PUR) document
// Domain: strategic
//
// Social impact and stakeholder value beyond profit
type PURDocument struct {
	BaseBSpecDocument




}

// NewPURDocument creates a new PUR document with defaults
func NewPURDocument(id, title, owner string) *PURDocument {
	doc := &PURDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypePUR,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainStrategic
	doc.Domain = &domain





	return doc
}

// Validate validates the PUR document and returns any validation errors
func (d *PURDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypePUR {
		errors = append(errors, "document type must be PUR")
	}

	// Document type specific validation




	return errors
}