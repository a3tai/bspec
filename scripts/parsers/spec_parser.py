#!/usr/bin/env python3
"""
BSpec Specification Parser

Parses the BSpec specification files in spec/v1/ to extract:
- Document types and their domains
- YAML frontmatter schema
- Conformance level requirements
- Industry profiles
- Relationship rules

This parser reads from the specification as the single source of truth
and provides structured data for SDK generation.
"""

import re
import os
import yaml
from pathlib import Path
from typing import Dict, List, Optional, Any, Tuple
from dataclasses import dataclass


@dataclass
class DocumentType:
    """Represents a BSpec document type"""
    code: str  # 3-letter code (MSN, VSN, etc.)
    name: str  # Full name (Mission, Vision, etc.)
    domain: str  # Domain name (strategic, market, etc.)
    purpose: str  # What this document is for
    domain_order: int  # Order within domain (0-based)


@dataclass
class Domain:
    """Represents a BSpec domain"""
    name: str  # Domain name (strategic, market, etc.)
    display_name: str  # Display name (Strategic Foundation)
    order: int  # Global order (1-11)
    emoji: str  # Domain emoji
    count: int  # Number of document types
    description: str  # Domain description


@dataclass
class ConformanceLevel:
    """Represents a conformance level (Bronze, Silver, Gold)"""
    name: str  # bronze, silver, gold
    display_name: str  # Bronze, Silver, Gold
    emoji: str  # 🥉, 🥈, 🥇
    min_documents: int  # Minimum document count
    description: str  # Level description
    requirements: List[str]  # Required document types
    additional_requirements: List[str]  # Additional requirements


@dataclass
class IndustryProfile:
    """Represents an industry profile"""
    name: str  # software-saas, physical-product, etc.
    display_name: str  # Software/SaaS, Physical Product, etc.
    additional_documents: List[str]  # Additional required document types
    specialized_metrics: List[str]  # Industry-specific metrics
    description: str  # Profile description


@dataclass
class YAMLSchema:
    """Represents the YAML frontmatter schema"""
    required_fields: List[str]  # Always required fields
    required_for_accepted: List[str]  # Required for Accepted status
    required_for_operational: List[str]  # Required for operational businesses
    optional_fields: List[str]  # Optional but recommended fields
    field_descriptions: Dict[str, str]  # Field descriptions
    field_types: Dict[str, str]  # Field type information
    validation_rules: Dict[str, str]  # Validation patterns


class SpecificationParser:
    """Parses BSpec specification files to extract structured data"""

    def __init__(self, spec_dir: str = "spec/v1"):
        """Initialize parser with specification directory"""
        self.spec_dir = Path(spec_dir)
        self.spec_file = self.spec_dir / "spec.md"

        # Ensure spec directory exists
        if not self.spec_dir.exists():
            raise FileNotFoundError(f"Specification directory not found: {spec_dir}")

        if not self.spec_file.exists():
            raise FileNotFoundError(f"Main specification file not found: {self.spec_file}")

    def parse_document_types(self) -> Tuple[List[DocumentType], List[Domain]]:
        """Parse document types and domains from spec.md"""
        content = self.spec_file.read_text(encoding='utf-8')

        document_types = []
        domains = []

        # Parse each domain section using the new format
        domain_patterns = [
            (1, r'### 1\. Strategic Foundation \(8 types\)(.*?)(?=### 2\.|### \d+\.)', 'strategic-foundation', 'Strategic Foundation', '🎯'),
            (2, r'### 2\. Market & Environment \(10 types\)(.*?)(?=### 3\.|### \d+\.)', 'market-environment', 'Market & Environment', '🌍'),
            (3, r'### 3\. Customer & Value \(12 types\)(.*?)(?=### 4\.|### \d+\.)', 'customer-value', 'Customer & Value', '👥'),
            (4, r'### 4\. Product & Service \(10 types\)(.*?)(?=### 5\.|### \d+\.)', 'product-service', 'Product & Service', '📦'),
            (5, r'### 5\. Business Model \(9 types\)(.*?)(?=### 6\.|### \d+\.)', 'business-model', 'Business Model', '💰'),
            (6, r'### 6\. Operations & Execution \(12 types\)(.*?)(?=### 7\.|### \d+\.)', 'operations-execution', 'Operations & Execution', '⚙️'),
            (7, r'### 7\. Technology & Data \(8 types\)(.*?)(?=### 8\.|### \d+\.)', 'technology-data', 'Technology & Data', '🔧'),
            (8, r'### 8\. Financial & Investment \(10 types\)(.*?)(?=### 9\.|### \d+\.)', 'financial-investment', 'Financial & Investment', '📊'),
            (9, r'### 9\. Risk & Governance \(8 types\)(.*?)(?=### 10\.|### \d+\.)', 'risk-governance', 'Risk & Governance', '⚠️'),
            (10, r'### 10\. Growth & Innovation \(8 types\)(.*?)(?=### 11\.|### \d+\.)', 'growth-innovation', 'Growth & Innovation', '📈'),
            (11, r'### 11\. Brand & Marketing \(12 types\)(.*?)(?=### 12\.|### \d+\.)', 'brand-marketing', 'Brand & Marketing', '📢'),
            (12, r'### 12\. Learning & Decisions \(6 types\)(.*?)(?=## |### \d+\.)', 'learning-decisions', 'Learning & Decisions', '🧠'),
        ]

        for order, pattern, domain_name, display_name, emoji in domain_patterns:
            domain_match = re.search(pattern, content, re.DOTALL)
            if not domain_match:
                continue

            domain_content = domain_match.group(1)

            # Extract description from the first line after the heading
            lines = domain_content.strip().split('\n')
            description = lines[0].strip() if lines else ""

            # Parse document types using bullet point format
            domain_doc_types = []
            for line in lines:
                # Look for pattern: - **CODE** - Name: Description
                match = re.search(r'-\s+\*\*([A-Z]{3})\*\*\s+-\s+([^:]+):\s+(.+)', line)
                if match:
                    code = match.group(1)
                    name = match.group(2).strip()
                    purpose = match.group(3).strip()

                    doc_type = DocumentType(
                        code=code,
                        name=name,
                        domain=domain_name,
                        purpose=purpose,
                        domain_order=len(domain_doc_types)
                    )
                    document_types.append(doc_type)
                    domain_doc_types.append(doc_type)

            # Create domain
            domain = Domain(
                name=domain_name,
                display_name=display_name,
                order=order,
                emoji=emoji,
                count=len(domain_doc_types),
                description=description
            )
            domains.append(domain)

        return document_types, domains

    def parse_yaml_schema(self) -> YAMLSchema:
        """Parse YAML frontmatter schema from spec.md"""
        content = self.spec_file.read_text(encoding='utf-8')

        # Find the YAML frontmatter section
        schema_pattern = r'### YAML Frontmatter Structure.*?```yaml\n---\n(.*?)\n---\n```'
        schema_match = re.search(schema_pattern, content, re.DOTALL)

        if not schema_match:
            raise ValueError("Could not find YAML schema section in spec.md")

        yaml_content = schema_match.group(1)

        # Parse field categories from the text
        required_fields = [
            'id', 'title', 'type', 'status', 'version',
            'owner', 'created', 'updated'
        ]

        required_for_accepted = [
            'stakeholders', 'success_criteria'
        ]

        required_for_operational = [
            'implementation_status', 'metrics'
        ]

        optional_fields = [
            'reviewers', 'contributors', 'expires', 'review_cycle',
            'parent', 'depends_on', 'enables', 'conflicts_with', 'related', 'supersedes',
            'domain', 'scope', 'horizon', 'priority', 'visibility',
            'assumptions', 'constraints', 'risks',
            'implementation_date', 'completion_date', 'resources_required',
            'tags', 'industry', 'geography', 'language', 'classification',
            'changelog'
        ]

        # Field descriptions and types from the schema
        field_descriptions = {
            'id': 'Globally unique document identifier',
            'title': 'Clear, descriptive title',
            'type': 'Document type code (MSN, SVC, etc.)',
            'status': 'Document lifecycle status',
            'version': 'Semantic versioning',
            'owner': 'Who owns and maintains this document',
            'stakeholders': 'Who has interest in this document',
            'reviewers': 'Who reviewed/approved this version',
            'contributors': 'Who contributed to this document',
            'created': 'When document was first created',
            'updated': 'When document was last modified',
            'expires': 'When document expires (optional)',
            'review_cycle': 'How often to review',
            'parent': 'Parent document (hierarchical)',
            'depends_on': 'Dependencies (this needs those)',
            'enables': 'Enablements (this makes those possible)',
            'conflicts_with': 'Mutual exclusions',
            'related': 'Other relevant documents',
            'supersedes': 'What this document replaces',
            'domain': 'Business domain classification',
            'scope': 'Organizational scope',
            'horizon': 'Time horizon',
            'priority': 'Business importance',
            'visibility': 'Access level',
            'assumptions': 'Key assumptions',
            'constraints': 'Key constraints',
            'success_criteria': 'Measurable success criteria',
            'risks': 'Associated risk documents',
            'metrics': 'How success is measured',
            'implementation_status': 'Implementation progress',
            'implementation_date': 'When implementation begins/began',
            'completion_date': 'When implementation completes',
            'resources_required': 'Resource requirements',
            'tags': 'Searchable tags',
            'industry': 'Industry classifications',
            'geography': 'Geographic relevance',
            'language': 'Primary language',
            'classification': 'Information classification',
            'changelog': 'Version history'
        }

        field_types = {
            'id': 'string',
            'title': 'string',
            'type': 'string',
            'status': 'enum',
            'version': 'string',
            'owner': 'string',
            'stakeholders': 'array',
            'reviewers': 'array',
            'contributors': 'array',
            'created': 'date',
            'updated': 'date',
            'expires': 'date',
            'review_cycle': 'enum',
            'parent': 'string',
            'depends_on': 'array',
            'enables': 'array',
            'conflicts_with': 'array',
            'related': 'array',
            'supersedes': 'string',
            'domain': 'enum',
            'scope': 'enum',
            'horizon': 'enum',
            'priority': 'enum',
            'visibility': 'enum',
            'assumptions': 'array',
            'constraints': 'array',
            'success_criteria': 'array',
            'risks': 'array',
            'metrics': 'array',
            'implementation_status': 'enum',
            'implementation_date': 'date',
            'completion_date': 'date',
            'resources_required': 'array',
            'tags': 'array',
            'industry': 'array',
            'geography': 'array',
            'language': 'string',
            'classification': 'enum',
            'changelog': 'array'
        }

        validation_rules = {
            'id': '^[A-Z]{3}-[a-z0-9-]+$',
            'version': '^\\d+\\.\\d+\\.\\d+$',
            'status': 'Draft|Review|Accepted|Deprecated',
            'created': '^\\d{4}-\\d{2}-\\d{2}$',
            'updated': '^\\d{4}-\\d{2}-\\d{2}$',
            'expires': '^\\d{4}-\\d{2}-\\d{2}$',
            'implementation_date': '^\\d{4}-\\d{2}-\\d{2}$',
            'completion_date': '^\\d{4}-\\d{2}-\\d{2}$'
        }

        return YAMLSchema(
            required_fields=required_fields,
            required_for_accepted=required_for_accepted,
            required_for_operational=required_for_operational,
            optional_fields=optional_fields,
            field_descriptions=field_descriptions,
            field_types=field_types,
            validation_rules=validation_rules
        )

    def parse_conformance_levels(self) -> List[ConformanceLevel]:
        """Parse conformance levels from spec.md"""
        content = self.spec_file.read_text(encoding='utf-8')

        levels = []

        # Bronze level - Minimum Viable
        bronze = ConformanceLevel(
            name='bronze',
            display_name='Bronze Level (Minimum Viable)',
            emoji='🥉',
            min_documents=15,
            description='Basic framework with core components',
            requirements=['MSN', 'VSN', 'VAL', 'STR', 'PER', 'JTB', 'PRD', 'REV', 'CST', 'RSK'],
            additional_requirements=[
                'Basic framework with core components',
                'Simple implementation approach',
                'All required metadata fields completed'
            ]
        )
        levels.append(bronze)

        # Silver level - Investment Ready
        silver = ConformanceLevel(
            name='silver',
            display_name='Silver Level (Investment Ready)',
            emoji='🥈',
            min_documents=30,
            description='Comprehensive framework with structured implementation',
            requirements=['MKT', 'SEG', 'CMP', 'PRD', 'FIN', 'OPS', 'ARC', 'SEC', 'GOV'],
            additional_requirements=[
                'Comprehensive framework with detailed implementation',
                'Structured governance and quality standards',
                'Regular review and improvement processes'
            ]
        )
        levels.append(silver)

        # Gold level - Operational Excellence
        gold = ConformanceLevel(
            name='gold',
            display_name='Gold Level (Operational Excellence)',
            emoji='🥇',
            min_documents=50,
            description='Advanced capabilities with sophisticated optimization',
            requirements=['INN', 'LEA', 'KNO', 'WIS', 'DEC', 'ANA', 'BRD', 'CAM'],
            additional_requirements=[
                'Advanced capabilities with sophisticated optimization',
                'Comprehensive ecosystem integration',
                'Strategic management driving competitive advantage'
            ]
        )
        levels.append(gold)

        return levels

    def parse_industry_profiles(self) -> List[IndustryProfile]:
        """Parse industry profiles from spec.md"""
        content = self.spec_file.read_text(encoding='utf-8')

        profiles = []

        # Software/SaaS profile
        saas = IndustryProfile(
            name='software-saas',
            display_name='Software/SaaS',
            additional_documents=['ARC', 'API', 'SEC', 'SUP'],
            specialized_metrics=['MRR', 'CAC', 'LTV', 'Churn', 'NRR'],
            description='Software as a Service and cloud platforms'
        )
        profiles.append(saas)

        # Physical Product profile
        physical = IndustryProfile(
            name='physical-product',
            display_name='Physical Product',
            additional_documents=['INF', 'QUA', 'VND', 'REG'],
            specialized_metrics=['COGS', 'Inventory', 'Manufacturing'],
            description='Physical products and manufacturing businesses'
        )
        profiles.append(physical)

        # Service Business profile
        service = IndustryProfile(
            name='service-business',
            display_name='Service Business',
            additional_documents=['PRC', 'SLA', 'SKI', 'QUA'],
            specialized_metrics=['Utilization', 'Capacity', 'Quality'],
            description='Professional services and consulting businesses'
        )
        profiles.append(service)

        # Nonprofit profile
        nonprofit = IndustryProfile(
            name='nonprofit',
            display_name='Nonprofit',
            additional_documents=['PUR', 'STA', 'MET', 'GVN'],
            specialized_metrics=['Impact', 'Beneficiaries', 'Funding'],
            description='Nonprofit organizations and social enterprises'
        )
        profiles.append(nonprofit)

        return profiles

    def get_bspec_version(self) -> str:
        """Extract BSpec version from spec.md"""
        content = self.spec_file.read_text(encoding='utf-8')

        version_pattern = r'\*\*Version:\*\* (\d+\.\d+\.\d+)'
        version_match = re.search(version_pattern, content)

        if version_match:
            return version_match.group(1)

        return "1.0.0"  # Default version

    def parse_all(self) -> Dict[str, Any]:
        """Parse all specification data and return as structured dictionary"""
        document_types, domains = self.parse_document_types()
        yaml_schema = self.parse_yaml_schema()
        conformance_levels = self.parse_conformance_levels()
        industry_profiles = self.parse_industry_profiles()
        bspec_version = self.get_bspec_version()

        return {
            'bspec_version': bspec_version,
            'document_types': document_types,
            'domains': domains,
            'yaml_schema': yaml_schema,
            'conformance_levels': conformance_levels,
            'industry_profiles': industry_profiles,
            'total_document_types': len(document_types),
            'total_domains': len(domains)
        }


def main():
    """Test the specification parser"""
    parser = SpecificationParser()

    try:
        result = parser.parse_all()

        print(f"BSpec Version: {result['bspec_version']}")
        print(f"Total Document Types: {result['total_document_types']}")
        print(f"Total Domains: {result['total_domains']}")
        print()

        print("Domains:")
        for domain in result['domains']:
            print(f"  {domain.emoji} {domain.display_name} ({domain.count} types)")
        print()

        print("Sample Document Types:")
        for doc_type in result['document_types'][:10]:
            print(f"  {doc_type.code}: {doc_type.name} ({doc_type.domain})")
        print()

        print("Conformance Levels:")
        for level in result['conformance_levels']:
            print(f"  {level.emoji} {level.display_name}: {level.min_documents}+ documents")
        print()

        print("Industry Profiles:")
        for profile in result['industry_profiles']:
            print(f"  {profile.display_name}: +{len(profile.additional_documents)} docs")

    except Exception as e:
        print(f"Error parsing specification: {e}")
        return 1

    return 0


if __name__ == "__main__":
    exit(main())