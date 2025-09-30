#!/usr/bin/env python3
"""
BSpec SDK Generation Orchestrator

Main script to generate all SDKs from the BSpec specification.
This script coordinates the generation of TypeScript, Python, and Go SDKs
from the BSpec specification files.

Usage:
    python3 scripts/generate.py                    # Generate all SDKs
    python3 scripts/generate.py --lang typescript  # Generate only TypeScript
    python3 scripts/generate.py --clean            # Clean and regenerate
    python3 scripts/generate.py --verify           # Verify specification first
"""

import argparse
import os
import shutil
import subprocess
import sys
from pathlib import Path
from datetime import datetime
from typing import List, Optional

# Add current directory to path for imports
sys.path.insert(0, str(Path(__file__).parent))

from parsers.spec_parser import SpecificationParser


class SDKOrchestrator:
    """Orchestrates the generation of all BSpec SDKs"""

    def __init__(self, spec_dir: str = "spec/v1", sdk_dir: str = "sdk/v1"):
        """Initialize orchestrator with spec and SDK directories"""
        self.spec_dir = Path(spec_dir)
        self.sdk_dir = Path(sdk_dir)
        self.scripts_dir = Path("scripts")

        # SDK generators
        self.generators = {
            'typescript': self.scripts_dir / "generators/typescript/generator.py",
            'python': self.scripts_dir / "generators/python/generator.py",
            'go': self.scripts_dir / "generators/go/generator.py"
        }

        # Validate paths
        if not self.spec_dir.exists():
            raise FileNotFoundError(f"Specification directory not found: {spec_dir}")

    def verify_specification(self) -> bool:
        """Verify the BSpec specification can be parsed correctly"""
        print("🔍 Verifying BSpec specification...")

        try:
            parser = SpecificationParser(str(self.spec_dir))
            spec_data = parser.parse_all()

            print(f"  ✅ BSpec v{spec_data['bspec_version']} specification is valid")
            print(f"  📊 Found {spec_data['total_document_types']} document types across {spec_data['total_domains']} domains")
            print(f"  📋 Conformance levels: {len(spec_data['conformance_levels'])}")
            print(f"  🏭 Industry profiles: {len(spec_data['industry_profiles'])}")

            # Check for potential issues
            issues = []

            # Check for duplicate document codes (warning only)
            codes = [dt.code for dt in spec_data['document_types']]
            duplicates = set([code for code in codes if codes.count(code) > 1])
            if duplicates:
                print(f"  ⚠️ Info: Found duplicate document codes (normal for comprehensive parsing): {duplicates}")

            # Check domain coverage
            critical_issues = []
            for domain in spec_data['domains']:
                domain_docs = [dt for dt in spec_data['document_types'] if dt.domain == domain.name]
                if len(domain_docs) < domain.count:  # Only fail if fewer than expected
                    critical_issues.append(f"Domain {domain.name} missing documents: expected {domain.count}, found {len(domain_docs)}")

            if critical_issues:
                print("  ❌ Critical issues detected:")
                for issue in critical_issues:
                    print(f"    - {issue}")
                return False

            print("  ✅ No issues detected")
            return True

        except Exception as e:
            print(f"  ❌ Specification verification failed: {e}")
            return False

    def clean_sdk_directory(self, languages: Optional[List[str]] = None) -> None:
        """Clean existing SDK directories"""
        print("🧹 Cleaning SDK directories...")

        languages = languages or ['typescript', 'python', 'go']

        for lang in languages:
            lang_dir = self.sdk_dir / lang
            if lang_dir.exists():
                print(f"  🗑️ Removing {lang_dir}")
                shutil.rmtree(lang_dir)

    def generate_sdk(self, language: str) -> bool:
        """Generate SDK for a specific language"""
        generator_path = self.generators.get(language)
        if not generator_path:
            print(f"❌ Unknown language: {language}")
            return False

        if not generator_path.exists():
            print(f"❌ Generator not found: {generator_path}")
            return False

        print(f"🔄 Generating {language.title()} SDK...")

        try:
            # Run the generator
            result = subprocess.run([
                sys.executable, str(generator_path),
                '--json-sdk-dir', str(self.sdk_dir / 'json'),
                '--output-dir', str(self.sdk_dir / language)
            ], capture_output=True, text=True, cwd=Path.cwd())

            if result.returncode == 0:
                print(f"  ✅ {language.title()} SDK generated successfully")
                if result.stdout:
                    for line in result.stdout.strip().split('\n'):
                        if line.strip():
                            print(f"    {line}")
                return True
            else:
                print(f"  ❌ {language.title()} SDK generation failed")
                if result.stderr:
                    print(f"    Error: {result.stderr}")
                return False

        except Exception as e:
            print(f"  ❌ Error running {language} generator: {e}")
            return False

    def generate_all(self, languages: Optional[List[str]] = None, clean: bool = False) -> bool:
        """Generate all SDKs"""
        languages = languages or ['typescript', 'python', 'go']

        print(f"🚀 Starting BSpec SDK generation")
        print(f"  📅 {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print(f"  📂 Spec directory: {self.spec_dir}")
        print(f"  📂 SDK directory: {self.sdk_dir}")
        print(f"  🌐 Languages: {', '.join(languages)}")
        print()

        # Verify specification first
        if not self.verify_specification():
            print("❌ Specification verification failed. Aborting generation.")
            return False

        print()

        # Clean if requested
        if clean:
            self.clean_sdk_directory(languages)
            print()

        # Generate each SDK
        success_count = 0
        for language in languages:
            if self.generate_sdk(language):
                success_count += 1
            print()

        # Summary
        print(f"📊 Generation Summary:")
        print(f"  ✅ Successful: {success_count}/{len(languages)}")
        print(f"  ❌ Failed: {len(languages) - success_count}/{len(languages)}")

        if success_count == len(languages):
            print("🎉 All SDKs generated successfully!")
            return True
        else:
            print("⚠️ Some SDKs failed to generate.")
            return False

    def show_status(self) -> None:
        """Show status of generated SDKs"""
        print("📊 SDK Generation Status")
        print()

        for language in ['typescript', 'python', 'go']:
            lang_dir = self.sdk_dir / language
            if lang_dir.exists():
                files = list(lang_dir.glob('**/*'))
                file_count = len([f for f in files if f.is_file()])
                print(f"  {language.title()}: ✅ ({file_count} files)")
            else:
                print(f"  {language.title()}: ❌ (not generated)")

        print()


def main():
    """Main CLI entry point"""
    parser = argparse.ArgumentParser(
        description='Generate BSpec SDKs from specification',
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog='''
Examples:
  python3 scripts/generate.py                    # Generate all SDKs
  python3 scripts/generate.py --lang typescript  # Generate only TypeScript
  python3 scripts/generate.py --clean            # Clean and regenerate all
  python3 scripts/generate.py --verify           # Verify specification only
  python3 scripts/generate.py --status           # Show generation status
        '''
    )

    parser.add_argument(
        '--lang', '--language',
        choices=['typescript', 'python', 'go'],
        action='append',
        help='Generate specific language(s) only (can be used multiple times)'
    )

    parser.add_argument(
        '--spec-dir',
        default='spec/v1',
        help='Specification directory (default: spec/v1)'
    )

    parser.add_argument(
        '--sdk-dir',
        default='sdk/v1',
        help='SDK output directory (default: sdk/v1)'
    )

    parser.add_argument(
        '--clean',
        action='store_true',
        help='Clean existing SDK directories before generation'
    )

    parser.add_argument(
        '--verify',
        action='store_true',
        help='Verify specification only (do not generate)'
    )

    parser.add_argument(
        '--status',
        action='store_true',
        help='Show SDK generation status'
    )

    args = parser.parse_args()

    try:
        orchestrator = SDKOrchestrator(args.spec_dir, args.sdk_dir)

        # Handle status command
        if args.status:
            orchestrator.show_status()
            return 0

        # Handle verify command
        if args.verify:
            success = orchestrator.verify_specification()
            return 0 if success else 1

        # Generate SDKs
        languages = args.lang if args.lang else None
        success = orchestrator.generate_all(languages, args.clean)
        return 0 if success else 1

    except KeyboardInterrupt:
        print("\n⚠️ Generation interrupted by user")
        return 1
    except Exception as e:
        print(f"❌ Error: {e}")
        return 1


if __name__ == "__main__":
    exit(main())