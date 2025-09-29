#!/usr/bin/env python3
"""Test script to verify generators work"""

import sys
import os
from pathlib import Path

# Add scripts to path
sys.path.insert(0, 'scripts')

def test_parser():
    """Test the specification parser"""
    try:
        from parsers.spec_parser import SpecificationParser

        parser = SpecificationParser('spec/v1')
        result = parser.parse_all()

        print(f"✅ Parser test successful!")
        print(f"   BSpec Version: {result['bspec_version']}")
        print(f"   Document Types: {result['total_document_types']}")
        print(f"   Domains: {result['total_domains']}")

        return result

    except Exception as e:
        print(f"❌ Parser test failed: {e}")
        import traceback
        traceback.print_exc()
        return None

def test_json_generator():
    """Test JSON generator"""
    try:
        from generators.json.generator import JSONGenerator

        # Create output directory
        output_dir = Path('sdk/v1/json')
        output_dir.mkdir(parents=True, exist_ok=True)

        generator = JSONGenerator('spec/v1', 'sdk/v1/json')
        generator.generate_all()

        print("✅ JSON Generator test successful!")
        return True

    except Exception as e:
        print(f"❌ JSON Generator test failed: {e}")
        import traceback
        traceback.print_exc()
        return False

if __name__ == "__main__":
    print("🧪 Testing BSpec generators...")

    # Test parser first
    result = test_parser()
    if result is None:
        print("Stopping - parser failed")
        sys.exit(1)

    print()

    # Test JSON generator
    if test_json_generator():
        print("\n🎉 All tests passed!")
    else:
        print("\n❌ Some tests failed")
        sys.exit(1)