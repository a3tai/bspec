# BSpec Semantic Versioning Specification

## Overview

This document defines the semantic versioning strategy for the BSpec ecosystem, including the core specification, SDKs, applications, and all related components. BSpec follows [Semantic Versioning 2.0.0](https://semver.org/) with additional domain-specific conventions.

## Version Format

All BSpec components use the format: `MAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]`

### Version Components

- **MAJOR**: Incompatible API changes or breaking specification changes
- **MINOR**: Backward-compatible functionality additions
- **PATCH**: Backward-compatible bug fixes
- **PRERELEASE** (optional): alpha, beta, rc (release candidate)
- **BUILD** (optional): Build metadata (commit hash, build date)

### Examples
```
1.0.0          # Stable release
1.1.0          # Minor feature addition
1.0.1          # Bug fix
2.0.0          # Breaking changes
1.2.0-beta.1   # Beta prerelease
1.0.0+20250130 # Build metadata
```

## Component Versioning Strategy

### 1. Core Specification (`spec/v1/`)

**Current Version**: `1.0.0`
**File**: `/spec/v1/version.txt`

#### Version Increment Rules

**Major (X.0.0)**:
- Breaking changes to document type schemas
- Removal of existing document types
- Incompatible changes to YAML frontmatter structure
- Fundamental changes to BSpec format specification

**Minor (x.Y.0)**:
- New document types added
- New optional fields in existing document types
- Enhanced validation rules (non-breaking)
- New relationship types or patterns

**Patch (x.y.Z)**:
- Documentation improvements
- Example updates
- Clarification of existing requirements
- Bug fixes in specifications

#### Prerelease Conventions
- `X.Y.Z-alpha.N`: Early development versions
- `X.Y.Z-beta.N`: Feature-complete, testing phase
- `X.Y.Z-rc.N`: Release candidates, final testing

### 2. Generated SDKs (`sdk/v1/`)

**Versioning Philosophy**: SDKs follow specification versions but may have independent patch versions for implementation fixes.

#### TypeScript SDK (`sdk/v1/typescript/`)
**Current Version**: `1.0.0`
**File**: `package.json`

**Version Strategy**: 
- Major/Minor versions match specification
- Patch versions independent for TypeScript-specific fixes
- Example: Spec `1.2.0` → SDK `1.2.3` (after 3 TypeScript-specific patches)

#### Python SDK (`sdk/v1/python/`)
**Current Version**: `1.0.0`
**Files**: `setup.py`, `bspec/__init__.py`

**Version Strategy**:
- Major/Minor versions match specification
- Patch versions for Python-specific improvements
- Version in both `setup.py` and `__init__.py` must match

#### Go SDK (`sdk/v1/go/`)
**Current Version**: `1.0.0`
**File**: `bspec.go`

**Version Strategy**:
- Uses Go modules versioning
- Tags: `v1.0.0`, `v1.1.0`, etc.
- Major/Minor aligned with specification

#### Rust SDK (`sdk/v1/rust/`)
**Current Version**: `1.0.0`
**File**: `Cargo.toml`

**Version Strategy**:
- Semantic versioning in Cargo.toml
- Major/Minor aligned with specification
- Independent patch versions for Rust-specific fixes

#### JSON Schemas (`sdk/v1/json/`)
**Current Version**: `1.0.0`
**File**: `package.json`

**Version Strategy**:
- Versions match specification exactly
- No independent patch versions (schemas are generated)

### 3. CLI Tool (`sdk/cli/`)

**Current Version**: `1.0.0`
**Files**: `internal/commands/version.go`, `Makefile`

#### Version Management
- Version defined in `version.go` as constants
- Build process injects Git commit and build date
- Major/Minor versions independent of specification
- Focuses on CLI functionality evolution

**Version Increment Rules**:

**Major**: Breaking CLI interface changes, removed commands
**Minor**: New commands, new features, enhanced functionality  
**Patch**: Bug fixes, performance improvements, minor enhancements

### 4. Applications (`apps/`)

#### Web Application (`apps/web/`)
**Current Version**: `0.0.1`
**File**: `package.json`

**Version Strategy**:
- Independent versioning focused on user-facing features
- Major: UI/UX breaking changes, architecture overhauls
- Minor: New features, pages, functionality
- Patch: Bug fixes, performance improvements

#### MCP Server (`apps/mcp/`)
**Current Version**: `1.0.0`
**Files**: `package.json`, `version.txt`

**Version Strategy**:
- Aligned with BSpec specification for API compatibility
- Major/Minor: Changes affecting MCP protocol compatibility
- Patch: Server-specific improvements, bug fixes

#### Documentation Site (`apps/docs/`)
**Current Version**: `0.0.1`
**File**: `package.json`

**Version Strategy**:
- Independent versioning for documentation site
- Major: Site architecture changes
- Minor: New documentation sections, features
- Patch: Content updates, fixes

## Versioning Workflows

### 1. Specification Updates

#### Process
1. Update `/spec/v1/version.txt`
2. Update specification files
3. Run SDK generation: `python3 scripts/generate.py`
4. Update all affected SDK versions
5. Test all generated SDKs
6. Create Git tag: `spec-v1.X.Y`

#### Version Coordination
```bash
# Example: Releasing BSpec 1.2.0
echo "1.2.0" > spec/v1/version.txt

# Update SDK versions to match
# TypeScript
jq '.version = "1.2.0"' sdk/v1/typescript/package.json > tmp && mv tmp sdk/v1/typescript/package.json

# Python  
sed -i '' 's/version="[0-9.]*"/version="1.2.0"/' sdk/v1/python/setup.py

# JSON Schema
jq '.version = "1.2.0"' sdk/v1/json/package.json > tmp && mv tmp sdk/v1/json/package.json
```

### 2. SDK-Specific Updates

#### Independent Patch Versions
When fixing SDK-specific bugs without specification changes:

```bash
# Example: TypeScript SDK patch update
# Specification stays at 1.2.0
# TypeScript SDK goes from 1.2.0 → 1.2.1

jq '.version = "1.2.1"' sdk/v1/typescript/package.json > tmp && mv tmp sdk/v1/typescript/package.json
```

### 3. Application Updates

#### Independent Versioning
Applications version independently based on user-facing changes:

```bash
# Example: Web app feature release
jq '.version = "0.2.0"' apps/web/package.json > tmp && mv tmp apps/web/package.json

# Example: MCP server aligned with spec
jq '.version = "1.2.0"' apps/mcp/package.json > tmp && mv tmp apps/mcp/package.json
echo "1.2.0" > apps/mcp/version.txt
```

## Git Tagging Strategy

### Tag Format
- Specification: `spec-v{version}` (e.g., `spec-v1.2.0`)
- CLI: `cli-v{version}` (e.g., `cli-v1.1.0`)
- Web App: `web-v{version}` (e.g., `web-v0.2.0`)
- MCP Server: `mcp-v{version}` (e.g., `mcp-v1.2.0`)
- SDKs: `{language}-v{version}` (e.g., `typescript-v1.2.1`, `python-v1.2.0`)

### Release Tags
```bash
# Specification release
git tag -a spec-v1.2.0 -m "BSpec Specification v1.2.0"

# CLI release  
git tag -a cli-v1.1.0 -m "BSpec CLI v1.1.0"

# SDK releases
git tag -a typescript-v1.2.1 -m "TypeScript SDK v1.2.1"
git tag -a python-v1.2.0 -m "Python SDK v1.2.0"
```

## Version Synchronization

### Automated Version Management

Create a version management script:

```bash
# scripts/version-sync.py
```

### Version Dependencies

#### Critical Alignments
- **Specification ↔ JSON Schema**: Always synchronized
- **Specification ↔ MCP Server**: Major/Minor aligned
- **SDK Major/Minor ↔ Specification**: Aligned for compatibility

#### Independent Versioning
- **CLI**: Independent evolution
- **Web App**: Independent user-facing versioning
- **SDK Patches**: Language-specific improvements

## Breaking Change Policy

### Major Version Changes (X.0.0)

#### Specification Breaking Changes
- Document type schema modifications
- YAML frontmatter structure changes
- Relationship model changes
- Validation rule breaking changes

#### SDK Breaking Changes
- API interface modifications
- Method signature changes
- Class/struct name changes
- Required parameter additions

### Deprecation Process

#### 1. Deprecation Notice (Minor Release)
- Mark features as deprecated
- Add warnings to documentation
- Provide migration guidance

#### 2. Deprecation Period
- Minimum 6 months for specification
- Minimum 3 months for SDKs
- Clear migration timeline

#### 3. Removal (Major Release)
- Remove deprecated features
- Update all documentation
- Provide migration tools

## Release Process

### 1. Pre-Release Checklist
- [ ] All tests pass
- [ ] Documentation updated
- [ ] Version numbers synchronized
- [ ] Breaking changes documented
- [ ] Migration guides prepared

### 2. Release Steps
1. Update version files
2. Generate updated SDKs
3. Test all components
4. Update CHANGELOG.md
5. Create Git tags
6. Build and publish packages
7. Deploy applications
8. Update documentation

### 3. Post-Release
- [ ] Verify deployments
- [ ] Test published packages
- [ ] Update project documentation
- [ ] Announce release

## Backward Compatibility

### Specification Compatibility
- **Major versions**: Breaking changes allowed
- **Minor versions**: Only additions, no removals
- **Patch versions**: No functionality changes

### SDK Compatibility
- **Major versions**: Breaking API changes allowed
- **Minor versions**: New features, backward compatible
- **Patch versions**: Bug fixes only

### Data Compatibility
- **Document format**: Backward compatible within major versions
- **YAML structure**: Strict compatibility within major versions
- **Validation**: Enhanced rules okay, stricter rules require major version

## Version Automation

### Build-Time Version Injection

#### CLI Tool
```makefile
# Makefile
VERSION := $(shell cat version.txt)
BUILD_DATE := $(shell date -u +"%Y-%m-%d %H:%M:%S UTC")
GIT_COMMIT := $(shell git rev-parse --short HEAD)

build:
    go build -ldflags "-X 'main.Version=$(VERSION)' -X 'main.BuildDate=$(BUILD_DATE)' -X 'main.GitCommit=$(GIT_COMMIT)'"
```

#### SDKs
```python
# scripts/update-versions.py
def sync_versions():
    spec_version = read_spec_version()
    update_sdk_versions(spec_version)
    update_package_files(spec_version)
```

### CI/CD Integration

#### Version Validation
```yaml
# .github/workflows/version-check.yml
- name: Validate Version Consistency
  run: python scripts/validate-versions.py
```

#### Automated Releases
```yaml
# .github/workflows/release.yml
- name: Create Release Tags
  run: |
    git tag -a spec-v${{ env.SPEC_VERSION }}
    git tag -a typescript-v${{ env.TS_VERSION }}
    git push --tags
```

## Documentation Requirements

### Version Documentation
- **CHANGELOG.md**: All version changes documented
- **MIGRATION.md**: Breaking change migration guides  
- **README.md**: Current version information

### API Documentation
- **SDK docs**: Version-specific API documentation
- **Specification**: Version history and evolution

## Monitoring and Metrics

### Version Tracking
- Usage statistics by version
- Migration completion rates
- Support ticket correlation with versions

### Health Metrics
- Build success rates by version
- Test pass rates across versions
- Deployment success tracking

---

**Document Version**: 1.0.0  
**Last Updated**: 2025-01-30  
**Next Review**: 2025-04-30