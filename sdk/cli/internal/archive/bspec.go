package archive

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// BSpecArchive represents a .bspec file structure
type BSpecArchive struct {
	Manifest  Manifest                   `json:"manifest"`
	Documents map[string]BSpecDocument   `json:"documents"`
	Assets    map[string][]byte          `json:"assets"`
	Computed  map[string]interface{}     `json:"computed"`
}

// Manifest represents the manifest.json structure in a .bspec file
type Manifest struct {
	FormatVersion   string    `json:"format_version"`
	BSpecVersion    string    `json:"bspec_version"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Author          string    `json:"author"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	TotalDocuments  int       `json:"total_documents"`
	DocumentTypes   []string  `json:"document_types"`
	Domains         []string  `json:"domains"`
	ConformanceLevel string   `json:"conformance_level"`
	IndustryProfile string    `json:"industry_profile"`
}

// BSpecDocument represents a BSpec document with frontmatter and content
type BSpecDocument struct {
	ID       string                 `json:"id"`
	Title    string                 `json:"title"`
	Type     string                 `json:"type"`
	Status   string                 `json:"status"`
	Version  string                 `json:"version"`
	Owner    string                 `json:"owner"`
	Created  string                 `json:"created"`
	Updated  string                 `json:"updated"`
	Domain   string                 `json:"domain,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Content  string                 `json:"content"`
}

// Extract extracts a .bspec file to a directory structure
func Extract(bspecPath, outputDir string) error {
	// Open the .bspec file
	file, err := os.Open(bspecPath)
	if err != nil {
		return fmt.Errorf("failed to open .bspec file: %w", err)
	}
	defer file.Close()

	// Create gzip reader
	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzReader.Close()

	// Create tar reader
	tarReader := tar.NewReader(gzReader)

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Extract files
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar header: %w", err)
		}

		// Get the file path
		targetPath := filepath.Join(outputDir, header.Name)

		// Ensure the directory exists
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Extract the file
		switch header.Typeflag {
		case tar.TypeReg:
			// Regular file
			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", targetPath, err)
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("failed to copy file content: %w", err)
			}

			// Set file permissions
			if err := os.Chmod(targetPath, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to set file permissions: %w", err)
			}

		case tar.TypeDir:
			// Directory
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		}
	}

	return nil
}

// Pack creates a .bspec file from a directory structure
func Pack(sourceDir, bspecPath string) error {
	// Create the .bspec file
	file, err := os.Create(bspecPath)
	if err != nil {
		return fmt.Errorf("failed to create .bspec file: %w", err)
	}
	defer file.Close()

	// Create gzip writer
	gzWriter := gzip.NewWriter(file)
	defer gzWriter.Close()

	// Create tar writer
	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	// Walk the source directory
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == sourceDir {
			return nil
		}

		// Get relative path
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Create tar header
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return fmt.Errorf("failed to create tar header: %w", err)
		}

		// Set the name in the header
		header.Name = relPath

		// Write the header
		if err := tarWriter.WriteHeader(header); err != nil {
			return fmt.Errorf("failed to write tar header: %w", err)
		}

		// If it's a regular file, write the content
		if info.Mode().IsRegular() {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to open file: %w", err)
			}
			defer file.Close()

			if _, err := io.Copy(tarWriter, file); err != nil {
				return fmt.Errorf("failed to copy file content: %w", err)
			}
		}

		return nil
	})
}

// Read reads and parses a .bspec file into a BSpecArchive structure
func Read(bspecPath string) (*BSpecArchive, error) {
	// Create temporary directory for extraction
	tempDir, err := os.MkdirTemp("", "bspec-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Extract the archive
	if err := Extract(bspecPath, tempDir); err != nil {
		return nil, fmt.Errorf("failed to extract archive: %w", err)
	}

	// Read the manifest
	manifestPath := filepath.Join(tempDir, "manifest.json")
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest: %w", err)
	}

	var manifest Manifest
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	archive := &BSpecArchive{
		Manifest:  manifest,
		Documents: make(map[string]BSpecDocument),
		Assets:    make(map[string][]byte),
		Computed:  make(map[string]interface{}),
	}

	// Read documents
	documentsDir := filepath.Join(tempDir, "documents")
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
				doc, err := parseDocument(content)
				if err != nil {
					return fmt.Errorf("failed to parse document %s: %w", path, err)
				}

				relPath, _ := filepath.Rel(documentsDir, path)
				archive.Documents[relPath] = *doc
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read documents: %w", err)
		}
	}

	// Read assets
	assetsDir := filepath.Join(tempDir, "assets")
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
				archive.Assets[relPath] = content
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read assets: %w", err)
		}
	}

	// Read computed data
	computedDir := filepath.Join(tempDir, "computed")
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
				archive.Computed[strings.TrimSuffix(relPath, ".json")] = data
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read computed data: %w", err)
		}
	}

	return archive, nil
}

// parseDocument parses a markdown document with YAML frontmatter
func parseDocument(content []byte) (*BSpecDocument, error) {
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
	doc := &BSpecDocument{
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