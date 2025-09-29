"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.985243
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class VSNDocument(BaseBSpecDocument):
    """
    Vision (VSN)
    Domain: strategic

    The future state the organization aims to create
    """

    # Type constraint
    type: Literal[DocumentType.VSN] = DocumentType.VSN

    # Domain constraint
    domain: Literal[BusinessDomain.STRATEGIC] = BusinessDomain.STRATEGIC

    # Strategic foundation documents are required for all conformance levels
    strategic_foundation: bool = True




    def __post_init__(self):
        """Post-initialization for VSN documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.VSN
        self.domain = BusinessDomain.STRATEGIC

    def validate(self) -> list[str]:
        """Validate VSN document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation
        # Strategic foundation documents require success criteria
        if not self.success_criteria or len(self.success_criteria) == 0:
            errors.append("Strategic foundation documents must have success_criteria defined")




        return errors