"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.994211
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class CMPDocument(BaseBSpecDocument):
    """
    Compliance (CMP)
    Domain: risk

    Regulatory obligations and adherence requirements
    """

    # Type constraint
    type: Literal[DocumentType.CMP] = DocumentType.CMP

    # Domain constraint
    domain: Literal[BusinessDomain.RISK] = BusinessDomain.RISK





    def __post_init__(self):
        """Post-initialization for CMP documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.CMP
        self.domain = BusinessDomain.RISK

    def validate(self) -> list[str]:
        """Validate CMP document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors