"""
Generated BSpec Document Type Class

GENERATED CODE - DO NOT EDIT MANUALLY
Generated from BSpec v1.0.0 specification
Generated at: 2025-09-28T15:10:35.994016
Generator: python-generator v1.0.0
"""

from dataclasses import dataclass
from typing import Literal
from .base import BaseBSpecDocument, DocumentType, BusinessDomain


@dataclass
class RSKDocument(BaseBSpecDocument):
    """
    Risks (RSK)
    Domain: risk

    Identified threats to business success
    """

    # Type constraint
    type: Literal[DocumentType.RSK] = DocumentType.RSK

    # Domain constraint
    domain: Literal[BusinessDomain.RISK] = BusinessDomain.RISK




    # Risk management documents require mitigation planning
    risk_management: bool = True

    def __post_init__(self):
        """Post-initialization for RSK documents"""
        super().__post_init__()

        # Ensure type is correctly set
        self.type = DocumentType.RSK
        self.domain = BusinessDomain.RISK

    def validate(self) -> list[str]:
        """Validate RSK document and return list of validation errors"""
        errors = super().validate()

        # Document type specific validation



        # Risk management documents should reference each other
        if self.type == DocumentType.RSK and not any(ref.startswith('MIT-') for ref in (self.related or [])):
            errors.append("Risk documents should reference corresponding mitigation documents")
        elif self.type == DocumentType.MIT and not any(ref.startswith('RSK-') for ref in (self.related or [])):
            errors.append("Mitigation documents should reference corresponding risk documents")

        return errors