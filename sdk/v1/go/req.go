// Package bspec provides REQ document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.611506
// Generator: go-generator v1.0.0

package bspec

// REQDocument represents a Requirements (REQ) document
// Domain: product
//
// Functional and non-functional specifications
type REQDocument struct {
	BaseBSpecDocument




}

// NewREQDocument creates a new REQ document with defaults
func NewREQDocument(id, title, owner string) *REQDocument {
	doc := &REQDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeREQ,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainProduct
	doc.Domain = &domain





	return doc
}

// Validate validates the REQ document and returns any validation errors
func (d *REQDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeREQ {
		errors = append(errors, "document type must be REQ")
	}

	// Document type specific validation




	return errors
}