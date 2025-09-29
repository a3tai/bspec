"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.990507
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class UNTDocument(BaseBSpecDocument):
    """
    Unit Economics (UNT)
    Domain: business-model

    Per-customer or per-unit financial metrics
    """

    # Type constraint
    type: Literal[DocumentType.UNT] = DocumentType.UNT

    # Domain constraint
    domain: Literal[BusinessDomain.BUSINESS_MODEL] = BusinessDomain.BUSINESS_MODEL





    def __post_init__(self):
        """Post-initialization for UNT documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.UNT
        self.domain = BusinessDomain.BUSINESS_MODEL

    def validate(self) -> list[str]:
        """Validate UNT document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors