"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.985978
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class THYDocument(BaseBSpecDocument):
    """
    Theory of Change (THY)
    Domain: strategic

    Logic model connecting activities to outcomes
    """

    # Type constraint
    type: Literal[DocumentType.THY] = DocumentType.THY

    # Domain constraint
    domain: Literal[BusinessDomain.STRATEGIC] = BusinessDomain.STRATEGIC





    def __post_init__(self):
        """Post-initialization for THY documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.THY
        self.domain = BusinessDomain.STRATEGIC

    def validate(self) -> list[str]:
        """Validate THY document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors