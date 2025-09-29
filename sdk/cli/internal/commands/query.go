package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/a3tai/bspec/cli/internal/archive"
	"github.com/a3tai/bspec/cli/internal/output"
	"github.com/a3tai/bspec/cli/internal/query"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query <bspec-file>",
	Short: "Query BSpec documents with structured queries",
	Long: `Query BSpec documents using structured query syntax.

Examples:
  # Find all documents of type MSN
  bspec query project.bspec --type=MSN

  # Find documents by domain
  bspec query project.bspec --domain=strategic

  # Find documents by status
  bspec query project.bspec --status=Draft

  # Find documents by owner
  bspec query project.bspec --owner="John Doe"

  # Text search in document content
  bspec query project.bspec --search="business model"

  # Combine multiple filters
  bspec query project.bspec --type=CAP --domain=product --status=Accepted

  # Select specific fields only
  bspec query project.bspec --fields=id,title,type,status

  # Limit results
  bspec query project.bspec --limit=10

  # Use JSON query (advanced)
  bspec query project.bspec --json='{"type":"MSN","domain":"strategic"}'`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := args[0]

		// Check if path exists
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			return fmt.Errorf("path does not exist: %s", inputPath)
		}

		// Read the archive (either .bspec file or directory)
		arch, err := readArchiveFromPath(inputPath)
		if err != nil {
			return fmt.Errorf("failed to read archive: %w", err)
		}

		// Build query from flags or JSON
		q, err := buildQueryFromFlags(cmd)
		if err != nil {
			return fmt.Errorf("failed to build query: %w", err)
		}

		// Validate query
		if err := query.ValidateQuery(q); err != nil {
			return fmt.Errorf("invalid query: %w", err)
		}

		// Execute query
		qe := query.NewQueryEngine(arch)
		result, err := qe.Execute(q)
		if err != nil {
			return fmt.Errorf("query execution failed: %w", err)
		}

		// Format and display results
		outputFormat := viper.GetString("output")
		prettyFlag, _ := cmd.Flags().GetBool("pretty")

		formatter, err := output.NewFormatter(outputFormat, prettyFlag)
		if err != nil {
			return fmt.Errorf("failed to create formatter: %w", err)
		}

		formattedResult, err := formatter.FormatQueryResult(result)
		if err != nil {
			return fmt.Errorf("failed to format results: %w", err)
		}

		fmt.Print(formattedResult)
		return nil
	},
}

func buildQueryFromFlags(cmd *cobra.Command) (query.Query, error) {
	var q query.Query

	// Check if JSON query is provided
	jsonQuery, _ := cmd.Flags().GetString("json")
	if jsonQuery != "" {
		if err := json.Unmarshal([]byte(jsonQuery), &q); err != nil {
			return q, fmt.Errorf("invalid JSON query: %w", err)
		}
		return q, nil
	}

	// Build query from individual flags
	if cmd.Flags().Changed("type") {
		q.Type, _ = cmd.Flags().GetString("type")
	}

	if cmd.Flags().Changed("domain") {
		q.Domain, _ = cmd.Flags().GetString("domain")
	}

	if cmd.Flags().Changed("status") {
		q.Status, _ = cmd.Flags().GetString("status")
	}

	if cmd.Flags().Changed("owner") {
		q.Owner, _ = cmd.Flags().GetString("owner")
	}

	if cmd.Flags().Changed("search") {
		q.Search, _ = cmd.Flags().GetString("search")
	}

	if cmd.Flags().Changed("tags") {
		tagsStr, _ := cmd.Flags().GetString("tags")
		if tagsStr != "" {
			q.Tags = strings.Split(tagsStr, ",")
			// Trim whitespace
			for i, tag := range q.Tags {
				q.Tags[i] = strings.TrimSpace(tag)
			}
		}
	}

	if cmd.Flags().Changed("fields") {
		fieldsStr, _ := cmd.Flags().GetString("fields")
		if fieldsStr != "" {
			q.Fields = strings.Split(fieldsStr, ",")
			// Trim whitespace
			for i, field := range q.Fields {
				q.Fields[i] = strings.TrimSpace(field)
			}
		}
	}

	if cmd.Flags().Changed("limit") {
		q.Limit, _ = cmd.Flags().GetInt("limit")
	}

	if cmd.Flags().Changed("sort-by") {
		q.SortBy, _ = cmd.Flags().GetString("sort-by")
	}

	if cmd.Flags().Changed("sort-order") {
		q.SortOrder, _ = cmd.Flags().GetString("sort-order")
	}

	// Handle metadata filters
	metadataFlags, _ := cmd.Flags().GetStringSlice("metadata")
	if len(metadataFlags) > 0 {
		q.Metadata = make(map[string]string)
		for _, meta := range metadataFlags {
			parts := strings.SplitN(meta, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				q.Metadata[key] = value
			}
		}
	}

	return q, nil
}

// readArchiveFromPath reads an archive from either a .bspec file or a directory
func readArchiveFromPath(inputPath string) (*archive.BSpecArchive, error) {
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to stat path: %w", err)
	}

	if fileInfo.IsDir() {
		// Read from directory structure
		return readArchiveFromDirectory(inputPath)
	} else {
		// Read from .bspec file
		return archive.Read(inputPath)
	}
}

// readArchiveFromDirectory reads an archive from a directory structure
func readArchiveFromDirectory(dirPath string) (*archive.BSpecArchive, error) {
	// Read manifest.json
	manifestPath := filepath.Join(dirPath, "manifest.json")
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest.json: %w", err)
	}

	var manifest archive.Manifest
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest.json: %w", err)
	}

	arch := &archive.BSpecArchive{
		Manifest:  manifest,
		Documents: make(map[string]archive.BSpecDocument),
		Assets:    make(map[string][]byte),
		Computed:  make(map[string]interface{}),
	}

	// Read documents from documents/ directory
	documentsDir := filepath.Join(dirPath, "documents")
	if _, err := os.Stat(documentsDir); err == nil {
		err := filepath.Walk(documentsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(path, ".md") {
				content, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("failed to read document %s: %w", path, err)
				}

				// Parse frontmatter and content
				doc, err := parseDocumentFromPath(content)
				if err != nil {
					return fmt.Errorf("failed to parse document %s: %w", path, err)
				}

				relPath, _ := filepath.Rel(documentsDir, path)
				arch.Documents[relPath] = *doc
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read documents: %w", err)
		}
	}

	// Read assets from assets/ directory
	assetsDir := filepath.Join(dirPath, "assets")
	if _, err := os.Stat(assetsDir); err == nil {
		err := filepath.Walk(assetsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				content, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("failed to read asset %s: %w", path, err)
				}

				relPath, _ := filepath.Rel(assetsDir, path)
				arch.Assets[relPath] = content
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read assets: %w", err)
		}
	}

	// Read computed data from computed/ directory
	computedDir := filepath.Join(dirPath, "computed")
	if _, err := os.Stat(computedDir); err == nil {
		err := filepath.Walk(computedDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.HasSuffix(path, ".json") {
				content, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("failed to read computed file %s: %w", path, err)
				}

				var data interface{}
				if err := json.Unmarshal(content, &data); err != nil {
					return fmt.Errorf("failed to parse computed file %s: %w", path, err)
				}

				relPath, _ := filepath.Rel(computedDir, path)
				arch.Computed[strings.TrimSuffix(relPath, ".json")] = data
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read computed data: %w", err)
		}
	}

	return arch, nil
}

// parseDocumentFromPath parses a markdown document with YAML frontmatter
func parseDocumentFromPath(content []byte) (*archive.BSpecDocument, error) {
	contentStr := string(content)

	// Check for frontmatter
	if !strings.HasPrefix(contentStr, "---\n") {
		return nil, fmt.Errorf("document missing YAML frontmatter")
	}

	// Find the end of frontmatter
	parts := strings.SplitN(contentStr[4:], "\n---\n", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid frontmatter format")
	}

	frontmatter := parts[0]
	markdownContent := parts[1]

	// Parse frontmatter as JSON for simplicity (in real implementation, use YAML parser)
	var metadata map[string]interface{}

	// For now, create a basic document structure
	// In a real implementation, you'd use a proper YAML parser
	doc := &archive.BSpecDocument{
		Content:  markdownContent,
		Metadata: metadata,
	}

	// Extract basic fields from frontmatter (simplified parsing)
	lines := strings.Split(frontmatter, "\n")
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(strings.Trim(parts[1], "\"'"))

				switch key {
				case "id":
					doc.ID = value
				case "title":
					doc.Title = value
				case "type":
					doc.Type = value
				case "status":
					doc.Status = value
				case "version":
					doc.Version = value
				case "owner":
					doc.Owner = value
				case "created":
					doc.Created = value
				case "updated":
					doc.Updated = value
				case "domain":
					doc.Domain = value
				}
			}
		}
	}

	return doc, nil
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Query filter flags
	queryCmd.Flags().StringP("type", "t", "", "Filter by document type")
	queryCmd.Flags().StringP("domain", "d", "", "Filter by domain")
	queryCmd.Flags().StringP("status", "", "", "Filter by status")
	queryCmd.Flags().StringP("owner", "", "", "Filter by owner")
	queryCmd.Flags().StringP("search", "", "", "Text search in content")
	queryCmd.Flags().StringP("tags", "", "", "Filter by tags (comma-separated)")

	// Output control flags
	queryCmd.Flags().StringP("fields", "f", "", "Select specific fields (comma-separated)")
	queryCmd.Flags().IntP("limit", "l", 0, "Limit number of results")

	// Sorting flags
	queryCmd.Flags().StringP("sort-by", "", "", "Sort by field")
	queryCmd.Flags().StringP("sort-order", "", "asc", "Sort order (asc|desc)")

	// Advanced flags
	queryCmd.Flags().StringSliceP("metadata", "m", []string{}, "Filter by metadata (key=value)")
	queryCmd.Flags().StringP("json", "j", "", "JSON query string (advanced)")
	queryCmd.Flags().BoolP("pretty", "p", true, "Pretty print output")
}