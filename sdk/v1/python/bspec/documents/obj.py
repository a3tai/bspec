"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.985686
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class OBJDocument(BaseBSpecDocument):
    """
    Objectives (OBJ)
    Domain: strategic

    Specific, measurable goals with timeframes (OKRs)
    """

    # Type constraint
    type: Literal[DocumentType.OBJ] = DocumentType.OBJ

    # Domain constraint
    domain: Literal[BusinessDomain.STRATEGIC] = BusinessDomain.STRATEGIC





    def __post_init__(self):
        """Post-initialization for OBJ documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.OBJ
        self.domain = BusinessDomain.STRATEGIC

    def validate(self) -> list[str]:
        """Validate OBJ document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors