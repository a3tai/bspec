// Package bspec provides PAI document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.609899
// Generator: go-generator v1.0.0

package bspec

// PAIDocument represents a Pain Points (PAI) document
// Domain: customer
//
// Customer problems and frustrations being solved
type PAIDocument struct {
	BaseBSpecDocument




}

// NewPAIDocument creates a new PAI document with defaults
func NewPAIDocument(id, title, owner string) *PAIDocument {
	doc := &PAIDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypePAI,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainCustomer
	doc.Domain = &domain





	return doc
}

// Validate validates the PAI document and returns any validation errors
func (d *PAIDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypePAI {
		errors = append(errors, "document type must be PAI")
	}

	// Document type specific validation




	return errors
}