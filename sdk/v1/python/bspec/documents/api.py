"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.992473
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class APIDocument(BaseBSpecDocument):
    """
    APIs (API)
    Domain: technology

    Interface specifications, protocols, and integrations
    """

    # Type constraint
    type: Literal[DocumentType.API] = DocumentType.API

    # Domain constraint
    domain: Literal[BusinessDomain.TECHNOLOGY] = BusinessDomain.TECHNOLOGY





    def __post_init__(self):
        """Post-initialization for API documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.API
        self.domain = BusinessDomain.TECHNOLOGY

    def validate(self) -> list[str]:
        """Validate API document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors