//! C-compatible interface for CLI and native app integration
//!
//! This module provides C-compatible functions that can be called from
//! Go CLI tools, native macOS apps, and other languages.
//!
//! **GENERATED CODE - DO NOT EDIT MANUALLY**
//! Generated at: 2025-09-28T16:00:48.188037

use std::ffi::{CStr, CString};
use std::os::raw::{c_char, c_int, c_void};
use std::ptr;
use std::collections::HashMap;

use crate::archive::BSpecArchive;
use crate::base::{BSpecDocument, DocumentType, DocumentStatus};
use crate::error::{BSpecError, Result};

/// Opaque pointer to BSpecArchive for C interface
pub struct Archive {
    inner: BSpecArchive,
}

/// Opaque pointer to query results for C interface
pub struct QueryResult {
    documents: Vec<String>, // JSON representations
    current_index: usize,
}

/// Error codes for C interface
#[repr(C)]
pub enum BSpecErrorCode {
    Success = 0,
    InvalidPath = 1,
    FileNotFound = 2,
    InvalidArchive = 3,
    InvalidQuery = 4,
    SerializationError = 5,
    ValidationError = 6,
    UnknownError = 99,
}

impl From<BSpecError> for BSpecErrorCode {
    fn from(error: BSpecError) -> Self {
        match error {
            BSpecError::Io(_) => BSpecErrorCode::FileNotFound,
            BSpecError::InvalidArchive(_) => BSpecErrorCode::InvalidArchive,
            BSpecError::InvalidDocumentType(_) => BSpecErrorCode::InvalidQuery,
            BSpecError::InvalidDocumentStatus(_) => BSpecErrorCode::InvalidQuery,
            BSpecError::ValidationError(_) => BSpecErrorCode::ValidationError,
            BSpecError::Serialization(_) => BSpecErrorCode::SerializationError,
            _ => BSpecErrorCode::UnknownError,
        }
    }
}

/// Open a .bspec archive file
///
/// # Safety
/// The path must be a valid C string pointer
#[no_mangle]
pub unsafe extern "C" fn bspec_open_archive(path: *const c_char) -> *mut Archive {
    if path.is_null() {
        return ptr::null_mut();
    }

    let c_str = match CStr::from_ptr(path).to_str() {
        Ok(s) => s,
        Err(_) => return ptr::null_mut(),
    };

    match BSpecArchive::load_from_file(c_str) {
        Ok(archive) => {
            let boxed = Box::new(Archive { inner: archive });
            Box::into_raw(boxed)
        }
        Err(_) => ptr::null_mut(),
    }
}

/// Close and free a BSpec archive
///
/// # Safety
/// The archive pointer must be valid and not used after this call
#[no_mangle]
pub unsafe extern "C" fn bspec_close_archive(archive: *mut Archive) {
    if !archive.is_null() {
        let _ = Box::from_raw(archive);
    }
}

/// Create a new empty archive
///
/// # Safety
/// The name must be a valid C string pointer
#[no_mangle]
pub unsafe extern "C" fn bspec_create_archive(name: *const c_char) -> *mut Archive {
    if name.is_null() {
        return ptr::null_mut();
    }

    let c_str = match CStr::from_ptr(name).to_str() {
        Ok(s) => s,
        Err(_) => return ptr::null_mut(),
    };

    let archive = BSpecArchive::new(c_str.to_string());
    let boxed = Box::new(Archive { inner: archive });
    Box::into_raw(boxed)
}

/// Save archive to file
///
/// # Safety
/// Both pointers must be valid
#[no_mangle]
pub unsafe extern "C" fn bspec_save_archive(
    archive: *mut Archive,
    path: *const c_char,
) -> BSpecErrorCode {
    if archive.is_null() || path.is_null() {
        return BSpecErrorCode::InvalidPath;
    }

    let archive_ref = &mut (*archive).inner;
    let c_str = match CStr::from_ptr(path).to_str() {
        Ok(s) => s,
        Err(_) => return BSpecErrorCode::InvalidPath,
    };

    match archive_ref.save_to_file(c_str) {
        Ok(()) => BSpecErrorCode::Success,
        Err(e) => e.into(),
    }
}

/// Query documents in archive
///
/// # Safety
/// Both pointers must be valid
#[no_mangle]
pub unsafe extern "C" fn bspec_query_documents(
    archive: *mut Archive,
    query_json: *const c_char,
) -> *mut QueryResult {
    if archive.is_null() || query_json.is_null() {
        return ptr::null_mut();
    }

    let archive_ref = &(*archive).inner;
    let query_str = match CStr::from_ptr(query_json).to_str() {
        Ok(s) => s,
        Err(_) => return ptr::null_mut(),
    };

    // Parse query JSON
    let query: serde_json::Value = match serde_json::from_str(query_str) {
        Ok(q) => q,
        Err(_) => return ptr::null_mut(),
    };

    let documents = execute_query(archive_ref, &query);

    let result = QueryResult {
        documents,
        current_index: 0,
    };

    Box::into_raw(Box::new(result))
}

/// Get the number of results in a query
///
/// # Safety
/// The result pointer must be valid
#[no_mangle]
pub unsafe extern "C" fn bspec_query_result_count(result: *mut QueryResult) -> c_int {
    if result.is_null() {
        return 0;
    }

    (*result).documents.len() as c_int
}

/// Get a document JSON from query results by index
///
/// # Safety
/// The result pointer must be valid and index must be within bounds
#[no_mangle]
pub unsafe extern "C" fn bspec_get_document_json(
    result: *mut QueryResult,
    index: usize,
) -> *const c_char {
    if result.is_null() {
        return ptr::null();
    }

    let result_ref = &(*result);
    if index >= result_ref.documents.len() {
        return ptr::null();
    }

    // Create a C string that will be owned by the caller
    match CString::new(result_ref.documents[index].clone()) {
        Ok(c_string) => c_string.into_raw(),
        Err(_) => ptr::null(),
    }
}

/// Free a string returned by bspec_get_document_json
///
/// # Safety
/// The string must have been returned by bspec_get_document_json
#[no_mangle]
pub unsafe extern "C" fn bspec_free_string(s: *mut c_char) {
    if !s.is_null() {
        let _ = CString::from_raw(s);
    }
}

/// Free query results
///
/// # Safety
/// The result pointer must be valid and not used after this call
#[no_mangle]
pub unsafe extern "C" fn bspec_free_query_result(result: *mut QueryResult) {
    if !result.is_null() {
        let _ = Box::from_raw(result);
    }
}

/// Get archive statistics as JSON
///
/// # Safety
/// The archive pointer must be valid
#[no_mangle]
pub unsafe extern "C" fn bspec_get_archive_stats(archive: *mut Archive) -> *const c_char {
    if archive.is_null() {
        return ptr::null();
    }

    let archive_ref = &(*archive).inner;

    let stats = serde_json::json!({
        "name": archive_ref.manifest.name,
        "total_documents": archive_ref.manifest.total_documents,
        "document_types": archive_ref.manifest.document_types,
        "domains": archive_ref.manifest.domains,
        "conformance_level": archive_ref.manifest.conformance_level,
        "industry_profile": archive_ref.manifest.industry_profile,
        "created_at": archive_ref.manifest.created_at,
        "updated_at": archive_ref.manifest.updated_at
    });

    match CString::new(stats.to_string()) {
        Ok(c_string) => c_string.into_raw(),
        Err(_) => ptr::null(),
    }
}

/// Execute a query against the archive
fn execute_query(archive: &BSpecArchive, query: &serde_json::Value) -> Vec<String> {
    let mut results = Vec::new();

    // Get all documents first
    let all_docs = archive.get_all_documents();

    for doc in all_docs {
        let mut matches = true;

        // Check type filter
        if let Some(doc_type) = query.get("type").and_then(|v| v.as_str()) {
            if doc.document_type().as_str() != doc_type {
                matches = false;
            }
        }

        // Check status filter
        if let Some(status) = query.get("status").and_then(|v| v.as_str()) {
            if doc.metadata().status.to_string().to_lowercase() != status.to_lowercase() {
                matches = false;
            }
        }

        // Check domain filter
        if let Some(domain) = query.get("domain").and_then(|v| v.as_str()) {
            if doc.document_type().domain().as_str() != domain {
                matches = false;
            }
        }

        // Check owner filter
        if let Some(owner) = query.get("owner").and_then(|v| v.as_str()) {
            if !doc.metadata().owner.contains(owner) {
                matches = false;
            }
        }

        // Check content search
        if let Some(search) = query.get("search").and_then(|v| v.as_str()) {
            let search_lower = search.to_lowercase();
            if !doc.content().to_lowercase().contains(&search_lower)
                && !doc.metadata().title.to_lowercase().contains(&search_lower)
            {
                matches = false;
            }
        }

        if matches {
            if let Ok(json) = doc.to_json() {
                results.push(json);
            }
        }
    }

    // Apply limit if specified
    if let Some(limit) = query.get("limit").and_then(|v| v.as_u64()) {
        results.truncate(limit as usize);
    }

    results
}

/// Rust-side utilities for working with the C interface
impl Archive {
    /// Get a reference to the inner archive
    pub fn inner(&self) -> &BSpecArchive {
        &self.inner
    }

    /// Get a mutable reference to the inner archive
    pub fn inner_mut(&mut self) -> &mut BSpecArchive {
        &mut self.inner
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::ffi::CString;
    use tempfile::NamedTempFile;

    #[test]
    fn test_c_interface_basic() {
        unsafe {
            // Create archive
            let name = CString::new("Test Archive").unwrap();
            let archive = bspec_create_archive(name.as_ptr());
            assert!(!archive.is_null());

            // Get stats
            let stats_ptr = bspec_get_archive_stats(archive);
            assert!(!stats_ptr.is_null());

            let stats_str = CStr::from_ptr(stats_ptr).to_str().unwrap();
            assert!(stats_str.contains("Test Archive"));

            // Clean up
            bspec_free_string(stats_ptr as *mut c_char);
            bspec_close_archive(archive);
        }
    }

    #[test]
    fn test_query_interface() {
        unsafe {
            // Create archive with test data
            let name = CString::new("Test").unwrap();
            let archive = bspec_create_archive(name.as_ptr());

            // Query (empty archive)
            let query = CString::new(r#"{"type": "MSN"}"#).unwrap();
            let result = bspec_query_documents(archive, query.as_ptr());
            assert!(!result.is_null());

            let count = bspec_query_result_count(result);
            assert_eq!(count, 0);

            // Clean up
            bspec_free_query_result(result);
            bspec_close_archive(archive);
        }
    }
}