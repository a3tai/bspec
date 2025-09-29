"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.993217
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class FORDocument(BaseBSpecDocument):
    """
    Forecasts (FOR)
    Domain: financial

    Forward-looking financial predictions and scenarios
    """

    # Type constraint
    type: Literal[DocumentType.FOR] = DocumentType.FOR

    # Domain constraint
    domain: Literal[BusinessDomain.FINANCIAL] = BusinessDomain.FINANCIAL





    def __post_init__(self):
        """Post-initialization for FOR documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.FOR
        self.domain = BusinessDomain.FINANCIAL

    def validate(self) -> list[str]:
        """Validate FOR document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors