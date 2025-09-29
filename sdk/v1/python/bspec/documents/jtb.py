"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.987153
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class JTBDocument(BaseBSpecDocument):
    """
    Jobs-to-be-Done (JTB)
    Domain: customer

    Specific outcomes customers hire products to achieve
    """

    # Type constraint
    type: Literal[DocumentType.JTB] = DocumentType.JTB

    # Domain constraint
    domain: Literal[BusinessDomain.CUSTOMER] = BusinessDomain.CUSTOMER


    # Customer understanding documents require specific validation
    customer_focused: bool = True



    def __post_init__(self):
        """Post-initialization for JTB documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.JTB
        self.domain = BusinessDomain.CUSTOMER

    def validate(self) -> list[str]:
        """Validate JTB document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation

        # Customer understanding documents should have related documents
        if not self.related or len(self.related) == 0:
            errors.append("Customer understanding documents should reference related personas or jobs-to-be-done")



        return errors