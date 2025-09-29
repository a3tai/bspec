#!/usr/bin/env python3
"""
BSpec TypeScript SDK Generator

Generates TypeScript SDK interfaces and types from the BSpec JSON SDK
using the JSON data as the source of truth.
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

class TypeScriptGenerator:
    """Generates TypeScript SDK from BSpec JSON SDK"""

    def __init__(self, json_sdk_dir: str = "sdk/v1/json", output_dir: str = "sdk/v1/typescript"):
        """Initialize generator with JSON SDK input and TypeScript output paths"""
        self.json_sdk_dir = Path(json_sdk_dir)
        self.output_dir = Path(output_dir)

    def generate_all(self) -> None:
        """Generate complete TypeScript SDK from JSON SDK"""
        print("🔄 Loading BSpec JSON SDK...")
        json_data = self._load_json_data()

        print("🎯 Generating TypeScript SDK...")

        # Create output directories
        self.output_dir.mkdir(parents=True, exist_ok=True)
        (self.output_dir / "interfaces").mkdir(exist_ok=True)
        (self.output_dir / "types").mkdir(exist_ok=True)

        # Generate base interface
        self._generate_base_interface(json_data)

        # Generate document type interfaces
        self._generate_document_interfaces(json_data)

        # Generate index file
        self._generate_index(json_data)

        # Generate enums and constants
        self._generate_constants(json_data)

        # Copy JSON data file
        self._copy_json_data(json_data)

        # Generate version file
        self._generate_version_file(json_data)

        # Generate package.json
        self._generate_package_json(json_data)

        print(f"✅ TypeScript SDK generated in {self.output_dir}")

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
                "metadata": {},
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
            json_data["total_document_types"] = len(document_types)
            json_data["total_domains"] = len(domains)

            return json_data

        finally:
            # Clean up temporary directory
            shutil.rmtree(temp_dir)

    def _generate_base_interface(self, json_data: Dict[str, Any]) -> None:
        """Generate base document interface"""
        content = f'''/**
 * BSpec TypeScript SDK - Base Interfaces
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v{json_data['bspec_version']} specification
 * Generated at: {datetime.now().isoformat()}
 * Generator: typescript-generator-v1.0.0
 */

// Base document interface
export interface BaseBSpecDocument {{
  id: string;
  title: string;
  type: DocumentType;
  status: DocumentStatus;
  version: string;
  owner: string;
  domain: BusinessDomain;
  depends_on?: string[];
  enables?: string[];
  related?: string[];
  stakeholders?: string[];
  priority?: Priority;
  review_cycle?: string;
  constraints?: string[];
  assumptions?: string[];
  success_criteria?: string[];
}}

// Document types
export type DocumentType = {' | '.join([f'"{dt["code"]}"' for dt in json_data['document_types']])};

// Document status
export type DocumentStatus = "Draft" | "Review" | "Approved" | "Active" | "Deprecated";

// Business domains
export type BusinessDomain = {' | '.join([f'"{domain["name"]}"' for domain in json_data['domains']])};

// Priority levels
export type Priority = "Critical" | "High" | "Medium" | "Low";

// Conformance levels
export type ConformanceLevel = {' | '.join([f'"{level["name"]}"' for level in json_data['conformance_levels']])};

// Changelog entry
export interface ChangelogEntry {{
  version: string;
  date: string;
  changes: string[];
  author: string;
}}

// BSpec metadata
export interface BSpecMetadata {{
  version: string;
  generated_at: string;
  generator: string;
  total_document_types: number;
  total_domains: number;
}}
'''

        output_file = self.output_dir / "base.ts"
        output_file.write_text(content, encoding='utf-8')
        print(f"  📄 Generated base.ts")

    def _generate_document_interfaces(self, json_data: Dict[str, Any]) -> None:
        """Generate individual document type interfaces"""

        for doc_type in json_data['document_types']:
            interface_name = f"{doc_type['code']}Document"

            content = f'''/**
 * {doc_type['name']} Document Interface
 *
 * GENERATED CODE - DO NOT EDIT MANUALLY
 * Generated from BSpec v{json_data['bspec_version']} specification
 */

import {{ BaseBSpecDocument }} from '../base';

/**
 * {doc_type['name']} ({doc_type['code']})
 *
 * Domain: {doc_type['domain']}
 * Purpose: {doc_type['purpose']}
 */
export interface {interface_name} extends BaseBSpecDocument {{
  type: "{doc_type['code']}";
  domain: "{doc_type['domain']}";

  // Content sections
  executive_summary?: string;
  framework?: any;
  implementation?: any;
  validation?: any;
  quality_standards?: {{
    bronze?: string[];
    silver?: string[];
    gold?: string[];
  }};

  // Domain-specific fields can be added here
  [key: string]: any;
}}

export default {interface_name};
'''

            output_file = self.output_dir / "interfaces" / f"{doc_type['code'].lower()}.ts"
            output_file.write_text(content, encoding='utf-8')

        print(f"  📄 Generated {len(json_data['document_types'])} document interfaces")

    def _generate_index(self, json_data: Dict[str, Any]) -> None:
        """Generate index.ts file that exports everything"""
        lines = [
            "/**",
            " * BSpec TypeScript SDK",
            " *",
            " * GENERATED CODE - DO NOT EDIT MANUALLY",
            f" * Generated from BSpec v{json_data['bspec_version']} specification",
            f" * Generated at: {datetime.now().isoformat()}",
            " * Generator: typescript-generator-v1.0.0",
            " */",
            "",
            "// Export base interface and types",
            "export * from './base';",
            "export * from './constants';",
            "",
            "// Export document type interfaces",
        ]

        # Add exports for each document type interface
        for doc_type in json_data['document_types']:
            lines.append(f"export * from './interfaces/{doc_type['code'].lower()}';")

        # Add utility exports
        lines.extend([
            "",
            "// Re-export main types for convenience",
            "export type {",
            "  BaseBSpecDocument,",
            "  DocumentType,",
            "  DocumentStatus,",
            "  BusinessDomain,",
            "  ConformanceLevel,",
            "  Priority,",
            "  ChangelogEntry,",
            "  BSpecMetadata",
            "} from './base';",
            "",
            "// Document type union",
            "export type AnyBSpecDocument = " + " | ".join([
                f"{doc_type['code']}Document" for doc_type in json_data['document_types']
            ]) + ";",
            "",
            "// Export JSON data",
            "export { default as BSpecData } from './bspec-data.json';",
            ""
        ])

        output = "\\n".join(lines)
        output_file = self.output_dir / "index.ts"
        output_file.write_text(output, encoding='utf-8')
        print(f"  📄 Generated index.ts")

    def _generate_constants(self, json_data: Dict[str, Any]) -> None:
        """Generate constants and enum files"""
        lines = [
            "/**",
            " * BSpec Constants and Enums",
            " *",
            " * GENERATED CODE - DO NOT EDIT MANUALLY",
            f" * Generated from BSpec v{json_data['bspec_version']} specification",
            f" * Generated at: {datetime.now().isoformat()}",
            " */",
            "",
            "// Document type codes",
            "export const DOCUMENT_TYPES = {",
        ]

        # Add document type constants
        for doc_type in json_data['document_types']:
            lines.append(f"  {doc_type['code']}: '{doc_type['code']}' as const,")

        lines.extend([
            "} as const;",
            "",
            "// Domain names",
            "export const DOMAINS = {",
        ])

        # Add domain constants
        for domain in json_data['domains']:
            domain_key = domain['name'].upper().replace(' ', '_').replace('&', 'AND')
            lines.append(f"  {domain_key}: '{domain['name']}' as const,")

        lines.extend([
            "} as const;",
            "",
            "// Conformance levels",
            "export const CONFORMANCE_LEVELS = {",
        ])

        # Add conformance level constants
        for level in json_data['conformance_levels']:
            level_key = level['name'].upper()
            lines.append(f"  {level_key}: '{level['name']}' as const,")

        lines.extend([
            "} as const;",
            "",
            "// Document type metadata",
            "export const DOCUMENT_TYPE_INFO = {",
        ])

        # Add document type info
        for doc_type in json_data['document_types']:
            lines.append(f"  {doc_type['code']}: {{")
            lines.append(f"    code: '{doc_type['code']}',")
            lines.append(f"    name: '{doc_type['name']}',")
            lines.append(f"    purpose: '{doc_type['purpose']}',")
            lines.append(f"    domain: '{doc_type['domain']}'")
            lines.append(f"  }},")

        lines.extend([
            "} as const;",
            ""
        ])

        output = "\\n".join(lines)
        output_file = self.output_dir / "constants.ts"
        output_file.write_text(output, encoding='utf-8')
        print(f"  📄 Generated constants.ts")

    def _copy_json_data(self, json_data: Dict[str, Any]) -> None:
        """Copy JSON data file to TypeScript SDK"""
        output_file = self.output_dir / "bspec-data.json"
        with open(output_file, 'w', encoding='utf-8') as f:
            json.dump(json_data, f, indent=2, ensure_ascii=False)

        print(f"  📄 Generated bspec-data.json")

    def _generate_version_file(self, json_data: Dict[str, Any]) -> None:
        """Generate version.txt file"""
        version_content = f"{json_data['bspec_version']}\\n"

        output_file = self.output_dir / "version.txt"
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(version_content)

        print(f"  📄 Generated version.txt")

    def _generate_package_json(self, json_data: Dict[str, Any]) -> None:
        """Generate package.json for npm publishing"""
        package_info = {
            "name": "@bspec/typescript-sdk",
            "version": json_data['bspec_version'],
            "description": "BSpec TypeScript SDK - Universal Business Specification Standard",
            "main": "index.js",
            "types": "index.d.ts",
            "files": [
                "**/*.ts",
                "**/*.js",
                "**/*.d.ts",
                "*.json",
                "version.txt",
                "README.md"
            ],
            "scripts": {
                "build": "tsc",
                "test": "jest"
            },
            "keywords": [
                "bspec",
                "business-specification",
                "typescript",
                "types",
                "documentation"
            ],
            "author": "BSpec Foundation",
            "license": "CC-BY-4.0",
            "repository": {
                "type": "git",
                "url": "https://github.com/a3tai/bspec.git",
                "directory": "sdk/v1/typescript"
            },
            "bugs": {
                "url": "https://github.com/a3tai/bspec/issues"
            },
            "homepage": "https://bspec.dev",
            "devDependencies": {
                "typescript": "^5.0.0",
                "@types/node": "^20.0.0"
            },
            "generated": {
                "at": datetime.now().isoformat(),
                "from": f"BSpec v{json_data['bspec_version']}",
                "generator": "typescript-generator-v1.0.0"
            }
        }

        output_file = self.output_dir / "package.json"
        with open(output_file, 'w', encoding='utf-8') as f:
            json.dump(package_info, f, indent=2, ensure_ascii=False)

        print(f"  📄 Generated package.json")

        # Generate README
        readme_content = f"""# BSpec TypeScript SDK

TypeScript interfaces and types for the BSpec v{json_data['bspec_version']} Universal Business Specification Standard.

## Installation

```bash
npm install @bspec/typescript-sdk
```

## Usage

```typescript
import {{ BaseBSpecDocument, MSNDocument, DOCUMENT_TYPES }} from '@bspec/typescript-sdk';

// Use type-safe document interfaces
const mission: MSNDocument = {{
  id: 'MSN-company-mission',
  title: 'Company Mission Statement',
  type: DOCUMENT_TYPES.MSN,
  status: 'Approved',
  version: '1.0.0',
  owner: 'executive-team',
  domain: 'Strategic Foundation',
  // ... other fields
}};
```

## Contents

- `base.ts` - Base interfaces and types
- `interfaces/` - Individual document type interfaces
- `constants.ts` - Type-safe constants and enums
- `bspec-data.json` - Complete specification data
- `version.txt` - BSpec version number

## Document Types

This SDK includes TypeScript interfaces for all {json_data['total_document_types']} BSpec document types across {json_data['total_domains']} business domains.

## Generated

- **From**: BSpec v{json_data['bspec_version']} specification
- **At**: {datetime.now().isoformat()}
- **Generator**: typescript-generator-v1.0.0

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.
"""

        readme_file = self.output_dir / "README.md"
        with open(readme_file, 'w', encoding='utf-8') as f:
            f.write(readme_content)

        print(f"  📄 Generated README.md")


def main():
    """Generate TypeScript SDK from JSON SDK"""
    import argparse

    parser = argparse.ArgumentParser(description='Generate TypeScript SDK from BSpec JSON SDK')
    parser.add_argument('--json-sdk-dir', default='sdk/v1/json', help='JSON SDK directory')
    parser.add_argument('--output-dir', default='sdk/v1/typescript', help='Output directory')

    args = parser.parse_args()

    try:
        generator = TypeScriptGenerator(args.json_sdk_dir, args.output_dir)
        generator.generate_all()
        return 0
    except Exception as e:
        print(f"❌ Error generating TypeScript SDK: {e}")
        return 1


if __name__ == "__main__":
    exit(main())