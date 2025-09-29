"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.994928
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class GRWDocument(BaseBSpecDocument):
    """
    Growth Model (GRW)
    Domain: growth

    How the business scales customers and revenue
    """

    # Type constraint
    type: Literal[DocumentType.GRW] = DocumentType.GRW

    # Domain constraint
    domain: Literal[BusinessDomain.GROWTH] = BusinessDomain.GROWTH





    def __post_init__(self):
        """Post-initialization for GRW documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.GRW
        self.domain = BusinessDomain.GROWTH

    def validate(self) -> list[str]:
        """Validate GRW document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors