package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/a3tai/bspec/cli/internal/archive"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initialize a new BSpec project structure",
	Long: `Initialize a new BSpec project with the standard directory structure.

This creates:
  - manifest.json: Project metadata
  - documents/: Directory for BSpec documents
  - assets/: Directory for static assets
  - computed/: Directory for computed analysis files
  - README.md: Project documentation

Examples:
  bspec init myproject                           # Create myproject/ directory
  bspec init myproject --author "John Doe"      # Set author
  bspec init myproject --conformance bronze     # Set conformance level
  bspec init myproject --industry software-saas # Set industry profile`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		// Check if directory already exists
		if _, err := os.Stat(projectName); err == nil {
			force, _ := cmd.Flags().GetBool("force")
			if !force {
				return fmt.Errorf("directory already exists: %s (use --force to overwrite)", projectName)
			}
			// Remove existing directory
			if err := os.RemoveAll(projectName); err != nil {
				return fmt.Errorf("failed to remove existing directory: %w", err)
			}
		}

		// Create project directory
		if err := os.MkdirAll(projectName, 0755); err != nil {
			return fmt.Errorf("failed to create project directory: %w", err)
		}

		// Get options from flags
		author, _ := cmd.Flags().GetString("author")
		description, _ := cmd.Flags().GetString("description")
		conformance, _ := cmd.Flags().GetString("conformance")
		industry, _ := cmd.Flags().GetString("industry")

		// Set defaults
		if author == "" {
			author = "Unknown"
		}
		if description == "" {
			description = fmt.Sprintf("BSpec project: %s", projectName)
		}
		if conformance == "" {
			conformance = "bronze"
		}
		if industry == "" {
			industry = "software-saas"
		}

		// Create directory structure
		dirs := []string{
			"documents",
			"assets",
			"computed",
		}

		for _, dir := range dirs {
			dirPath := filepath.Join(projectName, dir)
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dir, err)
			}
		}

		// Create manifest.json
		if err := createManifest(projectName, projectName, description, author, conformance, industry); err != nil {
			return fmt.Errorf("failed to create manifest: %w", err)
		}

		// Create sample documents
		createSamples, _ := cmd.Flags().GetBool("samples")
		if createSamples {
			if err := createSampleDocuments(projectName); err != nil {
				return fmt.Errorf("failed to create sample documents: %w", err)
			}
		}

		// Create README.md
		if err := createReadme(projectName, projectName, description); err != nil {
			return fmt.Errorf("failed to create README: %w", err)
		}

		// Create .gitignore
		if err := createGitignore(projectName); err != nil {
			return fmt.Errorf("failed to create .gitignore: %w", err)
		}

		fmt.Printf("Successfully initialized BSpec project: %s\n", projectName)

		// Show what was created
		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			if err := showInitSummary(projectName); err != nil {
				return fmt.Errorf("failed to show init summary: %w", err)
			}
		}

		return nil
	},
}

func createManifest(projectDir, name, description, author, conformance, industry string) error {
	manifest := archive.Manifest{
		FormatVersion:    "1.0.0",
		BSpecVersion:     "1.0.0",
		Name:             name,
		Description:      description,
		Author:           author,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		TotalDocuments:   0,
		DocumentTypes:    []string{},
		Domains:          []string{},
		ConformanceLevel: conformance,
		IndustryProfile:  industry,
	}

	manifestPath := filepath.Join(projectDir, "manifest.json")
	manifestData, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal manifest: %w", err)
	}

	return os.WriteFile(manifestPath, manifestData, 0644)
}

func createSampleDocuments(projectDir string) error {
	samples := []struct {
		filename string
		content  string
	}{
		{
			filename: "MSN-mission-v1.0.0.md",
			content: `---
id: msn-mission-001
title: Organization Mission Statement
type: MSN
status: Draft
version: 1.0.0
owner: Leadership Team
created: ` + time.Now().Format("2006-01-02") + `
updated: ` + time.Now().Format("2006-01-02") + `
domain: strategic
---

# Organization Mission Statement

## Purpose

Define the core purpose and mission of the organization.

## Mission

[Define your organization's mission here]

## Vision

[Define your organization's vision here]

## Values

[List your organization's core values here]
`,
		},
		{
			filename: "VSN-vision-v1.0.0.md",
			content: `---
id: vsn-vision-001
title: Organization Vision
type: VSN
status: Draft
version: 1.0.0
owner: Leadership Team
created: ` + time.Now().Format("2006-01-02") + `
updated: ` + time.Now().Format("2006-01-02") + `
domain: strategic
related:
  - msn-mission-001
---

# Organization Vision

## Vision Statement

[Define your organization's vision statement here]

## Strategic Goals

[List key strategic goals that support the vision]

## Success Metrics

[Define how success will be measured]
`,
		},
	}

	documentsDir := filepath.Join(projectDir, "documents")
	for _, sample := range samples {
		filePath := filepath.Join(documentsDir, sample.filename)
		if err := os.WriteFile(filePath, []byte(sample.content), 0644); err != nil {
			return fmt.Errorf("failed to create sample document %s: %w", sample.filename, err)
		}
	}

	return nil
}

func createReadme(projectDir, name, description string) error {
	content := fmt.Sprintf("# %s\n\n%s\n\n## About\n\nThis is a BSpec (Business Specification Standard) project that documents business requirements, architecture, and processes.\n\n## Structure\n\n- **documents/**: BSpec document files (.md)\n- **assets/**: Static assets (images, diagrams, etc.)\n- **computed/**: Computed analysis files (.json)\n- **manifest.json**: Project metadata\n\n## Getting Started\n\n1. Add your BSpec documents to the documents/ directory\n2. Include any assets in the assets/ directory\n3. Use the BSpec CLI to query and manage documents:\n\n```bash\n# Query documents\nbspec query . --type=MSN\n\n# Pack into .bspec file\nbspec pack . %s.bspec\n\n# Extract from .bspec file\nbspec extract %s.bspec\n```\n\n## Documentation\n\nFor more information about BSpec, visit: https://bspec.dev\n\n## License\n\n[Add your license information here]\n", name, description, name, name)

	readmePath := filepath.Join(projectDir, "README.md")
	return os.WriteFile(readmePath, []byte(content), 0644)
}

func createGitignore(projectDir string) error {
	content := `# BSpec generated files
*.bspec
computed/

# OS files
.DS_Store
Thumbs.db

# Editor files
.vscode/
.idea/
*.swp
*.swo
*~

# Logs
*.log

# Temporary files
*.tmp
*.temp
`

	gitignorePath := filepath.Join(projectDir, ".gitignore")
	return os.WriteFile(gitignorePath, []byte(content), 0644)
}

func showInitSummary(projectDir string) error {
	fmt.Printf("\nProject Structure:\n")
	fmt.Printf("%s/\n", projectDir)
	fmt.Printf("├── manifest.json\n")
	fmt.Printf("├── README.md\n")
	fmt.Printf("├── .gitignore\n")
	fmt.Printf("├── documents/\n")

	// Check for sample documents
	documentsDir := filepath.Join(projectDir, "documents")
	files, err := os.ReadDir(documentsDir)
	if err == nil {
		for _, file := range files {
			if !file.IsDir() {
				fmt.Printf("│   ├── %s\n", file.Name())
			}
		}
	}

	fmt.Printf("├── assets/\n")
	fmt.Printf("└── computed/\n")

	fmt.Printf("\nNext Steps:\n")
	fmt.Printf("1. cd %s\n", projectDir)
	fmt.Printf("2. Edit documents in the documents/ directory\n")
	fmt.Printf("3. Use 'bspec query .' to explore your documents\n")
	fmt.Printf("4. Use 'bspec pack .' to create a .bspec archive\n")

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Add flags
	initCmd.Flags().StringP("author", "a", "", "Project author")
	initCmd.Flags().StringP("description", "d", "", "Project description")
	initCmd.Flags().StringP("conformance", "c", "bronze", "Conformance level (bronze|silver|gold)")
	initCmd.Flags().StringP("industry", "i", "software-saas", "Industry profile (software-saas|physical-product|service-business|nonprofit)")
	initCmd.Flags().BoolP("samples", "s", false, "Create sample documents")
	initCmd.Flags().BoolP("force", "f", false, "Overwrite existing directory")
	initCmd.Flags().BoolP("verbose", "v", false, "Show detailed summary")
}