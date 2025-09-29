// Package bspec provides REP document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616431
// Generator: go-generator v1.0.0

package bspec

// REPDocument represents a Reporting (REP) document
// Domain: financial
//
// Financial reporting and dashboard requirements
type REPDocument struct {
	BaseBSpecDocument




}

// NewREPDocument creates a new REP document with defaults
func NewREPDocument(id, title, owner string) *REPDocument {
	doc := &REPDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeREP,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainFinancial
	doc.Domain = &domain





	return doc
}

// Validate validates the REP document and returns any validation errors
func (d *REPDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeREP {
		errors = append(errors, "document type must be REP")
	}

	// Document type specific validation




	return errors
}