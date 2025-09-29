package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/a3tai/bspec/cli/internal/archive"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract <bspec-file> [output-directory]",
	Short: "Extract a .bspec file to a directory structure",
	Long: `Extract a .bspec file to a directory structure on disk.

The extracted directory will contain:
  - manifest.json: Archive metadata
  - documents/: BSpec document files (.md)
  - assets/: Static assets (images, etc.)
  - computed/: Computed analysis files (.json)

Examples:
  bspec extract project.bspec                    # Extract to project/
  bspec extract project.bspec ./extracted       # Extract to ./extracted/
  bspec extract project.bspec --force           # Overwrite existing directory`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		bspecFile := args[0]

		// Determine output directory
		var outputDir string
		if len(args) > 1 {
			outputDir = args[1]
		} else {
			// Default: use base name without extension
			baseName := filepath.Base(bspecFile)
			if ext := filepath.Ext(baseName); ext != "" {
				baseName = baseName[:len(baseName)-len(ext)]
			}
			outputDir = baseName
		}

		// Check if input file exists
		if _, err := os.Stat(bspecFile); os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", bspecFile)
		}

		// Check if output directory exists
		force, _ := cmd.Flags().GetBool("force")
		if _, err := os.Stat(outputDir); err == nil {
			if !force {
				return fmt.Errorf("output directory already exists: %s (use --force to overwrite)", outputDir)
			}
			// Remove existing directory
			if err := os.RemoveAll(outputDir); err != nil {
				return fmt.Errorf("failed to remove existing directory: %w", err)
			}
		}

		// Extract the archive
		if err := archive.Extract(bspecFile, outputDir); err != nil {
			return fmt.Errorf("failed to extract archive: %w", err)
		}

		fmt.Printf("Successfully extracted %s to %s\n", bspecFile, outputDir)

		// Show what was extracted
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			if err := showExtractionSummary(outputDir); err != nil {
				return fmt.Errorf("failed to show extraction summary: %w", err)
			}
		}

		return nil
	},
}

func showExtractionSummary(outputDir string) error {
	fmt.Printf("\nExtraction Summary:\n")

	// Check manifest
	manifestPath := filepath.Join(outputDir, "manifest.json")
	if _, err := os.Stat(manifestPath); err == nil {
		fmt.Printf("  ✓ manifest.json\n")
	}

	// Count documents
	documentsDir := filepath.Join(outputDir, "documents")
	if info, err := os.Stat(documentsDir); err == nil && info.IsDir() {
		docCount := 0
		err := filepath.Walk(documentsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".md" {
				docCount++
			}
			return nil
		})
		if err != nil {
			return err
		}
		fmt.Printf("  ✓ documents/ (%d files)\n", docCount)
	}

	// Count assets
	assetsDir := filepath.Join(outputDir, "assets")
	if info, err := os.Stat(assetsDir); err == nil && info.IsDir() {
		assetCount := 0
		err := filepath.Walk(assetsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				assetCount++
			}
			return nil
		})
		if err != nil {
			return err
		}
		fmt.Printf("  ✓ assets/ (%d files)\n", assetCount)
	}

	// Count computed files
	computedDir := filepath.Join(outputDir, "computed")
	if info, err := os.Stat(computedDir); err == nil && info.IsDir() {
		computedCount := 0
		err := filepath.Walk(computedDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".json" {
				computedCount++
			}
			return nil
		})
		if err != nil {
			return err
		}
		fmt.Printf("  ✓ computed/ (%d files)\n", computedCount)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// Add flags
	extractCmd.Flags().BoolP("force", "f", false, "Overwrite existing output directory")
	extractCmd.Flags().BoolP("verbose", "v", false, "Show detailed extraction summary")
}