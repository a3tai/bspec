// Package bspec provides USE document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.609546
// Generator: go-generator v1.0.0

package bspec

// USEDocument represents a Use Cases (USE) document
// Domain: customer
//
// Specific scenarios where customers apply the solution
type USEDocument struct {
	BaseBSpecDocument




}

// NewUSEDocument creates a new USE document with defaults
func NewUSEDocument(id, title, owner string) *USEDocument {
	doc := &USEDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeUSE,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainCustomer
	doc.Domain = &domain





	return doc
}

// Validate validates the USE document and returns any validation errors
func (d *USEDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeUSE {
		errors = append(errors, "document type must be USE")
	}

	// Document type specific validation




	return errors
}