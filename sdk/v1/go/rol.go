// Package bspec provides ROL document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.613993
// Generator: go-generator v1.0.0

package bspec

// ROLDocument represents a Roles (ROL) document
// Domain: operations
//
// Individual position definitions and responsibilities
type ROLDocument struct {
	BaseBSpecDocument




}

// NewROLDocument creates a new ROL document with defaults
func NewROLDocument(id, title, owner string) *ROLDocument {
	doc := &ROLDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeROL,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainOperations
	doc.Domain = &domain





	return doc
}

// Validate validates the ROL document and returns any validation errors
func (d *ROLDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeROL {
		errors = append(errors, "document type must be ROL")
	}

	// Document type specific validation




	return errors
}