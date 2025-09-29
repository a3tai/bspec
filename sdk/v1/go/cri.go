// Package bspec provides CRI document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.617172
// Generator: go-generator v1.0.0

package bspec

// CRIDocument represents a Crisis Management (CRI) document
// Domain: risk
//
// Crisis response and business continuity plans
type CRIDocument struct {
	BaseBSpecDocument




}

// NewCRIDocument creates a new CRI document with defaults
func NewCRIDocument(id, title, owner string) *CRIDocument {
	doc := &CRIDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeCRI,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainRisk
	doc.Domain = &domain





	return doc
}

// Validate validates the CRI document and returns any validation errors
func (d *CRIDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeCRI {
		errors = append(errors, "document type must be CRI")
	}

	// Document type specific validation




	return errors
}