// Package bspec provides OBJ document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.607530
// Generator: go-generator v1.0.0

package bspec

// OBJDocument represents a Objectives (OBJ) document
// Domain: strategic
//
// Specific, measurable goals with timeframes (OKRs)
type OBJDocument struct {
	BaseBSpecDocument




}

// NewOBJDocument creates a new OBJ document with defaults
func NewOBJDocument(id, title, owner string) *OBJDocument {
	doc := &OBJDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeOBJ,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainStrategic
	doc.Domain = &domain





	return doc
}

// Validate validates the OBJ document and returns any validation errors
func (d *OBJDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeOBJ {
		errors = append(errors, "document type must be OBJ")
	}

	// Document type specific validation




	return errors
}