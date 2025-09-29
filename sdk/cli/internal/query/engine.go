package query

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/a3tai/bspec/cli/internal/archive"
)

// QueryEngine provides structured querying capabilities for BSpec documents
type QueryEngine struct {
	archive *archive.BSpecArchive
}

// NewQueryEngine creates a new query engine for the given archive
func NewQueryEngine(arch *archive.BSpecArchive) *QueryEngine {
	return &QueryEngine{
		archive: arch,
	}
}

// Query represents a structured query
type Query struct {
	Type      string            `json:"type,omitempty"`       // Filter by document type
	Domain    string            `json:"domain,omitempty"`     // Filter by domain
	Status    string            `json:"status,omitempty"`     // Filter by status
	Owner     string            `json:"owner,omitempty"`      // Filter by owner
	Tags      []string          `json:"tags,omitempty"`       // Filter by tags
	Search    string            `json:"search,omitempty"`     // Text search in content
	Fields    []string          `json:"fields,omitempty"`     // Select specific fields
	Limit     int               `json:"limit,omitempty"`      // Limit results
	SortBy    string            `json:"sort_by,omitempty"`    // Sort field
	SortOrder string            `json:"sort_order,omitempty"` // asc or desc
	Metadata  map[string]string `json:"metadata,omitempty"`   // Filter by metadata
}

// QueryResult represents the result of a query
type QueryResult struct {
	Documents []archive.BSpecDocument `json:"documents"`
	Total     int                     `json:"total"`
	Query     Query                   `json:"query"`
}

// Execute executes a query against the archive
func (qe *QueryEngine) Execute(q Query) (*QueryResult, error) {
	var results []archive.BSpecDocument

	// Apply filters
	for _, doc := range qe.archive.Documents {
		if qe.matchesQuery(doc, q) {
			results = append(results, doc)
		}
	}

	// Apply sorting
	if q.SortBy != "" {
		results = qe.sortDocuments(results, q.SortBy, q.SortOrder)
	}

	// Apply limit
	if q.Limit > 0 && len(results) > q.Limit {
		results = results[:q.Limit]
	}

	// Apply field selection
	if len(q.Fields) > 0 {
		results = qe.selectFields(results, q.Fields)
	}

	return &QueryResult{
		Documents: results,
		Total:     len(results),
		Query:     q,
	}, nil
}

// matchesQuery checks if a document matches the query criteria
func (qe *QueryEngine) matchesQuery(doc archive.BSpecDocument, q Query) bool {
	// Type filter
	if q.Type != "" && !strings.EqualFold(doc.Type, q.Type) {
		return false
	}

	// Domain filter
	if q.Domain != "" && !strings.EqualFold(doc.Domain, q.Domain) {
		return false
	}

	// Status filter
	if q.Status != "" && !strings.EqualFold(doc.Status, q.Status) {
		return false
	}

	// Owner filter
	if q.Owner != "" && !strings.Contains(strings.ToLower(doc.Owner), strings.ToLower(q.Owner)) {
		return false
	}

	// Text search filter
	if q.Search != "" {
		searchTerm := strings.ToLower(q.Search)
		searchIn := strings.ToLower(doc.Title + " " + doc.Content)
		if !strings.Contains(searchIn, searchTerm) {
			return false
		}
	}

	// Tags filter (if implemented in metadata)
	if len(q.Tags) > 0 {
		// This would need to be implemented based on how tags are stored in metadata
		// For now, skip this filter
	}

	// Metadata filters
	if len(q.Metadata) > 0 {
		for key, value := range q.Metadata {
			if docValue, exists := doc.Metadata[key]; exists {
				if docValueStr, ok := docValue.(string); ok {
					if !strings.Contains(strings.ToLower(docValueStr), strings.ToLower(value)) {
						return false
					}
				}
			} else {
				return false
			}
		}
	}

	return true
}

// sortDocuments sorts documents by the specified field
func (qe *QueryEngine) sortDocuments(docs []archive.BSpecDocument, sortBy, order string) []archive.BSpecDocument {
	// Simple sorting implementation
	// In a real implementation, you'd use sort.Slice with proper field comparisons

	// For now, return as-is (sorting would be implemented here)
	return docs
}

// selectFields selects only the specified fields from documents
func (qe *QueryEngine) selectFields(docs []archive.BSpecDocument, fields []string) []archive.BSpecDocument {
	// Create field set for quick lookup
	fieldSet := make(map[string]bool)
	for _, field := range fields {
		fieldSet[field] = true
	}

	var result []archive.BSpecDocument
	for _, doc := range docs {
		newDoc := archive.BSpecDocument{}

		// Copy only requested fields
		if fieldSet["id"] || fieldSet["*"] {
			newDoc.ID = doc.ID
		}
		if fieldSet["title"] || fieldSet["*"] {
			newDoc.Title = doc.Title
		}
		if fieldSet["type"] || fieldSet["*"] {
			newDoc.Type = doc.Type
		}
		if fieldSet["status"] || fieldSet["*"] {
			newDoc.Status = doc.Status
		}
		if fieldSet["version"] || fieldSet["*"] {
			newDoc.Version = doc.Version
		}
		if fieldSet["owner"] || fieldSet["*"] {
			newDoc.Owner = doc.Owner
		}
		if fieldSet["created"] || fieldSet["*"] {
			newDoc.Created = doc.Created
		}
		if fieldSet["updated"] || fieldSet["*"] {
			newDoc.Updated = doc.Updated
		}
		if fieldSet["domain"] || fieldSet["*"] {
			newDoc.Domain = doc.Domain
		}
		if fieldSet["content"] || fieldSet["*"] {
			newDoc.Content = doc.Content
		}
		if fieldSet["metadata"] || fieldSet["*"] {
			newDoc.Metadata = doc.Metadata
		}

		result = append(result, newDoc)
	}

	return result
}

// GetDocumentTypes returns all unique document types in the archive
func (qe *QueryEngine) GetDocumentTypes() []string {
	typeSet := make(map[string]bool)
	for _, doc := range qe.archive.Documents {
		if doc.Type != "" {
			typeSet[doc.Type] = true
		}
	}

	var types []string
	for docType := range typeSet {
		types = append(types, docType)
	}
	return types
}

// GetDomains returns all unique domains in the archive
func (qe *QueryEngine) GetDomains() []string {
	domainSet := make(map[string]bool)
	for _, doc := range qe.archive.Documents {
		if doc.Domain != "" {
			domainSet[doc.Domain] = true
		}
	}

	var domains []string
	for domain := range domainSet {
		domains = append(domains, domain)
	}
	return domains
}

// GetOwners returns all unique owners in the archive
func (qe *QueryEngine) GetOwners() []string {
	ownerSet := make(map[string]bool)
	for _, doc := range qe.archive.Documents {
		if doc.Owner != "" {
			ownerSet[doc.Owner] = true
		}
	}

	var owners []string
	for owner := range ownerSet {
		owners = append(owners, owner)
	}
	return owners
}

// ValidateQuery validates a query structure
func ValidateQuery(q Query) error {
	// Sort order validation
	if q.SortOrder != "" && q.SortOrder != "asc" && q.SortOrder != "desc" {
		return fmt.Errorf("invalid sort order: %s (must be 'asc' or 'desc')", q.SortOrder)
	}

	// Limit validation
	if q.Limit < 0 {
		return fmt.Errorf("limit cannot be negative")
	}

	// Search validation (basic regex check)
	if q.Search != "" {
		// Check if it's a valid regex pattern
		_, err := regexp.Compile(q.Search)
		if err != nil {
			// If not a valid regex, treat as literal string (no error)
		}
	}

	return nil
}

// GetStats returns statistics about the archive
func (qe *QueryEngine) GetStats() map[string]interface{} {
	stats := make(map[string]interface{})

	stats["total_documents"] = len(qe.archive.Documents)
	stats["total_assets"] = len(qe.archive.Assets)
	stats["document_types"] = qe.GetDocumentTypes()
	stats["domains"] = qe.GetDomains()
	stats["owners"] = qe.GetOwners()

	// Status distribution
	statusCounts := make(map[string]int)
	for _, doc := range qe.archive.Documents {
		statusCounts[doc.Status]++
	}
	stats["status_distribution"] = statusCounts

	// Type distribution
	typeCounts := make(map[string]int)
	for _, doc := range qe.archive.Documents {
		typeCounts[doc.Type]++
	}
	stats["type_distribution"] = typeCounts

	return stats
}