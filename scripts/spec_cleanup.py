#!/usr/bin/env python3
"""
BSpec Specification Cleanup Script

Helps automate the process of upgrading specifications to professional standards.
"""

import os
import re
import sys
from pathlib import Path
from typing import Dict, List, Tuple

# Professional language replacements
LANGUAGE_IMPROVEMENTS = {
    r'\bfigure out\b': 'determine',
    r'\bdeal with\b': 'address',
    r'\bmake sure\b': 'ensure',
    r'\bhelp\b': 'enable',
    r'\bget\b': 'obtain',
    r'\breally important\b': 'critical',
    r'\bvery important\b': 'essential',
    r'\bcompanies\b': 'organizations',
    r'\bThis document tells you\b': 'This specification defines',
    r'\bYou can\b': 'Organizations may',
    r'\bYou should\b': 'Organizations should',
    r'\bYou must\b': 'Organizations must',
}

def add_abstract_section(content: str, doc_type: str, full_name: str, domain: str) -> str:
    """Add professional abstract section after the document header."""
    
    # Find the position after the header but before "Purpose and Scope"
    header_end = content.find('## Purpose and Scope')
    if header_end == -1:
        return content
    
    abstract = f"""## Abstract

This specification defines the {full_name} document type within the BSpec 1.0 Universal Business Specification Standard. It establishes normative requirements, structured templates, and implementation guidance for organizations documenting {full_name.lower()} within the {domain} domain. This specification enables systematic, machine-readable documentation that supports strategic planning, operational execution, and organizational alignment.

"""
    
    return content[:header_end] + abstract + content[header_end:]

def improve_purpose_and_scope(content: str) -> str:
    """Make Purpose and Scope section more formal and normative."""
    
    # Find the Purpose and Scope section
    pattern = r'(## Purpose and Scope\n\n)(.*?)(\n\n##)'
    match = re.search(pattern, content, re.DOTALL)
    
    if not match:
        return content
    
    current_text = match.group(2)
    
    # Make it more normative - this is a simple example
    # In practice, each spec would need custom attention
    improved_text = current_text
    for old, new in LANGUAGE_IMPROVEMENTS.items():
        improved_text = re.sub(old, new, improved_text, flags=re.IGNORECASE)
    
    return content.replace(match.group(0), match.group(1) + improved_text + match.group(3))

def standardize_metadata_schema(content: str) -> str:
    """Ensure metadata schema follows the standard format."""
    
    # This is a placeholder - each domain might need custom metadata standardization
    return content

def add_validation_checklist(content: str) -> str:
    """Add comprehensive validation checklist if missing."""
    
    if "## Validation Checklist" in content:
        return content
    
    # Add before the end of the document
    validation_section = """
## Validation Checklist

- [ ] Document header contains all required metadata fields
- [ ] Abstract clearly explains the document type purpose and scope
- [ ] Purpose and Scope section uses normative language
- [ ] Metadata schema is valid YAML with all required fields
- [ ] Content template is complete and actionable
- [ ] Quality standards progress logically from Bronze to Gold
- [ ] Implementation guidelines cover creation, review, and measurement
- [ ] All examples are realistic and industry-appropriate
- [ ] Language is professional and appropriate for executive audience
- [ ] Document follows BSpec writing standards consistently
"""
    
    return content + validation_section

def cleanup_specification_file(file_path: Path) -> bool:
    """Clean up a single specification file."""
    
    print(f"Processing: {file_path}")
    
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Extract document info from filename and content
        doc_type = file_path.stem.replace('-spec', '')
        domain = file_path.parent.name
        
        # Extract full name from content
        full_name_match = re.search(r'\*\*Document Type Name:\*\* (.+)', content)
        full_name = full_name_match.group(1) if full_name_match else doc_type
        
        # Apply improvements
        original_content = content
        content = add_abstract_section(content, doc_type, full_name, domain)
        content = improve_purpose_and_scope(content)
        content = standardize_metadata_schema(content)
        content = add_validation_checklist(content)
        
        # Apply language improvements
        for old, new in LANGUAGE_IMPROVEMENTS.items():
            content = re.sub(old, new, content, flags=re.IGNORECASE)
        
        # Only write if changes were made
        if content != original_content:
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(content)
            print(f"  ✅ Updated {file_path}")
            return True
        else:
            print(f"  ⏭️  No changes needed for {file_path}")
            return False
            
    except Exception as e:
        print(f"  ❌ Error processing {file_path}: {e}")
        return False

def cleanup_domain(domain_dir: Path) -> Dict[str, int]:
    """Clean up all specifications in a domain directory."""
    
    results = {"processed": 0, "updated": 0, "errors": 0}
    
    if not domain_dir.exists() or not domain_dir.is_dir():
        print(f"Domain directory not found: {domain_dir}")
        return results
    
    spec_files = list(domain_dir.glob("*-spec.md"))
    print(f"\n📁 Processing {domain_dir.name} domain ({len(spec_files)} files)")
    
    for spec_file in spec_files:
        results["processed"] += 1
        try:
            if cleanup_specification_file(spec_file):
                results["updated"] += 1
        except Exception as e:
            print(f"Error processing {spec_file}: {e}")
            results["errors"] += 1
    
    return results

def main():
    """Main cleanup function."""
    
    if len(sys.argv) > 1:
        domain_name = sys.argv[1]
        spec_dir = Path(__file__).parent.parent / "spec" / "v1"
        domain_dir = spec_dir / domain_name
        results = cleanup_domain(domain_dir)
    else:
        print("BSpec Specification Cleanup")
        print("Usage: python spec_cleanup.py <domain-name>")
        print("\nAvailable domains:")
        spec_dir = Path(__file__).parent.parent / "spec" / "v1"
        for domain_dir in spec_dir.iterdir():
            if domain_dir.is_dir() and not domain_dir.name.startswith('.'):
                spec_count = len(list(domain_dir.glob("*-spec.md")))
                if spec_count > 0:
                    print(f"  - {domain_dir.name} ({spec_count} specs)")
        return
    
    print(f"\n📊 Results:")
    print(f"  Processed: {results['processed']}")
    print(f"  Updated: {results['updated']}")
    print(f"  Errors: {results['errors']}")

if __name__ == "__main__":
    main()