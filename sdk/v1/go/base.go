// Package bspec provides Go types and utilities for BSpec documents
//
// GENERATED CODE - DO NOT EDIT MANUALLY
// Generated from BSpec v1.0.0 specification
// Generated at: 2025-09-28T15:12:38.601773
// Generator: go-generator v1.0.0

package bspec

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// DocumentStatus represents the lifecycle status of a document
type DocumentStatus string

const (
	DocumentStatusDraft      DocumentStatus = "Draft"
	DocumentStatusReview     DocumentStatus = "Review"
	DocumentStatusAccepted   DocumentStatus = "Accepted"
	DocumentStatusDeprecated DocumentStatus = "Deprecated"
)

// ReviewCycle represents how often a document should be reviewed
type ReviewCycle string

const (
	ReviewCycleMonthly   ReviewCycle = "monthly"
	ReviewCycleQuarterly ReviewCycle = "quarterly"
	ReviewCycleAnnually  ReviewCycle = "annually"
)

// BusinessDomain represents the business domain classification
type BusinessDomain string

const (
BusinessDomainStrategic BusinessDomain = "strategic"
BusinessDomainMarket BusinessDomain = "market"
BusinessDomainCustomer BusinessDomain = "customer"
BusinessDomainProduct BusinessDomain = "product"
BusinessDomainBusinessModel BusinessDomain = "business-model"
BusinessDomainOperations BusinessDomain = "operations"
BusinessDomainTechnology BusinessDomain = "technology"
BusinessDomainFinancial BusinessDomain = "financial"
BusinessDomainRisk BusinessDomain = "risk"
BusinessDomainGrowth BusinessDomain = "growth"
)

// OrganizationalScope represents the organizational scope
type OrganizationalScope string

const (
	OrganizationalScopeGlobal   OrganizationalScope = "global"
	OrganizationalScopeRegional OrganizationalScope = "regional"
	OrganizationalScopeDivision OrganizationalScope = "division"
	OrganizationalScopeTeam     OrganizationalScope = "team"
)

// TimeHorizon represents the time horizon for planning
type TimeHorizon string

const (
	TimeHorizonImmediate TimeHorizon = "immediate"
	TimeHorizonShort     TimeHorizon = "short"
	TimeHorizonMedium    TimeHorizon = "medium"
	TimeHorizonLong      TimeHorizon = "long"
)

// Priority represents the business importance
type Priority string

const (
	PriorityCritical Priority = "critical"
	PriorityHigh     Priority = "high"
	PriorityMedium   Priority = "medium"
	PriorityLow      Priority = "low"
)

// Visibility represents the access level
type Visibility string

const (
	VisibilityPublic     Visibility = "public"
	VisibilityInternal   Visibility = "internal"
	VisibilityRestricted Visibility = "restricted"
)

// ImplementationStatus represents the implementation progress
type ImplementationStatus string

const (
	ImplementationStatusNotStarted ImplementationStatus = "not_started"
	ImplementationStatusInProgress ImplementationStatus = "in_progress"
	ImplementationStatusCompleted  ImplementationStatus = "completed"
	ImplementationStatusOnHold     ImplementationStatus = "on_hold"
)

// Classification represents the information classification
type Classification string

const (
	ClassificationPublic       Classification = "public"
	ClassificationInternal     Classification = "internal"
	ClassificationConfidential Classification = "confidential"
)

// DocumentType represents the BSpec document type codes
type DocumentType string

const (
DocumentTypeMSN DocumentType = "MSN"
DocumentTypeVSN DocumentType = "VSN"
DocumentTypeVAL DocumentType = "VAL"
DocumentTypeSTR DocumentType = "STR"
DocumentTypeOBJ DocumentType = "OBJ"
DocumentTypeMOT DocumentType = "MOT"
DocumentTypePUR DocumentType = "PUR"
DocumentTypeTHY DocumentType = "THY"
DocumentTypeMKT DocumentType = "MKT"
DocumentTypeSEG DocumentType = "SEG"
DocumentTypeCMP DocumentType = "CMP"
DocumentTypePOS DocumentType = "POS"
DocumentTypeTRN DocumentType = "TRN"
DocumentTypeECO DocumentType = "ECO"
DocumentTypeOPP DocumentType = "OPP"
DocumentTypeTHR DocumentType = "THR"
DocumentTypeREG DocumentType = "REG"
DocumentTypeMAC DocumentType = "MAC"
DocumentTypePER DocumentType = "PER"
DocumentTypeJTB DocumentType = "JTB"
DocumentTypeCJM DocumentType = "CJM"
DocumentTypeUSE DocumentType = "USE"
DocumentTypeSTO DocumentType = "STO"
DocumentTypePAI DocumentType = "PAI"
DocumentTypeGAI DocumentType = "GAI"
DocumentTypeEMP DocumentType = "EMP"
DocumentTypeFEE DocumentType = "FEE"
DocumentTypeINT DocumentType = "INT"
DocumentTypeSUR DocumentType = "SUR"
DocumentTypeBEH DocumentType = "BEH"
DocumentTypePRD DocumentType = "PRD"
DocumentTypeSVC DocumentType = "SVC"
DocumentTypeFEA DocumentType = "FEA"
DocumentTypeROD DocumentType = "ROD"
DocumentTypeREQ DocumentType = "REQ"
DocumentTypeQUA DocumentType = "QUA"
DocumentTypeUXD DocumentType = "UXD"
DocumentTypePER DocumentType = "PER"
DocumentTypeINT DocumentType = "INT"
DocumentTypeSUP DocumentType = "SUP"
DocumentTypeBMC DocumentType = "BMC"
DocumentTypeREV DocumentType = "REV"
DocumentTypePRC DocumentType = "PRC"
DocumentTypeCST DocumentType = "CST"
DocumentTypeCHN DocumentType = "CHN"
DocumentTypeREL DocumentType = "REL"
DocumentTypeRES DocumentType = "RES"
DocumentTypeACT DocumentType = "ACT"
DocumentTypePRT DocumentType = "PRT"
DocumentTypeUNT DocumentType = "UNT"
DocumentTypeLTV DocumentType = "LTV"
DocumentTypeCAC DocumentType = "CAC"
DocumentTypePRC DocumentType = "PRC"
DocumentTypeWFL DocumentType = "WFL"
DocumentTypeORG DocumentType = "ORG"
DocumentTypeROL DocumentType = "ROL"
DocumentTypeTEA DocumentType = "TEA"
DocumentTypeSKI DocumentType = "SKI"
DocumentTypePOL DocumentType = "POL"
DocumentTypeSLA DocumentType = "SLA"
DocumentTypeVND DocumentType = "VND"
DocumentTypeFAC DocumentType = "FAC"
DocumentTypeTOO DocumentType = "TOO"
DocumentTypeCAP DocumentType = "CAP"
DocumentTypeARC DocumentType = "ARC"
DocumentTypeSYS DocumentType = "SYS"
DocumentTypeDAT DocumentType = "DAT"
DocumentTypeAPI DocumentType = "API"
DocumentTypeINF DocumentType = "INF"
DocumentTypeSEC DocumentType = "SEC"
DocumentTypeDEV DocumentType = "DEV"
DocumentTypeANA DocumentType = "ANA"
DocumentTypeFIN DocumentType = "FIN"
DocumentTypeBUD DocumentType = "BUD"
DocumentTypeFOR DocumentType = "FOR"
DocumentTypeFND DocumentType = "FND"
DocumentTypeINV DocumentType = "INV"
DocumentTypeVAL DocumentType = "VAL"
DocumentTypeMET DocumentType = "MET"
DocumentTypeREP DocumentType = "REP"
DocumentTypeAUD DocumentType = "AUD"
DocumentTypeTAX DocumentType = "TAX"
DocumentTypeRSK DocumentType = "RSK"
DocumentTypeMIT DocumentType = "MIT"
DocumentTypeCMP DocumentType = "CMP"
DocumentTypeGVN DocumentType = "GVN"
DocumentTypeCTL DocumentType = "CTL"
DocumentTypeCRI DocumentType = "CRI"
DocumentTypeETH DocumentType = "ETH"
DocumentTypeSTA DocumentType = "STA"
DocumentTypeGTM DocumentType = "GTM"
DocumentTypeGRW DocumentType = "GRW"
DocumentTypeSCL DocumentType = "SCL"
DocumentTypeEXP DocumentType = "EXP"
DocumentTypeINN DocumentType = "INN"
DocumentTypeRND DocumentType = "RND"
DocumentTypeACQ DocumentType = "ACQ"
DocumentTypeEXP DocumentType = "EXP"
)

// ConformanceLevel represents the BSpec conformance levels
type ConformanceLevel string

const (
ConformanceLevelBronze ConformanceLevel = "bronze"
ConformanceLevelSilver ConformanceLevel = "silver"
ConformanceLevelGold ConformanceLevel = "gold"
)

// IndustryProfile represents the BSpec industry profiles
type IndustryProfile string

const (
IndustryProfileSoftwareSaas IndustryProfile = "software-saas"
IndustryProfilePhysicalProduct IndustryProfile = "physical-product"
IndustryProfileServiceBusiness IndustryProfile = "service-business"
IndustryProfileNonprofit IndustryProfile = "nonprofit"
)

// ChangelogEntry represents a single entry in the document version history
type ChangelogEntry struct {
	Version         string `json:"version" yaml:"version"`                                 // Version number
	Date            string `json:"date" yaml:"date"`                                       // Change date (YYYY-MM-DD)
	Author          string `json:"author" yaml:"author"`                                   // Author name
	Changes         string `json:"changes" yaml:"changes"`                                 // Description of changes
	BreakingChanges bool   `json:"breaking_changes" yaml:"breaking_changes"`               // Whether changes are breaking
}

// BaseBSpecDocument defines the universal YAML frontmatter schema
// All BSpec documents extend this base structure
type BaseBSpecDocument struct {
	// === CORE IDENTITY ===
	ID      string         `json:"id" yaml:"id"`           // Globally unique document identifier
	Title   string         `json:"title" yaml:"title"`     // Clear, descriptive title
	Type    DocumentType   `json:"type" yaml:"type"`       // Document type code
	Status  DocumentStatus `json:"status" yaml:"status"`   // Document lifecycle status
	Version string         `json:"version" yaml:"version"` // Semantic versioning

	// === OWNERSHIP & RESPONSIBILITY ===
	Owner        string   `json:"owner" yaml:"owner"`                               // Who owns and maintains this document
	Stakeholders []string `json:"stakeholders,omitempty" yaml:"stakeholders,flow"` // Who has interest in this document
	Reviewers    []string `json:"reviewers,omitempty" yaml:"reviewers,flow"`       // Who reviewed/approved this version
	Contributors []string `json:"contributors,omitempty" yaml:"contributors,flow"` // Who contributed to this document

	// === TEMPORAL METADATA ===
	Created     string       `json:"created" yaml:"created"`                     // When document was first created (YYYY-MM-DD)
	Updated     string       `json:"updated" yaml:"updated"`                     // When document was last modified (YYYY-MM-DD)
	Expires     *string      `json:"expires,omitempty" yaml:"expires,omitempty"` // When document expires (YYYY-MM-DD)
	ReviewCycle *ReviewCycle `json:"review_cycle,omitempty" yaml:"review_cycle,omitempty"` // How often to review

	// === RELATIONSHIP GRAPH ===
	Parent        *string  `json:"parent,omitempty" yaml:"parent,omitempty"`               // Parent document (hierarchical)
	DependsOn     []string `json:"depends_on,omitempty" yaml:"depends_on,flow"`           // Dependencies (this needs those)
	Enables       []string `json:"enables,omitempty" yaml:"enables,flow"`                 // Enablements (this makes those possible)
	ConflictsWith []string `json:"conflicts_with,omitempty" yaml:"conflicts_with,flow"`   // Mutual exclusions
	Related       []string `json:"related,omitempty" yaml:"related,flow"`                 // Other relevant documents
	Supersedes    *string  `json:"supersedes,omitempty" yaml:"supersedes,omitempty"`      // What this document replaces

	// === BUSINESS CONTEXT ===
	Domain     *BusinessDomain      `json:"domain,omitempty" yaml:"domain,omitempty"`         // Business domain classification
	Scope      *OrganizationalScope `json:"scope,omitempty" yaml:"scope,omitempty"`           // Organizational scope
	Horizon    *TimeHorizon         `json:"horizon,omitempty" yaml:"horizon,omitempty"`       // Time horizon
	Priority   *Priority            `json:"priority,omitempty" yaml:"priority,omitempty"`     // Business importance
	Visibility *Visibility          `json:"visibility,omitempty" yaml:"visibility,omitempty"` // Access level

	// === VALIDATION & MEASUREMENT ===
	Assumptions      []string `json:"assumptions,omitempty" yaml:"assumptions,flow"`           // Key assumptions
	Constraints      []string `json:"constraints,omitempty" yaml:"constraints,flow"`           // Key constraints
	SuccessCriteria  []string `json:"success_criteria,omitempty" yaml:"success_criteria,flow"` // Measurable success criteria
	Risks            []string `json:"risks,omitempty" yaml:"risks,flow"`                       // Associated risk documents
	Metrics          []string `json:"metrics,omitempty" yaml:"metrics,flow"`                   // How success is measured

	// === IMPLEMENTATION ===
	ImplementationStatus *ImplementationStatus `json:"implementation_status,omitempty" yaml:"implementation_status,omitempty"` // Implementation progress
	ImplementationDate   *string               `json:"implementation_date,omitempty" yaml:"implementation_date,omitempty"`     // When implementation begins/began (YYYY-MM-DD)
	CompletionDate       *string               `json:"completion_date,omitempty" yaml:"completion_date,omitempty"`             // When implementation completes (YYYY-MM-DD)
	ResourcesRequired    []string              `json:"resources_required,omitempty" yaml:"resources_required,flow"`            // Resource requirements

	// === METADATA & DISCOVERY ===
	Tags           []string        `json:"tags,omitempty" yaml:"tags,flow"`                       // Searchable tags
	Industry       []string        `json:"industry,omitempty" yaml:"industry,flow"`               // Industry classifications
	Geography      []string        `json:"geography,omitempty" yaml:"geography,flow"`             // Geographic relevance
	Language       *string         `json:"language,omitempty" yaml:"language,omitempty"`          // Primary language
	Classification *Classification `json:"classification,omitempty" yaml:"classification,omitempty"` // Information classification

	// === CHANGE TRACKING ===
	Changelog []ChangelogEntry `json:"changelog,omitempty" yaml:"changelog,omitempty"` // Version history

	// === DOCUMENT CONTENT ===
	Content string `json:"content" yaml:"content"` // Markdown content of the document
}

// Validate validates the document and returns any validation errors
func (d *BaseBSpecDocument) Validate() []string {
	var errors []string

	// Required field validation
	if d.ID == "" {
		errors = append(errors, "id is required")
	}
	if d.Title == "" {
		errors = append(errors, "title is required")
	}
	if d.Owner == "" {
		errors = append(errors, "owner is required")
	}
	if d.Created == "" {
		errors = append(errors, "created date is required")
	}
	if d.Updated == "" {
		errors = append(errors, "updated date is required")
	}

	// ID format validation
	if d.ID != "" && !strings.HasPrefix(d.ID, string(d.Type)+"-") {
		errors = append(errors, fmt.Sprintf("id must start with document type '%s-'", d.Type))
	}

	// Version format validation (semantic versioning)
	if d.Version != "" {
		versionRegex := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
		if !versionRegex.MatchString(d.Version) {
			errors = append(errors, "version must follow semantic versioning (e.g., '1.0.0')")
		}
	}

	// Date format validation (basic YYYY-MM-DD check)
	dateFields := map[string]string{
		"created":             d.Created,
		"updated":             d.Updated,
	}
	if d.Expires != nil {
		dateFields["expires"] = *d.Expires
	}
	if d.ImplementationDate != nil {
		dateFields["implementation_date"] = *d.ImplementationDate
	}
	if d.CompletionDate != nil {
		dateFields["completion_date"] = *d.CompletionDate
	}

	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	for fieldName, fieldValue := range dateFields {
		if fieldValue != "" && !dateRegex.MatchString(fieldValue) {
			errors = append(errors, fmt.Sprintf("%s must be in YYYY-MM-DD format", fieldName))
		}
	}

	return errors
}

// IsValid returns true if the document passes validation
func (d *BaseBSpecDocument) IsValid() bool {
	return len(d.Validate()) == 0
}

// ToJSON serializes the document to JSON
func (d *BaseBSpecDocument) ToJSON() ([]byte, error) {
	return json.Marshal(d)
}

// FromJSON deserializes the document from JSON
func (d *BaseBSpecDocument) FromJSON(data []byte) error {
	return json.Unmarshal(data, d)
}

// GetDomain returns the business domain, defaulting based on document type if not set
func (d *BaseBSpecDocument) GetDomain() BusinessDomain {
	if d.Domain != nil {
		return *d.Domain
	}

	// Default domain mapping based on document type
	switch d.Type {
case DocumentTypeMSN:
		return BusinessDomainStrategic
case DocumentTypeVSN:
		return BusinessDomainStrategic
case DocumentTypeVAL:
		return BusinessDomainStrategic
case DocumentTypeSTR:
		return BusinessDomainStrategic
case DocumentTypeOBJ:
		return BusinessDomainStrategic
case DocumentTypeMOT:
		return BusinessDomainStrategic
case DocumentTypePUR:
		return BusinessDomainStrategic
case DocumentTypeTHY:
		return BusinessDomainStrategic
case DocumentTypeMKT:
		return BusinessDomainMarket
case DocumentTypeSEG:
		return BusinessDomainMarket
case DocumentTypeCMP:
		return BusinessDomainMarket
case DocumentTypePOS:
		return BusinessDomainMarket
case DocumentTypeTRN:
		return BusinessDomainMarket
case DocumentTypeECO:
		return BusinessDomainMarket
case DocumentTypeOPP:
		return BusinessDomainMarket
case DocumentTypeTHR:
		return BusinessDomainMarket
case DocumentTypeREG:
		return BusinessDomainMarket
case DocumentTypeMAC:
		return BusinessDomainMarket
case DocumentTypePER:
		return BusinessDomainCustomer
case DocumentTypeJTB:
		return BusinessDomainCustomer
case DocumentTypeCJM:
		return BusinessDomainCustomer
case DocumentTypeUSE:
		return BusinessDomainCustomer
case DocumentTypeSTO:
		return BusinessDomainCustomer
case DocumentTypePAI:
		return BusinessDomainCustomer
case DocumentTypeGAI:
		return BusinessDomainCustomer
case DocumentTypeEMP:
		return BusinessDomainCustomer
case DocumentTypeFEE:
		return BusinessDomainCustomer
case DocumentTypeINT:
		return BusinessDomainCustomer
case DocumentTypeSUR:
		return BusinessDomainCustomer
case DocumentTypeBEH:
		return BusinessDomainCustomer
case DocumentTypePRD:
		return BusinessDomainProduct
case DocumentTypeSVC:
		return BusinessDomainProduct
case DocumentTypeFEA:
		return BusinessDomainProduct
case DocumentTypeROD:
		return BusinessDomainProduct
case DocumentTypeREQ:
		return BusinessDomainProduct
case DocumentTypeQUA:
		return BusinessDomainProduct
case DocumentTypeUXD:
		return BusinessDomainProduct
case DocumentTypePER:
		return BusinessDomainProduct
case DocumentTypeINT:
		return BusinessDomainProduct
case DocumentTypeSUP:
		return BusinessDomainProduct
case DocumentTypeBMC:
		return BusinessDomainBusinessModel
case DocumentTypeREV:
		return BusinessDomainBusinessModel
case DocumentTypePRC:
		return BusinessDomainBusinessModel
case DocumentTypeCST:
		return BusinessDomainBusinessModel
case DocumentTypeCHN:
		return BusinessDomainBusinessModel
case DocumentTypeREL:
		return BusinessDomainBusinessModel
case DocumentTypeRES:
		return BusinessDomainBusinessModel
case DocumentTypeACT:
		return BusinessDomainBusinessModel
case DocumentTypePRT:
		return BusinessDomainBusinessModel
case DocumentTypeUNT:
		return BusinessDomainBusinessModel
case DocumentTypeLTV:
		return BusinessDomainBusinessModel
case DocumentTypeCAC:
		return BusinessDomainBusinessModel
case DocumentTypePRC:
		return BusinessDomainOperations
case DocumentTypeWFL:
		return BusinessDomainOperations
case DocumentTypeORG:
		return BusinessDomainOperations
case DocumentTypeROL:
		return BusinessDomainOperations
case DocumentTypeTEA:
		return BusinessDomainOperations
case DocumentTypeSKI:
		return BusinessDomainOperations
case DocumentTypePOL:
		return BusinessDomainOperations
case DocumentTypeSLA:
		return BusinessDomainOperations
case DocumentTypeVND:
		return BusinessDomainOperations
case DocumentTypeFAC:
		return BusinessDomainOperations
case DocumentTypeTOO:
		return BusinessDomainOperations
case DocumentTypeCAP:
		return BusinessDomainOperations
case DocumentTypeARC:
		return BusinessDomainTechnology
case DocumentTypeSYS:
		return BusinessDomainTechnology
case DocumentTypeDAT:
		return BusinessDomainTechnology
case DocumentTypeAPI:
		return BusinessDomainTechnology
case DocumentTypeINF:
		return BusinessDomainTechnology
case DocumentTypeSEC:
		return BusinessDomainTechnology
case DocumentTypeDEV:
		return BusinessDomainTechnology
case DocumentTypeANA:
		return BusinessDomainTechnology
case DocumentTypeFIN:
		return BusinessDomainFinancial
case DocumentTypeBUD:
		return BusinessDomainFinancial
case DocumentTypeFOR:
		return BusinessDomainFinancial
case DocumentTypeFND:
		return BusinessDomainFinancial
case DocumentTypeINV:
		return BusinessDomainFinancial
case DocumentTypeVAL:
		return BusinessDomainFinancial
case DocumentTypeMET:
		return BusinessDomainFinancial
case DocumentTypeREP:
		return BusinessDomainFinancial
case DocumentTypeAUD:
		return BusinessDomainFinancial
case DocumentTypeTAX:
		return BusinessDomainFinancial
case DocumentTypeRSK:
		return BusinessDomainRisk
case DocumentTypeMIT:
		return BusinessDomainRisk
case DocumentTypeCMP:
		return BusinessDomainRisk
case DocumentTypeGVN:
		return BusinessDomainRisk
case DocumentTypeCTL:
		return BusinessDomainRisk
case DocumentTypeCRI:
		return BusinessDomainRisk
case DocumentTypeETH:
		return BusinessDomainRisk
case DocumentTypeSTA:
		return BusinessDomainRisk
case DocumentTypeGTM:
		return BusinessDomainGrowth
case DocumentTypeGRW:
		return BusinessDomainGrowth
case DocumentTypeSCL:
		return BusinessDomainGrowth
case DocumentTypeEXP:
		return BusinessDomainGrowth
case DocumentTypeINN:
		return BusinessDomainGrowth
case DocumentTypeRND:
		return BusinessDomainGrowth
case DocumentTypeACQ:
		return BusinessDomainGrowth
case DocumentTypeEXP:
		return BusinessDomainGrowth
default:
		return BusinessDomainStrategic // Default fallback
	}
}

