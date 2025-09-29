// Package bspec provides PER document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.611841
// Generator: go-generator v1.0.0

package bspec

// PERDocument represents a Performance (PER) document
// Domain: product
//
// Speed, reliability, and performance characteristics
type PERDocument struct {
	BaseBSpecDocument


	// Customer understanding documents require specific validation
	CustomerFocused bool `json:"customer_focused" yaml:"customer_focused"`


}

// NewPERDocument creates a new PER document with defaults
func NewPERDocument(id, title, owner string) *PERDocument {
	doc := &PERDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypePER,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainProduct
	doc.Domain = &domain


	doc.CustomerFocused = true



	return doc
}

// Validate validates the PER document and returns any validation errors
func (d *PERDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypePER {
		errors = append(errors, "document type must be PER")
	}

	// Document type specific validation

	// Customer understanding documents should have related documents
	if len(d.Related) == 0 {
		errors = append(errors, "Customer understanding documents should reference related personas or jobs-to-be-done")
	}



	return errors
}