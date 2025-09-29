// Package bspec provides GAI document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.610007
// Generator: go-generator v1.0.0

package bspec

// GAIDocument represents a Gains (GAI) document
// Domain: customer
//
// Benefits and value customers receive
type GAIDocument struct {
	BaseBSpecDocument




}

// NewGAIDocument creates a new GAI document with defaults
func NewGAIDocument(id, title, owner string) *GAIDocument {
	doc := &GAIDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeGAI,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainCustomer
	doc.Domain = &domain





	return doc
}

// Validate validates the GAI document and returns any validation errors
func (d *GAIDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeGAI {
		errors = append(errors, "document type must be GAI")
	}

	// Document type specific validation




	return errors
}