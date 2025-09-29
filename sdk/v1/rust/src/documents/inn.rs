//! Innovation (INN) document implementation
//!
//! Future product and service development pipeline
//!
//! **GENERATED CODE - DO NOT EDIT MANUALLY**
//! Generated at: 2025-09-28T16:00:48.181703

use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

use crate::base::{BSpecDocument, DocumentMetadata, DocumentStatus, DocumentType, Domain};
use crate::error::{BSpecError, Result};

/// Innovation document
///
/// Future product and service development pipeline
///
/// **Domain:** Growth
/// **Conformance Requirements:**#[derive(Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct InnDocument {
    /// Common document metadata
    #[serde(flatten)]
    pub metadata: DocumentMetadata,

    /// Document content in Markdown format
    pub content: String,}

impl InnDocument {
    /// Create a new Innovation document
    pub fn new(id: String, title: String, owner: String) -> Self {
        let mut metadata = DocumentMetadata::default();
        metadata.id = id;
        metadata.title = title;
        metadata.owner = owner;
        metadata.domain = Domain::Growth;

        Self {
            metadata,
            content: String::new(),        }
    }
    /// Set the main content of the document
    pub fn set_content(&mut self, content: String) {
        self.content = content;
        self.touch();
    }

    /// Update document status
    pub fn set_status(&mut self, status: DocumentStatus) {
        self.metadata.status = status;
        self.touch();
    }

    /// Add a related document ID
    pub fn add_related(&mut self, doc_id: String) {
        if !self.metadata.related.contains(&doc_id) {
            self.metadata.related.push(doc_id);
            self.touch();
        }
    }

    /// Add a tag
    pub fn add_tag(&mut self, tag: String) {
        if !self.metadata.tags.contains(&tag) {
            self.metadata.tags.push(tag);
            self.touch();
        }
    }
}

impl BSpecDocument for InnDocument {
    fn document_type(&self) -> DocumentType {
        DocumentType::Inn
    }

    fn metadata(&self) -> &DocumentMetadata {
        &self.metadata
    }

    fn metadata_mut(&mut self) -> &mut DocumentMetadata {
        &mut self.metadata
    }

    fn content(&self) -> &str {
        &self.content
    }

    fn set_content(&mut self, content: String) {
        self.set_content(content);
    }

    #[cfg(feature = "validation")]
    fn validate(&self) -> Result<()> {
        // Validate required fields
        if self.metadata.id.is_empty() {
            return Err(BSpecError::ValidationError("id cannot be empty".to_string()));
        }

        if self.metadata.title.is_empty() {
            return Err(BSpecError::ValidationError("title cannot be empty".to_string()));
        }

        if self.metadata.owner.is_empty() {
            return Err(BSpecError::ValidationError("owner cannot be empty".to_string()));
        }
        // Validate ID format (should be INN-something)
        if !self.metadata.id.starts_with("INN-") {
            return Err(BSpecError::ValidationError(
                format!("ID should start with 'INN-', got '{}'", self.metadata.id)
            ));
        }

        // Validate version format (semantic versioning)
        if !is_valid_semver(&self.metadata.version) {
            return Err(BSpecError::ValidationError(
                format!("Invalid semantic version: {}", self.metadata.version)
            ));
        }

        Ok(())
    }
}

impl Default for InnDocument {
    fn default() -> Self {
        Self::new(
            "INN-default".to_string(),
            "New Innovation".to_string(),
            "Unknown".to_string(),
        )
    }
}

/// Validate semantic version format
fn is_valid_semver(version: &str) -> bool {
    use regex::Regex;
    lazy_static::lazy_static! {
        static ref SEMVER_REGEX: Regex = Regex::new(r"^\d+\.\d+\.\d+(-[\w\.-]+)?(\+[\w\.-]+)?$").unwrap();
    }
    SEMVER_REGEX.is_match(version)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_new_document() {
        let doc = InnDocument::new(
            "INN-test-001".to_string(),
            "Test Innovation".to_string(),
            "Test Owner".to_string(),
        );

        assert_eq!(doc.metadata.id, "INN-test-001");
        assert_eq!(doc.metadata.title, "Test Innovation");
        assert_eq!(doc.metadata.owner, "Test Owner");
        assert_eq!(doc.document_type(), DocumentType::Inn);
    }

    #[test]
    fn test_validation() {
        let mut doc = InnDocument::default();

        // Should fail validation initially
        #[cfg(feature = "validation")]
        assert!(doc.validate().is_err());

        // Set required fields
        doc.metadata.id = "INN-valid-001".to_string();
        doc.metadata.title = "Valid Document".to_string();
        doc.metadata.owner = "Valid Owner".to_string();
        // Should pass validation now
        #[cfg(feature = "validation")]
        assert!(doc.validate().is_ok());
    }

    #[test]
    fn test_serialization() {
        let doc = InnDocument::new(
            "INN-ser-001".to_string(),
            "Serialization Test".to_string(),
            "Test User".to_string(),
        );

        // Test JSON serialization
        let json = doc.to_json().unwrap();
        assert!(json.contains("INN-ser-001"));

        // Test YAML serialization
        let yaml = doc.to_yaml().unwrap();
        assert!(yaml.contains("INN-ser-001"));

        // Test deserialization
        let deserialized: InnDocument = serde_json::from_str(&json).unwrap();
        assert_eq!(deserialized.metadata.id, doc.metadata.id);
    }
}