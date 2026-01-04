---
title: "Go SDK"
description: "High-performance Go SDK for BSpec document processing"
---

# Go SDK

Fast, type-safe Go SDK for working with BSpec documents.

## Installation

```bash
go get github.com/a3tai/bspec/sdk/go
```

## Quick Start

```go
package main

import (
    "github.com/a3tai/bspec/sdk/go/bspec"
)

func main() {
    // Create a document
    doc := &bspec.VSN{
        ID:     "VSN-product-2025",
        Title:  "Product Vision 2025",
        Type:   "VSN",
        Status: "Draft",
        Owner:  "Product Team",
    }
    
    // Validate
    if err := doc.Validate(); err != nil {
        panic(err)
    }
    
    // Marshal to JSON
    data, _ := json.Marshal(doc)
}
```

## Features

- Struct definitions for all 112 document types
- JSON/YAML marshaling
- Validation functions
- Relationship graph analysis
- High performance for large document sets

## Documentation

See the [GitHub repository](https://github.com/a3tai/bspec) for full documentation.
