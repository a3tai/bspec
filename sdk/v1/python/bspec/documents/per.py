"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.989289
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class PERDocument(BaseBSpecDocument):
    """
    Performance (PER)
    Domain: product

    Speed, reliability, and performance characteristics
    """

    # Type constraint
    type: Literal[DocumentType.PER] = DocumentType.PER

    # Domain constraint
    domain: Literal[BusinessDomain.PRODUCT] = BusinessDomain.PRODUCT


    # Customer understanding documents require specific validation
    customer_focused: bool = True



    def __post_init__(self):
        """Post-initialization for PER documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.PER
        self.domain = BusinessDomain.PRODUCT

    def validate(self) -> list[str]:
        """Validate PER document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation

        # Customer understanding documents should have related documents
        if not self.related or len(self.related) == 0:
            errors.append("Customer understanding documents should reference related personas or jobs-to-be-done")



        return errors