"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.992012
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class CAPDocument(BaseBSpecDocument):
    """
    Capabilities (CAP)
    Domain: operations

    Organizational capabilities and competencies
    """

    # Type constraint
    type: Literal[DocumentType.CAP] = DocumentType.CAP

    # Domain constraint
    domain: Literal[BusinessDomain.OPERATIONS] = BusinessDomain.OPERATIONS





    def __post_init__(self):
        """Post-initialization for CAP documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.CAP
        self.domain = BusinessDomain.OPERATIONS

    def validate(self) -> list[str]:
        """Validate CAP document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation




        return errors