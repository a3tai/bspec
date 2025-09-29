# BSpec Rust SDK

Rust SDK for the BSpec v1.0.0 Universal Business Specification Standard.

## Installation

Add to your `Cargo.toml`:

```toml
[dependencies]
bspec = "1.0.0"
```

## Usage

### Basic Usage

```rust
use bspec::{BSpec, load_from_json_str};
use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Load BSpec from JSON file
    let json_content = fs::read_to_string("bspec-complete.json")?;
    let bspec = load_from_json_str(&json_content)?;

    // Get basic information
    println!("BSpec v{}", bspec.version());
    println!("Total domains: {}", bspec.statistics().total_domains);
    println!("Total document types: {}", bspec.statistics().total_document_types);

    Ok(())
}
```

### Working with Domains

```rust
// Get specific domain
if let Some(strategic_foundation) = bspec.get_domain("Strategic Foundation") {
    println!("Domain: {} - {}", strategic_foundation.display_name, strategic_foundation.description);
}

// Get document types for a domain
let strategic_doc_types = bspec.get_document_types_for_domain("Strategic Foundation");
for doc_type in strategic_doc_types {
    println!("{}: {} - {}", doc_type.code, doc_type.name, doc_type.purpose);
}
```

### Working with Document Types

```rust
// Get specific document type
if let Some(mission_type) = bspec.get_document_type("MSN") {
    println!("{}: {}", mission_type.name, mission_type.purpose);
}

// Search document types
let customer_types = bspec.search_document_types("customer");
println!("Found {} customer-related types", customer_types.len());
```

### Working with Files

```rust
// Get specific file
if let Some(spec_file) = bspec.get_file("spec.md") {
    println!("File content length: {}", spec_file.content.len());
}

// Get files by type
let markdown_files = bspec.get_files_by_type("markdown");
println!("Found {} markdown files", markdown_files.len());

// Get document files (markdown with frontmatter)
let document_files = bspec.get_document_files();
println!("Found {} document files", document_files.len());
```

### Serialization

```rust
// Convert back to JSON
let json_string = bspec.to_json()?;
let json_bytes = bspec.to_json_bytes()?;

// Save to file
fs::write("output.json", json_bytes)?;
```

## Generated Information

- **From**: BSpec v1.0.0 JSON SDK
- **Generated**: 2025-09-28T21:08:19.445400
- **Generator**: rust-generator-v1.0.0
- **Total Files**: 112
- **Total Document Types**: 112
- **Total Domains**: 12

## License

CC BY 4.0 - See [BSpec Foundation](https://bspec.dev) for details.
