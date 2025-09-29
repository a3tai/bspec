"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.987462
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class STODocument(BaseBSpecDocument):
    """
    User Stories (STO)
    Domain: customer

    Individual requirements from user perspective
    """

    # Type constraint
    type: Literal[DocumentType.STO] = DocumentType.STO

    # Domain constraint
    domain: Literal[BusinessDomain.CUSTOMER] = BusinessDomain.CUSTOMER





    def __post_init__(self):
        """Post-initialization for STO documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.STO
        self.domain = BusinessDomain.CUSTOMER

    def validate(self) -> list[str]:
        """Validate STO document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors