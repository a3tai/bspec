// Package bspec provides GVN document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616993
// Generator: go-generator v1.0.0

package bspec

// GVNDocument represents a Governance (GVN) document
// Domain: risk
//
// Decision-making structure and oversight
type GVNDocument struct {
	BaseBSpecDocument




}

// NewGVNDocument creates a new GVN document with defaults
func NewGVNDocument(id, title, owner string) *GVNDocument {
	doc := &GVNDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeGVN,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainRisk
	doc.Domain = &domain





	return doc
}

// Validate validates the GVN document and returns any validation errors
func (d *GVNDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeGVN {
		errors = append(errors, "document type must be GVN")
	}

	// Document type specific validation




	return errors
}