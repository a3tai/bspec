"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.993310
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class FNDDocument(BaseBSpecDocument):
    """
    Funding (FND)
    Domain: financial

    Capital requirements and financing strategy
    """

    # Type constraint
    type: Literal[DocumentType.FND] = DocumentType.FND

    # Domain constraint
    domain: Literal[BusinessDomain.FINANCIAL] = BusinessDomain.FINANCIAL





    def __post_init__(self):
        """Post-initialization for FND documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.FND
        self.domain = BusinessDomain.FINANCIAL

    def validate(self) -> list[str]:
        """Validate FND document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors