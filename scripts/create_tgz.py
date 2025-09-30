#!/usr/bin/env python3
"""
Create comprehensive bspec-v1.0.0.tgz archive from spec/v1
Includes all specification files in both Markdown and JSON formats
"""

import tarfile
import os
import json
import tempfile
import shutil
import re
from pathlib import Path
from typing import Dict, Any, List
import yaml

def parse_markdown_to_json(md_path: Path) -> Dict[str, Any]:
    """Convert a markdown specification document to structured JSON"""
    with open(md_path, 'r', encoding='utf-8') as f:
        content = f.read()

    # Extract metadata from the content
    lines = content.split('\n')
    title = ""
    doc_type = ""
    domain = ""
    version = ""
    status = ""

    # Parse the structured header information
    for line in lines[:20]:  # Check first 20 lines for metadata
        if line.startswith('# ') and not title:
            title = line[2:].strip()
        elif line.startswith('**Document Type Code:**'):
            doc_type = line.split(':', 1)[1].strip()
        elif line.startswith('**Domain:**'):
            domain = line.split(':', 1)[1].strip()
        elif line.startswith('**Version:**'):
            version = line.split(':', 1)[1].strip()
        elif line.startswith('**Status:**'):
            status = line.split(':', 1)[1].strip()

    # Structure the JSON output
    json_doc = {
        "title": title,
        "type": doc_type,
        "domain": domain,
        "version": version,
        "status": status,
        "source_file": md_path.name,
        "content": {
            "raw_markdown": content,
            "sections": []
        }
    }

    # Extract sections (lines starting with ##)
    current_section = None
    section_content = []

    for line in lines:
        if line.startswith('## '):
            if current_section:
                json_doc["content"]["sections"].append({
                    "heading": current_section,
                    "content": '\n'.join(section_content).strip()
                })
            current_section = line[3:].strip()
            section_content = []
        else:
            section_content.append(line)

    # Add the last section
    if current_section:
        json_doc["content"]["sections"].append({
            "heading": current_section,
            "content": '\n'.join(section_content).strip()
        })

    return json_doc

def copy_directory_structure(src: Path, dst: Path):
    """Copy all files from src to dst, preserving directory structure"""
    if not src.exists():
        raise FileNotFoundError(f"Source directory not found: {src}")

    dst.mkdir(parents=True, exist_ok=True)

    total_files = 0
    processed_files = 0

    # Count total files first
    for item in src.rglob('*'):
        if item.is_file():
            total_files += 1

    print(f"Processing {total_files} files from {src}...")

    for item in src.rglob('*'):
        if item.is_file():
            # Calculate relative path from source
            rel_path = item.relative_to(src)
            dest_path = dst / rel_path

            # Create parent directories if they don't exist
            dest_path.parent.mkdir(parents=True, exist_ok=True)

            # Copy file
            shutil.copy2(item, dest_path)
            processed_files += 1

            if processed_files % 10 == 0 or processed_files == total_files:
                print(f"  Copied {processed_files}/{total_files} files...")

    return processed_files

def create_json_versions(markdown_dir: Path, json_dir: Path):
    """Create JSON versions of all markdown files"""
    json_dir.mkdir(parents=True, exist_ok=True)

    md_files = list(markdown_dir.rglob('*.md'))
    print(f"Converting {len(md_files)} markdown files to JSON...")

    processed = 0
    for md_file in md_files:
        try:
            # Calculate relative path and create JSON equivalent
            rel_path = md_file.relative_to(markdown_dir)
            json_path = json_dir / rel_path.with_suffix('.json')

            # Create parent directories
            json_path.parent.mkdir(parents=True, exist_ok=True)

            # Convert to JSON
            json_data = parse_markdown_to_json(md_file)

            # Write JSON file
            with open(json_path, 'w', encoding='utf-8') as f:
                json.dump(json_data, f, indent=2, ensure_ascii=False)

            processed += 1
            if processed % 10 == 0 or processed == len(md_files):
                print(f"  Converted {processed}/{len(md_files)} files...")

        except Exception as e:
            print(f"  Warning: Failed to convert {md_file}: {e}")

    return processed

def create_manifest(work_dir: Path, md_files: int, json_files: int) -> Path:
    """Create a manifest file describing the archive contents"""
    manifest = {
        "name": "BSpec v1.0 Specification",
        "version": "1.0.0",
        "created": "2024-01-01",
        "description": "Complete Business Specification Standard v1.0",
        "contents": {
            "markdown_files": md_files,
            "json_files": json_files,
            "total_files": md_files + json_files
        },
        "structure": {
            "markdown/": "Original specification files in Markdown format",
            "json/": "Structured JSON versions of all specification files",
            "version.txt": "Version identifier",
            "manifest.json": "This manifest file"
        }
    }

    manifest_path = work_dir / "manifest.json"
    with open(manifest_path, 'w', encoding='utf-8') as f:
        json.dump(manifest, f, indent=2, ensure_ascii=False)

    return manifest_path

def create_tgz():
    base_path = Path.cwd()  # Use current working directory
    spec_dir = base_path / "spec/v1"
    output_dir = base_path / "sdk/v1/json"
    tgz_path = output_dir / "bspec-v1-0-0.tgz"

    output_dir.mkdir(parents=True, exist_ok=True)

    print(f"🚀 Creating comprehensive BSpec v1.0 archive...")
    print(f"   Source: {spec_dir}")
    print(f"   Output: {tgz_path}")

    # Create temporary working directory
    with tempfile.TemporaryDirectory() as temp_dir:
        work_dir = Path(temp_dir)
        markdown_dir = work_dir / "markdown"
        json_dir = work_dir / "json"

        try:
            # Copy all files from spec/v1 to markdown/
            print("\n📁 Copying specification files...")
            md_files = copy_directory_structure(spec_dir, markdown_dir)

            # Create JSON versions
            print("\n🔄 Converting to JSON format...")
            json_files = create_json_versions(markdown_dir, json_dir)

            # Copy version.txt to root if it exists
            version_file = spec_dir / "version.txt"
            if version_file.exists():
                shutil.copy2(version_file, work_dir / "version.txt")

            # Create manifest
            print("\n📋 Creating manifest...")
            create_manifest(work_dir, md_files, json_files)

            # Create the archive
            print(f"\n📦 Creating archive: {tgz_path}")
            with tarfile.open(tgz_path, "w:gz") as tar:
                # Add all files from work directory
                for item in work_dir.rglob('*'):
                    if item.is_file():
                        arcname = item.relative_to(work_dir)
                        tar.add(item, arcname=str(arcname))

            # Verify archive was created
            if tgz_path.exists():
                size = tgz_path.stat().st_size
                print(f"\n✅ Successfully created {tgz_path.name} ({size:,} bytes)")

                # Show archive contents summary
                print("\n📊 Archive contents:")
                with tarfile.open(tgz_path, "r:gz") as tar:
                    members = tar.getmembers()
                    dirs = [m for m in members if m.isdir()]
                    files = [m for m in members if m.isfile()]

                    print(f"   📂 Directories: {len(dirs)}")
                    print(f"   📄 Files: {len(files)}")
                    print(f"   💾 Total size: {sum(m.size for m in files):,} bytes")

                    # Show directory structure
                    print(f"\n📋 Directory structure:")
                    root_items = set()
                    for member in members[:20]:  # Show first 20 items
                        path_parts = Path(member.name).parts
                        if path_parts:
                            root_items.add(path_parts[0])

                    for item in sorted(root_items):
                        print(f"   📁 {item}")

                return True
            else:
                print("❌ Failed to create archive")
                return False

        except Exception as e:
            print(f"❌ Error creating archive: {e}")
            return False

if __name__ == "__main__":
    create_tgz()