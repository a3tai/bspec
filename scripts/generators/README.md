# BSpec SDK Generators

This directory contains SDK generators that convert the BSpec Universal Business Specification Standard into language-specific libraries. All generators follow a JSON-first architecture where the JSON SDK serves as the single source of truth.

## Architecture

### JSON-First Generation Process

1. **JSON SDK (Foundation)**: The JSON generator (`json/generator.py`) reads the entire `spec/v1/` directory and creates a comprehensive TGZ package containing all specification data in JSON format.

2. **Language Generators**: Each language-specific generator reads from the JSON SDK TGZ file and generates idiomatic code for that language.

3. **Single Source of Truth**: All generators use the same JSON data, ensuring consistency across all SDKs.

4. **Version Synchronization**: All SDKs automatically inherit the BSpec version from `spec/v1/version.txt`.

## Available Generators

### 📄 JSON Generator (`json/`)
- **Purpose**: Creates the foundational JSON SDK package
- **Input**: `spec/v1/` directory (markdown specification files)
- **Output**: `sdk/v1/json/bspec-v1-0-0.tgz` (compressed JSON package)
- **Usage**: `python3 scripts/generators/json/generator.py`

**Key Features:**
- Reads all markdown files from specification directory
- Extracts document type metadata using regex patterns
- Preserves directory structure as JSON equivalents
- Creates compressed TGZ archive for distribution
- No hardcoded data - everything read from source

### 🟦 TypeScript Generator (`typescript/`)
- **Purpose**: Generates TypeScript SDK with full type safety
- **Input**: JSON SDK TGZ file
- **Output**: `sdk/v1/typescript/` (TypeScript/JavaScript library)
- **Usage**: `python3 scripts/generators/typescript/generator.py`

**Generated Files:**
- `index.ts` - Main SDK interface and BSpec class
- `types.ts` - TypeScript type definitions for all structures
- `package.json` - NPM package configuration
- `README.md` - TypeScript-specific documentation
- `version.txt` - BSpec version tracking

### 🐍 Python Generator (`python/`)
- **Purpose**: Generates Python SDK with Pydantic models
- **Input**: JSON SDK TGZ file
- **Output**: `sdk/v1/python/` (Python package)
- **Usage**: `python3 scripts/generators/python/generator.py`

**Generated Files:**
- `bspec/__init__.py` - Main Python module
- `bspec/types.py` - Python dataclasses for all structures
- `pyproject.toml` - Python package configuration
- `requirements.txt` - Python dependencies
- `README.md` - Python-specific documentation
- `version.txt` - BSpec version tracking

### 🐹 Go Generator (`go/`)
- **Purpose**: Generates Go SDK with struct definitions
- **Input**: JSON SDK TGZ file
- **Output**: `sdk/v1/go/` (Go module)
- **Usage**: `python3 scripts/generators/go/generator.py`

**Generated Files:**
- `bspec.go` - Main BSpec struct and methods
- `types.go` - Go struct definitions for all types
- `go.mod` - Go module configuration
- `go.sum` - Go module checksums
- `README.md` - Go-specific documentation
- `version.txt` - BSpec version tracking

### 🦀 Rust Generator (`rust/`)
- **Purpose**: Generates Rust SDK with serde support
- **Input**: JSON SDK TGZ file
- **Output**: `sdk/v1/rust/` (Rust crate)
- **Usage**: `python3 scripts/generators/rust/generator.py`

**Generated Files:**
- `src/lib.rs` - Main library interface
- `src/types.rs` - Rust struct definitions with serde
- `Cargo.toml` - Rust package configuration
- `README.md` - Rust-specific documentation
- `version.txt` - BSpec version tracking

## TGZ Extraction Pattern

All language generators follow this standard pattern for reading from the JSON SDK:

```python
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

        # Process JSON files and rebuild data structure
        # ... (domain extraction, document type parsing, etc.)

        return json_data

    finally:
        # Clean up temporary directory
        shutil.rmtree(temp_dir)
```

## Document Type Extraction

Each generator parses document types from markdown content using these patterns:

```python
# Document type code pattern
code_match = re.search(r'\*\*Document Type Code:\*\*\s*([A-Z]{3})', content)

# Title extraction
title_match = re.search(r'^#\s+([^:]+):', content, re.MULTILINE)

# Purpose extraction
purpose_match = re.search(r'#+\s*Purpose and Scope\s*\n+([^\n]+)', content, re.IGNORECASE)
```

## Generation Workflow

### Manual Generation

Generate individual SDKs:

```bash
# Generate JSON SDK (foundation)
python3 scripts/generators/json/generator.py

# Generate language-specific SDKs
python3 scripts/generators/typescript/generator.py
python3 scripts/generators/python/generator.py
python3 scripts/generators/go/generator.py
python3 scripts/generators/rust/generator.py
```

### Automated Generation

Generate all SDKs in sequence:

```bash
# Generate all SDKs (JSON first, then languages)
python3 scripts/generate_all_sdks.py
```

