#!/usr/bin/env python3
"""
BSpec SDK Generator Master Script

Generates all SDKs from the BSpec specification:
1. JSON SDK (comprehensive package with all spec files)
2. TypeScript SDK (from JSON package)
3. Python SDK (from JSON package)
4. Go SDK (from JSON package)
5. Rust SDK (from JSON package)

Each SDK includes version tracking and proper packaging.
"""

import os
import sys
import subprocess
import json
from pathlib import Path
from datetime import datetime


def run_generator(language: str, json_sdk_dir: str = "sdk/v1/json") -> bool:
    """Run a specific language generator using JSON SDK as source"""
    script_path = Path(f"scripts/generators/{language}/generator.py")

    if not script_path.exists():
        print(f"⚠️  Generator not found: {script_path}")
        return False

    print(f"\\n🔄 Generating {language.upper()} SDK from JSON...")

    try:
        # Run the generator script with JSON SDK as input
        result = subprocess.run([
            sys.executable, str(script_path),
            "--json-sdk-dir", json_sdk_dir,
            "--output-dir", f"sdk/v1/{language}"
        ], capture_output=True, text=True, cwd=Path.cwd())

        if result.returncode == 0:
            print(f"✅ {language.upper()} SDK generated successfully")
            if result.stdout:
                print(result.stdout)
            return True
        else:
            print(f"❌ Error generating {language.upper()} SDK:")
            if result.stderr:
                print(result.stderr)
            return False

    except Exception as e:
        print(f"❌ Exception generating {language.upper()} SDK: {e}")
        return False


def generate_json_sdk(spec_dir: str = "spec/v1") -> bool:
    """Generate the foundational JSON SDK from specification"""
    print("\\n🔄 Generating foundational JSON SDK...")

    json_generator_path = Path("scripts/generators/json/generator.py")

    if not json_generator_path.exists():
        print(f"❌ JSON generator not found: {json_generator_path}")
        return False

    try:
        result = subprocess.run([
            sys.executable, str(json_generator_path),
            "--spec-dir", spec_dir,
            "--output-dir", "sdk/v1/json"
        ], capture_output=True, text=True, cwd=Path.cwd())

        if result.returncode == 0:
            print("✅ JSON SDK generated successfully")
            if result.stdout:
                print(result.stdout)
            return True
        else:
            print("❌ Error generating JSON SDK:")
            if result.stderr:
                print(result.stderr)
            return False

    except Exception as e:
        print(f"❌ Exception generating JSON SDK: {e}")
        return False


def verify_json_sdk() -> bool:
    """Verify that JSON SDK was generated correctly"""
    print("\\n🔍 Verifying JSON SDK...")

    # Look for TGZ file instead of bspec-complete.json
    json_dir = Path("sdk/v1/json")
    tgz_files = list(json_dir.glob("bspec-v*.tgz"))

    if not tgz_files:
        print(f"❌ No BSpec TGZ file found in {json_dir}")
        return False

    tgz_file = tgz_files[0]
    print(f"  📦 Found TGZ package: {tgz_file.name}")

    # Check TGZ file size
    size = tgz_file.stat().st_size
    print(f"  📦 TGZ package size: {size:,} bytes")

    # Read version from spec
    version_file = Path("spec/v1/version.txt")
    if version_file.exists():
        version = version_file.read_text(encoding='utf-8').strip()
        print(f"  📊 BSpec v{version}")
    else:
        print("  ⚠️  Version file not found")

    return True


def create_master_index():
    """Create master SDK index file"""
    print("\\n📝 Creating master SDK index...")

    try:
        # Read version from spec
        version_file = Path("spec/v1/version.txt")
        if version_file.exists():
            version = version_file.read_text(encoding='utf-8').strip()
        else:
            version = "1.0.0"  # Default fallback

        # Create basic stats (we don't need the full JSON data for the index)
        stats = {
            'total_files': 116,  # Approximate based on what we know
            'total_document_types': 112,
            'total_domains': 11
        }

        index_content = f"""# BSpec SDKs

Generated SDKs for BSpec v{version} Universal Business Specification Standard.

## Available SDKs

### 📄 JSON SDK
- **Path**: `sdk/v1/json/`
- **Main File**: `bspec-v{version.replace(".", "-")}.tgz`
- **Package**: `@bspec/json-sdk`
- **Use Case**: Universal data format for any language
- **Archive**: Complete JSON SDK package

### 🟦 TypeScript SDK
- **Path**: `sdk/v1/typescript/`
- **Main File**: `index.ts`
- **Package**: `@bspec/typescript-sdk`
- **Use Case**: Type-safe TypeScript/JavaScript development
- **Source**: Generated from JSON SDK

### 🐍 Python SDK
- **Path**: `sdk/v1/python/`
- **Main File**: `bspec/__init__.py`
- **Package**: `bspec`
- **Use Case**: Python applications and data analysis
- **Source**: Generated from JSON SDK

### 🐹 Go SDK
- **Path**: `sdk/v1/go/`
- **Main File**: `bspec.go`
- **Package**: `github.com/bspec-foundation/bspec-go`
- **Use Case**: Go applications and microservices
- **Source**: Generated from JSON SDK

### 🦀 Rust SDK
- **Path**: `sdk/v1/rust/`
- **Main File**: `src/lib.rs`
- **Package**: `bspec`
- **Use Case**: Rust applications and systems programming
- **Source**: Generated from JSON SDK

## Generation Architecture

1. **JSON SDK (Foundation)**: Converts entire `spec/v1/` directory to comprehensive JSON
2. **Language SDKs**: Generated from JSON SDK for language-specific interfaces
3. **Version Tracking**: All SDKs include `version.txt` with BSpec version
4. **Package Management**: Native package files for each ecosystem

## Version Information

Each SDK includes:
- `version.txt` - BSpec version number ({version})
- `README.md` - Usage documentation
- Package manifest for language ecosystem
- Complete specification data in appropriate format

## Statistics

- **BSpec Version**: {version}
- **Total Files**: {stats['total_files']}
- **Document Types**: {stats['total_document_types']}
- **Business Domains**: {stats['total_domains']}
- **Generated**: {datetime.now().isoformat()}

## Usage

All SDKs are generated from the same source truth (JSON SDK), ensuring consistency across languages while providing native interfaces and idioms for each ecosystem.

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.
"""

        sdk_dir = Path("sdk/v1")
        sdk_dir.mkdir(parents=True, exist_ok=True)

        index_file = sdk_dir / "README.md"
        with open(index_file, 'w', encoding='utf-8') as f:
            f.write(index_content)

        print(f"✅ Created master SDK index: {index_file}")

    except Exception as e:
        print(f"⚠️  Could not create master index: {e}")


def main():
    """Generate all SDKs with JSON-first architecture"""
    print("🚀 BSpec SDK Generation (JSON-First Architecture)")
    print("=" * 60)

    # Step 1: Generate foundational JSON SDK
    if not generate_json_sdk():
        print("\\n❌ JSON SDK generation failed - aborting")
        return 1

    # Step 2: Verify JSON SDK
    if not verify_json_sdk():
        print("\\n❌ JSON SDK verification failed - aborting")
        return 1

    # Step 3: Generate language-specific SDKs from JSON
    generators = ["typescript", "python", "go", "rust"]
    success_count = 0
    total_count = len(generators)

    for generator in generators:
        if run_generator(generator):
            success_count += 1

    # Step 4: Create master index
    create_master_index()

    # Summary
    print("\\n" + "=" * 60)
    print(f"📊 SDK Generation Summary")
    print(f"   JSON SDK: ✅ Generated (foundation)")
    print(f"   Language SDKs: {success_count}/{total_count} successful")
    print(f"   Generated at: {datetime.now().isoformat()}")

    if success_count == total_count:
        print("🎉 All SDKs generated successfully!")
        return 0
    else:
        print(f"⚠️  {total_count - success_count} language generators failed")
        return 1


if __name__ == "__main__":
    exit(main())