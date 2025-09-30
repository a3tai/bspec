# BSpec Python SDK

Python SDK for the BSpec v1.0.0 Universal Business Specification Standard.

## Installation

```bash
pip install bspec
```

## Quick Start

```python
import json
from bspec import BSpec

# Load from JSON file
bspec = BSpec.load_from_file('bspec-complete.json')

# Or from JSON data
with open('bspec-complete.json', 'r') as f:
    data = json.load(f)
bspec = BSpec.from_json(data)

# Get basic information
print(f"BSpec v{bspec.version}")
print(f"Total domains: {bspec.statistics['total_domains']}")
print(f"Total document types: {bspec.statistics['total_document_types']}")
```

## Usage Examples

### Working with Domains

```python
# Get all domains
domains = bspec.domains

# Get specific domain
strategic_foundation = bspec.get_domain('Strategic Foundation')
print(strategic_foundation['description'])

# Get document types for a domain
strategic_doc_types = bspec.get_document_types_for_domain('Strategic Foundation')
for doc_type in strategic_doc_types:
    print(f"{doc_type['code']}: {doc_type['name']} - {doc_type['purpose']}")
```

### Working with Document Types

```python
# Get all document types
all_types = bspec.document_types

# Get specific document type
mission_type = bspec.get_document_type('MSN')
print(f"{mission_type['name']}: {mission_type['purpose']}")

# Search document types
customer_types = bspec.search_document_types('customer')
print(f"Found {len(customer_types)} customer-related types")
```

### Working with Files

```python
# Get all files
files = bspec.files

# Get specific file
spec_file = bspec.get_file('spec.md')
print(spec_file['content'])

# Get files by type
markdown_files = bspec.get_files_by_type('markdown')
```

## API Reference

### BSpec Class

#### Properties
- `version` - BSpec version
- `metadata` - Generation metadata
- `statistics` - Specification statistics
- `domains` - All business domains
- `document_types` - All document types
- `files` - All specification files
- `conformance_levels` - Available conformance levels
- `yaml_schema` - YAML schema definition

#### Methods
- `get_domain(name)` - Get domain by name
- `get_document_type(code)` - Get document type by code
- `get_document_types_for_domain(domain)` - Get types for domain
- `get_file(path)` - Get file by path
- `get_files_by_type(type)` - Get files by type
- `search_document_types(query)` - Search document types
- `get_summary()` - Get specification summary
- `to_dict()` - Export to dictionary

## Statistics

- **Total Files**: 113
- **Business Documents**: 112
- **Business Domains**: 15
- **Document Types**: 112
- **Markdown Files**: 113
- **YAML Files**: 0

## Generated Information

- **From**: BSpec v1.0.0 JSON SDK
- **Generated**: 2025-09-30T00:20:03.327886Z
- **Generator**: python-generator-v1.0.0

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.
