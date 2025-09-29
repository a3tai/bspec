// Package bspec provides TAX document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616631
// Generator: go-generator v1.0.0

package bspec

// TAXDocument represents a Tax Strategy (TAX) document
// Domain: financial
//
// Tax planning, structure, and compliance
type TAXDocument struct {
	BaseBSpecDocument




}

// NewTAXDocument creates a new TAX document with defaults
func NewTAXDocument(id, title, owner string) *TAXDocument {
	doc := &TAXDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeTAX,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the TAX document and returns any validation errors
func (d *TAXDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeTAX {
		errors = append(errors, "document type must be TAX")
	}

	// Document type specific validation




	return errors
}