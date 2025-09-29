"""
BSpec Python SDK - Main BSpec Class
Generated from BSpec v1.0.0 JSON SDK
"""

import json
from typing import List, Optional, Dict, Any
from .types import (
    BSpecData, BSpecDomain, BSpecDocumentType, BSpecFile,
    BSpecStatistics, BusinessDomain, ConformanceLevel
)


class BSpec:
    """Main BSpec class for working with BSpec data"""

    def __init__(self, data: Dict[str, Any]):
        """Initialize BSpec with JSON data"""
        self._data = data

    @classmethod
    def from_json(cls, json_data: Dict[str, Any]) -> 'BSpec':
        """Create BSpec instance from JSON data"""
        return cls(json_data)

    @classmethod
    def load_from_file(cls, file_path: str) -> 'BSpec':
        """Load BSpec from JSON file"""
        with open(file_path, 'r', encoding='utf-8') as f:
            data = json.load(f)
        return cls.from_json(data)

    @property
    def version(self) -> str:
        """Get BSpec version"""
        return self._data['metadata']['bspec_version']

    @property
    def metadata(self) -> Dict[str, Any]:
        """Get generation metadata"""
        return self._data['metadata']

    @property
    def statistics(self) -> Dict[str, Any]:
        """Get specification statistics"""
        return self._data['statistics']

    @property
    def domains(self) -> List[Dict[str, Any]]:
        """Get all business domains"""
        return self._data['domains']

    @property
    def document_types(self) -> List[Dict[str, Any]]:
        """Get all document types"""
        return self._data['document_types']

    @property
    def files(self) -> Dict[str, Any]:
        """Get all files in the specification"""
        return self._data['files']

    @property
    def conformance_levels(self) -> List[Dict[str, Any]]:
        """Get conformance levels"""
        return self._data['conformance_levels']

    @property
    def yaml_schema(self) -> Dict[str, Any]:
        """Get YAML schema definition"""
        return self._data['yaml_schema']

    def get_domain(self, name: str) -> Optional[Dict[str, Any]]:
        """Get a specific domain by name"""
        for domain in self.domains:
            if domain['name'] == name:
                return domain
        return None

    def get_document_type(self, code: str) -> Optional[Dict[str, Any]]:
        """Get a specific document type by code"""
        for doc_type in self.document_types:
            if doc_type['code'] == code:
                return doc_type
        return None

    def get_document_types_for_domain(self, domain_name: str) -> List[Dict[str, Any]]:
        """Get document types for a specific domain"""
        domain = self.get_domain(domain_name)
        if not domain:
            return []

        return [dt for dt in self.document_types
                if dt['code'] in domain['document_types']]

    def get_file(self, path: str) -> Optional[Dict[str, Any]]:
        """Get a specific file by path"""
        return self.files.get(path)

    def get_files_by_type(self, file_type: str) -> List[Dict[str, Any]]:
        """Get all files of a specific type"""
        return [f for f in self.files.values() if f.get('type') == file_type]

    def search_document_types(self, query: str) -> List[Dict[str, Any]]:
        """Search for document types by name, purpose, or code"""
        query_lower = query.lower()
        results = []

        for doc_type in self.document_types:
            if (query_lower in doc_type['name'].lower() or
                query_lower in doc_type['purpose'].lower() or
                query_lower in doc_type['code'].lower()):
                results.append(doc_type)

        return results

    def get_summary(self) -> Dict[str, Any]:
        """Get summary information about the specification"""
        return {
            'version': self.version,
            'statistics': self.statistics,
            'domains': [{
                'name': d['name'],
                'description': d['description'],
                'document_count': d['document_count']
            } for d in self.domains],
            'conformance_levels': [cl['name'] for cl in self.conformance_levels],
            'generated_at': self.metadata['generated_at'],
            'generator': self.metadata['generator']
        }

    def to_dict(self) -> Dict[str, Any]:
        """Export to dictionary"""
        return self._data
