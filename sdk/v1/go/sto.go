// Package bspec provides STO document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.609737
// Generator: go-generator v1.0.0

package bspec

// STODocument represents a User Stories (STO) document
// Domain: customer
//
// Individual requirements from user perspective
type STODocument struct {
	BaseBSpecDocument




}

// NewSTODocument creates a new STO document with defaults
func NewSTODocument(id, title, owner string) *STODocument {
	doc := &STODocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeSTO,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainCustomer
	doc.Domain = &domain





	return doc
}

// Validate validates the STO document and returns any validation errors
func (d *STODocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeSTO {
		errors = append(errors, "document type must be STO")
	}

	// Document type specific validation




	return errors
}