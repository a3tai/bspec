// Package bspec provides CMP document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616924
// Generator: go-generator v1.0.0

package bspec

// CMPDocument represents a Compliance (CMP) document
// Domain: risk
//
// Regulatory obligations and adherence requirements
type CMPDocument struct {
	BaseBSpecDocument




}

// NewCMPDocument creates a new CMP document with defaults
func NewCMPDocument(id, title, owner string) *CMPDocument {
	doc := &CMPDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeCMP,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainRisk
	doc.Domain = &domain





	return doc
}

// Validate validates the CMP document and returns any validation errors
func (d *CMPDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeCMP {
		errors = append(errors, "document type must be CMP")
	}

	// Document type specific validation




	return errors
}