// Package bspec provides CST document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.612528
// Generator: go-generator v1.0.0

package bspec

// CSTDocument represents a Cost Structure (CST) document
// Domain: business-model
//
// Major cost categories and cost drivers
type CSTDocument struct {
	BaseBSpecDocument



	// Business model documents are critical for financial validation
	BusinessModelCore bool `json:"business_model_core" yaml:"business_model_core"`

}

// NewCSTDocument creates a new CST document with defaults
func NewCSTDocument(id, title, owner string) *CSTDocument {
	doc := &CSTDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeCST,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainBusinessModel
	doc.Domain = &domain



	doc.BusinessModelCore = true


	return doc
}

// Validate validates the CST document and returns any validation errors
func (d *CSTDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeCST {
		errors = append(errors, "document type must be CST")
	}

	// Document type specific validation


	// Business model documents require metrics
	if len(d.Metrics) == 0 {
		errors = append(errors, "Business model documents must have metrics defined for measurement")
	}


	return errors
}