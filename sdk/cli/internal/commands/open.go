package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/a3tai/bspec/cli/internal/archive"
	"github.com/a3tai/bspec/cli/internal/output"
	"github.com/a3tai/bspec/cli/internal/query"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open <bspec-file>",
	Short: "Open and display information about a .bspec file",
	Long: `Open a .bspec file and display its contents or information.

Examples:
  bspec open myproject.bspec                    # Show archive info
  bspec open myproject.bspec --info             # Show detailed archive info
  bspec open myproject.bspec --list             # List all documents
  bspec open myproject.bspec --stats            # Show statistics`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		bspecFile := args[0]

		// Check if file exists
		if _, err := os.Stat(bspecFile); os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", bspecFile)
		}

		// Read the archive
		arch, err := archive.Read(bspecFile)
		if err != nil {
			return fmt.Errorf("failed to read .bspec file: %w", err)
		}

		// Get output format
		outputFormat := viper.GetString("output")
		prettyFlag, _ := cmd.Flags().GetBool("pretty")

		formatter, err := output.NewFormatter(outputFormat, prettyFlag)
		if err != nil {
			return fmt.Errorf("failed to create formatter: %w", err)
		}

		// Determine what to display
		showInfo, _ := cmd.Flags().GetBool("info")
		showList, _ := cmd.Flags().GetBool("list")
		showStats, _ := cmd.Flags().GetBool("stats")

		switch {
		case showStats:
			return displayStats(arch, formatter)
		case showList:
			return displayDocumentList(arch, formatter)
		case showInfo:
			return displayArchiveInfo(arch, formatter)
		default:
			// Default: show basic info
			return displayArchiveInfo(arch, formatter)
		}
	},
}

func displayArchiveInfo(arch *archive.BSpecArchive, formatter *output.Formatter) error {
	result, err := formatter.FormatArchiveInfo(arch)
	if err != nil {
		return fmt.Errorf("failed to format archive info: %w", err)
	}

	fmt.Print(result)
	return nil
}

func displayDocumentList(arch *archive.BSpecArchive, formatter *output.Formatter) error {
	// Create a simple list of documents
	docs := make([]archive.BSpecDocument, 0, len(arch.Documents))
	for _, doc := range arch.Documents {
		docs = append(docs, doc)
	}

	// Format as a query result for consistent output
	queryResult := &query.QueryResult{
		Documents: docs,
		Total:     len(docs),
	}

	result, err := formatter.FormatQueryResult(queryResult)
	if err != nil {
		return fmt.Errorf("failed to format document list: %w", err)
	}

	fmt.Print(result)
	return nil
}

func displayStats(arch *archive.BSpecArchive, formatter *output.Formatter) error {
	// Create query engine to get stats
	qe := query.NewQueryEngine(arch)
	stats := qe.GetStats()

	result, err := formatter.FormatStats(stats)
	if err != nil {
		return fmt.Errorf("failed to format stats: %w", err)
	}

	fmt.Print(result)
	return nil
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Add flags
	openCmd.Flags().BoolP("info", "i", false, "Show detailed archive information")
	openCmd.Flags().BoolP("list", "l", false, "List all documents in the archive")
	openCmd.Flags().BoolP("stats", "s", false, "Show archive statistics")
	openCmd.Flags().BoolP("pretty", "p", true, "Pretty print output (JSON only)")
}