"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.990916
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class WFLDocument(BaseBSpecDocument):
    """
    Workflows (WFL)
    Domain: operations

    Detailed task sequences within processes
    """

    # Type constraint
    type: Literal[DocumentType.WFL] = DocumentType.WFL

    # Domain constraint
    domain: Literal[BusinessDomain.OPERATIONS] = BusinessDomain.OPERATIONS





    def __post_init__(self):
        """Post-initialization for WFL documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.WFL
        self.domain = BusinessDomain.OPERATIONS

    def validate(self) -> list[str]:
        """Validate WFL document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors