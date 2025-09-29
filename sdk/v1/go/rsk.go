// Package bspec provides RSK document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616728
// Generator: go-generator v1.0.0

package bspec

// RSKDocument represents a Risks (RSK) document
// Domain: risk
//
// Identified threats to business success
type RSKDocument struct {
	BaseBSpecDocument




	// Risk management documents require mitigation planning
	RiskManagement bool `json:"risk_management" yaml:"risk_management"`
}

// NewRSKDocument creates a new RSK document with defaults
func NewRSKDocument(id, title, owner string) *RSKDocument {
	doc := &RSKDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeRSK,
			Owner: owner,
		},
	}

	// Set document type specific defaults
	doc.SetDefaults()
	domain := BusinessDomainRisk
	doc.Domain = &domain




	doc.RiskManagement = true

	return doc
}

// Validate validates the RSK document and returns any validation errors
func (d *RSKDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeRSK {
		errors = append(errors, "document type must be RSK")
	}

	// Document type specific validation



	// Risk management documents should reference each other
	if d.Type == DocumentTypeRSK {
		hasRelatedMitigation := false
		for _, ref := range d.Related {
			if len(ref) >= 4 && ref[:4] == "MIT-" {
				hasRelatedMitigation = true
				break
			}
		}
		if !hasRelatedMitigation {
			errors = append(errors, "Risk documents should reference corresponding mitigation documents")
		}
	}

	return errors
}