"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.987258
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class CJMDocument(BaseBSpecDocument):
    """
    Customer Journey (CJM)
    Domain: customer

    End-to-end experience from awareness to advocacy
    """

    # Type constraint
    type: Literal[DocumentType.CJM] = DocumentType.CJM

    # Domain constraint
    domain: Literal[BusinessDomain.CUSTOMER] = BusinessDomain.CUSTOMER





    def __post_init__(self):
        """Post-initialization for CJM documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.CJM
        self.domain = BusinessDomain.CUSTOMER

    def validate(self) -> list[str]:
        """Validate CJM document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors