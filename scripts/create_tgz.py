#!/usr/bin/env python3
"""
Create BSpec TGZ package from existing JSON SDK
"""

import os
import tarfile
import json
from pathlib import Path

def create_bspec_tgz():
    """Create bspec-v1.0.0.tgz from JSON SDK content"""

    base_dir = Path("/Users/steverude/Documents/BSpec-Foundations")
    json_sdk_dir = base_dir / "sdk/v1/json"

    # Create output directory if it doesn't exist
    json_sdk_dir.mkdir(parents=True, exist_ok=True)

    # Create the TGZ file
    tgz_path = json_sdk_dir / "bspec-v1.0.0.tgz"

    print(f"Creating {tgz_path}...")

    with tarfile.open(tgz_path, "w:gz") as tar:
        # Add the complete JSON file
        json_file = json_sdk_dir / "bspec-complete.json"
        if json_file.exists():
            tar.add(json_file, arcname="bspec-complete.json")
            print(f"Added bspec-complete.json")

        # Add version.txt
        version_file = json_sdk_dir / "version.txt"
        if version_file.exists():
            tar.add(version_file, arcname="version.txt")
            print(f"Added version.txt")

        # Add package.json if it exists
        package_file = json_sdk_dir / "package.json"
        if package_file.exists():
            tar.add(package_file, arcname="package.json")
            print(f"Added package.json")

    if tgz_path.exists():
        size = tgz_path.stat().st_size
        print(f"✅ Created {tgz_path} ({size:,} bytes)")
        return str(tgz_path)
    else:
        print("❌ Failed to create TGZ file")
        return None

if __name__ == "__main__":
    create_bspec_tgz()