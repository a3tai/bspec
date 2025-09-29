"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.987680
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class GAIDocument(BaseBSpecDocument):
    """
    Gains (GAI)
    Domain: customer

    Benefits and value customers receive
    """

    # Type constraint
    type: Literal[DocumentType.GAI] = DocumentType.GAI

    # Domain constraint
    domain: Literal[BusinessDomain.CUSTOMER] = BusinessDomain.CUSTOMER





    def __post_init__(self):
        """Post-initialization for GAI documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.GAI
        self.domain = BusinessDomain.CUSTOMER

    def validate(self) -> list[str]:
        """Validate GAI document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors