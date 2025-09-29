#!/usr/bin/env python3
"""
Manually create bspec-v1.0.0.tgz archive
"""

import tarfile
import os
from pathlib import Path

def create_tgz():
    base_path = Path("/Users/steverude/Documents/BSpec-Foundations")
    json_dir = base_path / "sdk/v1/json"
    tgz_path = json_dir / "bspec-v1.0.0.tgz"

    print(f"Creating TGZ archive at: {tgz_path}")

    # Files to include in the archive
    files_to_add = [
        "bspec-complete.json",
        "version.txt",
        "package.json"
    ]

    with tarfile.open(tgz_path, "w:gz") as tar:
        for filename in files_to_add:
            file_path = json_dir / filename
            if file_path.exists():
                print(f"Adding {filename} to archive...")
                tar.add(file_path, arcname=filename)
            else:
                print(f"Warning: {filename} not found at {file_path}")

    # Check if archive was created
    if tgz_path.exists():
        size = tgz_path.stat().st_size
        print(f"✅ Successfully created bspec-v1.0.0.tgz ({size:,} bytes)")

        # Verify archive contents
        print("\nArchive contents:")
        with tarfile.open(tgz_path, "r:gz") as tar:
            for member in tar.getmembers():
                print(f"  - {member.name} ({member.size} bytes)")

        return True
    else:
        print("❌ Failed to create archive")
        return False

if __name__ == "__main__":
    create_tgz()