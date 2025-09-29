"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.986080
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class MKTDocument(BaseBSpecDocument):
    """
    Market Definition (MKT)
    Domain: market

    TAM/SAM/SOM, market boundaries and sizing
    """

    # Type constraint
    type: Literal[DocumentType.MKT] = DocumentType.MKT

    # Domain constraint
    domain: Literal[BusinessDomain.MARKET] = BusinessDomain.MARKET





    def __post_init__(self):
        """Post-initialization for MKT documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.MKT
        self.domain = BusinessDomain.MARKET

    def validate(self) -> list[str]:
        """Validate MKT document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors