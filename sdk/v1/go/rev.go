// Package bspec provides REV document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.612239
// Generator: go-generator v1.0.0

package bspec

// REVDocument represents a Revenue Streams (REV) document
// Domain: business-model
//
// How money is generated from customers
type REVDocument struct {
	BaseBSpecDocument



	// Business model documents are critical for financial validation
	BusinessModelCore bool `json:"business_model_core" yaml:"business_model_core"`

}

// NewREVDocument creates a new REV document with defaults
func NewREVDocument(id, title, owner string) *REVDocument {
	doc := &REVDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeREV,
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

// Validate validates the REV document and returns any validation errors
func (d *REVDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeREV {
		errors = append(errors, "document type must be REV")
	}

	// Document type specific validation


	// Business model documents require metrics
	if len(d.Metrics) == 0 {
		errors = append(errors, "Business model documents must have metrics defined for measurement")
	}


	return errors
}