"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.985786
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class MOTDocument(BaseBSpecDocument):
    """
    Moats (MOT)
    Domain: strategic

    Competitive advantages that protect market position
    """

    # Type constraint
    type: Literal[DocumentType.MOT] = DocumentType.MOT

    # Domain constraint
    domain: Literal[BusinessDomain.STRATEGIC] = BusinessDomain.STRATEGIC





    def __post_init__(self):
        """Post-initialization for MOT documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.MOT
        self.domain = BusinessDomain.STRATEGIC

    def validate(self) -> list[str]:
        """Validate MOT document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors