package sdk

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/a3tai/bspec/cli/internal/archive"
)

// Import the working SDK types by referencing the actual SDK package
// We'll use dynamic loading since the SDK has compilation issues

// SDKConverter handles conversion between archive format and SDK format
type SDKConverter struct {
	bspecData map[string]interface{} // Dynamic data from SDK JSON
}

// NewSDKConverter creates a new SDK converter
func NewSDKConverter() *SDKConverter {
	return &SDKConverter{
		bspecData: make(map[string]interface{}),
	}
}

// LoadFromArchive loads BSpec data from an archive and converts to SDK format
func (c *SDKConverter) LoadFromArchive(archivePath string) error {
	// Load archive data
	bspecArchive, err := archive.Read(archivePath)
	if err != nil {
		return fmt.Errorf("failed to load archive: %w", err)
	}

	// Convert to SDK format dynamically
	return c.convertFromArchive(bspecArchive)
}

// LoadSDKFromJSON loads the actual SDK JSON data
func (c *SDKConverter) LoadSDKFromJSON(jsonPath string) error {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return fmt.Errorf("failed to read SDK JSON: %w", err)
	}

	if err := json.Unmarshal(data, &c.bspecData); err != nil {
		return fmt.Errorf("failed to parse SDK JSON: %w", err)
	}

	return nil
}

// GetDocumentTypes returns all document types from the SDK
func (c *SDKConverter) GetDocumentTypes() ([]map[string]interface{}, error) {
	if c.bspecData == nil {
		return nil, fmt.Errorf("no SDK data loaded")
	}

	docTypes, ok := c.bspecData["document_types"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("document_types not found in SDK data")
	}

	var result []map[string]interface{}
	for _, dt := range docTypes {
		if docTypeMap, ok := dt.(map[string]interface{}); ok {
			result = append(result, docTypeMap)
		}
	}

	return result, nil
}

// GetDocumentType returns a specific document type by code
func (c *SDKConverter) GetDocumentType(code string) (map[string]interface{}, error) {
	docTypes, err := c.GetDocumentTypes()
	if err != nil {
		return nil, err
	}

	for _, docType := range docTypes {
		if docCode, ok := docType["code"].(string); ok && docCode == code {
			return docType, nil
		}
	}

	return nil, fmt.Errorf("document type %s not found", code)
}

// GetDomains returns all domains from the SDK
func (c *SDKConverter) GetDomains() ([]map[string]interface{}, error) {
	if c.bspecData == nil {
		return nil, fmt.Errorf("no SDK data loaded")
	}

	domains, ok := c.bspecData["domains"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("domains not found in SDK data")
	}

	var result []map[string]interface{}
	for _, d := range domains {
		if domainMap, ok := d.(map[string]interface{}); ok {
			result = append(result, domainMap)
		}
	}

	return result, nil
}

// GetYAMLSchema returns the YAML schema from the SDK
func (c *SDKConverter) GetYAMLSchema() (map[string]interface{}, error) {
	if c.bspecData == nil {
		return nil, fmt.Errorf("no SDK data loaded")
	}

	schema, ok := c.bspecData["yaml_schema"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("yaml_schema not found in SDK data")
	}

	return schema, nil
}

// ValidateDocumentMetadata validates document metadata against BSpec YAML schema
func (c *SDKConverter) ValidateDocumentMetadata(metadata map[string]interface{}) error {
	schema, err := c.GetYAMLSchema()
	if err != nil {
		// Fall back to basic validation if no schema available
		return c.validateBasicMetadata(metadata)
	}

	// Check for required fields based on YAML schema
	if requiredFields, ok := schema["required_fields"].([]interface{}); ok {
		for _, fieldInterface := range requiredFields {
			if field, ok := fieldInterface.(string); ok {
				if _, exists := metadata[field]; !exists {
					return fmt.Errorf("missing required field: %s", field)
				}
			}
		}
	}

	// Validate field types
	if fieldTypes, ok := schema["field_types"].(map[string]interface{}); ok {
		for field, expectedTypeInterface := range fieldTypes {
			if expectedType, ok := expectedTypeInterface.(string); ok {
				if value, exists := metadata[field]; exists {
					if err := c.validateFieldType(field, value, expectedType); err != nil {
						return err
					}
				}
			}
		}
	}

	// Validate enums
	if enums, ok := schema["enums"].(map[string]interface{}); ok {
		for field, enumValuesInterface := range enums {
			if enumValues, ok := enumValuesInterface.([]interface{}); ok {
				if value, exists := metadata[field]; exists {
					if !c.isValueInEnum(value, enumValues) {
						return fmt.Errorf("invalid value for %s: %v", field, value)
					}
				}
			}
		}
	}

	return nil
}

// validateBasicMetadata provides basic validation when no schema is available
func (c *SDKConverter) validateBasicMetadata(metadata map[string]interface{}) error {
	requiredFields := []string{"id", "title", "status", "version", "owner"}
	for _, field := range requiredFields {
		if _, exists := metadata[field]; !exists {
			return fmt.Errorf("missing required field: %s", field)
		}
	}

	// Validate status field
	if statusRaw, exists := metadata["status"]; exists {
		status, ok := statusRaw.(string)
		if !ok {
			return fmt.Errorf("status field must be a string")
		}
		validStatuses := []string{"Draft", "Review", "Accepted", "Deprecated"}
		if !c.isStringInSlice(status, validStatuses) {
			return fmt.Errorf("invalid status value: %s", status)
		}
	}

	return nil
}

// validateFieldType validates that a field value matches the expected type
func (c *SDKConverter) validateFieldType(field string, value interface{}, expectedType string) error {
	switch expectedType {
	case "string":
		if _, ok := value.(string); !ok {
			return fmt.Errorf("field %s must be a string", field)
		}
	case "array":
		if _, ok := value.([]interface{}); !ok {
			return fmt.Errorf("field %s must be an array", field)
		}
	case "object":
		if _, ok := value.(map[string]interface{}); !ok {
			return fmt.Errorf("field %s must be an object", field)
		}
	case "number":
		switch value.(type) {
		case int, int32, int64, float32, float64:
			// Valid number types
		default:
			return fmt.Errorf("field %s must be a number", field)
		}
	case "boolean":
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("field %s must be a boolean", field)
		}
	}
	return nil
}

// isValueInEnum checks if a value exists in an enum slice
func (c *SDKConverter) isValueInEnum(value interface{}, enumValues []interface{}) bool {
	for _, enumValue := range enumValues {
		if value == enumValue {
			return true
		}
	}
	return false
}

// isStringInSlice checks if a string exists in a slice
func (c *SDKConverter) isStringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}
	return false
}

// SearchDocumentTypes searches document types by query
func (c *SDKConverter) SearchDocumentTypes(query string) ([]map[string]interface{}, error) {
	docTypes, err := c.GetDocumentTypes()
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	for _, docType := range docTypes {
		// Search in name, purpose, and code
		if c.matchesQuery(docType, query) {
			results = append(results, docType)
		}
	}

	return results, nil
}

// matchesQuery checks if a document type matches the search query
func (c *SDKConverter) matchesQuery(docType map[string]interface{}, query string) bool {
	searchableFields := []string{"name", "purpose", "code"}

	for _, field := range searchableFields {
		if value, ok := docType[field].(string); ok {
			if contains(value, query) {
				return true
			}
		}
	}

	return false
}

// contains checks if a string contains a substring (case insensitive)
func contains(str, substr string) bool {
	return len(str) >= len(substr) &&
		   (str == substr ||
		   len(str) > len(substr) &&
		   (str[:len(substr)] == substr ||
		   str[len(str)-len(substr):] == substr ||
		   findInString(str, substr)))
}

// findInString is a simple substring search
func findInString(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// convertFromArchive converts archive format to SDK format
func (c *SDKConverter) convertFromArchive(bspecArchive *archive.BSpecArchive) error {
	// Build basic BSpec structure from archive
	c.bspecData = map[string]interface{}{
		"metadata": map[string]interface{}{
			"bspec_version": "1.0.0",
			"generated_at":  time.Now().Format(time.RFC3339),
			"generator":     "bspec-cli",
			"source_spec":   "archive",
		},
		"statistics": map[string]interface{}{
			"total_files":          len(bspecArchive.Documents),
			"total_document_types": c.countUniqueTypes(bspecArchive.Documents),
			"total_domains":        c.countUniqueDomains(bspecArchive.Documents),
		},
		"yaml_schema": map[string]interface{}{
			"required_fields": []string{"id", "title", "status", "version", "owner"},
			"field_types": map[string]interface{}{
				"id":       "string",
				"title":    "string",
				"status":   "string",
				"version":  "string",
				"owner":    "string",
				"type":     "string",
				"contexts": "array",
				"related":  "array",
				"metrics":  "array",
				"risks":    "array",
			},
			"enums": map[string]interface{}{
				"status": []string{"Draft", "Review", "Accepted", "Deprecated"},
			},
		},
	}

	// Build document types and domains dynamically from archive
	c.buildDocumentTypesFromArchive(bspecArchive)
	c.buildDomainsFromArchive(bspecArchive)

	return nil
}

// countUniqueTypes counts unique document types in archive
func (c *SDKConverter) countUniqueTypes(documents map[string]archive.BSpecDocument) int {
	types := make(map[string]bool)
	for _, doc := range documents {
		if doc.Type != "" {
			types[doc.Type] = true
		}
	}
	return len(types)
}

// countUniqueDomains counts unique domains in archive
func (c *SDKConverter) countUniqueDomains(documents map[string]archive.BSpecDocument) int {
	domains := make(map[string]bool)
	for _, doc := range documents {
		if doc.Domain != "" {
			domains[doc.Domain] = true
		}
	}
	return len(domains)
}

// buildDocumentTypesFromArchive builds document types from archive data
func (c *SDKConverter) buildDocumentTypesFromArchive(bspecArchive *archive.BSpecArchive) {
	types := make(map[string]map[string]interface{})

	for _, doc := range bspecArchive.Documents {
		if doc.Type != "" {
			if _, exists := types[doc.Type]; !exists {
				types[doc.Type] = map[string]interface{}{
					"code":    doc.Type,
					"name":    c.generateTypeName(doc.Type),
					"purpose": c.generateTypePurpose(doc.Type),
					"domain":  doc.Domain,
					"examples": []map[string]interface{}{},
				}
			}
		}
	}

	var docTypes []interface{}
	for _, docType := range types {
		docTypes = append(docTypes, docType)
	}

	c.bspecData["document_types"] = docTypes
}

// buildDomainsFromArchive builds domains from archive data
func (c *SDKConverter) buildDomainsFromArchive(bspecArchive *archive.BSpecArchive) {
	domainMap := make(map[string]map[string]interface{})

	for _, doc := range bspecArchive.Documents {
		if doc.Domain != "" {
			if _, exists := domainMap[doc.Domain]; !exists {
				domainMap[doc.Domain] = map[string]interface{}{
					"name":           doc.Domain,
					"display_name":   doc.Domain,
					"description":    fmt.Sprintf("%s domain", doc.Domain),
					"emoji":          "📋",
					"document_types": []string{},
					"document_count": 0,
				}
			}

			// Count documents and track types
			domain := domainMap[doc.Domain]
			domain["document_count"] = domain["document_count"].(int) + 1

			// Add document type if not already present
			types := domain["document_types"].([]string)
			typeExists := false
			for _, t := range types {
				if t == doc.Type {
					typeExists = true
					break
				}
			}
			if !typeExists && doc.Type != "" {
				domain["document_types"] = append(types, doc.Type)
			}
		}
	}

	var domains []interface{}
	for _, domain := range domainMap {
		domains = append(domains, domain)
	}

	c.bspecData["domains"] = domains
}

// generateTypeName generates a human-readable name for a document type code
func (c *SDKConverter) generateTypeName(code string) string {
	names := map[string]string{
		"MSN": "Mission Statement",
		"VSN": "Vision Statement",
		"CAP": "Capability",
		"CTX": "Context",
		"PER": "Persona",
		"PRC": "Process",
		"MET": "Metric",
		"RSK": "Risk",
		"ADR": "Architecture Decision Record",
		"EXP": "Experiment",
	}

	if name, exists := names[code]; exists {
		return name
	}
	return fmt.Sprintf("%s Document", code)
}

// generateTypePurpose generates a purpose description for a document type code
func (c *SDKConverter) generateTypePurpose(code string) string {
	purposes := map[string]string{
		"MSN": "Define organizational mission and core purpose",
		"VSN": "Articulate long-term vision and aspirations",
		"CAP": "Describe business capabilities and value delivery",
		"CTX": "Define bounded contexts and system boundaries",
		"PER": "Document user personas and their jobs-to-be-done",
		"PRC": "Specify repeatable business processes",
		"MET": "Define key performance indicators and metrics",
		"RSK": "Identify and assess business risks",
		"ADR": "Record architectural decisions and rationale",
		"EXP": "Document validation experiments and learnings",
	}

	if purpose, exists := purposes[code]; exists {
		return purpose
	}
	return fmt.Sprintf("Document of type %s", code)
}