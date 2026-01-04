#!/usr/bin/env python3
"""
Generate web-friendly markdown from BSpec specification files.
Converts spec/v1/ document type specs into apps/web/docs/ for the documentation site.
"""

import os
import re
from pathlib import Path
from typing import List, Dict, Tuple

# Domain mapping from directory names to friendly names
DOMAIN_NAMES = {
    "strategic-foundation": "Strategic Foundation",
    "market-environment": "Market Environment",
    "customer-value": "Customer Value",
    "product-service": "Product & Service",
    "business-model": "Business Model",
    "operations-execution": "Operations & Execution",
    "technology-data": "Technology & Data",
    "financial-investment": "Financial & Investment",
    "risk-governance": "Risk & Governance",
    "growth-innovation": "Growth & Innovation",
    "learning-decisions": "Learning & Decisions",
    "brand-marketing": "Brand & Marketing"
}

def get_spec_dir() -> Path:
    """Get the spec/v1 directory path."""
    return Path(__file__).parent.parent / "spec" / "v1"

def get_web_docs_dir() -> Path:
    """Get the apps/web/docs directory path."""
    return Path(__file__).parent.parent / "apps" / "web" / "docs"

def extract_document_code(filename: str) -> str:
    """Extract document type code from filename (e.g., VSN from VSN-spec.md)."""
    match = re.match(r'([A-Z]{3})-spec\.md', filename)
    return match.group(1) if match else None

def parse_spec_file(file_path: Path) -> Dict:
    """Parse a spec file and extract key information."""
    content = file_path.read_text()
    
    # Extract title (first # heading)
    title_match = re.search(r'^# (.+)$', content, re.MULTILINE)
    title = title_match.group(1) if title_match else file_path.stem
    
    # Extract document type name and code
    code_match = re.search(r'\*\*Document Type Code:\*\* ([A-Z]{3})', content)
    name_match = re.search(r'\*\*Document Type Name:\*\* (.+)$', content, re.MULTILINE)
    domain_match = re.search(r'\*\*Domain:\*\* (.+)$', content, re.MULTILINE)
    
    return {
        'title': title,
        'code': code_match.group(1) if code_match else None,
        'name': name_match.group(1) if name_match else None,
        'domain': domain_match.group(1) if domain_match else None,
        'content': content,
        'file_path': file_path
    }

def convert_spec_to_web_markdown(spec: Dict) -> str:
    """Convert a spec document to web-friendly markdown."""
    content = spec['content']
    
    # Add VitePress frontmatter
    frontmatter = f"""---
title: "{spec['code']}: {spec['name']}"
description: "BSpec {spec['code']} document type specification"
---

"""
    
    # Add breadcrumb navigation
    breadcrumb = f"""# {spec['code']}: {spec['name']}

::: tip Document Type
**Code**: {spec['code']}  
**Name**: {spec['name']}  
**Domain**: {spec['domain']}
:::

"""
    
    # Remove the original title and metadata section (first few lines)
    # Keep everything from "## Abstract" onward
    abstract_pos = content.find('## Abstract')
    if abstract_pos > 0:
        main_content = content[abstract_pos:]
    else:
        main_content = content
    
    # Enhance code blocks for VitePress
    main_content = enhance_code_blocks(main_content)
    
    # Add footer with related docs
    footer = f"""

---

## Related Documents

- [Back to Document Types](/docs/document-types)
- [Domain: {spec['domain']}](/docs/domains/{domain_to_slug(spec['domain'])})
- [Full Specification](/spec/v1-0-0)

"""
    
    return frontmatter + breadcrumb + main_content + footer

def enhance_code_blocks(content: str) -> str:
    """Enhance code blocks for better VitePress display."""
    # Add language hints to code blocks if missing
    content = re.sub(r'```\n---', '```yaml\n---', content)
    content = re.sub(r'```\n#', '```markdown\n#', content)
    
    # Fix nested code blocks by indenting them inside markdown code blocks
    # This prevents VitePress from trying to parse nested code fences
    lines = content.split('\n')
    result = []
    in_markdown_block = False
    markdown_block_depth = 0
    
    for i, line in enumerate(lines):
        # Detect markdown code block start
        if line.strip() == '```markdown':
            in_markdown_block = True
            markdown_block_depth = 0
            result.append(line)
            continue
        
        # Detect markdown code block end
        if in_markdown_block and line.strip() == '```' and markdown_block_depth == 0:
            in_markdown_block = False
            result.append(line)
            continue
        
        # Inside markdown block, handle nested code blocks
        if in_markdown_block:
            # Check for nested code fence start
            if line.strip().startswith('```'):
                if markdown_block_depth == 0:
                    markdown_block_depth = 1
                else:
                    markdown_block_depth = 0
                # Indent nested code fence markers
                result.append('    ' + line)
            else:
                # If we're inside a nested block, indent its content
                if markdown_block_depth > 0:
                    result.append('    ' + line)
                else:
                    result.append(line)
        else:
            result.append(line)
    
    return '\n'.join(result)

def domain_to_slug(domain: str) -> str:
    """Convert domain name to URL slug."""
    return domain.lower().replace(' & ', '-').replace(' ', '-')

