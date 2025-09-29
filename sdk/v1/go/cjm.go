// Package bspec provides CJM document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.609444
// Generator: go-generator v1.0.0

package bspec

// CJMDocument represents a Customer Journey (CJM) document
// Domain: customer
//
// End-to-end experience from awareness to advocacy
type CJMDocument struct {
	BaseBSpecDocument




}

// NewCJMDocument creates a new CJM document with defaults
func NewCJMDocument(id, title, owner string) *CJMDocument {
	doc := &CJMDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeCJM,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainCustomer
	doc.Domain = &domain





	return doc
}

// Validate validates the CJM document and returns any validation errors
func (d *CJMDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeCJM {
		errors = append(errors, "document type must be CJM")
	}

	// Document type specific validation




	return errors
}