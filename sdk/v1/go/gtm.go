// Package bspec provides GTM document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.617461
// Generator: go-generator v1.0.0

package bspec

// GTMDocument represents a Go-to-Market (GTM) document
// Domain: growth
//
// Launch and customer acquisition strategy
type GTMDocument struct {
	BaseBSpecDocument




}

// NewGTMDocument creates a new GTM document with defaults
func NewGTMDocument(id, title, owner string) *GTMDocument {
	doc := &GTMDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeGTM,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainGrowth
	doc.Domain = &domain





	return doc
}

// Validate validates the GTM document and returns any validation errors
func (d *GTMDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeGTM {
		errors = append(errors, "document type must be GTM")
	}

	// Document type specific validation




	return errors
}