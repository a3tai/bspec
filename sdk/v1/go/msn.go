// Package bspec provides MSN document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.607078
// Generator: go-generator v1.0.0

package bspec

// MSNDocument represents a Mission (MSN) document
// Domain: strategic
//
// Why the organization exists and its core purpose
type MSNDocument struct {
	BaseBSpecDocument

	// Strategic foundation documents are required for all conformance levels
	StrategicFoundation bool `json:"strategic_foundation" yaml:"strategic_foundation"`



}

// NewMSNDocument creates a new MSN document with defaults
func NewMSNDocument(id, title, owner string) *MSNDocument {
	doc := &MSNDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeMSN,
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

// Validate validates the MSN document and returns any validation errors
func (d *MSNDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeMSN {
		errors = append(errors, "document type must be MSN")
	}

	// Document type specific validation
	// Strategic foundation documents require success criteria
	if len(d.SuccessCriteria) == 0 {
		errors = append(errors, "Strategic foundation documents must have success_criteria defined")
	}




	return errors
}