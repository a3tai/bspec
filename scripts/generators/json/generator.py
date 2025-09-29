#!/usr/bin/env python3
"""
BSpec JSON SDK Generator

Converts the entire BSpec specification directory (spec/v1/) to JSON format
and packages it as a compressed archive for use by other SDK generators.
"""

import os
import sys
import json
import tarfile
import yaml
import shutil
import tempfile
from pathlib import Path
from datetime import datetime
from typing import Dict, Any, List
import frontmatter


class JSONGenerator:
    """Generates comprehensive JSON SDK from BSpec specification directory"""

    def __init__(self, spec_dir: str = "spec/v1", output_dir: str = "sdk/v1/json"):
        """Initialize generator with input and output paths"""
        self.spec_dir = Path(spec_dir)
        self.output_dir = Path(output_dir)
        self.temp_dir = None  # Will be created as needed

    def _read_version(self) -> str:
        """Read BSpec version from spec/v1/version.txt"""
        version_file = self.spec_dir / "version.txt"
        if not version_file.exists():
            raise FileNotFoundError(f"Version file not found: {version_file}")

        version = version_file.read_text(encoding='utf-8').strip()
        return version

    def generate_all(self) -> None:
        """Generate complete JSON SDK and package as tgz"""
        print("🔄 Converting BSpec specification to JSON...")

        if not self.spec_dir.exists():
            raise FileNotFoundError(f"Specification directory not found: {self.spec_dir}")

        # Read version from spec
        version = self._read_version()
        print(f"  📋 BSpec version: {version}")

        # Create output directory
        self.output_dir.mkdir(parents=True, exist_ok=True)

        # Create temporary directory for JSON files
        self.temp_dir = Path(tempfile.mkdtemp(prefix="bspec_json_"))
        print(f"  📁 Using temp directory: {self.temp_dir}")

        try:
            # Convert all specification files to JSON
            spec_data = self._convert_spec_to_json(version)

            # Generate individual JSON files with same directory structure
            self._generate_individual_json_files(spec_data)

            # Generate version file
            self._generate_version(version)

            # Generate README
            self._generate_readme(spec_data)

            # Create the packaged tgz file
            tgz_file = self._create_package(version)

            print(f"✅ JSON SDK generated")
            print(f"📦 Package created: {tgz_file}")

        finally:
            # Clean up temp directory
            if self.temp_dir and self.temp_dir.exists():
                shutil.rmtree(self.temp_dir)
                print(f"  🧹 Cleaned up temp directory")

    def _convert_spec_to_json(self, version: str) -> Dict[str, Any]:
        """Convert entire specification directory to JSON structure"""
        print("  📖 Reading all specification files...")

        spec_data = {
            "metadata": {
                "bspec_version": version,
                "generated_at": datetime.now().isoformat() + "Z",
                "generator": f"json-generator-v{version}",
                "source_directory": str(self.spec_dir)
            },
            "files": {},
            "directory_structure": [],
            "document_types": []
        }

        # Recursively read all files in the spec directory
        total_files = 0
        for file_path in self.spec_dir.rglob("*"):
            if file_path.is_file() and file_path.suffix in ['.md', '.txt', '.yaml', '.yml']:
                total_files += 1
                relative_path = file_path.relative_to(self.spec_dir)

                try:
                    with open(file_path, 'r', encoding='utf-8') as f:
                        content = f.read()

                    # Try to parse frontmatter
                    has_frontmatter = False
                    frontmatter_data = {}
                    parsed_content = content

                    if content.startswith('---'):
                        try:
                            post = frontmatter.loads(content)
                            has_frontmatter = True
                            frontmatter_data = post.metadata
                            parsed_content = post.content
                        except:
                            # If frontmatter parsing fails, use raw content
                            pass

                    # Store file data
                    file_data = {
                        "path": str(relative_path),
                        "name": file_path.name,
                        "type": self._get_file_type(file_path),
                        "extension": file_path.suffix,
                        "has_frontmatter": has_frontmatter,
                        "frontmatter": frontmatter_data,
                        "content": content,
                        "parsed_content": parsed_content if has_frontmatter else ""
                    }

                    spec_data["files"][str(relative_path)] = file_data

                    # Add to directory structure
                    path_parts = list(relative_path.parts)
                    if path_parts not in spec_data["directory_structure"]:
                        spec_data["directory_structure"].append(path_parts)

                    # Extract document type info from any markdown file
                    if file_path.suffix == '.md':
                        doc_type = self._extract_document_type(content, str(relative_path))
                        if doc_type:
                            spec_data["document_types"].append(doc_type)

                except Exception as e:
                    print(f"  ⚠️  Failed to read {relative_path}: {e}")

        # Update statistics
        spec_data["statistics"] = {
            "total_files": total_files,
            "total_document_types": len(spec_data["document_types"])
        }

        print(f"  📊 Processed {total_files} files")
        print(f"  📋 Found {len(spec_data['document_types'])} document types")

        return spec_data

    def _get_file_type(self, file_path: Path) -> str:
        """Determine file type based on extension"""
        extension = file_path.suffix.lower()
        if extension == '.md':
            return 'markdown'
        elif extension in ['.yml', '.yaml']:
            return 'yaml'
        elif extension == '.txt':
            return 'text'
        else:
            return 'unknown'

    def _extract_document_type(self, content: str, file_path: str) -> Dict[str, Any]:
        """Extract document type information from markdown file"""
        doc_type = {
            "code": "",
            "name": "",
            "purpose": "",
            "domain": "",
            "file_path": file_path
        }

        lines = content.split('\n')

        # Parse content for document type details
        for i, line in enumerate(lines):
            line = line.strip()
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

        # If we didn't get domain from content, extract from directory path
        if not doc_type["domain"]:
            path_parts = file_path.split('/')
            if len(path_parts) > 1:
                domain_dir = path_parts[0]
                doc_type["domain"] = self._format_domain_name(domain_dir)

        return doc_type if doc_type["code"] else None

    def _format_domain_name(self, domain_dir: str) -> str:
        """Convert directory name to formatted domain name"""
        return domain_dir.replace('-', ' ').title()

    def _generate_domains(self, document_types: List[Dict[str, Any]]) -> List[Dict[str, Any]]:
        """Generate domain information from document types"""
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

        return list(domain_map.values())


    def _generate_conformance_levels(self) -> List[Dict[str, Any]]:
        """Generate conformance level definitions from spec files if available"""
        # Return empty list - conformance levels should come from spec files
        return []

    def _generate_yaml_schema(self) -> Dict[str, Any]:
        """Generate YAML schema definition from spec files if available"""
        # Return empty dict - schema should come from actual spec files
        return {}

    def _generate_individual_json_files(self, spec_data: Dict[str, Any]) -> None:
        """Generate individual JSON files maintaining spec/v1 directory structure"""
        print("  📁 Creating JSON files with directory structure...")

        files_generated = 0

        for file_path, file_data in spec_data["files"].items():
            if file_data["type"] == "markdown":
                # Create the JSON equivalent path
                json_path = Path(file_path).with_suffix('.json')
                output_file = self.temp_dir / json_path

                # Create directories if they don't exist
                output_file.parent.mkdir(parents=True, exist_ok=True)

                # Convert file data to clean JSON structure
                json_content = {
                    "metadata": {
                        "source_file": file_path,
                        "file_name": file_data["name"],
                        "generated_at": spec_data["metadata"]["generated_at"],
                        "bspec_version": spec_data["metadata"]["bspec_version"]
                    }
                }

                # Add frontmatter if it exists
                if file_data["has_frontmatter"] and file_data["frontmatter"]:
                    json_content["frontmatter"] = file_data["frontmatter"]

                # Add the content
                json_content["content"] = file_data["content"]

                # If this file has parsed content (from frontmatter), add it
                if file_data["parsed_content"]:
                    json_content["parsed_content"] = file_data["parsed_content"]

                # Write the JSON file
                with open(output_file, 'w', encoding='utf-8') as f:
                    json.dump(json_content, f, indent=2, ensure_ascii=False)

                files_generated += 1

        print(f"  📄 Generated {files_generated} JSON files")

    def _create_package(self, version: str) -> Path:
        """Create TGZ package of the JSON SDK"""
        version_clean = version.replace(".", "-")  # Use hyphens instead of dots
        tgz_file = self.output_dir / f"bspec-v{version_clean}.tgz"

        with tarfile.open(tgz_file, "w:gz") as tar:
            # Add all files from temp directory
            for file_path in self.temp_dir.rglob("*"):
                if file_path.is_file():
                    # Get relative path from temp_dir
                    arcname = file_path.relative_to(self.temp_dir)
                    tar.add(file_path, arcname=str(arcname))

        print(f"  📦 Created package: {tgz_file}")
        return tgz_file

    def _generate_version(self, version: str) -> None:
        """Generate version.txt file"""
        version_file = self.temp_dir / "version.txt"
        version_file.write_text(version, encoding='utf-8')
        print(f"  📄 Generated version.txt")

    def _generate_readme(self, spec_data: Dict[str, Any]) -> None:
        """Generate README.md for the JSON SDK"""
        stats = spec_data["statistics"]
        version = spec_data["metadata"]["bspec_version"]

        readme_content = f"""# BSpec JSON SDK

JSON SDK for the BSpec v{version} Universal Business Specification Standard.

## Overview

This package contains the complete BSpec specification converted to JSON format with the same directory structure as the original `spec/v1` directory. Each markdown file has been converted to its corresponding JSON representation.

## Contents

- Directory structure matching `spec/v1/` with `.json` files instead of `.md`
- `version.txt` - BSpec version identifier
- `README.md` - This documentation
- `bspec-v{version}.tgz` - Compressed archive with all files

## Directory Structure

The JSON files maintain the same organization as the original specification:

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
```

## JSON File Format

Each JSON file contains:

```json
{{
  "metadata": {{
    "source_file": "path/to/original.md",
    "file_name": "original-filename.md",
    "generated_at": "ISO-timestamp",
    "bspec_version": "1.0.0"
  }},
  "frontmatter": {{ ... }},  // If the file had YAML frontmatter
  "content": "Full markdown content...",
  "parsed_content": "Content without frontmatter..."  // If applicable
}}
```

## Statistics

- **BSpec Version**: {version}
- **Total Files**: {stats['total_files']}
- **Document Types**: {stats['total_document_types']}
- **Generated**: {spec_data['metadata']['generated_at']}

## Usage

This JSON SDK serves as the foundation for all language-specific SDKs:

- TypeScript SDK
- Python SDK
- Go SDK
- Rust SDK

Each language SDK reads from this JSON package to provide native interfaces while maintaining the original directory structure.

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.
"""

        readme_file = self.temp_dir / "README.md"
        readme_file.write_text(readme_content, encoding='utf-8')
        print(f"  📄 Generated README.md")


def main():
    """Generate JSON SDK from BSpec specification"""
    import argparse

    parser = argparse.ArgumentParser(description='Generate JSON SDK from BSpec specification')
    parser.add_argument('--spec-dir', default='spec/v1', help='Specification directory')
    parser.add_argument('--output-dir', default='sdk/v1/json', help='Output directory')

    args = parser.parse_args()

    try:
        generator = JSONGenerator(args.spec_dir, args.output_dir)
        generator.generate_all()
        return 0
    except Exception as e:
        print(f"❌ Error generating JSON SDK: {e}")
        return 1


if __name__ == "__main__":
    exit(main())