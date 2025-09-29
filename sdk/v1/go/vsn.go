// Package bspec provides VSN document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.607213
// Generator: go-generator v1.0.0

package bspec

// VSNDocument represents a Vision (VSN) document
// Domain: strategic
//
// The future state the organization aims to create
type VSNDocument struct {
	BaseBSpecDocument

	// Strategic foundation documents are required for all conformance levels
	StrategicFoundation bool `json:"strategic_foundation" yaml:"strategic_foundation"`



}

// NewVSNDocument creates a new VSN document with defaults
func NewVSNDocument(id, title, owner string) *VSNDocument {
	doc := &VSNDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeVSN,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainStrategic
	doc.Domain = &domain

	doc.StrategicFoundation = true




	return doc
}

// Validate validates the VSN document and returns any validation errors
func (d *VSNDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeVSN {
		errors = append(errors, "document type must be VSN")
	}

	// Document type specific validation
	// Strategic foundation documents require success criteria
	if len(d.SuccessCriteria) == 0 {
		errors = append(errors, "Strategic foundation documents must have success_criteria defined")
	}




	return errors
}