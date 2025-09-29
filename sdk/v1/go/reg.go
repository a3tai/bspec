// Package bspec provides REG document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.609043
// Generator: go-generator v1.0.0

package bspec

// REGDocument represents a Regulatory Environment (REG) document
// Domain: market
//
// Laws, regulations, and compliance requirements
type REGDocument struct {
	BaseBSpecDocument




}

// NewREGDocument creates a new REG document with defaults
func NewREGDocument(id, title, owner string) *REGDocument {
	doc := &REGDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeREG,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainMarket
	doc.Domain = &domain





	return doc
}

// Validate validates the REG document and returns any validation errors
func (d *REGDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeREG {
		errors = append(errors, "document type must be REG")
	}

	// Document type specific validation




	return errors
}