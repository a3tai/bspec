// Package bspec provides JTB document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.609354
// Generator: go-generator v1.0.0

package bspec

// JTBDocument represents a Jobs-to-be-Done (JTB) document
// Domain: customer
//
// Specific outcomes customers hire products to achieve
type JTBDocument struct {
	BaseBSpecDocument


	// Customer understanding documents require specific validation
	CustomerFocused bool `json:"customer_focused" yaml:"customer_focused"`


}

// NewJTBDocument creates a new JTB document with defaults
func NewJTBDocument(id, title, owner string) *JTBDocument {
	doc := &JTBDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeJTB,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainCustomer
	doc.Domain = &domain


	doc.CustomerFocused = true



	return doc
}

// Validate validates the JTB document and returns any validation errors
func (d *JTBDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeJTB {
		errors = append(errors, "document type must be JTB")
	}

	// Document type specific validation

	// Customer understanding documents should have related documents
	if len(d.Related) == 0 {
		errors = append(errors, "Customer understanding documents should reference related personas or jobs-to-be-done")
	}



	return errors
}