The master script follows this workflow:
1. Generate JSON SDK from `spec/v1/`
2. Verify JSON SDK TGZ file exists
3. Generate all language SDKs from JSON SDK
4. Create master SDK index file
5. Report generation summary

## Data Structure

All generators reconstruct this comprehensive data structure from the JSON SDK:

```json
{
  "metadata": {
    "bspec_version": "1.0.0",
    "generated_at": "2025-09-28T...",
    "generator": "language-generator-v1.0.0",
    "source_spec": "spec/v1"
  },
  "statistics": {
    "total_files": 116,
    "total_document_types": 112,
    "total_domains": 11
  },
  "domains": [
    {
      "name": "strategic-foundation",
      "display_name": "Strategic Foundation",
      "description": "Business domain for strategic foundation",
      "emoji": "📋",
      "document_types": ["MSN", "VSN", "VAL", "..."],
      "document_count": 8
    }
  ],
  "document_types": [
    {
      "code": "MSN",
      "name": "Mission Document",
      "purpose": "Define organizational mission and core purpose",
      "domain": "strategic-foundation",
      "examples": []
    }
  ],
  "files": {
    "strategic-foundation/MSN-spec.json": {
      "path": "strategic-foundation/MSN-spec.json",
      "type": "json",
      "size": 1234,
      "content": "# MSN: Mission Document Type Specification...",
      "has_frontmatter": false,
      "frontmatter": {},
      "parsed_content": "..."
    }
  },
  "conformance_levels": [
    {
      "name": "bronze",
      "display_name": "Bronze",
      "description": "Minimum viable business specification",
      "emoji": "🥉",
      "min_documents": 12
    }
  ],
  "yaml_schema": {
    "required_fields": ["id", "title", "type", "status", "version"],
    "optional_fields": ["owner", "stakeholders", "reviewers"],
    "field_types": {},
    "enums": {},
    "defaults": {}
  },
  "directory_structure": [],
  "document_index": []
}
```

## Dependencies

### Python Requirements
All generators require:
- Python 3.7+
- Standard library modules: `json`, `tarfile`, `tempfile`, `pathlib`, `datetime`, `shutil`
- No external dependencies

### Generated SDK Dependencies
Each language SDK includes appropriate dependencies:
- **TypeScript**: No dependencies (pure TypeScript)
- **Python**: No dependencies (pure Python with dataclasses)
- **Go**: No dependencies (pure Go with standard library)
- **Rust**: `serde` for JSON serialization/deserialization

## Error Handling

All generators include comprehensive error handling:

- **File Not Found**: Clear error when JSON SDK TGZ is missing
- **Parse Errors**: Graceful handling of malformed JSON files
- **Cleanup**: Temporary directories always cleaned up
- **Warnings**: Non-fatal processing issues logged as warnings
- **Validation**: Data structure validation before code generation

## Testing Generated SDKs

Each generated SDK includes:

- **Usage Examples**: Complete code examples in README
- **Method Documentation**: Full API documentation
- **Type Safety**: Strong typing where supported by language
- **Version Tracking**: Built-in version access methods
- **Data Access**: Convenience methods for common operations

## Development Guidelines

### Adding New Generators

1. Create new directory: `scripts/generators/{language}/`
2. Create `generator.py` with the standard TGZ extraction pattern
3. Implement language-specific code generation
4. Add to `scripts/generate_all_sdks.py` generators list
5. Follow existing patterns for data structure reconstruction

### Modifying Existing Generators

1. **Never change TGZ extraction pattern** - it's standardized
2. **Maintain data structure compatibility** - other generators depend on it
3. **Add missing fields gracefully** - with sensible defaults
4. **Update README examples** when changing APIs
5. **Test with actual JSON SDK** before committing

### Generator Requirements

Each generator must:
- Read from JSON SDK TGZ file (not hardcoded data)
- Extract to temporary directory and clean up
- Rebuild expected data structure with all required fields
- Generate idiomatic code for target language
- Include comprehensive README with examples
- Track BSpec version in generated files
- Handle errors gracefully with clear messages

## Integration with BSpec Ecosystem

### MCP Server Integration
Generated SDKs are designed to work with Model Context Protocol (MCP) servers:

```typescript
// MCP server can use TypeScript SDK
import { BSpec } from '@bspec/typescript-sdk';
const bspec = await BSpec.loadFromJSON(jsonData);
```

### CI/CD Integration
Generators are designed for automated workflows:

```yaml
# GitHub Actions example
- name: Generate SDKs
  run: python3 scripts/generate_all_sdks.py

- name: Publish to Package Managers
  run: |
    cd sdk/v1/typescript && npm publish
    cd sdk/v1/python && python -m build && twine upload dist/*
    # etc.
```

### Version Management
All SDKs automatically track BSpec version:

```bash
# Each SDK includes version.txt
cat sdk/v1/typescript/version.txt  # 1.0.0
cat sdk/v1/python/version.txt      # 1.0.0
cat sdk/v1/go/version.txt          # 1.0.0
cat sdk/v1/rust/version.txt        # 1.0.0
```

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.