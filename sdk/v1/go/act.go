// Package bspec provides ACT document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.613122
// Generator: go-generator v1.0.0

package bspec

// ACTDocument represents a Key Activities (ACT) document
// Domain: business-model
//
// Essential processes for value creation
type ACTDocument struct {
	BaseBSpecDocument




}

// NewACTDocument creates a new ACT document with defaults
func NewACTDocument(id, title, owner string) *ACTDocument {
	doc := &ACTDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeACT,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainBusinessModel
	doc.Domain = &domain





	return doc
}

// Validate validates the ACT document and returns any validation errors
func (d *ACTDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeACT {
		errors = append(errors, "document type must be ACT")
	}

	// Document type specific validation




	return errors
}