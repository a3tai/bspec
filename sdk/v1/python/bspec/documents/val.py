"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.993498
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class VALDocument(BaseBSpecDocument):
    """
    Valuation (VAL)
    Domain: financial

    Business valuation models and assumptions
    """

    # Type constraint
    type: Literal[DocumentType.VAL] = DocumentType.VAL

    # Domain constraint
    domain: Literal[BusinessDomain.FINANCIAL] = BusinessDomain.FINANCIAL

    # Strategic foundation documents are required for all conformance levels
    strategic_foundation: bool = True




    def __post_init__(self):
        """Post-initialization for VAL documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.VAL
        self.domain = BusinessDomain.FINANCIAL

    def validate(self) -> list[str]:
        """Validate VAL document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation
        # Strategic foundation documents require success criteria
        if not self.success_criteria or len(self.success_criteria) == 0:
            errors.append("Strategic foundation documents must have success_criteria defined")




        return errors