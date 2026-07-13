# BSpec SDKs

Generated SDKs for BSpec v1.1.2 Universal Business Specification Standard.

## Available SDKs

### 📄 JSON SDK
- **Path**: `sdk/v1/json/`
- **Main File**: `bspec-v1-1-2.tgz`
- **Package**: `@bspec/json-sdk`
- **Use Case**: Universal data format for any language
- **Archive**: Complete JSON SDK package

### 🟦 TypeScript SDK
- **Path**: `sdk/v1/typescript/`
- **Main File**: `index.ts`
- **Package**: `@bspec/typescript-sdk`
- **Use Case**: Type-safe TypeScript/JavaScript development
- **Source**: Generated from JSON SDK

### 🐍 Python SDK
- **Path**: `sdk/v1/python/`
- **Main File**: `bspec/__init__.py`
- **Package**: `bspec`
- **Use Case**: Python applications and data analysis
- **Source**: Generated from JSON SDK

### 🐹 Go SDK
- **Path**: `sdk/v1/go/`
- **Main File**: `bspec.go`
- **Package**: `github.com/bspec-foundation/bspec-go`
- **Use Case**: Go applications and microservices
- **Source**: Generated from JSON SDK

### 🦀 Rust SDK
- **Path**: `sdk/v1/rust/`
- **Main File**: `src/lib.rs`
- **Package**: `bspec`
- **Use Case**: Rust applications and systems programming
- **Source**: Generated from JSON SDK

## Generation Architecture

1. **JSON SDK (Foundation)**: Converts entire `spec/v1/` directory to comprehensive JSON
2. **Language SDKs**: Generated from JSON SDK for language-specific interfaces
3. **Version Tracking**: All SDKs include `version.txt` with BSpec version
4. **Package Management**: Native package files for each ecosystem

## Version Information

Each SDK includes:
- `version.txt` - BSpec version number (1.1.2)
- `README.md` - Usage documentation
- Package manifest for language ecosystem
- Complete specification data in appropriate format

## Statistics

- **BSpec Version**: 1.1.2
- **Total Files**: 116
- **Document Types**: 112
- **Business Domains**: 11
- **Generated**: 2026-07-12T17:47:25.391781

## Usage

All SDKs are generated from the same source truth (JSON SDK), ensuring consistency across languages while providing native interfaces and idioms for each ecosystem.

## License

MIT License - see the repository LICENSE file for details.