def generate_domain_overview(domain_slug: str, domain_name: str, doc_types: List[Dict]) -> str:
    """Generate an overview page for a domain."""
    content = f"""---
title: "{domain_name}"
description: "BSpec document types in the {domain_name} domain"
---

# {domain_name}

## Overview

This domain contains {len(doc_types)} document types that cover {domain_name.lower()} aspects of your business.

## Document Types

"""
    
    for doc in sorted(doc_types, key=lambda x: x['code']):
        content += f"""### [{doc['code']}: {doc['name']}](/docs/types/{doc['code']})

{doc['content'][:200].split('##')[0].strip()}...

"""
    
    content += """
---

[← Back to All Document Types](/docs/document-types)
"""
    
    return content

def generate_document_types_index(all_specs: Dict[str, List[Dict]]) -> str:
    """Generate the main document types index page."""
    content = """---
title: "Document Types"
description: "All 82 BSpec document types organized by domain"
---

# Document Types

BSpec defines 82 standardized document types across 11 business domains. Each document type has a specific purpose, structure, and set of relationships.

## Overview by Domain

"""
    
    total = 0
    for domain_slug, docs in sorted(all_specs.items()):
        domain_name = DOMAIN_NAMES.get(domain_slug, domain_slug)
        doc_count = len(docs)
        total += doc_count
        
        content += f"""
### [{domain_name}](/docs/domains/{domain_slug}) ({doc_count} types)

"""
        
        # List document types in this domain
        for doc in sorted(docs, key=lambda x: x['code']):
            content += f"- **[{doc['code']}](/docs/types/{doc['code']})** - {doc['name']}\n"
    
    content += f"""

## Total: {total} Document Types

Each document type includes:
- **Metadata Schema** - Required fields and relationships
- **Content Structure** - Template and sections
- **Quality Standards** - Bronze, Silver, and Gold levels
- **Success Patterns** - Best practices and examples
- **Common Pitfalls** - What to avoid

## Quick Reference

### By Document Code

"""
    
    # Create alphabetical index
    all_docs = []
    for docs in all_specs.values():
        all_docs.extend(docs)
    
    for doc in sorted(all_docs, key=lambda x: x['code']):
        domain_name = DOMAIN_NAMES.get(domain_to_slug(doc['domain']), doc['domain'])
        content += f"**[{doc['code']}](/docs/types/{doc['code']})** - {doc['name']} ({domain_name})  \n"
    
    content += """

---

## Learn More

- [Core Concepts](/docs/concepts) - Understand the BSpec fundamentals
- [Quick Start](/docs/quickstart) - Create your first document
- [Full Specification](/spec/v1-0-0) - Complete technical reference
"""
    
    return content

def main():
    """Main generation function."""
    print("🔨 Generating web documentation from BSpec specification...")
    
    spec_dir = get_spec_dir()
    web_docs_dir = get_web_docs_dir()
    
    # Create output directories
    types_dir = web_docs_dir / "types"
    domains_dir = web_docs_dir / "domains"
    types_dir.mkdir(parents=True, exist_ok=True)
    domains_dir.mkdir(parents=True, exist_ok=True)
    
    # Process all spec files organized by domain
    all_specs: Dict[str, List[Dict]] = {}
    
    for domain_dir in spec_dir.iterdir():
        if not domain_dir.is_dir() or domain_dir.name.startswith('.'):
            continue
        
        domain_slug = domain_dir.name
        domain_name = DOMAIN_NAMES.get(domain_slug, domain_slug)
        domain_specs = []
        
        print(f"\n📁 Processing {domain_name}...")
        
        for spec_file in domain_dir.glob("*-spec.md"):
            code = extract_document_code(spec_file.name)
            if not code:
                continue
            
            print(f"  ├─ {code}")
            
            # Parse and convert spec
            spec = parse_spec_file(spec_file)
            domain_specs.append(spec)
            
            # Generate individual document type page
            web_content = convert_spec_to_web_markdown(spec)
            output_file = types_dir / f"{code}.md"
            output_file.write_text(web_content)
        
        if domain_specs:
            all_specs[domain_slug] = domain_specs
            
            # Generate domain overview page
            domain_content = generate_domain_overview(domain_slug, domain_name, domain_specs)
            domain_file = domains_dir / f"{domain_slug}.md"
            domain_file.write_text(domain_content)
            print(f"  └─ Generated {len(domain_specs)} document types")
    
    # Generate main index
    print(f"\n📚 Generating document types index...")
    index_content = generate_document_types_index(all_specs)
    index_file = web_docs_dir / "document-types.md"
    index_file.write_text(index_content)
    
    # Summary
    total_types = sum(len(docs) for docs in all_specs.values())
    print(f"\n✅ Generated {total_types} document type pages")
    print(f"✅ Generated {len(all_specs)} domain overview pages")
    print(f"✅ Generated 1 main index page")
    print(f"\n📂 Output: {web_docs_dir}")
    print(f"\n🚀 Ready to build and deploy!")

if __name__ == "__main__":
    main()
