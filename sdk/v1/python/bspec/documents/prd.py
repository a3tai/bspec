"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.988469
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class PRDDocument(BaseBSpecDocument):
    """
    Products (PRD)
    Domain: product

    Physical or digital products offered
    """

    # Type constraint
    type: Literal[DocumentType.PRD] = DocumentType.PRD

    # Domain constraint
    domain: Literal[BusinessDomain.PRODUCT] = BusinessDomain.PRODUCT





    def __post_init__(self):
        """Post-initialization for PRD documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.PRD
        self.domain = BusinessDomain.PRODUCT

    def validate(self) -> list[str]:
        """Validate PRD document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors