//! BSpec Rust SDK v1.0.0
//!
//! Rust SDK for the BSpec Universal Business Specification Standard.
//! Generated from BSpec JSON SDK at 2025-09-28T21:08:19.444990

pub mod types;
pub mod bspec;

pub use bspec::{BSpec, BSpecError};
pub use types::*;

/// BSpec SDK version
pub const VERSION: &str = "1.0.0";

/// Load BSpec from JSON bytes
pub fn load_from_json(json_data: &[u8]) -> Result<BSpec, BSpecError> {
    BSpec::from_json(json_data)
}

/// Load BSpec from JSON string
pub fn load_from_json_str(json_str: &str) -> Result<BSpec, BSpecError> {
    BSpec::from_json(json_str.as_bytes())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_version() {
        assert_eq!(VERSION, "1.0.0");
    }
}
