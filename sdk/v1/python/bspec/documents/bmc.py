"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.989587
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class BMCDocument(BaseBSpecDocument):
    """
    Business Model Canvas (BMC)
    Domain: business-model

    Complete business model overview
    """

    # Type constraint
    type: Literal[DocumentType.BMC] = DocumentType.BMC

    # Domain constraint
    domain: Literal[BusinessDomain.BUSINESS_MODEL] = BusinessDomain.BUSINESS_MODEL





    def __post_init__(self):
        """Post-initialization for BMC documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.BMC
        self.domain = BusinessDomain.BUSINESS_MODEL

    def validate(self) -> list[str]:
        """Validate BMC document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors