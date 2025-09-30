#!/usr/bin/env python3
"""
BSpec Version Synchronization Script

This script manages version synchronization across all BSpec components,
ensuring proper semantic versioning and component alignment.

Usage:
    python3 scripts/version-sync.py --help
    python3 scripts/version-sync.py --sync-all
    python3 scripts/version-sync.py --update-spec 1.2.0
    python3 scripts/version-sync.py --validate
"""

import json
import os
import re
import sys
import argparse
from pathlib import Path
from typing import Dict, List, Tuple, Optional
import subprocess

# Base directory for the BSpec project
BASE_DIR = Path(__file__).parent.parent
SPEC_VERSION_FILE = BASE_DIR / "spec" / "v1" / "version.txt"

class VersionManager:
    """Manages version synchronization across BSpec components."""
    
    def __init__(self):
        self.spec_version = self.read_spec_version()
        self.component_versions = {}
        
    def read_spec_version(self) -> str:
        """Read the current specification version."""
        if not SPEC_VERSION_FILE.exists():
            raise FileNotFoundError(f"Spec version file not found: {SPEC_VERSION_FILE}")
        
        return SPEC_VERSION_FILE.read_text().strip()
    
    def update_spec_version(self, version: str) -> None:
        """Update the specification version."""
        if not self.is_valid_semver(version):
            raise ValueError(f"Invalid semantic version: {version}")
        
        SPEC_VERSION_FILE.write_text(version)
        self.spec_version = version
        print(f"✓ Updated specification version to {version}")
    
    def is_valid_semver(self, version: str) -> bool:
        """Validate semantic version format."""
        pattern = r'^(\d+)\.(\d+)\.(\d+)(?:-([a-zA-Z0-9\-\.]+))?(?:\+([a-zA-Z0-9\-\.]+))?$'
        return bool(re.match(pattern, version))
    
    def parse_semver(self, version: str) -> Tuple[int, int, int, Optional[str], Optional[str]]:
        """Parse semantic version into components."""
        pattern = r'^(\d+)\.(\d+)\.(\d+)(?:-([a-zA-Z0-9\-\.]+))?(?:\+([a-zA-Z0-9\-\.]+))?$'
        match = re.match(pattern, version)
        if not match:
            raise ValueError(f"Invalid semantic version: {version}")
        
        major, minor, patch, prerelease, build = match.groups()
        return int(major), int(minor), int(patch), prerelease, build
    
    def get_component_versions(self) -> Dict[str, str]:
        """Get current versions of all BSpec components."""
        versions = {}
        
        # TypeScript SDK
        ts_package = BASE_DIR / "sdk" / "v1" / "typescript" / "package.json"
        if ts_package.exists():
            data = json.loads(ts_package.read_text())
            versions["typescript-sdk"] = data.get("version", "unknown")
        
        # Python SDK
        python_setup = BASE_DIR / "sdk" / "v1" / "python" / "setup.py"
        if python_setup.exists():
            content = python_setup.read_text()
            match = re.search(r'version=["\']([^"\']+)["\']', content)
            versions["python-sdk"] = match.group(1) if match else "unknown"
        
        # Go SDK (from version constant)
        go_bspec = BASE_DIR / "sdk" / "v1" / "go" / "bspec.go"
        if go_bspec.exists():
            content = go_bspec.read_text()
            match = re.search(r'Version\s*=\s*["\']([^"\']+)["\']', content)
            versions["go-sdk"] = match.group(1) if match else "unknown"
        
        # Rust SDK
        rust_cargo = BASE_DIR / "sdk" / "v1" / "rust" / "Cargo.toml"
        if rust_cargo.exists():
            content = rust_cargo.read_text()
            match = re.search(r'version\s*=\s*["\']([^"\']+)["\']', content)
            versions["rust-sdk"] = match.group(1) if match else "unknown"
        
        # JSON Schema
        json_package = BASE_DIR / "sdk" / "v1" / "json" / "package.json"
        if json_package.exists():
            data = json.loads(json_package.read_text())
            versions["json-schema"] = data.get("version", "unknown")
        
        # CLI Tool
        cli_version = BASE_DIR / "sdk" / "cli" / "internal" / "commands" / "version.go"
        if cli_version.exists():
            content = cli_version.read_text()
            match = re.search(r'Version\s*=\s*["\']([^"\']+)["\']', content)
            versions["cli"] = match.group(1) if match else "unknown"
        
        # Web App
        web_package = BASE_DIR / "apps" / "web" / "package.json"
        if web_package.exists():
            data = json.loads(web_package.read_text())
            versions["web-app"] = data.get("version", "unknown")
        
        # MCP Server
        mcp_package = BASE_DIR / "apps" / "mcp" / "package.json"
        if mcp_package.exists():
            data = json.loads(mcp_package.read_text())
            versions["mcp-server"] = data.get("version", "unknown")
        
        # Documentation Site
        docs_package = BASE_DIR / "apps" / "docs" / "package.json"
        if docs_package.exists():
            data = json.loads(docs_package.read_text())
            versions["docs-site"] = data.get("version", "unknown")
        
        # Specification
        versions["specification"] = self.spec_version
        
        return versions
    
    def update_typescript_sdk_version(self, version: str) -> None:
        """Update TypeScript SDK version."""
        package_path = BASE_DIR / "sdk" / "v1" / "typescript" / "package.json"
        if not package_path.exists():
            print(f"⚠ TypeScript package.json not found: {package_path}")
            return
        
        data = json.loads(package_path.read_text())
        data["version"] = version
        
        with open(package_path, 'w') as f:
            json.dump(data, f, indent=2)
        
        print(f"✓ Updated TypeScript SDK to {version}")
    
    def update_python_sdk_version(self, version: str) -> None:
        """Update Python SDK version."""
        setup_path = BASE_DIR / "sdk" / "v1" / "python" / "setup.py"
        init_path = BASE_DIR / "sdk" / "v1" / "python" / "bspec" / "__init__.py"
        
        # Update setup.py
        if setup_path.exists():
            content = setup_path.read_text()
            content = re.sub(r'version=["\'][^"\']+["\']', f'version="{version}"', content)
            setup_path.write_text(content)
            print(f"✓ Updated Python setup.py to {version}")
        
        # Update __init__.py
        if init_path.exists():
            content = init_path.read_text()
            content = re.sub(r'__version__\s*=\s*["\'][^"\']+["\']', f'__version__ = "{version}"', content)
            init_path.write_text(content)
            print(f"✓ Updated Python __init__.py to {version}")
    
    def update_go_sdk_version(self, version: str) -> None:
        """Update Go SDK version."""
        bspec_path = BASE_DIR / "sdk" / "v1" / "go" / "bspec.go"
        if not bspec_path.exists():
            print(f"⚠ Go bspec.go not found: {bspec_path}")
            return
        
        content = bspec_path.read_text()
        content = re.sub(r'Version\s*=\s*["\'][^"\']+["\']', f'Version = "{version}"', content)
        bspec_path.write_text(content)
        print(f"✓ Updated Go SDK to {version}")
    
    def update_rust_sdk_version(self, version: str) -> None:
        """Update Rust SDK version."""
        cargo_path = BASE_DIR / "sdk" / "v1" / "rust" / "Cargo.toml"
        if not cargo_path.exists():
            print(f"⚠ Rust Cargo.toml not found: {cargo_path}")
            return
        
        content = cargo_path.read_text()
        content = re.sub(r'version\s*=\s*["\'][^"\']+["\']', f'version = "{version}"', content, count=1)
        cargo_path.write_text(content)
        print(f"✓ Updated Rust SDK to {version}")
    
    def update_json_schema_version(self, version: str) -> None:
        """Update JSON Schema version."""
        package_path = BASE_DIR / "sdk" / "v1" / "json" / "package.json"
        if not package_path.exists():
            print(f"⚠ JSON package.json not found: {package_path}")
            return
        
        data = json.loads(package_path.read_text())
        data["version"] = version
        
        with open(package_path, 'w') as f:
            json.dump(data, f, indent=2)
        
        print(f"✓ Updated JSON Schema to {version}")
    
    def update_cli_version(self, version: str) -> None:
        """Update CLI tool version."""
        version_path = BASE_DIR / "sdk" / "cli" / "internal" / "commands" / "version.go"
        if not version_path.exists():
            print(f"⚠ CLI version.go not found: {version_path}")
            return
        
        content = version_path.read_text()
        content = re.sub(r'Version\s*=\s*["\'][^"\']+["\']', f'Version   = "{version}"', content)
        version_path.write_text(content)
        print(f"✓ Updated CLI to {version}")
    
    def update_mcp_server_version(self, version: str) -> None:
        """Update MCP Server version."""
        package_path = BASE_DIR / "apps" / "mcp" / "package.json"
        version_path = BASE_DIR / "apps" / "mcp" / "version.txt"
        
        # Update package.json
        if package_path.exists():
            data = json.loads(package_path.read_text())
            data["version"] = version
            
            with open(package_path, 'w') as f:
                json.dump(data, f, indent=2)
            print(f"✓ Updated MCP Server package.json to {version}")
        
        # Update version.txt
        if version_path.exists():
            version_path.write_text(version)
            print(f"✓ Updated MCP Server version.txt to {version}")
    
    def update_web_app_version(self, version: str) -> None:
        """Update Web App version."""
        package_path = BASE_DIR / "apps" / "web" / "package.json"
        if not package_path.exists():
            print(f"⚠ Web App package.json not found: {package_path}")
            return
        
        data = json.loads(package_path.read_text())
        data["version"] = version
        
        with open(package_path, 'w') as f:
            json.dump(data, f, indent=2)
        
        print(f"✓ Updated Web App to {version}")
    
    def update_docs_site_version(self, version: str) -> None:
        """Update Documentation Site version."""
        package_path = BASE_DIR / "apps" / "docs" / "package.json"
        if not package_path.exists():
            print(f"⚠ Docs package.json not found: {package_path}")
            return
        
        data = json.loads(package_path.read_text())
        data["version"] = version
        
        with open(package_path, 'w') as f:
            json.dump(data, f, indent=2)
        
        print(f"✓ Updated Documentation Site to {version}")
    
    def sync_sdk_versions_to_spec(self) -> None:
        """Synchronize SDK versions to match specification version."""
        print(f"🔄 Synchronizing SDK versions to specification {self.spec_version}")
        
        # Components that should match spec version exactly
        self.update_typescript_sdk_version(self.spec_version)
        self.update_python_sdk_version(self.spec_version)
        self.update_go_sdk_version(self.spec_version)
        self.update_rust_sdk_version(self.spec_version)
        self.update_json_schema_version(self.spec_version)
        
        # MCP Server follows spec for compatibility
        self.update_mcp_server_version(self.spec_version)
        
        print("✅ SDK synchronization complete")
    
    def validate_versions(self) -> List[str]:
        """Validate version consistency and return issues."""
        issues = []
        versions = self.get_component_versions()
        
        # Check that critical components align with specification
        aligned_components = [
            "typescript-sdk", "python-sdk", "go-sdk", 
            "rust-sdk", "json-schema", "mcp-server"
        ]
        
        for component in aligned_components:
            if component in versions:
                component_version = versions[component]
                if component_version != self.spec_version:
                    # For SDKs, major.minor should match at minimum
                    spec_major, spec_minor, spec_patch, _, _ = self.parse_semver(self.spec_version)
                    try:
                        comp_major, comp_minor, comp_patch, _, _ = self.parse_semver(component_version)
                        if comp_major != spec_major or comp_minor != spec_minor:
                            issues.append(f"{component}: {component_version} (expected major.minor {spec_major}.{spec_minor}.x)")
                    except ValueError:
                        issues.append(f"{component}: invalid version format '{component_version}'")
        
        # Validate semantic version format
        for component, version in versions.items():
            if version != "unknown" and not self.is_valid_semver(version):
                issues.append(f"{component}: invalid semantic version format '{version}'")
        
        return issues
    
    def print_version_report(self) -> None:
        """Print a formatted version report."""
        versions = self.get_component_versions()
        
        print("\n🔍 BSpec Component Versions")
        print("=" * 50)
        
        # Group by category
        categories = {
            "Core": ["specification"],
            "SDKs": ["typescript-sdk", "python-sdk", "go-sdk", "rust-sdk", "json-schema"],
            "Tools": ["cli"],
            "Applications": ["web-app", "mcp-server", "docs-site"]
        }
        
        for category, components in categories.items():
            print(f"\n{category}:")
            for component in components:
                if component in versions:
                    version = versions[component]
                    status = "✓" if version != "unknown" else "⚠"
                    print(f"  {status} {component:<20} {version}")
        
        # Check for issues
        issues = self.validate_versions()
        if issues:
            print(f"\n⚠ Version Issues ({len(issues)}):")
            for issue in issues:
                print(f"  • {issue}")
        else:
            print(f"\n✅ All versions are consistent")
    
    def create_git_tags(self, push: bool = False) -> None:
        """Create Git tags for current versions."""
        versions = self.get_component_versions()
        
        tag_mappings = {
            "specification": "spec",
            "typescript-sdk": "typescript",
            "python-sdk": "python",
            "go-sdk": "go",
            "rust-sdk": "rust",
            "json-schema": "json",
            "cli": "cli",
            "web-app": "web",
            "mcp-server": "mcp",
            "docs-site": "docs"
        }
        
        tags_created = []
        
        for component, version in versions.items():
            if component in tag_mappings and version != "unknown":
                tag_name = f"{tag_mappings[component]}-v{version}"
                try:
                    # Create annotated tag
                    subprocess.run([
                        "git", "tag", "-a", tag_name, 
                        "-m", f"Release {component} v{version}"
                    ], check=True, cwd=BASE_DIR)
                    tags_created.append(tag_name)
                    print(f"✓ Created tag: {tag_name}")
                except subprocess.CalledProcessError:
                    print(f"⚠ Failed to create tag: {tag_name} (may already exist)")
        
        if push and tags_created:
            try:
                subprocess.run(["git", "push", "--tags"], check=True, cwd=BASE_DIR)
                print(f"✓ Pushed {len(tags_created)} tags to remote")
            except subprocess.CalledProcessError:
                print("⚠ Failed to push tags to remote")


def main():
    parser = argparse.ArgumentParser(description="BSpec Version Synchronization Tool")
    parser.add_argument("--sync-all", action="store_true", 
                       help="Synchronize all SDK versions to match specification")
    parser.add_argument("--update-spec", type=str, metavar="VERSION",
                       help="Update specification to new version and sync SDKs")
    parser.add_argument("--validate", action="store_true",
                       help="Validate version consistency across components")
    parser.add_argument("--report", action="store_true",
                       help="Display version report for all components")
    parser.add_argument("--create-tags", action="store_true",
                       help="Create Git tags for current versions")
    parser.add_argument("--push-tags", action="store_true",
                       help="Create and push Git tags for current versions")
    parser.add_argument("--update-component", type=str, nargs=2, 
                       metavar=("COMPONENT", "VERSION"),
                       help="Update specific component version")
    
    args = parser.parse_args()
    
    if len(sys.argv) == 1:
        parser.print_help()
        return
    
    try:
        vm = VersionManager()
        
        if args.update_spec:
            print(f"🔄 Updating specification to {args.update_spec}")
            vm.update_spec_version(args.update_spec)
            vm.sync_sdk_versions_to_spec()
        
        elif args.sync_all:
            print("🔄 Synchronizing all SDK versions to specification")
            vm.sync_sdk_versions_to_spec()
        
        elif args.update_component:
            component, version = args.update_component
            print(f"🔄 Updating {component} to {version}")
            
            if component == "typescript-sdk":
                vm.update_typescript_sdk_version(version)
            elif component == "python-sdk":
                vm.update_python_sdk_version(version)
            elif component == "go-sdk":
                vm.update_go_sdk_version(version)
            elif component == "rust-sdk":
                vm.update_rust_sdk_version(version)
            elif component == "json-schema":
                vm.update_json_schema_version(version)
            elif component == "cli":
                vm.update_cli_version(version)
            elif component == "web-app":
                vm.update_web_app_version(version)
            elif component == "mcp-server":
                vm.update_mcp_server_version(version)
            elif component == "docs-site":
                vm.update_docs_site_version(version)
            else:
                print(f"❌ Unknown component: {component}")
                return
        
        elif args.validate:
            issues = vm.validate_versions()
            if issues:
                print("❌ Version validation failed:")
                for issue in issues:
                    print(f"  • {issue}")
                sys.exit(1)
            else:
                print("✅ All versions are consistent")
        
        elif args.create_tags:
            print("🏷️  Creating Git tags")
            vm.create_git_tags(push=False)
        
        elif args.push_tags:
            print("🏷️  Creating and pushing Git tags")
            vm.create_git_tags(push=True)
        
        elif args.report:
            vm.print_version_report()
        
    except Exception as e:
        print(f"❌ Error: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()