// Package bspec provides SKI document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.614191
// Generator: go-generator v1.0.0

package bspec

// SKIDocument represents a Skills (SKI) document
// Domain: operations
//
// Required capabilities and competency frameworks
type SKIDocument struct {
	BaseBSpecDocument




}

// NewSKIDocument creates a new SKI document with defaults
func NewSKIDocument(id, title, owner string) *SKIDocument {
	doc := &SKIDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeSKI,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainOperations
	doc.Domain = &domain





	return doc
}

// Validate validates the SKI document and returns any validation errors
func (d *SKIDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeSKI {
		errors = append(errors, "document type must be SKI")
	}

	// Document type specific validation




	return errors
}