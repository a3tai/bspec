"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.993708
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class REPDocument(BaseBSpecDocument):
    """
    Reporting (REP)
    Domain: financial

    Financial reporting and dashboard requirements
    """

    # Type constraint
    type: Literal[DocumentType.REP] = DocumentType.REP

    # Domain constraint
    domain: Literal[BusinessDomain.FINANCIAL] = BusinessDomain.FINANCIAL





    def __post_init__(self):
        """Post-initialization for REP documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.REP
        self.domain = BusinessDomain.FINANCIAL

    def validate(self) -> list[str]:
        """Validate REP document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors