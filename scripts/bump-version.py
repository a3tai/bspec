#!/usr/bin/env python3
"""
BSpec Version Bump Script

Simple script to bump semantic versions across all BSpec components.

Usage:
    python3 scripts/bump-version.py patch    # 1.0.0 -> 1.0.1
    python3 scripts/bump-version.py minor    # 1.0.1 -> 1.1.0  
    python3 scripts/bump-version.py major    # 1.1.0 -> 2.0.0
    python3 scripts/bump-version.py --show   # Show current version
"""

import argparse
import json
import re
import sys
from pathlib import Path
from typing import Tuple, List, Dict, Any


class VersionBumper:
    """Handles version bumping across all BSpec components"""
    
    def __init__(self, base_dir: Path = None):
        self.base_dir = base_dir or Path(__file__).parent.parent
        self.version_file = self.base_dir / "spec" / "v1" / "version.txt"
        
    def get_current_version(self) -> str:
        """Get the current BSpec version from spec/v1/version.txt"""
        if not self.version_file.exists():
            raise FileNotFoundError(f"Version file not found: {self.version_file}")
        
        return self.version_file.read_text().strip()
    
    def parse_version(self, version: str) -> Tuple[int, int, int]:
        """Parse semantic version string into major, minor, patch tuple"""
        match = re.match(r'^(\d+)\.(\d+)\.(\d+)$', version)
        if not match:
            raise ValueError(f"Invalid semantic version: {version}")
        
        return tuple(map(int, match.groups()))
    
    def bump_version(self, version: str, bump_type: str) -> str:
        """Bump version according to type (major, minor, patch)"""
        major, minor, patch = self.parse_version(version)
        
        if bump_type == "major":
            return f"{major + 1}.0.0"
        elif bump_type == "minor":
            return f"{major}.{minor + 1}.0"
        elif bump_type == "patch":
            return f"{major}.{minor}.{patch + 1}"
        else:
            raise ValueError(f"Invalid bump type: {bump_type}. Use: major, minor, patch")
    
    def update_version_file(self, new_version: str) -> None:
        """Update the main version file"""
        print(f"📝 Updating {self.version_file.relative_to(self.base_dir)}")
        self.version_file.write_text(new_version)
    
    def update_package_json_files(self, new_version: str) -> List[Path]:
        """Update all package.json files with new version"""
        updated_files = []
        
        package_files = [
            self.base_dir / "sdk" / "v1" / "typescript" / "package.json",
            self.base_dir / "sdk" / "v1" / "json" / "package.json", 
            self.base_dir / "apps" / "web" / "package.json",
            self.base_dir / "apps" / "docs" / "package.json",
            self.base_dir / "apps" / "mcp" / "package.json"
        ]
        
        for package_file in package_files:
            if package_file.exists():
                print(f"📝 Updating {package_file.relative_to(self.base_dir)}")
                
                with open(package_file, 'r') as f:
                    data = json.load(f)
                
                data["version"] = new_version
                
                with open(package_file, 'w') as f:
                    json.dump(data, f, indent=2)
                    f.write('\n')  # Add trailing newline
                
                updated_files.append(package_file)
        
        return updated_files
    
    def update_go_files(self, new_version: str) -> List[Path]:
        """Update Go files with version constants"""
        updated_files = []
        
        go_files = [
            self.base_dir / "sdk" / "v1" / "go" / "bspec.go",
            self.base_dir / "sdk" / "cli" / "internal" / "commands" / "version.go"
        ]
        
        for go_file in go_files:
            if go_file.exists():
                print(f"📝 Updating {go_file.relative_to(self.base_dir)}")
                
                content = go_file.read_text()
                
                # Update version constant
                content = re.sub(
                    r'(Version\s*=\s*)["\'][\d.]+["\']',
                    rf'\g<1>"{new_version}"',
                    content
                )
                
                go_file.write_text(content)
                updated_files.append(go_file)
        
        return updated_files
    
    def update_python_files(self, new_version: str) -> List[Path]:
        """Update Python files with version info"""
        updated_files = []
        
        python_files = [
            self.base_dir / "sdk" / "v1" / "python" / "bspec" / "__init__.py",
            self.base_dir / "sdk" / "v1" / "python" / "setup.py"
        ]
        
        for py_file in python_files:
            if py_file.exists():
                print(f"📝 Updating {py_file.relative_to(self.base_dir)}")
                
                content = py_file.read_text()
                
                # Update __version__ or version=
                content = re.sub(
                    r'(__version__\s*=\s*)["\'][\d.]+["\']',
                    rf'\g<1>"{new_version}"',
                    content
                )
                content = re.sub(
                    r'(version\s*=\s*)["\'][\d.]+["\']',
                    rf'\g<1>"{new_version}"',
                    content
                )
                
                py_file.write_text(content)
                updated_files.append(py_file)
        
        return updated_files
    
    def update_rust_files(self, new_version: str) -> List[Path]:
        """Update Rust Cargo.toml files"""
        updated_files = []
        
        cargo_file = self.base_dir / "sdk" / "v1" / "rust" / "Cargo.toml"
        
        if cargo_file.exists():
            print(f"📝 Updating {cargo_file.relative_to(self.base_dir)}")
            
            content = cargo_file.read_text()
            
            # Update version in [package] section
            content = re.sub(
                r'(^version\s*=\s*")[^"]+(")',
                rf'\g<1>{new_version}\g<2>',
                content,
                flags=re.MULTILINE
            )
            
            cargo_file.write_text(content)
            updated_files.append(cargo_file)
        
        return updated_files
    
    def update_all_versions(self, new_version: str) -> Dict[str, List[Path]]:
        """Update version in all components"""
        print(f"🚀 Updating all components to version {new_version}")
        
        updated_files = {
            "version_file": [self.version_file],
            "package_json": [],
            "go": [],
            "python": [],
            "rust": []
        }
        
        # Update main version file
        self.update_version_file(new_version)
        
        # Update all component files
        updated_files["package_json"] = self.update_package_json_files(new_version)
        updated_files["go"] = self.update_go_files(new_version)
        updated_files["python"] = self.update_python_files(new_version)
        updated_files["rust"] = self.update_rust_files(new_version)
        
        return updated_files
    
    def show_current_version(self) -> None:
        """Display current version information"""
        current = self.get_current_version()
        major, minor, patch = self.parse_version(current)
        
        print(f"📊 Current BSpec Version: {current}")
        print(f"   Major: {major}")
        print(f"   Minor: {minor}")  
        print(f"   Patch: {patch}")
        print()
        print(f"Next versions would be:")
        print(f"   Major: {self.bump_version(current, 'major')}")
        print(f"   Minor: {self.bump_version(current, 'minor')}")
        print(f"   Patch: {self.bump_version(current, 'patch')}")


def main():
    parser = argparse.ArgumentParser(
        description="Bump BSpec version across all components",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  python3 scripts/bump-version.py patch    # Bump patch version
  python3 scripts/bump-version.py minor    # Bump minor version  
  python3 scripts/bump-version.py major    # Bump major version
  python3 scripts/bump-version.py --show   # Show current version
        """
    )
    
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument(
        "bump_type", 
        nargs="?",
        choices=["major", "minor", "patch"],
        help="Type of version bump to perform"
    )
    group.add_argument(
        "--show", 
        action="store_true",
        help="Show current version information"
    )
    
    args = parser.parse_args()
    
    try:
        bumper = VersionBumper()
        
        if args.show:
            bumper.show_current_version()
            return
        
        current_version = bumper.get_current_version()
        new_version = bumper.bump_version(current_version, args.bump_type)
        
        print(f"🔄 Version Bump: {current_version} → {new_version}")
        print()
        
        updated_files = bumper.update_all_versions(new_version)
        
        # Summary
        total_files = sum(len(files) for files in updated_files.values())
        print()
        print(f"✅ Successfully updated {total_files} files")
        print(f"📦 New version: {new_version}")
        print()
        print("Next steps:")
        print("1. Review the changes")
        print("2. Commit the version bump: git add -A && git commit -m 'bump: version {}'".format(new_version))
        print("3. Tag the release: git tag v{}".format(new_version))
        print("4. Push changes: git push && git push --tags")
        
    except Exception as e:
        print(f"❌ Error: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()