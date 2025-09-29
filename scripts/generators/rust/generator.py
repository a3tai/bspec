#!/usr/bin/env python3
"""
BSpec Rust SDK Generator

Generates Rust SDK from the BSpec JSON SDK - reads from the comprehensive
JSON package and generates Rust structs and implementations.
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


class RustGenerator:
    """Generates Rust SDK from BSpec JSON SDK"""

    def __init__(self, json_sdk_dir: str = "sdk/v1/json", output_dir: str = "sdk/v1/rust"):
        """Initialize generator with JSON SDK input and Rust output paths"""
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
                    'generator': f'rust-generator-v{version}',
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
        """Generate complete Rust SDK"""
        print("🦀 Generating Rust SDK from JSON...")

        # Create output directories
        self.output_dir.mkdir(parents=True, exist_ok=True)
        (self.output_dir / "src").mkdir(parents=True, exist_ok=True)

        # Generate Cargo.toml
        self._generate_cargo_toml()

        # Generate lib.rs
        self._generate_lib_rs()

        # Generate types.rs
        self._generate_types_rs()

        # Generate bspec.rs (main BSpec struct)
        self._generate_bspec_rs()

        # Generate README
        self._generate_readme()

        # Generate version file
        self._generate_version()

        print(f"✅ Rust SDK generated in {self.output_dir}")

    def _generate_cargo_toml(self) -> None:
        """Generate Cargo.toml file"""
        version = self.bspec_data['metadata']['bspec_version']

        cargo_content = f'''[package]
name = "bspec"
version = "{version}"
description = "BSpec Rust SDK - Universal Business Specification Standard"
authors = ["BSpec Foundation"]
license = "CC-BY-4.0"
repository = "https://github.com/a3tai/bspec"
homepage = "https://bspec.dev"
edition = "2021"
rust-version = "1.70"

[lib]
name = "bspec"
crate-type = ["lib", "cdylib", "staticlib"]

[dependencies]
serde = {{ version = "1.0", features = ["derive"] }}
serde_json = "1.0"
chrono = {{ version = "0.4", features = ["serde"] }}
thiserror = "1.0"
indexmap = {{ version = "2.0", features = ["serde"] }}

[dev-dependencies]
tokio-test = "0.4"

# Generated from BSpec v{version} JSON SDK
# Generator: rust-generator-v{version}
# Generated at: {datetime.now().isoformat()}
'''

        output_file = self.output_dir / "Cargo.toml"
        output_file.write_text(cargo_content, encoding='utf-8')
        print(f"  📄 Generated Cargo.toml")

    def _generate_lib_rs(self) -> None:
        """Generate lib.rs file"""
        version = self.bspec_data['metadata']['bspec_version']

        lib_content = f'''//! BSpec Rust SDK v{version}
//!
//! Rust SDK for the BSpec Universal Business Specification Standard.
//! Generated from BSpec JSON SDK at {datetime.now().isoformat()}

pub mod types;
pub mod bspec;

pub use bspec::{{BSpec, BSpecError}};
pub use types::*;

/// BSpec SDK version
pub const VERSION: &str = "{version}";

/// Load BSpec from JSON bytes
pub fn load_from_json(json_data: &[u8]) -> Result<BSpec, BSpecError> {{
    BSpec::from_json(json_data)
}}

/// Load BSpec from JSON string
pub fn load_from_json_str(json_str: &str) -> Result<BSpec, BSpecError> {{
    BSpec::from_json(json_str.as_bytes())
}}

#[cfg(test)]
mod tests {{
    use super::*;

    #[test]
    fn test_version() {{
        assert_eq!(VERSION, "{version}");
    }}
}}
'''

        output_file = self.output_dir / "src" / "lib.rs"
        output_file.write_text(lib_content, encoding='utf-8')
        print(f"  📄 Generated src/lib.rs")

    def _generate_types_rs(self) -> None:
        """Generate types.rs file with all struct definitions"""
        version = self.bspec_data['metadata']['bspec_version']

        types_content = f'''//! Types for BSpec v{version}
//! Generated from BSpec JSON SDK at {datetime.now().isoformat()}

use serde::{{Deserialize, Serialize}};
use std::collections::HashMap;
use chrono::{{DateTime, Utc}};

/// Metadata contains generation metadata
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Metadata {{
    pub bspec_version: String,
    pub generated_at: DateTime<Utc>,
    pub generator: String,
    pub source_spec: String,
}}

/// Statistics contains specification statistics
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Statistics {{
    pub total_files: usize,
    pub total_document_types: usize,
    pub total_domains: usize,
}}

/// Domain represents a business domain
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Domain {{
    pub name: String,
    pub display_name: String,
    pub description: String,
    pub emoji: String,
    pub document_types: Vec<String>,
    pub document_count: usize,
}}

/// DocumentType represents a BSpec document type
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DocumentType {{
    pub code: String,
    pub name: String,
    pub purpose: String,
    pub domain: String,
    pub examples: Vec<DocumentExample>,
}}

/// DocumentExample represents an example document
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DocumentExample {{
    pub id: String,
    pub title: String,
    pub description: String,
}}

/// File represents a file in the specification
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct File {{
    pub path: String,
    #[serde(rename = "type")]
    pub file_type: String,
    pub size: usize,
    pub content: String,
    pub has_frontmatter: bool,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub frontmatter: Option<serde_json::Value>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub parsed_content: Option<String>,
}}

/// ConformanceLevel represents a conformance level
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConformanceLevel {{
    pub name: String,
    pub display_name: String,
    pub description: String,
    pub emoji: String,
    pub min_documents: usize,
}}

/// YAMLSchema represents the YAML schema definition
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct YAMLSchema {{
    pub required_fields: Vec<String>,
    pub optional_fields: Vec<String>,
    pub field_types: HashMap<String, String>,
    pub enums: HashMap<String, Vec<String>>,
    pub defaults: HashMap<String, serde_json::Value>,
}}

/// DocumentIndexEntry represents an entry in the document index
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DocumentIndexEntry {{
    pub id: String,
    pub title: String,
    #[serde(rename = "type")]
    pub entry_type: String,
    pub path: String,
    pub domain: String,
    pub owner: String,
    pub status: String,
    pub version: String,
}}

/// BSpecData represents the complete BSpec specification data
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct BSpecData {{
    pub metadata: Metadata,
    pub statistics: Statistics,
    pub domains: Vec<Domain>,
    pub document_types: Vec<DocumentType>,
    pub files: HashMap<String, File>,
    pub conformance_levels: Vec<ConformanceLevel>,
    pub yaml_schema: YAMLSchema,
    pub directory_structure: Vec<Vec<String>>,
    pub document_index: Vec<DocumentIndexEntry>,
}}
'''

        output_file = self.output_dir / "src" / "types.rs"
        output_file.write_text(types_content, encoding='utf-8')
        print(f"  📄 Generated src/types.rs")

    def _generate_bspec_rs(self) -> None:
        """Generate bspec.rs file with main BSpec implementation"""
        version = self.bspec_data['metadata']['bspec_version']

        bspec_content = f'''//! Main BSpec implementation for v{version}
//! Generated from BSpec JSON SDK at {datetime.now().isoformat()}

use crate::types::*;
use serde_json;
use std::collections::HashMap;
use thiserror::Error;

/// BSpec error types
#[derive(Error, Debug)]
pub enum BSpecError {{
    #[error("JSON parsing error: {{0}}")]
    JsonError(#[from] serde_json::Error),

    #[error("Domain not found: {{0}}")]
    DomainNotFound(String),

    #[error("Document type not found: {{0}}")]
    DocumentTypeNotFound(String),

    #[error("File not found: {{0}}")]
    FileNotFound(String),
}}

/// Main BSpec structure for working with BSpec data
#[derive(Debug, Clone)]
pub struct BSpec {{
    data: BSpecData,
}}

impl BSpec {{
    /// Create BSpec instance from JSON data
    pub fn from_json(json_data: &[u8]) -> Result<Self, BSpecError> {{
        let data: BSpecData = serde_json::from_slice(json_data)?;
        Ok(Self {{ data }})
    }}

    /// Get BSpec version
    pub fn version(&self) -> &str {{
        &self.data.metadata.bspec_version
    }}

    /// Get generation metadata
    pub fn metadata(&self) -> &Metadata {{
        &self.data.metadata
    }}

    /// Get specification statistics
    pub fn statistics(&self) -> &Statistics {{
        &self.data.statistics
    }}

    /// Get all business domains
    pub fn domains(&self) -> &[Domain] {{
        &self.data.domains
    }}

    /// Get all document types
    pub fn document_types(&self) -> &[DocumentType] {{
        &self.data.document_types
    }}

    /// Get all files in the specification
    pub fn files(&self) -> &HashMap<String, File> {{
        &self.data.files
    }}

    /// Get conformance levels
    pub fn conformance_levels(&self) -> &[ConformanceLevel] {{
        &self.data.conformance_levels
    }}

    /// Get YAML schema definition
    pub fn yaml_schema(&self) -> &YAMLSchema {{
        &self.data.yaml_schema
    }}

    /// Get directory structure
    pub fn directory_structure(&self) -> &[Vec<String>] {{
        &self.data.directory_structure
    }}

    /// Get document index
    pub fn document_index(&self) -> &[DocumentIndexEntry] {{
        &self.data.document_index
    }}

    /// Get a specific domain by name
    pub fn get_domain(&self, name: &str) -> Option<&Domain> {{
        self.data.domains.iter().find(|domain| domain.name == name)
    }}

    /// Get a specific document type by code
    pub fn get_document_type(&self, code: &str) -> Option<&DocumentType> {{
        self.data.document_types.iter().find(|doc_type| doc_type.code == code)
    }}

    /// Get document types for a specific domain
    pub fn get_document_types_for_domain(&self, domain_name: &str) -> Vec<&DocumentType> {{
        if let Some(domain) = self.get_domain(domain_name) {{
            self.data.document_types
                .iter()
                .filter(|doc_type| domain.document_types.contains(&doc_type.code))
                .collect()
        }} else {{
            Vec::new()
        }}
    }}

    /// Get a specific file by path
    pub fn get_file(&self, path: &str) -> Option<&File> {{
        self.data.files.get(path)
    }}

    /// Get all files of a specific type
    pub fn get_files_by_type(&self, file_type: &str) -> Vec<&File> {{
        self.data.files
            .values()
            .filter(|file| file.file_type == file_type)
            .collect()
    }}

    /// Get markdown files with frontmatter
    pub fn get_document_files(&self) -> Vec<&File> {{
        self.data.files
            .values()
            .filter(|file| file.file_type == "markdown" && file.has_frontmatter)
            .collect()
    }}

    /// Search for document types by name, purpose, or code
    pub fn search_document_types(&self, query: &str) -> Vec<&DocumentType> {{
        let query_lower = query.to_lowercase();
        self.data.document_types
            .iter()
            .filter(|doc_type| {{
                doc_type.name.to_lowercase().contains(&query_lower) ||
                doc_type.purpose.to_lowercase().contains(&query_lower) ||
                doc_type.code.to_lowercase().contains(&query_lower)
            }})
            .collect()
    }}

    /// Get domains containing a specific document type
    pub fn get_domains_for_document_type(&self, code: &str) -> Vec<&Domain> {{
        self.data.domains
            .iter()
            .filter(|domain| domain.document_types.contains(&code.to_string()))
            .collect()
    }}

    /// Get conformance level by name
    pub fn get_conformance_level(&self, name: &str) -> Option<&ConformanceLevel> {{
        self.data.conformance_levels
            .iter()
            .find(|level| level.name == name)
    }}

    /// Get required YAML fields
    pub fn get_required_fields(&self) -> &[String] {{
        &self.data.yaml_schema.required_fields
    }}

    /// Get optional YAML fields
    pub fn get_optional_fields(&self) -> &[String] {{
        &self.data.yaml_schema.optional_fields
    }}

    /// Convert back to JSON
    pub fn to_json(&self) -> Result<String, BSpecError> {{
        Ok(serde_json::to_string_pretty(&self.data)?)
    }}

    /// Convert back to JSON bytes
    pub fn to_json_bytes(&self) -> Result<Vec<u8>, BSpecError> {{
        Ok(serde_json::to_vec_pretty(&self.data)?)
    }}
}}

#[cfg(test)]
mod tests {{
    use super::*;

    #[test]
    fn test_bspec_creation() {{
        // This would need actual JSON data to test properly
        // In a real implementation, you'd load test data
    }}
}}
'''

        output_file = self.output_dir / "src" / "bspec.rs"
        output_file.write_text(bspec_content, encoding='utf-8')
        print(f"  📄 Generated src/bspec.rs")

    def _generate_readme(self) -> None:
        """Generate README.md for the Rust SDK"""
        version = self.bspec_data['metadata']['bspec_version']
        stats = self.bspec_data['statistics']

        readme_content = f'''# BSpec Rust SDK

Rust SDK for the BSpec v{version} Universal Business Specification Standard.

## Installation

Add to your `Cargo.toml`:

```toml
[dependencies]
bspec = "{version}"
```

## Usage

### Basic Usage

```rust
use bspec::{{BSpec, load_from_json_str}};
use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {{
    // Load BSpec from JSON file
    let json_content = fs::read_to_string("bspec-complete.json")?;
    let bspec = load_from_json_str(&json_content)?;

    // Get basic information
    println!("BSpec v{{}}", bspec.version());
    println!("Total domains: {{}}", bspec.statistics().total_domains);
    println!("Total document types: {{}}", bspec.statistics().total_document_types);

    Ok(())
}}
```

### Working with Domains

```rust
// Get specific domain
if let Some(strategic_foundation) = bspec.get_domain("Strategic Foundation") {{
    println!("Domain: {{}} - {{}}", strategic_foundation.display_name, strategic_foundation.description);
}}

// Get document types for a domain
let strategic_doc_types = bspec.get_document_types_for_domain("Strategic Foundation");
for doc_type in strategic_doc_types {{
    println!("{{}}: {{}} - {{}}", doc_type.code, doc_type.name, doc_type.purpose);
}}
```

### Working with Document Types

```rust
// Get specific document type
if let Some(mission_type) = bspec.get_document_type("MSN") {{
    println!("{{}}: {{}}", mission_type.name, mission_type.purpose);
}}

// Search document types
let customer_types = bspec.search_document_types("customer");
println!("Found {{}} customer-related types", customer_types.len());
```

### Working with Files

```rust
// Get specific file
if let Some(spec_file) = bspec.get_file("spec.md") {{
    println!("File content length: {{}}", spec_file.content.len());
}}

// Get files by type
let markdown_files = bspec.get_files_by_type("markdown");
println!("Found {{}} markdown files", markdown_files.len());

// Get document files (markdown with frontmatter)
let document_files = bspec.get_document_files();
println!("Found {{}} document files", document_files.len());
```

### Serialization

```rust
// Convert back to JSON
let json_string = bspec.to_json()?;
let json_bytes = bspec.to_json_bytes()?;

// Save to file
fs::write("output.json", json_bytes)?;
```

## Generated Information

- **From**: BSpec v{version} JSON SDK
- **Generated**: {datetime.now().isoformat()}
- **Generator**: rust-generator-v{version}
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
    """Generate Rust SDK from BSpec JSON SDK"""
    import argparse

    parser = argparse.ArgumentParser(description='Generate Rust SDK from BSpec JSON SDK')
    parser.add_argument('--json-sdk-dir', default='sdk/v1/json', help='JSON SDK directory')
    parser.add_argument('--output-dir', default='sdk/v1/rust', help='Output directory')

    args = parser.parse_args()

    try:
        generator = RustGenerator(args.json_sdk_dir, args.output_dir)
        generator.generate_all()
        return 0
    except Exception as e:
        print(f"❌ Error generating Rust SDK: {e}")
        return 1


if __name__ == "__main__":
    exit(main())