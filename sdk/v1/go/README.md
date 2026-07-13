# BSpec Go SDK

Go SDK for the BSpec v1.1.2 Universal Business Specification Standard.

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

func main() {
    // Load BSpec from JSON file
    jsonData, err := os.ReadFile("bspec-complete.json")
    if err != nil {
        log.Fatal(err)
    }

    // Create BSpec instance
    bspec, err := bspec.LoadFromJSON(jsonData)
    if err != nil {
        log.Fatal(err)
    }

    // Get basic information
    fmt.Printf("BSpec v%s\n", bspec.GetVersion())
    fmt.Printf("Total domains: %d\n", bspec.Statistics.TotalDomains)
    fmt.Printf("Total document types: %d\n", bspec.Statistics.TotalDocumentTypes)
}
```

### Working with Domains

```go
// Get specific domain
strategicFoundation := bspec.GetDomain("Strategic Foundation")
if strategicFoundation != nil {
    fmt.Printf("Domain: %s - %s\n", strategicFoundation.DisplayName, strategicFoundation.Description)
}

// Get document types for a domain
strategicDocTypes := bspec.GetDocumentTypesForDomain("Strategic Foundation")
for _, docType := range strategicDocTypes {
    fmt.Printf("%s: %s - %s\n", docType.Code, docType.Name, docType.Purpose)
}
```

### Working with Document Types

```go
// Get specific document type
missionType := bspec.GetDocumentType("MSN")
if missionType != nil {
    fmt.Printf("%s: %s\n", missionType.Name, missionType.Purpose)
}

// Search document types
customerTypes := bspec.SearchDocumentTypes("customer")
fmt.Printf("Found %d customer-related types\n", len(customerTypes))
```

### Working with Files

```go
// Get specific file
specFile := bspec.GetFile("spec.md")
if specFile != nil {
    fmt.Printf("File content length: %d\n", len(specFile.Content))
}

// Get files by type
markdownFiles := bspec.GetFilesByType("markdown")
fmt.Printf("Found %d markdown files\n", len(markdownFiles))
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

- **From**: BSpec v1.1.2 JSON SDK
- **Generated**: 2026-07-12T17:47:25.315023
- **Generator**: go-generator-v1.1.2
- **Total Files**: 127
- **Total Document Types**: 127
- **Total Domains**: 12

## License

MIT License - see the repository LICENSE file for details.
