"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.992265
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class SYSDocument(BaseBSpecDocument):
    """
    Systems (SYS)
    Domain: technology

    Software systems, platforms, and applications
    """

    # Type constraint
    type: Literal[DocumentType.SYS] = DocumentType.SYS

    # Domain constraint
    domain: Literal[BusinessDomain.TECHNOLOGY] = BusinessDomain.TECHNOLOGY





    def __post_init__(self):
        """Post-initialization for SYS documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.SYS
        self.domain = BusinessDomain.TECHNOLOGY

    def validate(self) -> list[str]:
        """Validate SYS document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors