"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.988983
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class REQDocument(BaseBSpecDocument):
    """
    Requirements (REQ)
    Domain: product

    Functional and non-functional specifications
    """

    # Type constraint
    type: Literal[DocumentType.REQ] = DocumentType.REQ

    # Domain constraint
    domain: Literal[BusinessDomain.PRODUCT] = BusinessDomain.PRODUCT





    def __post_init__(self):
        """Post-initialization for REQ documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.REQ
        self.domain = BusinessDomain.PRODUCT

    def validate(self) -> list[str]:
        """Validate REQ document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors