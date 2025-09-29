// Package bspec provides MET document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616337
// Generator: go-generator v1.0.0

package bspec

// METDocument represents a Metrics (MET) document
// Domain: financial
//
// Key performance indicators and measurement frameworks
type METDocument struct {
	BaseBSpecDocument




}

// NewMETDocument creates a new MET document with defaults
func NewMETDocument(id, title, owner string) *METDocument {
	doc := &METDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeMET,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the MET document and returns any validation errors
func (d *METDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeMET {
		errors = append(errors, "document type must be MET")
	}

	// Document type specific validation




	return errors
}