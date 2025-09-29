"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.992997
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class FINDocument(BaseBSpecDocument):
    """
    Financial Model (FIN)
    Domain: financial

    Comprehensive P&L, balance sheet, cash flow projections
    """

    # Type constraint
    type: Literal[DocumentType.FIN] = DocumentType.FIN

    # Domain constraint
    domain: Literal[BusinessDomain.FINANCIAL] = BusinessDomain.FINANCIAL





    def __post_init__(self):
        """Post-initialization for FIN documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.FIN
        self.domain = BusinessDomain.FINANCIAL

    def validate(self) -> list[str]:
        """Validate FIN document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors