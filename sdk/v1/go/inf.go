// Package bspec provides INF document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.615407
// Generator: go-generator v1.0.0

package bspec

// INFDocument represents a Infrastructure (INF) document
// Domain: technology
//
// Deployment, hosting, and runtime environment
type INFDocument struct {
	BaseBSpecDocument




}

// NewINFDocument creates a new INF document with defaults
func NewINFDocument(id, title, owner string) *INFDocument {
	doc := &INFDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeINF,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainTechnology
	doc.Domain = &domain





	return doc
}

// Validate validates the INF document and returns any validation errors
func (d *INFDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeINF {
		errors = append(errors, "document type must be INF")
	}

	// Document type specific validation




	return errors
}