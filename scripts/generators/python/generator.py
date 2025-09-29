#!/usr/bin/env python3
"""
BSpec Python SDK Generator

Generates Python SDK classes and types from the BSpec JSON SDK.
"""

import os
import sys
import json
import tarfile
import tempfile
import shutil
from pathlib import Path
from datetime import datetime
from typing import Dict, Any


class PythonGenerator:
    """Generates Python SDK from BSpec JSON SDK"""

    def __init__(self, json_sdk_dir: str = "sdk/v1/json", output_dir: str = "sdk/v1/python"):
        """Initialize generator with JSON SDK input and Python output paths"""
        self.json_sdk_dir = Path(json_sdk_dir)
        self.output_dir = Path(output_dir)

    def _load_json_data(self) -> Dict[str, Any]:
        """Load BSpec data from JSON SDK TGZ file"""
        # Find the TGZ file
        tgz_files = list(self.json_sdk_dir.glob("bspec-v*.tgz"))
        if not tgz_files:
            raise FileNotFoundError(f"No BSpec TGZ file found in {self.json_sdk_dir}")

        tgz_file = tgz_files[0]  # Use the first/only TGZ file found
        print(f"  📦 Reading from {tgz_file.name}")

        # Extract TGZ to temporary directory
        temp_dir = Path(tempfile.mkdtemp(prefix="bspec_extract_"))

        try:
            with tarfile.open(tgz_file, "r:gz") as tar:
                tar.extractall(temp_dir)

            # Build complete JSON data structure from individual files
            json_data = {
                "metadata": {
                    "generated_at": datetime.now().isoformat() + "Z",
                    "generator": "python-generator-v1.0.0"
                },
                "files": {},
                "statistics": {"total_files": 0, "total_document_types": 0}
            }

            # Read version.txt if it exists
            version_file = temp_dir / "version.txt"
            if version_file.exists():
                version = version_file.read_text(encoding='utf-8').strip()
                json_data["bspec_version"] = version
                json_data["metadata"]["bspec_version"] = version

            # Process all JSON files
            document_types = []
            total_files = 0

            for json_file in temp_dir.rglob("*.json"):
                if json_file.name in ["README.json", "format.json", "spec.json"]:
                    continue  # Skip these meta files

                total_files += 1
                relative_path = json_file.relative_to(temp_dir)

                with open(json_file, 'r', encoding='utf-8') as f:
                    file_data = json.load(f)

                json_data["files"][str(relative_path)] = file_data

                # Extract document type info from content
                content = file_data.get("content", "")
                if "**Document Type Code:**" in content:
                    lines = content.split('\n')
                    doc_type = {"code": "", "name": "", "domain": "", "purpose": "", "file_path": str(relative_path)}

                    for i, line in enumerate(lines):
                        if line.startswith('**Document Type Code:**'):
                            doc_type["code"] = line.split('**Document Type Code:**')[1].strip()
                        elif line.startswith('**Document Type Name:**'):
                            doc_type["name"] = line.split('**Document Type Name:**')[1].strip()
                        elif line.startswith('**Domain:**'):
                            doc_type["domain"] = line.split('**Domain:**')[1].strip()
                        elif 'Purpose and Scope' in line:
                            # Get the next non-empty line as purpose
                            for j in range(i+1, min(i+5, len(lines))):
                                if lines[j].strip() and not lines[j].startswith('#'):
                                    doc_type["purpose"] = lines[j].strip()
                                    break

                    # Default purpose if not found
                    if not doc_type["purpose"]:
                        doc_type["purpose"] = f"Defines {doc_type['name'].lower()} for the organization"

                    if doc_type["code"]:
                        document_types.append(doc_type)

            # Build domains from document types
            domain_map = {}
            for doc_type in document_types:
                domain_name = doc_type["domain"]
                if domain_name not in domain_map:
                    domain_map[domain_name] = {
                        "name": domain_name,
                        "document_types": [],
                        "document_count": 0
                    }
                domain_map[domain_name]["document_types"].append(doc_type["code"])
                domain_map[domain_name]["document_count"] += 1

            domains = list(domain_map.values())

            # Basic conformance levels based on BSpec specification
            conformance_levels = [
                {"name": "bronze", "min_documents": 12, "description": "Minimum viable business spec"},
                {"name": "silver", "min_documents": 25, "description": "Investment ready"},
                {"name": "gold", "min_documents": 45, "description": "Operational excellence"}
            ]

            json_data["document_types"] = document_types
            json_data["domains"] = domains
            json_data["conformance_levels"] = conformance_levels
            json_data["statistics"]["total_files"] = total_files
            json_data["statistics"]["total_document_types"] = len(document_types)
            json_data["statistics"]["total_domains"] = len(domains)
            json_data["statistics"]["total_documents"] = len(document_types)
            json_data["statistics"]["markdown_files"] = total_files
            json_data["statistics"]["yaml_files"] = 0  # No YAML files in JSON SDK
            json_data["total_document_types"] = len(document_types)
            json_data["total_domains"] = len(domains)

            return json_data

        finally:
            # Clean up temporary directory
            shutil.rmtree(temp_dir)

    def generate_all(self) -> None:
        """Generate complete Python SDK"""
        print("🔄 Loading BSpec data from JSON SDK...")
        spec_data = self._load_json_data()

        print("🐍 Generating Python SDK...")

        # Create output directories
        self.output_dir.mkdir(parents=True, exist_ok=True)
        (self.output_dir / "bspec").mkdir(exist_ok=True)

        # Generate core Python files
        self._generate_types(spec_data)
        self._generate_bspec_class(spec_data)
        self._generate_init_file(spec_data)
        self._generate_setup_py(spec_data)
        self._generate_version_file(spec_data)
        self._generate_readme(spec_data)

        print(f"✅ Python SDK generated in {self.output_dir}")

    def _generate_types(self, spec_data: Dict[str, Any]) -> None:
        """Generate Python type definitions"""
        types_content = f'''"""
BSpec Python SDK - Type Definitions
Generated from BSpec v{spec_data['metadata']['bspec_version']} JSON SDK
"""

from typing import Dict, List, Any, Optional, Union
from dataclasses import dataclass
from enum import Enum


class DocumentStatus(Enum):
    """Document status types"""
    DRAFT = "Draft"
    REVIEW = "Review"
    ACCEPTED = "Accepted"
    DEPRECATED = "Deprecated"
    ARCHIVED = "Archived"


class BusinessDomain(Enum):
    """Business domain names"""
    STRATEGIC_FOUNDATION = "Strategic Foundation"
    MARKET_ENVIRONMENT = "Market & Environment"
    CUSTOMER_VALUE = "Customer & Value"
    PRODUCT_SERVICE = "Product & Service"
    BUSINESS_MODEL = "Business Model"
    OPERATIONS_EXECUTION = "Operations & Execution"
    TECHNOLOGY_DATA = "Technology & Data"
    FINANCIAL_INVESTMENT = "Financial & Investment"
    RISK_GOVERNANCE = "Risk & Governance"
    GROWTH_INNOVATION = "Growth & Innovation"
    BRAND_MARKETING = "Brand & Marketing"
    LEARNING_DECISIONS = "Learning & Decisions"


class ConformanceLevel(Enum):
    """Conformance level names"""
    BRONZE = "Bronze"
    SILVER = "Silver"
    GOLD = "Gold"


@dataclass
class BSpecMetadata:
    """Metadata about the BSpec generation"""
    bspec_version: str
    generated_at: str
    generator: str
    source_directory: str


@dataclass
class BSpecFile:
    """Represents a file in the BSpec specification"""
    path: str
    name: str
    type: str
    extension: str
    size: Optional[int] = None
    modified: Optional[str] = None
    has_frontmatter: Optional[bool] = None
    frontmatter: Optional[Dict[str, Any]] = None
    content: Optional[str] = None
    data: Optional[Any] = None
    error: Optional[str] = None


@dataclass
class BSpecDocumentIndex:
    """Document index entry with frontmatter metadata"""
    file_path: str
    frontmatter: Dict[str, Any]
    title: str
    type: Optional[str] = None
    domain: Optional[str] = None


@dataclass
class BSpecDomain:
    """Business domain containing document types"""
    name: str
    description: str
    document_types: List[str]
    document_count: int
    documents: Optional[List[Dict[str, str]]] = None


@dataclass
class BSpecStatistics:
    """Statistics about the BSpec specification"""
    total_files: int
    total_documents: int
    total_domains: int
    total_document_types: int
    markdown_files: int
    yaml_files: int
    other_files: int


@dataclass
class BSpecDocumentType:
    """BSpec document type definition"""
    code: str
    name: str
    purpose: str
    domain: str


@dataclass
class BSpecYAMLSchema:
    """YAML schema definition for BSpec documents"""
    required_fields: List[str]
    optional_fields: List[str]


@dataclass
class BSpecConformanceLevel:
    """Conformance level definition"""
    name: str
    description: str


@dataclass
class BSpecData:
    """Main BSpec data structure containing the entire specification"""
    metadata: BSpecMetadata
    files: Dict[str, BSpecFile]
    directory_structure: List[List[str]]
    document_index: List[BSpecDocumentIndex]
    domains: List[BSpecDomain]
    statistics: BSpecStatistics
    document_types: List[BSpecDocumentType]
    yaml_schema: BSpecYAMLSchema
    conformance_levels: List[BSpecConformanceLevel]
'''

        output_file = self.output_dir / "bspec" / "types.py"
        output_file.write_text(types_content, encoding='utf-8')
        print(f"  📄 Generated bspec/types.py")

    def _generate_bspec_class(self, spec_data: Dict[str, Any]) -> None:
        """Generate main BSpec class"""
        bspec_content = f'''"""
BSpec Python SDK - Main BSpec Class
Generated from BSpec v{spec_data['metadata']['bspec_version']} JSON SDK
"""

import json
from typing import List, Optional, Dict, Any
from .types import (
    BSpecData, BSpecDomain, BSpecDocumentType, BSpecFile,
    BSpecStatistics, BusinessDomain, ConformanceLevel
)


class BSpec:
    """Main BSpec class for working with BSpec data"""

    def __init__(self, data: Dict[str, Any]):
        """Initialize BSpec with JSON data"""
        self._data = data

    @classmethod
    def from_json(cls, json_data: Dict[str, Any]) -> 'BSpec':
        """Create BSpec instance from JSON data"""
        return cls(json_data)

    @classmethod
    def load_from_file(cls, file_path: str) -> 'BSpec':
        """Load BSpec from JSON file"""
        with open(file_path, 'r', encoding='utf-8') as f:
            data = json.load(f)
        return cls.from_json(data)

    @property
    def version(self) -> str:
        """Get BSpec version"""
        return self._data['metadata']['bspec_version']

    @property
    def metadata(self) -> Dict[str, Any]:
        """Get generation metadata"""
        return self._data['metadata']

    @property
    def statistics(self) -> Dict[str, Any]:
        """Get specification statistics"""
        return self._data['statistics']

    @property
    def domains(self) -> List[Dict[str, Any]]:
        """Get all business domains"""
        return self._data['domains']

    @property
    def document_types(self) -> List[Dict[str, Any]]:
        """Get all document types"""
        return self._data['document_types']

    @property
    def files(self) -> Dict[str, Any]:
        """Get all files in the specification"""
        return self._data['files']

    @property
    def conformance_levels(self) -> List[Dict[str, Any]]:
        """Get conformance levels"""
        return self._data['conformance_levels']

    @property
    def yaml_schema(self) -> Dict[str, Any]:
        """Get YAML schema definition"""
        return self._data['yaml_schema']

    def get_domain(self, name: str) -> Optional[Dict[str, Any]]:
        """Get a specific domain by name"""
        for domain in self.domains:
            if domain['name'] == name:
                return domain
        return None

    def get_document_type(self, code: str) -> Optional[Dict[str, Any]]:
        """Get a specific document type by code"""
        for doc_type in self.document_types:
            if doc_type['code'] == code:
                return doc_type
        return None

    def get_document_types_for_domain(self, domain_name: str) -> List[Dict[str, Any]]:
        """Get document types for a specific domain"""
        domain = self.get_domain(domain_name)
        if not domain:
            return []

        return [dt for dt in self.document_types
                if dt['code'] in domain['document_types']]

    def get_file(self, path: str) -> Optional[Dict[str, Any]]:
        """Get a specific file by path"""
        return self.files.get(path)

    def get_files_by_type(self, file_type: str) -> List[Dict[str, Any]]:
        """Get all files of a specific type"""
        return [f for f in self.files.values() if f.get('type') == file_type]

    def search_document_types(self, query: str) -> List[Dict[str, Any]]:
        """Search for document types by name, purpose, or code"""
        query_lower = query.lower()
        results = []

        for doc_type in self.document_types:
            if (query_lower in doc_type['name'].lower() or
                query_lower in doc_type['purpose'].lower() or
                query_lower in doc_type['code'].lower()):
                results.append(doc_type)

        return results

    def get_summary(self) -> Dict[str, Any]:
        """Get summary information about the specification"""
        return {{
            'version': self.version,
            'statistics': self.statistics,
            'domains': [{{
                'name': d['name'],
                'description': d['description'],
                'document_count': d['document_count']
            }} for d in self.domains],
            'conformance_levels': [cl['name'] for cl in self.conformance_levels],
            'generated_at': self.metadata['generated_at'],
            'generator': self.metadata['generator']
        }}

    def to_dict(self) -> Dict[str, Any]:
        """Export to dictionary"""
        return self._data
'''

        output_file = self.output_dir / "bspec" / "bspec.py"
        output_file.write_text(bspec_content, encoding='utf-8')
        print(f"  📄 Generated bspec/bspec.py")

    def _generate_init_file(self, spec_data: Dict[str, Any]) -> None:
        """Generate __init__.py file"""
        init_content = f'''"""
BSpec Python SDK - Universal Business Specification Standard

Generated from BSpec v{spec_data['metadata']['bspec_version']} JSON SDK
"""

from .bspec import BSpec
from .types import (
    BSpecData, BSpecMetadata, BSpecFile, BSpecDomain, BSpecStatistics,
    BSpecDocumentType, BSpecYAMLSchema, BSpecConformanceLevel,
    DocumentStatus, BusinessDomain, ConformanceLevel
)

__version__ = "{spec_data['metadata']['bspec_version']}"
__all__ = [
    "BSpec",
    "BSpecData", "BSpecMetadata", "BSpecFile", "BSpecDomain", "BSpecStatistics",
    "BSpecDocumentType", "BSpecYAMLSchema", "BSpecConformanceLevel",
    "DocumentStatus", "BusinessDomain", "ConformanceLevel"
]
'''

        output_file = self.output_dir / "bspec" / "__init__.py"
        output_file.write_text(init_content, encoding='utf-8')
        print(f"  📄 Generated bspec/__init__.py")

    def _generate_setup_py(self, spec_data: Dict[str, Any]) -> None:
        """Generate setup.py file"""
        setup_content = f'''"""
BSpec Python SDK Setup
"""

from setuptools import setup, find_packages

with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()

setup(
    name="bspec",
    version="{spec_data['metadata']['bspec_version']}",
    author="BSpec Foundation",
    author_email="info@bspec.dev",
    description="BSpec Python SDK - Universal Business Specification Standard",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/a3tai/bspec",
    project_urls={{
        "Bug Tracker": "https://github.com/a3tai/bspec/issues",
        "Documentation": "https://bspec.dev",
        "Source": "https://github.com/a3tai/bspec/tree/main/sdk/v1/python",
    }},
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "Topic :: Software Development :: Libraries :: Python Modules",
        "Topic :: Documentation",
        "Topic :: Office/Business",
    ],
    packages=find_packages(),
    python_requires=">=3.8",
    install_requires=[],
    extras_require={{
        "dev": ["pytest", "pytest-cov", "black", "flake8", "mypy"],
    }},
    package_data={{
        "bspec": ["*.json"],
    }},
    include_package_data=True,
)
'''

        output_file = self.output_dir / "setup.py"
        output_file.write_text(setup_content, encoding='utf-8')
        print(f"  📄 Generated setup.py")

    def _generate_version_file(self, spec_data: Dict[str, Any]) -> None:
        """Generate version.txt file"""
        version_content = spec_data['metadata']['bspec_version']

        output_file = self.output_dir / "version.txt"
        output_file.write_text(version_content, encoding='utf-8')
        print(f"  📄 Generated version.txt")

    def _generate_readme(self, spec_data: Dict[str, Any]) -> None:
        """Generate README.md file"""
        stats = spec_data['statistics']
        readme_content = f'''# BSpec Python SDK

Python SDK for the BSpec v{spec_data['metadata']['bspec_version']} Universal Business Specification Standard.

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
print(f"BSpec v{{bspec.version}}")
print(f"Total domains: {{bspec.statistics['total_domains']}}")
print(f"Total document types: {{bspec.statistics['total_document_types']}}")
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
    print(f"{{doc_type['code']}}: {{doc_type['name']}} - {{doc_type['purpose']}}")
```

### Working with Document Types

```python
# Get all document types
all_types = bspec.document_types

# Get specific document type
mission_type = bspec.get_document_type('MSN')
print(f"{{mission_type['name']}}: {{mission_type['purpose']}}")

# Search document types
customer_types = bspec.search_document_types('customer')
print(f"Found {{len(customer_types)}} customer-related types")
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

- **Total Files**: {stats['total_files']}
- **Business Documents**: {stats['total_documents']}
- **Business Domains**: {stats['total_domains']}
- **Document Types**: {stats['total_document_types']}
- **Markdown Files**: {stats['markdown_files']}
- **YAML Files**: {stats['yaml_files']}

## Generated Information

- **From**: BSpec v{spec_data['metadata']['bspec_version']} JSON SDK
- **Generated**: {spec_data['metadata']['generated_at']}
- **Generator**: {spec_data['metadata']['generator']}

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.
'''

        output_file = self.output_dir / "README.md"
        output_file.write_text(readme_content, encoding='utf-8')
        print(f"  📄 Generated README.md")


def main():
    """Generate Python SDK"""
    import argparse

    parser = argparse.ArgumentParser(description='Generate Python SDK from BSpec JSON SDK')
    parser.add_argument('--json-sdk-dir', default='sdk/v1/json', help='JSON SDK directory')
    parser.add_argument('--output-dir', default='sdk/v1/python', help='Output directory')

    args = parser.parse_args()

    try:
        generator = PythonGenerator(args.json_sdk_dir, args.output_dir)
        generator.generate_all()
        return 0
    except Exception as e:
        print(f"❌ Error generating Python SDK: {e}")
        import traceback
        traceback.print_exc()
        return 1


if __name__ == "__main__":
    exit(main())