// Package bspec provides MAC document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.609151
// Generator: go-generator v1.0.0

package bspec

// MACDocument represents a Macro Environment (MAC) document
// Domain: market
//
// Economic, political, social, technological factors
type MACDocument struct {
	BaseBSpecDocument




}

// NewMACDocument creates a new MAC document with defaults
func NewMACDocument(id, title, owner string) *MACDocument {
	doc := &MACDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeMAC,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainMarket
	doc.Domain = &domain





	return doc
}

// Validate validates the MAC document and returns any validation errors
func (d *MACDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeMAC {
		errors = append(errors, "document type must be MAC")
	}

	// Document type specific validation




	return errors
}