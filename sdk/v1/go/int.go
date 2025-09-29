// Package bspec provides INT document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.611935
// Generator: go-generator v1.0.0

package bspec

// INTDocument represents a Integrations (INT) document
// Domain: product
//
// Connections with other systems and platforms
type INTDocument struct {
	BaseBSpecDocument




}

// NewINTDocument creates a new INT document with defaults
func NewINTDocument(id, title, owner string) *INTDocument {
	doc := &INTDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeINT,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainProduct
	doc.Domain = &domain





	return doc
}

// Validate validates the INT document and returns any validation errors
func (d *INTDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeINT {
		errors = append(errors, "document type must be INT")
	}

	// Document type specific validation




	return errors
}