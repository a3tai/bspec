"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.987580
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class PAIDocument(BaseBSpecDocument):
    """
    Pain Points (PAI)
    Domain: customer

    Customer problems and frustrations being solved
    """

    # Type constraint
    type: Literal[DocumentType.PAI] = DocumentType.PAI

    # Domain constraint
    domain: Literal[BusinessDomain.CUSTOMER] = BusinessDomain.CUSTOMER





    def __post_init__(self):
        """Post-initialization for PAI documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.PAI
        self.domain = BusinessDomain.CUSTOMER

    def validate(self) -> list[str]:
        """Validate PAI document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors