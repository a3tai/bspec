"""
BSpec Python SDK - Universal Business Specification Standard

Generated from BSpec v1.0.0 JSON SDK
"""

from .bspec import BSpec
from .types import (
    BSpecData, BSpecMetadata, BSpecFile, BSpecDomain, BSpecStatistics,
    BSpecDocumentType, BSpecYAMLSchema, BSpecConformanceLevel,
    DocumentStatus, BusinessDomain, ConformanceLevel
)

__version__ = "1.1.0"
__all__ = [
    "BSpec",
    "BSpecData", "BSpecMetadata", "BSpecFile", "BSpecDomain", "BSpecStatistics",
    "BSpecDocumentType", "BSpecYAMLSchema", "BSpecConformanceLevel",
    "DocumentStatus", "BusinessDomain", "ConformanceLevel"
]