// SetDefaults sets reasonable defaults for optional fields
func (d *BaseBSpecDocument) SetDefaults() {
	now := time.Now().Format("2006-01-02")

	if d.Created == "" {
		d.Created = now
	}
	if d.Updated == "" {
		d.Updated = now
	}
	if d.Status == "" {
		d.Status = DocumentStatusDraft
	}
	if d.Version == "" {
		d.Version = "1.0.0"
	}

	// Initialize slices if nil
	if d.Stakeholders == nil {
		d.Stakeholders = []string{}
	}
	if d.Reviewers == nil {
		d.Reviewers = []string{}
	}
	if d.Contributors == nil {
		d.Contributors = []string{}
	}
	if d.DependsOn == nil {
		d.DependsOn = []string{}
	}
	if d.Enables == nil {
		d.Enables = []string{}
	}
	if d.ConflictsWith == nil {
		d.ConflictsWith = []string{}
	}
	if d.Related == nil {
		d.Related = []string{}
	}
	if d.Assumptions == nil {
		d.Assumptions = []string{}
	}
	if d.Constraints == nil {
		d.Constraints = []string{}
	}
	if d.SuccessCriteria == nil {
		d.SuccessCriteria = []string{}
	}
	if d.Risks == nil {
		d.Risks = []string{}
	}
	if d.Metrics == nil {
		d.Metrics = []string{}
	}
	if d.ResourcesRequired == nil {
		d.ResourcesRequired = []string{}
	}
	if d.Tags == nil {
		d.Tags = []string{}
	}
	if d.Industry == nil {
		d.Industry = []string{}
	}
	if d.Geography == nil {
		d.Geography = []string{}
	}
	if d.Changelog == nil {
		d.Changelog = []ChangelogEntry{}
	}
}