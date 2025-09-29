// Package bspec provides BMC document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.612144
// Generator: go-generator v1.0.0

package bspec

// BMCDocument represents a Business Model Canvas (BMC) document
// Domain: business-model
//
// Complete business model overview
type BMCDocument struct {
	BaseBSpecDocument




}

// NewBMCDocument creates a new BMC document with defaults
func NewBMCDocument(id, title, owner string) *BMCDocument {
	doc := &BMCDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeBMC,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainBusinessModel
	doc.Domain = &domain





	return doc
}

// Validate validates the BMC document and returns any validation errors
func (d *BMCDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeBMC {
		errors = append(errors, "document type must be BMC")
	}

	// Document type specific validation




	return errors
}