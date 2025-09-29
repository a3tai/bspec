"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.990305
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class ACTDocument(BaseBSpecDocument):
    """
    Key Activities (ACT)
    Domain: business-model

    Essential processes for value creation
    """

    # Type constraint
    type: Literal[DocumentType.ACT] = DocumentType.ACT

    # Domain constraint
    domain: Literal[BusinessDomain.BUSINESS_MODEL] = BusinessDomain.BUSINESS_MODEL





    def __post_init__(self):
        """Post-initialization for ACT documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.ACT
        self.domain = BusinessDomain.BUSINESS_MODEL

    def validate(self) -> list[str]:
        """Validate ACT document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors