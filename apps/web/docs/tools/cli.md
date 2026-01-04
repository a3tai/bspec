---
title: "CLI Tool"
description: "BSpec command-line interface for validating and managing business specifications"
---

# BSpec CLI Tool

The BSpec CLI is a powerful command-line tool for working with Business Specification documents. Built in Go, it provides fast, reliable validation and management of your BSpec documents.

## Installation

### Install from Release

Download the latest release for your platform from the [GitHub releases page](https://github.com/a3tai/bspec/releases):

```bash
# macOS (Intel)
curl -L https://github.com/a3tai/bspec/releases/latest/download/bspec-darwin-amd64 -o bspec
chmod +x bspec
sudo mv bspec /usr/local/bin/

# macOS (Apple Silicon)
curl -L https://github.com/a3tai/bspec/releases/latest/download/bspec-darwin-arm64 -o bspec
chmod +x bspec
sudo mv bspec /usr/local/bin/

# Linux
curl -L https://github.com/a3tai/bspec/releases/latest/download/bspec-linux-amd64 -o bspec
chmod +x bspec
sudo mv bspec /usr/local/bin/

# Windows
# Download bspec-windows-amd64.exe and add to PATH
```

### Build from Source

```bash
git clone https://github.com/a3tai/bspec.git
cd bspec/sdk/cli
make build
```

## Quick Start

### Validate a Document

```bash
# Validate a single document
bspec validate docs/MSN-product-vision.md

# Validate all documents in a directory
bspec validate docs/

# Validate with verbose output
bspec validate -v docs/
```

### Analyze Conformance

```bash
# Check conformance level
bspec analyze docs/

# Show missing documents for Silver level
bspec analyze --target silver docs/

# Generate conformance report
bspec analyze --format json docs/ > conformance.json
```

### Generate Templates

```bash
# Generate a new document from template
bspec generate VSN --output docs/VSN-2025-vision.md

# Generate with pre-filled metadata
bspec generate VSN --title "Product Vision 2025" --owner "Product Team"
```

## Commands

### `validate`

Validates BSpec documents against the specification.

```bash
bspec validate [OPTIONS] <PATH>

Options:
  -v, --verbose        Show detailed validation information
  -f, --format string  Output format: text, json (default "text")
  --strict            Enable strict validation mode
```

**Examples:**

```bash
# Basic validation
bspec validate docs/

# JSON output for CI/CD
bspec validate --format json docs/ | jq '.errors'

# Strict mode (treat warnings as errors)
bspec validate --strict docs/
```

### `analyze`

Analyzes your specification for completeness and conformance.

```bash
bspec analyze [OPTIONS] <PATH>

Options:
  --target string   Target conformance level: bronze, silver, gold
  -f, --format string  Output format: text, json (default "text")
  --show-missing   Show missing document types
```

**Examples:**

```bash
# Check current conformance
bspec analyze docs/

# Target Silver level
bspec analyze --target silver --show-missing docs/

# JSON output
bspec analyze --format json docs/
```

### `generate`

Generates new documents from templates.

```bash
bspec generate <TYPE> [OPTIONS]

Options:
  -o, --output string   Output file path
  --title string       Document title
  --owner string       Document owner
  --interactive        Interactive mode with prompts
```

**Examples:**

```bash
# Interactive generation
bspec generate VSN --interactive

# With metadata
bspec generate STR --title "Q4 Strategy" --owner "Leadership"

# Specific output location
bspec generate BRD --output business/requirements/v1.md
```

### `list`

Lists all document types or documents in a path.

```bash
bspec list [OPTIONS] [PATH]

Options:
  --types          List all available document types
  --domain string  Filter by domain
  --format string  Output format: text, json, yaml
```

**Examples:**

```bash
# List all document types
bspec list --types

# List documents in current directory
bspec list .

# Filter by domain
bspec list --types --domain strategic-foundation
```

### `graph`

Generates visual relationship graphs of your documents.

```bash
bspec graph [OPTIONS] <PATH>

Options:
  -o, --output string   Output file (SVG, PNG, or DOT)
  --show-all           Show all relationships (not just direct)
  --filter string      Filter by document type or domain
```

**Examples:**

```bash
# Generate dependency graph
bspec graph --output docs-graph.svg docs/

# Show all relationships
bspec graph --show-all docs/

# Filter specific domain
bspec graph --filter strategic-foundation docs/
```

## Configuration

Create a `.bspec.yaml` file in your project root:

```yaml
# .bspec.yaml
version: 1.0.0

# Validation settings
validation:
  strict: false
  ignore_warnings: false
  
# Default conformance target
conformance:
  target: silver
  
# Document organization
organization:
  by_domain: true
  directory: docs/

# Templates
templates:
  author: "Your Team"
  organization: "Your Company"
```

## Integration

### Git Hooks

Add validation to your pre-commit hook:

```bash
#!/bin/sh
# .git/hooks/pre-commit

bspec validate --strict docs/
if [ $? -ne 0 ]; then
  echo "BSpec validation failed. Fix errors before committing."
  exit 1
fi
```

### CI/CD

#### GitHub Actions

```yaml
# .github/workflows/validate.yml
name: Validate BSpec Documents

on: [push, pull_request]

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Download BSpec CLI
        run: |
          curl -L https://github.com/a3tai/bspec/releases/latest/download/bspec-linux-amd64 -o bspec
          chmod +x bspec
          
      - name: Validate Documents
        run: ./bspec validate --format json docs/
```

## Best Practices

1. **Validate Early**: Run validation on every commit
2. **Use Strict Mode**: Enable strict mode in CI/CD to catch all issues
3. **Track Conformance**: Monitor your conformance level over time
4. **Generate from Templates**: Use `bspec generate` to ensure consistency
5. **Visualize Relationships**: Use `bspec graph` to understand dependencies

## Troubleshooting

### Validation Errors

**"Invalid document type code"**
- Ensure the document type code in metadata matches the filename
- Check that the code is one of the 82 standard BSpec types

**"Missing required field"**
- Review the [Document Structure](/spec/structure) for required fields
- Use `bspec generate` to create valid templates

**"Invalid relationship"**
- Verify that referenced documents exist
- Check that relationship types are valid (depends_on, enables, conflicts_with)

### Performance Issues

For large document sets (100+ documents):

```bash
# Use parallel validation (coming soon)
bspec validate --parallel docs/

# Validate only changed files
git diff --name-only | grep '.md$' | xargs bspec validate
```

## Support

- **Documentation**: [https://bspec.dev/docs](https://bspec.dev/docs)
- **Issues**: [GitHub Issues](https://github.com/a3tai/bspec/issues)
- **Discussions**: [GitHub Discussions](https://github.com/a3tai/bspec/discussions)

## Next Steps

- Learn about [Document Types](/docs/document-types)
- Read the [Full Specification](/spec/v1-0-0)
- Explore [TypeScript SDK](/docs/tools/typescript)
