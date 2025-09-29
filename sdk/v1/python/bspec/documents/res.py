"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.990210
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class RESDocument(BaseBSpecDocument):
    """
    Key Resources (RES)
    Domain: business-model

    Critical assets required for the business model
    """

    # Type constraint
    type: Literal[DocumentType.RES] = DocumentType.RES

    # Domain constraint
    domain: Literal[BusinessDomain.BUSINESS_MODEL] = BusinessDomain.BUSINESS_MODEL





    def __post_init__(self):
        """Post-initialization for RES documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.RES
        self.domain = BusinessDomain.BUSINESS_MODEL

    def validate(self) -> list[str]:
        """Validate RES document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors