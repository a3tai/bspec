//! Workflows (WFL) document implementation
//!
//! Detailed task sequences within processes
//!
//! **GENERATED CODE - DO NOT EDIT MANUALLY**
//! Generated at: 2025-09-28T16:00:48.173975

use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

use crate::base::{BSpecDocument, DocumentMetadata, DocumentStatus, DocumentType, Domain};
use crate::error::{BSpecError, Result};

/// Workflows document
///
/// Detailed task sequences within processes
///
/// **Domain:** Operations
/// **Conformance Requirements:**#[derive(Debug, Clone, PartialEq, Serialize, Deserialize)]
pub struct WflDocument {
    /// Common document metadata
    #[serde(flatten)]
    pub metadata: DocumentMetadata,

    /// Document content in Markdown format
    pub content: String,}

impl WflDocument {
    /// Create a new Workflows document
    pub fn new(id: String, title: String, owner: String) -> Self {
        let mut metadata = DocumentMetadata::default();
        metadata.id = id;
        metadata.title = title;
        metadata.owner = owner;
        metadata.domain = Domain::Operations;

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

impl BSpecDocument for WflDocument {
    fn document_type(&self) -> DocumentType {
        DocumentType::Wfl
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
        // Validate ID format (should be WFL-something)
        if !self.metadata.id.starts_with("WFL-") {
            return Err(BSpecError::ValidationError(
                format!("ID should start with 'WFL-', got '{}'", self.metadata.id)
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

impl Default for WflDocument {
    fn default() -> Self {
        Self::new(
            "WFL-default".to_string(),
            "New Workflows".to_string(),
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
        let doc = WflDocument::new(
            "WFL-test-001".to_string(),
            "Test Workflows".to_string(),
            "Test Owner".to_string(),
        );

        assert_eq!(doc.metadata.id, "WFL-test-001");
        assert_eq!(doc.metadata.title, "Test Workflows");
        assert_eq!(doc.metadata.owner, "Test Owner");
        assert_eq!(doc.document_type(), DocumentType::Wfl);
    }

    #[test]
    fn test_validation() {
        let mut doc = WflDocument::default();

        // Should fail validation initially
        #[cfg(feature = "validation")]
        assert!(doc.validate().is_err());

        // Set required fields
        doc.metadata.id = "WFL-valid-001".to_string();
        doc.metadata.title = "Valid Document".to_string();
        doc.metadata.owner = "Valid Owner".to_string();
        // Should pass validation now
        #[cfg(feature = "validation")]
        assert!(doc.validate().is_ok());
    }

    #[test]
    fn test_serialization() {
        let doc = WflDocument::new(
            "WFL-ser-001".to_string(),
            "Serialization Test".to_string(),
            "Test User".to_string(),
        );

        // Test JSON serialization
        let json = doc.to_json().unwrap();
        assert!(json.contains("WFL-ser-001"));

        // Test YAML serialization
        let yaml = doc.to_yaml().unwrap();
        assert!(yaml.contains("WFL-ser-001"));

        // Test deserialization
        let deserialized: WflDocument = serde_json::from_str(&json).unwrap();
        assert_eq!(deserialized.metadata.id, doc.metadata.id);
    }
}