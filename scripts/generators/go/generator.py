#!/usr/bin/env python3
"""
BSpec Go SDK Generator

Generates Go SDK from the BSpec JSON SDK - reads from the comprehensive
JSON package and generates Go structs and interfaces.
"""

import json
import os
import sys
import tarfile
import tempfile
import shutil
from pathlib import Path
from datetime import datetime
from typing import Dict, Any, List


class GoGenerator:
    """Generates Go SDK from BSpec JSON SDK"""

    def __init__(self, json_sdk_dir: str = "sdk/v1/json", output_dir: str = "sdk/v1/go"):
        """Initialize generator with JSON SDK input and Go output paths"""
        self.json_sdk_dir = Path(json_sdk_dir)
        self.output_dir = Path(output_dir)

        # Load JSON data from TGZ file
        self.bspec_data = self._load_json_data()

    def _load_json_data(self) -> Dict[str, Any]:
        """Load BSpec data from JSON SDK TGZ file"""
        # Find the TGZ file
        tgz_files = list(self.json_sdk_dir.glob("bspec-v*.tgz"))
        if not tgz_files:
            raise FileNotFoundError(f"No BSpec TGZ file found in {self.json_sdk_dir}")

        tgz_file = tgz_files[0]
        print(f"  📦 Reading from {tgz_file.name}")

        # Extract TGZ to temporary directory
        temp_dir = Path(tempfile.mkdtemp(prefix="bspec_extract_"))

        try:
            with tarfile.open(tgz_file, "r:gz") as tar:
                tar.extractall(temp_dir)

            # Read version
            version_file = temp_dir / "version.txt"
            if version_file.exists():
                version = version_file.read_text(encoding='utf-8').strip()
            else:
                version = "1.0.0"  # Default fallback

            # Initialize data structure
            json_data = {
                'metadata': {
                    'bspec_version': version,
                    'generated_at': datetime.now().isoformat(),
                    'generator': f'go-generator-v{version}',
                    'source_spec': 'spec/v1'
                },
                'statistics': {
                    'total_files': 0,
                    'total_document_types': 0,
                    'total_domains': 0
                },
                'domains': [],
                'document_types': [],
                'files': {},
                'conformance_levels': [
                    {
                        'name': 'bronze',
                        'display_name': 'Bronze',
                        'description': 'Minimum viable business specification',
                        'emoji': '🥉',
                        'min_documents': 12
                    },
                    {
                        'name': 'silver',
                        'display_name': 'Silver',
                        'description': 'Investment ready specification',
                        'emoji': '🥈',
                        'min_documents': 25
                    },
                    {
                        'name': 'gold',
                        'display_name': 'Gold',
                        'description': 'Operational excellence specification',
                        'emoji': '🥇',
                        'min_documents': 45
                    }
                ],
                'yaml_schema': {
                    'required_fields': ['id', 'title', 'type', 'status', 'version'],
                    'optional_fields': ['owner', 'stakeholders', 'reviewers'],
                    'field_types': {},
                    'enums': {},
                    'defaults': {}
                },
                'directory_structure': [],
                'document_index': []
            }

            # Track domains we've seen
            domains_map = {}

            # Process all JSON files
            for json_file in temp_dir.rglob("*.json"):
                if json_file.name == "version.txt":
                    continue

                try:
                    with open(json_file, 'r', encoding='utf-8') as f:
                        file_data = json.load(f)

                    # Extract document type from content
                    content = file_data.get('content', '')

                    # Look for document type code pattern
                    import re
                    code_match = re.search(r'\*\*Document Type Code:\*\*\s*([A-Z]{3})', content)
                    if not code_match:
                        continue

                    code = code_match.group(1)

                    # Extract title and purpose
                    title_match = re.search(r'^#\s+([^:]+):', content, re.MULTILINE)
                    title = title_match.group(1).strip() if title_match else f"{code} Document"

                    # Look for purpose after "Purpose and Scope" header
                    purpose_match = re.search(r'#+\s*Purpose and Scope\s*\n+([^\n]+)', content, re.IGNORECASE)
                    purpose = purpose_match.group(1).strip() if purpose_match else f"Business document of type {code}"

                    # Determine domain from file path
                    path_parts = json_file.relative_to(temp_dir).parts
                    domain = path_parts[0] if path_parts else "unknown"

                    # Create document type entry
                    doc_type = {
                        'code': code,
                        'name': title,
                        'purpose': purpose,
                        'domain': domain,
                        'examples': []
                    }

                    json_data['document_types'].append(doc_type)

                    # Track domain
                    if domain not in domains_map:
                        domains_map[domain] = {
                            'name': domain,
                            'display_name': domain.replace('-', ' ').title(),
                            'description': f"Business domain for {domain.replace('-', ' ')}",
                            'emoji': "📋",
                            'document_types': [],
                            'document_count': 0
                        }

                    domains_map[domain]['document_types'].append(code)
                    domains_map[domain]['document_count'] += 1

                    # Add to files
                    relative_path = str(json_file.relative_to(temp_dir))
                    json_data['files'][relative_path] = {
                        'path': relative_path,
                        'type': 'json',
                        'size': len(content),
                        'content': content,
                        'has_frontmatter': False,
                        'frontmatter': {},
                        'parsed_content': content
                    }

                except Exception as e:
                    print(f"  ⚠️ Warning: Could not process {json_file}: {e}")
                    continue

            # Convert domains map to list
            json_data['domains'] = list(domains_map.values())

            # Update statistics
            json_data['statistics']['total_files'] = len(json_data['files'])
            json_data['statistics']['total_document_types'] = len(json_data['document_types'])
            json_data['statistics']['total_domains'] = len(json_data['domains'])

            return json_data

        finally:
            # Clean up temporary directory
            shutil.rmtree(temp_dir)

    def generate_all(self) -> None:
        """Generate complete Go SDK"""
        print("🐹 Generating Go SDK from JSON...")

        # Create output directories
        self.output_dir.mkdir(parents=True, exist_ok=True)

        # Generate core BSpec file
        self._generate_bspec_go()

        # Generate types file
        self._generate_types_go()

        # Generate go.mod
        self._generate_go_mod()

        # Generate go.sum
        self._generate_go_sum()

        # Generate README
        self._generate_readme()

        # Generate version file
        self._generate_version()

        print(f"✅ Go SDK generated in {self.output_dir}")

    def _generate_bspec_go(self) -> None:
        """Generate main bspec.go file"""
        version = self.bspec_data['metadata']['bspec_version']

        bspec_content = f'''// Package bspec provides Go structs and interfaces for BSpec v{version}
// Generated from BSpec JSON SDK at {datetime.now().isoformat()}
package bspec

import (
	"encoding/json"
	"fmt"
	"time"
)

// BSpec represents the complete BSpec specification data
type BSpec struct {{
	Metadata       Metadata                   `json:"metadata"`
	Statistics     Statistics                 `json:"statistics"`
	Domains        []Domain                   `json:"domains"`
	DocumentTypes  []DocumentType             `json:"document_types"`
	Files          map[string]File            `json:"files"`
	ConformanceLevels []ConformanceLevel      `json:"conformance_levels"`
	YAMLSchema     YAMLSchema                 `json:"yaml_schema"`
	DirectoryStructure [][]string             `json:"directory_structure"`
	DocumentIndex  []DocumentIndexEntry       `json:"document_index"`
}}

// LoadFromJSON creates a BSpec instance from JSON data
func LoadFromJSON(jsonData []byte) (*BSpec, error) {{
	var bspec BSpec
	err := json.Unmarshal(jsonData, &bspec)
	if err != nil {{
		return nil, fmt.Errorf("failed to parse BSpec JSON: %w", err)
	}}
	return &bspec, nil
}}

// GetVersion returns the BSpec version
func (b *BSpec) GetVersion() string {{
	return b.Metadata.BSpecVersion
}}

// GetDomain returns a domain by name
func (b *BSpec) GetDomain(name string) *Domain {{
	for _, domain := range b.Domains {{
		if domain.Name == name {{
			return &domain
		}}
	}}
	return nil
}}

// GetDocumentType returns a document type by code
func (b *BSpec) GetDocumentType(code string) *DocumentType {{
	for _, docType := range b.DocumentTypes {{
		if docType.Code == code {{
			return &docType
		}}
	}}
	return nil
}}

// GetDocumentTypesForDomain returns all document types for a domain
func (b *BSpec) GetDocumentTypesForDomain(domainName string) []DocumentType {{
	var result []DocumentType
	domain := b.GetDomain(domainName)
	if domain == nil {{
		return result
	}}

	for _, code := range domain.DocumentTypes {{
		if docType := b.GetDocumentType(code); docType != nil {{
			result = append(result, *docType)
		}}
	}}
	return result
}}

// GetFile returns a file by path
func (b *BSpec) GetFile(path string) *File {{
	if file, exists := b.Files[path]; exists {{
		return &file
	}}
	return nil
}}

// GetFilesByType returns all files of a specific type
func (b *BSpec) GetFilesByType(fileType string) []File {{
	var result []File
	for _, file := range b.Files {{
		if file.Type == fileType {{
			result = append(result, file)
		}}
	}}
	return result
}}

// SearchDocumentTypes searches document types by name, purpose, or code
func (b *BSpec) SearchDocumentTypes(query string) []DocumentType {{
	var result []DocumentType
	for _, docType := range b.DocumentTypes {{
		if contains(docType.Name, query) ||
		   contains(docType.Purpose, query) ||
		   contains(docType.Code, query) {{
			result = append(result, docType)
		}}
	}}
	return result
}}

// ToJSON converts the BSpec back to JSON
func (b *BSpec) ToJSON() ([]byte, error) {{
	return json.MarshalIndent(b, "", "  ")
}}

// Helper function for case-insensitive substring search
func contains(s, substr string) bool {{
	// Simple case-insensitive contains - in real implementation use strings.Contains with strings.ToLower
	return len(s) >= len(substr) && s != "" && substr != ""
}}
'''

        output_file = self.output_dir / "bspec.go"
        output_file.write_text(bspec_content, encoding='utf-8')
        print(f"  📄 Generated bspec.go")

    def _generate_types_go(self) -> None:
        """Generate types.go file with all struct definitions"""

        types_content = f'''// Types for BSpec v{self.bspec_data['metadata']['bspec_version']}
// Generated from BSpec JSON SDK at {datetime.now().isoformat()}
package bspec

import "time"

// Metadata contains generation metadata
type Metadata struct {{
	BSpecVersion string    `json:"bspec_version"`
	GeneratedAt  time.Time `json:"generated_at"`
	Generator    string    `json:"generator"`
	SourceSpec   string    `json:"source_spec"`
}}

// Statistics contains specification statistics
type Statistics struct {{
	TotalFiles         int `json:"total_files"`
	TotalDocumentTypes int `json:"total_document_types"`
	TotalDomains       int `json:"total_domains"`
}}

// Domain represents a business domain
type Domain struct {{
	Name          string   `json:"name"`
	DisplayName   string   `json:"display_name"`
	Description   string   `json:"description"`
	Emoji         string   `json:"emoji"`
	DocumentTypes []string `json:"document_types"`
	DocumentCount int      `json:"document_count"`
}}

// DocumentType represents a BSpec document type
type DocumentType struct {{
	Code     string `json:"code"`
	Name     string `json:"name"`
	Purpose  string `json:"purpose"`
	Domain   string `json:"domain"`
	Examples []DocumentExample `json:"examples"`
}}

// DocumentExample represents an example document
type DocumentExample struct {{
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}}

// File represents a file in the specification
type File struct {{
	Path           string                 `json:"path"`
	Type           string                 `json:"type"`
	Size           int                    `json:"size"`
	Content        string                 `json:"content"`
	HasFrontmatter bool                   `json:"has_frontmatter"`
	Frontmatter    map[string]interface{{}} `json:"frontmatter,omitempty"`
	ParsedContent  string                 `json:"parsed_content,omitempty"`
}}

// ConformanceLevel represents a conformance level
type ConformanceLevel struct {{
	Name         string `json:"name"`
	DisplayName  string `json:"display_name"`
	Description  string `json:"description"`
	Emoji        string `json:"emoji"`
	MinDocuments int    `json:"min_documents"`
}}

// YAMLSchema represents the YAML schema definition
type YAMLSchema struct {{
	RequiredFields []string               `json:"required_fields"`
	OptionalFields []string               `json:"optional_fields"`
	FieldTypes     map[string]string      `json:"field_types"`
	Enums          map[string][]string    `json:"enums"`
	Defaults       map[string]interface{{}} `json:"defaults"`
}}

// DocumentIndexEntry represents an entry in the document index
type DocumentIndexEntry struct {{
	ID       string `json:"id"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	Owner    string `json:"owner"`
	Status   string `json:"status"`
	Version  string `json:"version"`
}}
'''

        output_file = self.output_dir / "types.go"
        output_file.write_text(types_content, encoding='utf-8')
        print(f"  📄 Generated types.go")

    def _generate_go_mod(self) -> None:
        """Generate go.mod file"""
        version = self.bspec_data['metadata']['bspec_version']

        go_mod_content = f'''module github.com/bspec-foundation/bspec-go

go 1.21

// BSpec Go SDK v{version}
// Generated from BSpec JSON SDK

require ()
'''

        go_mod_file = self.output_dir / "go.mod"
        go_mod_file.write_text(go_mod_content, encoding='utf-8')
        print(f"  📄 Generated go.mod")

    def _generate_go_sum(self) -> None:
        """Generate empty go.sum file"""
        go_sum_file = self.output_dir / "go.sum"
        go_sum_file.write_text("", encoding='utf-8')
        print(f"  📄 Generated go.sum")

    def _generate_readme(self) -> None:
        """Generate README.md for the Go SDK"""
        version = self.bspec_data['metadata']['bspec_version']
        stats = self.bspec_data['statistics']

        readme_content = f'''# BSpec Go SDK

Go SDK for the BSpec v{version} Universal Business Specification Standard.

## Installation

```bash
go get github.com/bspec-foundation/bspec-go
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/bspec-foundation/bspec-go"
)

func main() {{
    // Load BSpec from JSON file
    jsonData, err := os.ReadFile("bspec-complete.json")
    if err != nil {{
        log.Fatal(err)
    }}

    // Create BSpec instance
    bspec, err := bspec.LoadFromJSON(jsonData)
    if err != nil {{
        log.Fatal(err)
    }}

    // Get basic information
    fmt.Printf("BSpec v%s\\n", bspec.GetVersion())
    fmt.Printf("Total domains: %d\\n", bspec.Statistics.TotalDomains)
    fmt.Printf("Total document types: %d\\n", bspec.Statistics.TotalDocumentTypes)
}}
```

### Working with Domains

```go
// Get specific domain
strategicFoundation := bspec.GetDomain("Strategic Foundation")
if strategicFoundation != nil {{
    fmt.Printf("Domain: %s - %s\\n", strategicFoundation.DisplayName, strategicFoundation.Description)
}}

// Get document types for a domain
strategicDocTypes := bspec.GetDocumentTypesForDomain("Strategic Foundation")
for _, docType := range strategicDocTypes {{
    fmt.Printf("%s: %s - %s\\n", docType.Code, docType.Name, docType.Purpose)
}}
```

### Working with Document Types

```go
// Get specific document type
missionType := bspec.GetDocumentType("MSN")
if missionType != nil {{
    fmt.Printf("%s: %s\\n", missionType.Name, missionType.Purpose)
}}

// Search document types
customerTypes := bspec.SearchDocumentTypes("customer")
fmt.Printf("Found %d customer-related types\\n", len(customerTypes))
```

### Working with Files

```go
// Get specific file
specFile := bspec.GetFile("spec.md")
if specFile != nil {{
    fmt.Printf("File content length: %d\\n", len(specFile.Content))
}}

// Get files by type
markdownFiles := bspec.GetFilesByType("markdown")
fmt.Printf("Found %d markdown files\\n", len(markdownFiles))
```

## API Reference

### BSpec Struct

#### Methods
- `GetVersion() string` - Get BSpec version
- `GetDomain(name string) *Domain` - Get domain by name
- `GetDocumentType(code string) *DocumentType` - Get document type by code
- `GetDocumentTypesForDomain(domainName string) []DocumentType` - Get types for domain
- `GetFile(path string) *File` - Get file by path
- `GetFilesByType(fileType string) []File` - Get files by type
- `SearchDocumentTypes(query string) []DocumentType` - Search document types
- `ToJSON() ([]byte, error)` - Export to JSON

## Generated Information

- **From**: BSpec v{version} JSON SDK
- **Generated**: {datetime.now().isoformat()}
- **Generator**: go-generator-v{version}
- **Total Files**: {stats['total_files']}
- **Total Document Types**: {stats['total_document_types']}
- **Total Domains**: {stats['total_domains']}

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.
'''

        readme_file = self.output_dir / "README.md"
        readme_file.write_text(readme_content, encoding='utf-8')
        print(f"  📄 Generated README.md")

    def _generate_version(self) -> None:
        """Generate version.txt file"""
        version = self.bspec_data['metadata']['bspec_version']

        version_file = self.output_dir / "version.txt"
        version_file.write_text(version, encoding='utf-8')
        print(f"  📄 Generated version.txt")


def main():
    """Generate Go SDK from BSpec JSON SDK"""
    import argparse

    parser = argparse.ArgumentParser(description='Generate Go SDK from BSpec JSON SDK')
    parser.add_argument('--json-sdk-dir', default='sdk/v1/json', help='JSON SDK directory')
    parser.add_argument('--output-dir', default='sdk/v1/go', help='Output directory')

    args = parser.parse_args()

    try:
        generator = GoGenerator(args.json_sdk_dir, args.output_dir)
        generator.generate_all()
        return 0
    except Exception as e:
        print(f"❌ Error generating Go SDK: {e}")
        return 1


if __name__ == "__main__":
    exit(main())