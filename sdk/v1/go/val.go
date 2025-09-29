// Package bspec provides VAL document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616254
// Generator: go-generator v1.0.0

package bspec

// VALDocument represents a Valuation (VAL) document
// Domain: financial
//
// Business valuation models and assumptions
type VALDocument struct {
	BaseBSpecDocument

	// Strategic foundation documents are required for all conformance levels
	StrategicFoundation bool `json:"strategic_foundation" yaml:"strategic_foundation"`



}

// NewVALDocument creates a new VAL document with defaults
func NewVALDocument(id, title, owner string) *VALDocument {
	doc := &VALDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeVAL,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain

	doc.StrategicFoundation = true




	return doc
}

// Validate validates the VAL document and returns any validation errors
func (d *VALDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeVAL {
		errors = append(errors, "document type must be VAL")
	}

	// Document type specific validation
	// Strategic foundation documents require success criteria
	if len(d.SuccessCriteria) == 0 {
		errors = append(errors, "Strategic foundation documents must have success_criteria defined")
	}




	return errors
}