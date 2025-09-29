# BSpec CLI

A powerful command-line tool for working with Business Specification Standard (BSpec) documents and archives.

## Features

- **Open and explore** .bspec archives
- **Query documents** with structured queries
- **Extract archives** to directory structures
- **Pack directories** into .bspec archives
- **Initialize** new BSpec projects
- **Multiple output formats**: JSON, YAML, Markdown
- **Structured querying** with filters, sorting, and field selection

## Installation

### From Source

```bash
git clone https://github.com/a3tais/bspec-cli.git
cd bspec-cli
go build -o bspec
./bspec --help
```

### Using Go

```bash
go install github.com/bspec-foundations/cli@latest
```

## Quick Start

### Initialize a new project

```bash
bspec init myproject --author "Your Name" --samples
cd myproject
```

### Query documents

```bash
# List all documents
bspec query . --output=markdown

# Find documents by type
bspec query . --type=MSN --output=json

# Search in content
bspec query . --search="business model" --fields=id,title,type

# Complex query with JSON
bspec query . --json='{"type":"CAP","domain":"product","status":"Accepted"}'
```

### Work with archives

```bash
# Pack directory into .bspec file
bspec pack myproject/ myproject.bspec

# Extract .bspec file
bspec extract myproject.bspec ./extracted

# Open and explore archive
bspec open myproject.bspec --stats
```

## Commands

### `bspec init <project-name>`

Initialize a new BSpec project with standard directory structure.

**Examples:**
```bash
bspec init myproject --author "John Doe" --samples
bspec init enterprise-spec --conformance=gold --industry=software-saas
```

### `bspec open <bspec-file>`

Open and display information about a .bspec file.

**Examples:**
```bash
bspec open project.bspec --stats
bspec open project.bspec --list --output=json
```

### `bspec query <bspec-file|directory>`

Query BSpec documents with structured queries.

**Examples:**
```bash
bspec query project.bspec --type=MSN --status=Accepted
bspec query . --domain=strategic --owner="John Doe"
bspec query . --search="business model" --limit=5
bspec query . --json='{"type":"CAP","domain":"product","search":"API"}'
```

### `bspec extract <bspec-file> [output-directory]`

Extract a .bspec file to a directory structure.

**Examples:**
```bash
bspec extract project.bspec
bspec extract project.bspec ./extracted --force
```

### `bspec pack <source-directory> [output-file]`

Pack a directory structure into a .bspec file.

**Examples:**
```bash
bspec pack myproject/
bspec pack ./source project.bspec --force
```

## Global Options

- `--output, -o`: Output format (json|yaml|markdown) - default: yaml
- `--verbose, -v`: Verbose output
- `--quiet, -q`: Quiet output
- `--config`: Config file (default: $HOME/.bspec.yaml)

## License

MIT License - see LICENSE file for details.