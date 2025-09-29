"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.991716
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class FACDocument(BaseBSpecDocument):
    """
    Facilities (FAC)
    Domain: operations

    Physical spaces, equipment, and infrastructure
    """

    # Type constraint
    type: Literal[DocumentType.FAC] = DocumentType.FAC

    # Domain constraint
    domain: Literal[BusinessDomain.OPERATIONS] = BusinessDomain.OPERATIONS





    def __post_init__(self):
        """Post-initialization for FAC documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.FAC
        self.domain = BusinessDomain.OPERATIONS

    def validate(self) -> list[str]:
        """Validate FAC document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors