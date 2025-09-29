//! Types for BSpec v1.0.0
//! Generated from BSpec JSON SDK at 2025-09-28T21:08:19.445123

use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use chrono::{DateTime, Utc};

/// Metadata contains generation metadata
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Metadata {
    pub bspec_version: String,
    pub generated_at: DateTime<Utc>,
    pub generator: String,
    pub source_spec: String,
}

/// Statistics contains specification statistics
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Statistics {
    pub total_files: usize,
    pub total_document_types: usize,
    pub total_domains: usize,
}

/// Domain represents a business domain
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Domain {
    pub name: String,
    pub display_name: String,
    pub description: String,
    pub emoji: String,
    pub document_types: Vec<String>,
    pub document_count: usize,
}

/// DocumentType represents a BSpec document type
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DocumentType {
    pub code: String,
    pub name: String,
    pub purpose: String,
    pub domain: String,
    pub examples: Vec<DocumentExample>,
}

/// DocumentExample represents an example document
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DocumentExample {
    pub id: String,
    pub title: String,
    pub description: String,
}

/// File represents a file in the specification
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct File {
    pub path: String,
    #[serde(rename = "type")]
    pub file_type: String,
    pub size: usize,
    pub content: String,
    pub has_frontmatter: bool,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub frontmatter: Option<serde_json::Value>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub parsed_content: Option<String>,
}

/// ConformanceLevel represents a conformance level
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ConformanceLevel {
    pub name: String,
    pub display_name: String,
    pub description: String,
    pub emoji: String,
    pub min_documents: usize,
}

/// YAMLSchema represents the YAML schema definition
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct YAMLSchema {
    pub required_fields: Vec<String>,
    pub optional_fields: Vec<String>,
    pub field_types: HashMap<String, String>,
    pub enums: HashMap<String, Vec<String>>,
    pub defaults: HashMap<String, serde_json::Value>,
}

/// DocumentIndexEntry represents an entry in the document index
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DocumentIndexEntry {
    pub id: String,
    pub title: String,
    #[serde(rename = "type")]
    pub entry_type: String,
    pub path: String,
    pub domain: String,
    pub owner: String,
    pub status: String,
    pub version: String,
}

/// BSpecData represents the complete BSpec specification data
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct BSpecData {
    pub metadata: Metadata,
    pub statistics: Statistics,
    pub domains: Vec<Domain>,
    pub document_types: Vec<DocumentType>,
    pub files: HashMap<String, File>,
    pub conformance_levels: Vec<ConformanceLevel>,
    pub yaml_schema: YAMLSchema,
    pub directory_structure: Vec<Vec<String>>,
    pub document_index: Vec<DocumentIndexEntry>,
}
