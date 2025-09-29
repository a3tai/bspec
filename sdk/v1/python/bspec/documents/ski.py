"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.991298
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class SKIDocument(BaseBSpecDocument):
    """
    Skills (SKI)
    Domain: operations

    Required capabilities and competency frameworks
    """

    # Type constraint
    type: Literal[DocumentType.SKI] = DocumentType.SKI

    # Domain constraint
    domain: Literal[BusinessDomain.OPERATIONS] = BusinessDomain.OPERATIONS





    def __post_init__(self):
        """Post-initialization for SKI documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.SKI
        self.domain = BusinessDomain.OPERATIONS

    def validate(self) -> list[str]:
        """Validate SKI document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors