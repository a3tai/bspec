//! Main BSpec implementation for v1.0.0
//! Generated from BSpec JSON SDK at 2025-09-28T21:08:19.445250

use crate::types::*;
use serde_json;
use std::collections::HashMap;
use thiserror::Error;

/// BSpec error types
#[derive(Error, Debug)]
pub enum BSpecError {
    #[error("JSON parsing error: {0}")]
    JsonError(#[from] serde_json::Error),

    #[error("Domain not found: {0}")]
    DomainNotFound(String),

    #[error("Document type not found: {0}")]
    DocumentTypeNotFound(String),

    #[error("File not found: {0}")]
    FileNotFound(String),
}

/// Main BSpec structure for working with BSpec data
#[derive(Debug, Clone)]
pub struct BSpec {
    data: BSpecData,
}

impl BSpec {
    /// Create BSpec instance from JSON data
    pub fn from_json(json_data: &[u8]) -> Result<Self, BSpecError> {
        let data: BSpecData = serde_json::from_slice(json_data)?;
        Ok(Self { data })
    }

    /// Get BSpec version
    pub fn version(&self) -> &str {
        &self.data.metadata.bspec_version
    }

    /// Get generation metadata
    pub fn metadata(&self) -> &Metadata {
        &self.data.metadata
    }

    /// Get specification statistics
    pub fn statistics(&self) -> &Statistics {
        &self.data.statistics
    }

    /// Get all business domains
    pub fn domains(&self) -> &[Domain] {
        &self.data.domains
    }

    /// Get all document types
    pub fn document_types(&self) -> &[DocumentType] {
        &self.data.document_types
    }

    /// Get all files in the specification
    pub fn files(&self) -> &HashMap<String, File> {
        &self.data.files
    }

    /// Get conformance levels
    pub fn conformance_levels(&self) -> &[ConformanceLevel] {
        &self.data.conformance_levels
    }

    /// Get YAML schema definition
    pub fn yaml_schema(&self) -> &YAMLSchema {
        &self.data.yaml_schema
    }

    /// Get directory structure
    pub fn directory_structure(&self) -> &[Vec<String>] {
        &self.data.directory_structure
    }

    /// Get document index
    pub fn document_index(&self) -> &[DocumentIndexEntry] {
        &self.data.document_index
    }

    /// Get a specific domain by name
    pub fn get_domain(&self, name: &str) -> Option<&Domain> {
        self.data.domains.iter().find(|domain| domain.name == name)
    }

    /// Get a specific document type by code
    pub fn get_document_type(&self, code: &str) -> Option<&DocumentType> {
        self.data.document_types.iter().find(|doc_type| doc_type.code == code)
    }

    /// Get document types for a specific domain
    pub fn get_document_types_for_domain(&self, domain_name: &str) -> Vec<&DocumentType> {
        if let Some(domain) = self.get_domain(domain_name) {
            self.data.document_types
                .iter()
                .filter(|doc_type| domain.document_types.contains(&doc_type.code))
                .collect()
        } else {
            Vec::new()
        }
    }

    /// Get a specific file by path
    pub fn get_file(&self, path: &str) -> Option<&File> {
        self.data.files.get(path)
    }

    /// Get all files of a specific type
    pub fn get_files_by_type(&self, file_type: &str) -> Vec<&File> {
        self.data.files
            .values()
            .filter(|file| file.file_type == file_type)
            .collect()
    }

    /// Get markdown files with frontmatter
    pub fn get_document_files(&self) -> Vec<&File> {
        self.data.files
            .values()
            .filter(|file| file.file_type == "markdown" && file.has_frontmatter)
            .collect()
    }

    /// Search for document types by name, purpose, or code
    pub fn search_document_types(&self, query: &str) -> Vec<&DocumentType> {
        let query_lower = query.to_lowercase();
        self.data.document_types
            .iter()
            .filter(|doc_type| {
                doc_type.name.to_lowercase().contains(&query_lower) ||
                doc_type.purpose.to_lowercase().contains(&query_lower) ||
                doc_type.code.to_lowercase().contains(&query_lower)
            })
            .collect()
    }

    /// Get domains containing a specific document type
    pub fn get_domains_for_document_type(&self, code: &str) -> Vec<&Domain> {
        self.data.domains
            .iter()
            .filter(|domain| domain.document_types.contains(&code.to_string()))
            .collect()
    }

    /// Get conformance level by name
    pub fn get_conformance_level(&self, name: &str) -> Option<&ConformanceLevel> {
        self.data.conformance_levels
            .iter()
            .find(|level| level.name == name)
    }

    /// Get required YAML fields
    pub fn get_required_fields(&self) -> &[String] {
        &self.data.yaml_schema.required_fields
    }

    /// Get optional YAML fields
    pub fn get_optional_fields(&self) -> &[String] {
        &self.data.yaml_schema.optional_fields
    }

    /// Convert back to JSON
    pub fn to_json(&self) -> Result<String, BSpecError> {
        Ok(serde_json::to_string_pretty(&self.data)?)
    }

    /// Convert back to JSON bytes
    pub fn to_json_bytes(&self) -> Result<Vec<u8>, BSpecError> {
        Ok(serde_json::to_vec_pretty(&self.data)?)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_bspec_creation() {
        // This would need actual JSON data to test properly
        // In a real implementation, you'd load test data
    }
}
