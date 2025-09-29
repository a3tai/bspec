// Package bspec provides PRD document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.611012
// Generator: go-generator v1.0.0

package bspec

// PRDDocument represents a Products (PRD) document
// Domain: product
//
// Physical or digital products offered
type PRDDocument struct {
	BaseBSpecDocument




}

// NewPRDDocument creates a new PRD document with defaults
func NewPRDDocument(id, title, owner string) *PRDDocument {
	doc := &PRDDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypePRD,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainProduct
	doc.Domain = &domain





	return doc
}

// Validate validates the PRD document and returns any validation errors
func (d *PRDDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypePRD {
		errors = append(errors, "document type must be PRD")
	}

	// Document type specific validation




	return errors
}