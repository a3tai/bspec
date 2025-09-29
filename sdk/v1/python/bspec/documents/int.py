"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.989391
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class INTDocument(BaseBSpecDocument):
    """
    Integrations (INT)
    Domain: product

    Connections with other systems and platforms
    """

    # Type constraint
    type: Literal[DocumentType.INT] = DocumentType.INT

    # Domain constraint
    domain: Literal[BusinessDomain.PRODUCT] = BusinessDomain.PRODUCT





    def __post_init__(self):
        """Post-initialization for INT documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.INT
        self.domain = BusinessDomain.PRODUCT

    def validate(self) -> list[str]:
        """Validate INT document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors