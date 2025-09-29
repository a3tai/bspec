package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/a3tai/bspec/cli/internal/archive"
)

// packCmd represents the pack command
var packCmd = &cobra.Command{
	Use:   "pack <source-directory> [output-file]",
	Short: "Pack a directory structure into a .bspec file",
	Long: `Pack a directory structure into a .bspec file.

The source directory should contain:
  - manifest.json: Archive metadata (required)
  - documents/: BSpec document files (.md)
  - assets/: Static assets (images, etc.)
  - computed/: Computed analysis files (.json)

Examples:
  bspec pack myproject/                          # Create myproject.bspec
  bspec pack myproject/ project.bspec           # Create project.bspec
  bspec pack ./extracted output.bspec --force   # Overwrite existing file`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		sourceDir := args[0]

		// Determine output file
		var outputFile string
		if len(args) > 1 {
			outputFile = args[1]
		} else {
			// Default: use directory name with .bspec extension
			baseName := filepath.Base(sourceDir)
			if baseName == "." || baseName == ".." {
				baseName = "archive"
			}
			outputFile = baseName + ".bspec"
		}

		// Ensure .bspec extension
		if !strings.HasSuffix(outputFile, ".bspec") {
			outputFile += ".bspec"
		}

		// Check if source directory exists
		if info, err := os.Stat(sourceDir); os.IsNotExist(err) {
			return fmt.Errorf("source directory does not exist: %s", sourceDir)
		} else if !info.IsDir() {
			return fmt.Errorf("source is not a directory: %s", sourceDir)
		}

		// Validate directory structure
		if err := validateSourceDirectory(sourceDir); err != nil {
			return fmt.Errorf("invalid source directory: %w", err)
		}

		// Check if output file exists
		force, _ := cmd.Flags().GetBool("force")
		if _, err := os.Stat(outputFile); err == nil {
			if !force {
				return fmt.Errorf("output file already exists: %s (use --force to overwrite)", outputFile)
			}
		}

		// Pack the archive
		if err := archive.Pack(sourceDir, outputFile); err != nil {
			return fmt.Errorf("failed to pack archive: %w", err)
		}

		fmt.Printf("Successfully packed %s to %s\n", sourceDir, outputFile)

		// Show what was packed
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			if err := showPackingSummary(sourceDir, outputFile); err != nil {
				return fmt.Errorf("failed to show packing summary: %w", err)
			}
		}

		return nil
	},
}

func validateSourceDirectory(sourceDir string) error {
	// Check for required manifest.json
	manifestPath := filepath.Join(sourceDir, "manifest.json")
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		return fmt.Errorf("missing required manifest.json file")
	}

	// Warn about missing optional directories (but don't fail)
	documentsDir := filepath.Join(sourceDir, "documents")
	if _, err := os.Stat(documentsDir); os.IsNotExist(err) {
		fmt.Printf("Warning: documents/ directory not found\n")
	}

	return nil
}

func showPackingSummary(sourceDir, outputFile string) error {
	fmt.Printf("\nPacking Summary:\n")

	// Show source content
	fmt.Printf("Source: %s\n", sourceDir)

	// Check manifest
	manifestPath := filepath.Join(sourceDir, "manifest.json")
	if _, err := os.Stat(manifestPath); err == nil {
		fmt.Printf("  ✓ manifest.json\n")
	}

	// Count documents
	documentsDir := filepath.Join(sourceDir, "documents")
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
	assetsDir := filepath.Join(sourceDir, "assets")
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
	computedDir := filepath.Join(sourceDir, "computed")
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

	// Show output file info
	if info, err := os.Stat(outputFile); err == nil {
		fmt.Printf("\nOutput: %s (%.2f KB)\n", outputFile, float64(info.Size())/1024)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(packCmd)

	// Add flags
	packCmd.Flags().BoolP("force", "f", false, "Overwrite existing output file")
	packCmd.Flags().BoolP("verbose", "v", false, "Show detailed packing summary")
}