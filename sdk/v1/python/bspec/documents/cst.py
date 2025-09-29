"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.989899
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class CSTDocument(BaseBSpecDocument):
    """
    Cost Structure (CST)
    Domain: business-model

    Major cost categories and cost drivers
    """

    # Type constraint
    type: Literal[DocumentType.CST] = DocumentType.CST

    # Domain constraint
    domain: Literal[BusinessDomain.BUSINESS_MODEL] = BusinessDomain.BUSINESS_MODEL



    # Business model documents are critical for financial validation
    business_model_core: bool = True


    def __post_init__(self):
        """Post-initialization for CST documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.CST
        self.domain = BusinessDomain.BUSINESS_MODEL

    def validate(self) -> list[str]:
        """Validate CST document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation


        # Business model documents require metrics
        if not self.metrics or len(self.metrics) == 0:
            errors.append("Business model documents must have metrics defined for measurement")


        return errors