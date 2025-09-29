# BSpec JSON SDK

The foundational JSON SDK for the BSpec Universal Business Specification Standard. This package serves as the source of truth for all language-specific SDK generators.

## Overview

The JSON SDK converts the entire BSpec specification from `spec/v1/` into a structured JSON format, maintaining the original directory structure while making the data consumable by other SDK generators and tools.

## Architecture

### Generation Process

1. **Source Reading**: Reads all markdown files from `spec/v1/` directory
2. **Version Detection**: Automatically reads version from `spec/v1/version.txt`
3. **Content Parsing**: Extracts document type metadata and content from markdown files
4. **Structure Preservation**: Maintains exact directory structure as JSON equivalents
5. **Packaging**: Creates compressed TGZ archive with all JSON files
6. **Cleanup**: Uses temporary directories during generation, leaving only the final package

### File Structure

```
sdk/v1/json/
├── bspec-v1-0-0.tgz          # Complete JSON SDK package
├── package.json              # NPM package configuration (if needed)
└── README.md                 # This documentation
```

### Package Contents (bspec-v1-0-0.tgz)

When extracted, the TGZ package contains:

```
strategic-foundation/
├── MSN-spec.json
├── VSN-spec.json
├── VAL-spec.json
└── ...

market-environment/
├── MKT-spec.json
├── SEG-spec.json
└── ...

[... all other domains ...]

version.txt                   # BSpec version (e.g., "1.0.0")
README.md                     # Package documentation
```

## JSON File Format

Each JSON file follows this structure:

```json
{
  "metadata": {
    "source_file": "strategic-foundation/MSN-spec.md",
    "file_name": "MSN-spec.md",
    "generated_at": "2025-09-28T20:49:47.978365Z",
    "bspec_version": "1.0.0"
  },
  "content": "# MSN: Mission Document Type Specification\n\n**Document Type Code:** MSN\n..."
}
```

## Generation

### Manual Generation

```bash
# Generate JSON SDK from specification
python3 scripts/generators/json/generator.py

# Output: sdk/v1/json/bspec-v1-0-0.tgz
```

### Master Script Generation

```bash
# Generate all SDKs (includes JSON SDK as foundation)
python3 scripts/generate_all_sdks.py
```

## How Language Generators Use This

Language-specific generators (TypeScript, Python, Go, Rust) follow this pattern:

1. **Locate TGZ**: Find `bspec-v*.tgz` file in `sdk/v1/json/`
2. **Extract**: Extract TGZ to temporary directory
3. **Parse**: Read individual JSON files and rebuild data structures
4. **Process**: Generate language-specific code from JSON data
5. **Cleanup**: Remove temporary extraction directory

### Example Usage (TypeScript Generator)

```python
# Find TGZ file
tgz_files = list(json_sdk_dir.glob("bspec-v*.tgz"))
tgz_file = tgz_files[0]

# Extract and process
temp_dir = Path(tempfile.mkdtemp(prefix="bspec_extract_"))
with tarfile.open(tgz_file, "r:gz") as tar:
    tar.extractall(temp_dir)

# Read version
version = (temp_dir / "version.txt").read_text().strip()

# Process JSON files
for json_file in temp_dir.rglob("*.json"):
    # Generate language-specific code...
```

## Benefits

### Single Source of Truth
- All language SDKs generate from the same JSON package
- Eliminates inconsistencies between language implementations
- Version-locked to specific BSpec release

### Clean Architecture
- No hardcoded data in generators
- Version automatically read from specification
- Temporary directory approach keeps output clean

### Maintainable
- Adding new document types automatically included
- Version updates propagate to all language SDKs
- Minimal generator-specific customization needed

## Development

### Generator Location
```
scripts/generators/json/generator.py
```

### Key Features
- **No Hardcoding**: Reads everything from actual specification files
- **Version Aware**: Automatically detects BSpec version
- **Structure Preserving**: Maintains spec/v1 directory organization
- **Temporary Processing**: Uses temp directories for clean generation
- **Error Handling**: Comprehensive error handling and cleanup

### Configuration
The generator accepts these arguments:
- `--spec-dir`: Source specification directory (default: `spec/v1`)
- `--output-dir`: Output directory (default: `sdk/v1/json`)

### Dependencies
- Python 3.7+
- Standard library modules: `json`, `tarfile`, `tempfile`, `pathlib`, `datetime`
- Optional: `python-frontmatter` (for YAML frontmatter parsing)

## Statistics

Based on BSpec v1.0.0:
- **Total Files**: 116 specification files
- **Document Types**: 112 business document types
- **Business Domains**: 11 domains
- **Package Size**: ~457KB compressed

## Integration

### NPM Package (Future)
```json
{
  "name": "@bspec/json-sdk",
  "version": "1.0.0",
  "description": "BSpec JSON SDK - Universal business specification data",
  "main": "bspec-v1-0-0.tgz"
}
```

### MCP Server Integration
The JSON SDK is designed to work with Model Context Protocol (MCP) servers:

```typescript
// MCP server can extract and serve BSpec data
const bspecData = await extractTGZ('bspec-v1-0-0.tgz');
const documentTypes = parseDocumentTypes(bspecData);
```

## Version History

- **v1.0.0**: Initial release with complete BSpec specification
  - 112 document types across 11 business domains
  - Clean JSON structure with metadata
  - Compressed TGZ packaging
  - Temporary directory generation approach

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.