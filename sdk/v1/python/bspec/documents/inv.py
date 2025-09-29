"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.993401
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class INVDocument(BaseBSpecDocument):
    """
    Investment (INV)
    Domain: financial

    Capital allocation and investment decisions
    """

    # Type constraint
    type: Literal[DocumentType.INV] = DocumentType.INV

    # Domain constraint
    domain: Literal[BusinessDomain.FINANCIAL] = BusinessDomain.FINANCIAL





    def __post_init__(self):
        """Post-initialization for INV documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.INV
        self.domain = BusinessDomain.FINANCIAL

    def validate(self) -> list[str]:
        """Validate INV document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors