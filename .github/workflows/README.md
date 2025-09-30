# BSpec GitHub Actions Workflows

This directory contains three comprehensive workflows for the BSpec project:

## 🧪 test-version-bump.yml

**Purpose**: Comprehensive testing of the version bump script
**Triggers**: 
- Push/PR changes to `scripts/bump-version.py` or the workflow itself
- Manual dispatch with version bump type selection

**What it tests**:
- ✅ Script help and version display functionality
- ✅ All expected files exist in repository
- ✅ Patch, minor, and major version bumps work correctly
- ✅ All SDK files are updated with new versions
- ✅ All SDKs still build after version changes
- ✅ Error handling for invalid inputs and missing files
- ✅ Manual testing via workflow dispatch

## 🚀 release.yml

**Purpose**: Production release workflow using the version bump script
**Trigger**: Manual dispatch only (workflow_dispatch)

**What it does**:
1. **Version Management**: Bumps version across all components
2. **Git Operations**: Commits changes, creates tags, pushes to repository
3. **Release Creation**: Creates GitHub release with changelog
4. **Asset Generation**: Packages source code and all SDKs
5. **Build Validation**: Tests all SDKs build correctly with new version
6. **Package Publishing**: Ready for NPM/PyPI (disabled for safety)

**Options**:
- Version bump type (patch/minor/major)
- Create GitHub release (true/false)
- Publish packages (true/false - disabled by default)

## ⚡ ci.yml

**Purpose**: Continuous integration for all code changes
**Triggers**: Push/PR to main or develop branches

**What it validates**:
- 🏗️ **Repository Structure**: Required files and directories exist
- 🧪 **Version Script**: Functionality and logic testing
- 🔨 **Multi-Language Builds**: TypeScript, Python, Go, Rust, CLI
- 📝 **Code Quality**: Linting, formatting, syntax validation
- 🔒 **Security**: Checks for hardcoded secrets and best practices
- 📊 **Summary**: Complete CI results overview

## Usage Examples

### Run Version Bump Tests
Go to the Actions tab in GitHub and manually trigger "Test Version Bump Script" workflow with different bump types.

### Create a Release
1. Go to Actions tab
2. Select "Release BSpec" workflow  
3. Click "Run workflow"
4. Choose version bump type and options
5. Monitor the multi-job release process

### Monitor CI
CI runs automatically on all pushes and pull requests. Check the Actions tab for results.

## Workflow Features

### 🔄 **Version Synchronization**
All workflows use the centralized version bump script to ensure consistent versioning across:
- Python SDK (`bspec/__init__.py`, `setup.py`)
- TypeScript SDK (`package.json`)
- Go SDK (`bspec.go`) and CLI (`version.go`)
- Rust SDK (`Cargo.toml`)
- Web apps (`package.json` files)

### 🏗️ **Multi-Language Support**
Full environment setup for all BSpec languages:
- Node.js 20 with bun (following BSpec package management rules)
- Python 3.11
- Go 1.21
- Rust (stable toolchain)

### 🛡️ **Safety Features**
- Package publishing disabled by default (requires secrets)
- Comprehensive testing before any releases
- Error handling and rollback capabilities
- Security scanning for hardcoded secrets

### 📋 **Matrix Builds**
CI workflow uses matrix strategy to test all components in parallel:
- TypeScript SDK
- Python SDK  
- Go SDK
- Rust SDK
- Web application
- CLI tool

## Development Workflow

1. **Development**: Make changes, push to branch
2. **CI Validation**: Automatic testing via ci.yml
3. **Version Bump Testing**: Manual testing via test-version-bump.yml
4. **Release**: Production release via release.yml
5. **Post-Release**: Automated asset uploads and notifications

## Next Steps

- [ ] Add NPM_TOKEN secret for automatic NPM publishing
- [ ] Add PYPI_TOKEN secret for automatic PyPI publishing  
- [ ] Set up release notifications (Slack, email, etc.)
- [ ] Add performance benchmarking to CI
- [ ] Implement automated changelog generation
- [ ] Add deployment workflows for web applications

## Maintenance

These workflows are designed to be self-maintaining, but should be reviewed periodically for:
- GitHub Actions version updates
- Language version updates (Node.js, Python, Go, Rust)
- New SDK additions
- Security best practices updates