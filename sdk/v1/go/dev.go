// Package bspec provides DEV document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.615572
// Generator: go-generator v1.0.0

package bspec

// DEVDocument represents a Development (DEV) document
// Domain: technology
//
// Software development lifecycle and practices
type DEVDocument struct {
	BaseBSpecDocument




}

// NewDEVDocument creates a new DEV document with defaults
func NewDEVDocument(id, title, owner string) *DEVDocument {
	doc := &DEVDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeDEV,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainTechnology
	doc.Domain = &domain





	return doc
}

// Validate validates the DEV document and returns any validation errors
func (d *DEVDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeDEV {
		errors = append(errors, "document type must be DEV")
	}

	// Document type specific validation




	return errors
}