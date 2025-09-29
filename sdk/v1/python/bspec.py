"""
BSpec Python SDK

Python SDK for the BSpec v1.0.0 Universal Business Specification Standard.
Generated from BSpec JSON SDK.
"""

import json
from typing import Dict, List, Optional, Any
from dataclasses import dataclass
from datetime import datetime
from pathlib import Path


@dataclass
class BSpecMetadata:
    """BSpec generation metadata"""
    bspec_version: str
    generated_at: datetime
    generator: str
    source_directory: str


@dataclass
class BSpecStatistics:
    """BSpec specification statistics"""
    total_files: int
    total_document_types: int
    total_domains: int


@dataclass
class BSpecDomain:
    """Business domain definition"""
    name: str
    description: str
    document_types: List[str]
    document_count: int


@dataclass
class BSpecDocumentType:
    """Document type definition"""
    code: str
    name: str
    purpose: str
    domain: str
    examples: List[Dict[str, str]]


@dataclass
class BSpecFile:
    """File in the specification"""
    path: str
    name: str
    type: str
    extension: str
    has_frontmatter: bool
    frontmatter: Dict[str, Any]
    content: str


@dataclass
class BSpecConformanceLevel:
    """Conformance level definition"""
    name: str
    display_name: str
    description: str
    emoji: str
    min_documents: int


@dataclass
class BSpecYAMLSchema:
    """YAML schema definition"""
    required_fields: List[str]
    optional_fields: List[str]
    field_types: Dict[str, str]
    enums: Dict[str, List[str]]
    defaults: Dict[str, Any]


@dataclass
class BSpecDocumentIndexEntry:
    """Document index entry"""
    id: str
    title: str
    type: str
    path: str
    domain: str
    owner: str
    status: str
    version: str


class BSpec:
    """Main BSpec class for working with specification data"""

    def __init__(self, data: Dict[str, Any]):
        """Initialize BSpec from JSON data"""
        self.data = data
        self._metadata = None
        self._statistics = None
        self._domains = None
        self._document_types = None
        self._files = None
        self._conformance_levels = None
        self._yaml_schema = None
        self._document_index = None

    @classmethod
    def from_json_file(cls, json_file_path: str) -> 'BSpec':
        """Create BSpec instance from JSON file"""
        with open(json_file_path, 'r', encoding='utf-8') as f:
            data = json.load(f)
        return cls(data)

    @classmethod
    def from_json_string(cls, json_string: str) -> 'BSpec':
        """Create BSpec instance from JSON string"""
        data = json.loads(json_string)
        return cls(data)

    @property
    def metadata(self) -> BSpecMetadata:
        """Get metadata"""
        if self._metadata is None:
            meta = self.data['metadata']
            self._metadata = BSpecMetadata(
                bspec_version=meta['bspec_version'],
                generated_at=datetime.fromisoformat(meta['generated_at'].replace('Z', '+00:00')),
                generator=meta['generator'],
                source_directory=meta['source_directory']
            )
        return self._metadata

    @property
    def version(self) -> str:
        """Get BSpec version"""
        return self.metadata.bspec_version

    @property
    def statistics(self) -> BSpecStatistics:
        """Get statistics"""
        if self._statistics is None:
            stats = self.data['statistics']
            self._statistics = BSpecStatistics(
                total_files=stats['total_files'],
                total_document_types=stats['total_document_types'],
                total_domains=stats['total_domains']
            )
        return self._statistics

    @property
    def domains(self) -> List[BSpecDomain]:
        """Get all domains"""
        if self._domains is None:
            self._domains = [
                BSpecDomain(
                    name=domain['name'],
                    description=domain['description'],
                    document_types=domain['document_types'],
                    document_count=domain['document_count']
                )
                for domain in self.data['domains']
            ]
        return self._domains

    @property
    def document_types(self) -> List[BSpecDocumentType]:
        """Get all document types"""
        if self._document_types is None:
            self._document_types = [
                BSpecDocumentType(
                    code=doc['code'],
                    name=doc['name'],
                    purpose=doc['purpose'],
                    domain=doc['domain'],
                    examples=doc.get('examples', [])
                )
                for doc in self.data['document_types']
            ]
        return self._document_types

    @property
    def files(self) -> Dict[str, BSpecFile]:
        """Get all files"""
        if self._files is None:
            self._files = {}
            for path, file_data in self.data['files'].items():
                self._files[path] = BSpecFile(
                    path=file_data['path'],
                    name=file_data['name'],
                    type=file_data['type'],
                    extension=file_data['extension'],
                    has_frontmatter=file_data['has_frontmatter'],
                    frontmatter=file_data['frontmatter'],
                    content=file_data['content']
                )
        return self._files

    def get_domain(self, name: str) -> Optional[BSpecDomain]:
        """Get domain by name"""
        for domain in self.domains:
            if domain.name == name:
                return domain
        return None

    def get_document_type(self, code: str) -> Optional[BSpecDocumentType]:
        """Get document type by code"""
        for doc_type in self.document_types:
            if doc_type.code == code:
                return doc_type
        return None

    def get_document_types_for_domain(self, domain_name: str) -> List[BSpecDocumentType]:
        """Get all document types for a domain"""
        domain = self.get_domain(domain_name)
        if not domain:
            return []

        result = []
        for code in domain.document_types:
            doc_type = self.get_document_type(code)
            if doc_type:
                result.append(doc_type)
        return result

    def get_file(self, path: str) -> Optional[BSpecFile]:
        """Get file by path"""
        return self.files.get(path)

    def get_files_by_type(self, file_type: str) -> List[BSpecFile]:
        """Get files by type"""
        return [file for file in self.files.values() if file.type == file_type]

    def search_document_types(self, query: str) -> List[BSpecDocumentType]:
        """Search document types by name, purpose, or code"""
        query_lower = query.lower()
        result = []
        for doc_type in self.document_types:
            if (query_lower in doc_type.name.lower() or
                query_lower in doc_type.purpose.lower() or
                query_lower in doc_type.code.lower()):
                result.append(doc_type)
        return result

    def get_summary(self) -> Dict[str, Any]:
        """Get specification summary"""
        return {
            'version': self.version,
            'statistics': {
                'total_files': self.statistics.total_files,
                'total_document_types': self.statistics.total_document_types,
                'total_domains': self.statistics.total_domains
            },
            'domains': [
                {
                    'name': domain.name,
                    'description': domain.description,
                    'document_count': domain.document_count
                }
                for domain in self.domains
            ]
        }

    def to_json(self) -> str:
        """Export to JSON string"""
        return json.dumps(self.data, indent=2, ensure_ascii=False)


def load_bspec(json_file_path: str) -> BSpec:
    """Load BSpec from JSON file"""
    return BSpec.from_json_file(json_file_path)


def load_bspec_from_string(json_string: str) -> BSpec:
    """Load BSpec from JSON string"""
    return BSpec.from_json_string(json_string)