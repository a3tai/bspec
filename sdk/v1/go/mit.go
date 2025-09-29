// Package bspec provides MIT document type
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.616832
// Generator: go-generator v1.0.0

package bspec

// MITDocument represents a Mitigations (MIT) document
// Domain: risk
//
// Risk response strategies and controls
type MITDocument struct {
	BaseBSpecDocument




	// Risk management documents require mitigation planning
	RiskManagement bool `json:"risk_management" yaml:"risk_management"`
}

// NewMITDocument creates a new MIT document with defaults
func NewMITDocument(id, title, owner string) *MITDocument {
	doc := &MITDocument{
		BaseBSpecDocument: BaseBSpecDocument{
			ID:    id,
			Title: title,
			Type:  DocumentTypeMIT,
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

// Validate validates the MIT document and returns any validation errors
func (d *MITDocument) Validate() []string {
	errors := d.BaseBSpecDocument.Validate()

	// Ensure correct document type
	if d.Type != DocumentTypeMIT {
		errors = append(errors, "document type must be MIT")
	}

	// Document type specific validation



	// Risk management documents should reference each other
	if d.Type == DocumentTypeMIT {
		hasRelatedRisk := false
		for _, ref := range d.Related {
			if len(ref) >= 4 && ref[:4] == "RSK-" {
				hasRelatedRisk = true
				break
			}
		}
		if !hasRelatedRisk {
			errors = append(errors, "Mitigation documents should reference corresponding risk documents")
		}
	}

	return errors
}