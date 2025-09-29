"""
BSpec Python SDK - Type Definitions
Generated from BSpec v1.0.0 JSON SDK
"""

from typing import Dict, List, Any, Optional, Union
from dataclasses import dataclass
from enum import Enum


class DocumentStatus(Enum):
    """Document status types"""
    DRAFT = "Draft"
    REVIEW = "Review"
    ACCEPTED = "Accepted"
    DEPRECATED = "Deprecated"
    ARCHIVED = "Archived"


class BusinessDomain(Enum):
    """Business domain names"""
    STRATEGIC_FOUNDATION = "Strategic Foundation"
    MARKET_ENVIRONMENT = "Market & Environment"
    CUSTOMER_VALUE = "Customer & Value"
    PRODUCT_SERVICE = "Product & Service"
    BUSINESS_MODEL = "Business Model"
    OPERATIONS_EXECUTION = "Operations & Execution"
    TECHNOLOGY_DATA = "Technology & Data"
    FINANCIAL_INVESTMENT = "Financial & Investment"
    RISK_GOVERNANCE = "Risk & Governance"
    GROWTH_INNOVATION = "Growth & Innovation"
    BRAND_MARKETING = "Brand & Marketing"
    LEARNING_DECISIONS = "Learning & Decisions"


class ConformanceLevel(Enum):
    """Conformance level names"""
    BRONZE = "Bronze"
    SILVER = "Silver"
    GOLD = "Gold"


@dataclass
class BSpecMetadata:
    """Metadata about the BSpec generation"""
    bspec_version: str
    generated_at: str
    generator: str
    source_directory: str


@dataclass
class BSpecFile:
    """Represents a file in the BSpec specification"""
    path: str
    name: str
    type: str
    extension: str
    size: Optional[int] = None
    modified: Optional[str] = None
    has_frontmatter: Optional[bool] = None
    frontmatter: Optional[Dict[str, Any]] = None
    content: Optional[str] = None
    data: Optional[Any] = None
    error: Optional[str] = None


@dataclass
class BSpecDocumentIndex:
    """Document index entry with frontmatter metadata"""
    file_path: str
    frontmatter: Dict[str, Any]
    title: str
    type: Optional[str] = None
    domain: Optional[str] = None


@dataclass
class BSpecDomain:
    """Business domain containing document types"""
    name: str
    description: str
    document_types: List[str]
    document_count: int
    documents: Optional[List[Dict[str, str]]] = None


@dataclass
class BSpecStatistics:
    """Statistics about the BSpec specification"""
    total_files: int
    total_documents: int
    total_domains: int
    total_document_types: int
    markdown_files: int
    yaml_files: int
    other_files: int


@dataclass
class BSpecDocumentType:
    """BSpec document type definition"""
    code: str
    name: str
    purpose: str
    domain: str


@dataclass
class BSpecYAMLSchema:
    """YAML schema definition for BSpec documents"""
    required_fields: List[str]
    optional_fields: List[str]


@dataclass
class BSpecConformanceLevel:
    """Conformance level definition"""
    name: str
    description: str


@dataclass
class BSpecData:
    """Main BSpec data structure containing the entire specification"""
    metadata: BSpecMetadata
    files: Dict[str, BSpecFile]
    directory_structure: List[List[str]]
    document_index: List[BSpecDocumentIndex]
    domains: List[BSpecDomain]
    statistics: BSpecStatistics
    document_types: List[BSpecDocumentType]
    yaml_schema: BSpecYAMLSchema
    conformance_levels: List[BSpecConformanceLevel]
