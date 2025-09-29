"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.994825
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class GTMDocument(BaseBSpecDocument):
    """
    Go-to-Market (GTM)
    Domain: growth

    Launch and customer acquisition strategy
    """

    # Type constraint
    type: Literal[DocumentType.GTM] = DocumentType.GTM

    # Domain constraint
    domain: Literal[BusinessDomain.GROWTH] = BusinessDomain.GROWTH





    def __post_init__(self):
        """Post-initialization for GTM documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.GTM
        self.domain = BusinessDomain.GROWTH

    def validate(self) -> list[str]:
        """Validate GTM document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors