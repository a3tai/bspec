// Package bspec provides SUR document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.610813
// Generator: go-generator v1.0.0

package bspec

// SURDocument represents a Surveys (SUR) document
// Domain: customer
//
// Quantitative customer research and data
type SURDocument struct {
	BaseBSpecDocument




}

// NewSURDocument creates a new SUR document with defaults
func NewSURDocument(id, title, owner string) *SURDocument {
	doc := &SURDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeSUR,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainCustomer
	doc.Domain = &domain





	return doc
}

// Validate validates the SUR document and returns any validation errors
func (d *SURDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeSUR {
		errors = append(errors, "document type must be SUR")
	}

	// Document type specific validation




	return errors
}