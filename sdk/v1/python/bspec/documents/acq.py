"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.995442
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class ACQDocument(BaseBSpecDocument):
    """
    Acquisitions (ACQ)
    Domain: growth

    M&A strategy and integration planning
    """

    # Type constraint
    type: Literal[DocumentType.ACQ] = DocumentType.ACQ

    # Domain constraint
    domain: Literal[BusinessDomain.GROWTH] = BusinessDomain.GROWTH





    def __post_init__(self):
        """Post-initialization for ACQ documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.ACQ
        self.domain = BusinessDomain.GROWTH

    def validate(self) -> list[str]:
        """Validate ACQ document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors