# Releasing BSpec

This document describes how to cut a new release of BSpec.

## Quick Start

```bash
# From the repository root
./scripts/release.sh
```

This interactive script will guide you through the release process.

## Release Types

BSpec follows [Semantic Versioning](https://semver.org/):

- **Patch** (1.0.0 → 1.0.1): Bug fixes, documentation updates, minor changes
- **Minor** (1.0.0 → 1.1.0): New features, backward-compatible changes
- **Major** (1.0.0 → 2.0.0): Breaking changes, incompatible API changes

## Automated Release Process

### Using the Release Script

The recommended way to cut a release:

```bash
cd /path/to/bspec
./scripts/release.sh
```

The script will:
1. Check git status and branch
2. Show recent changes
3. Prompt for release type (patch/minor/major)
4. Run pre-release checks
5. Trigger the GitHub Actions release workflow

### Using GitHub Actions Directly

You can also trigger releases from the GitHub web interface:

1. Go to [Actions → Release BSpec](https://github.com/a3tai/bspec/actions/workflows/release.yml)
2. Click "Run workflow"
3. Select:
   - **version_bump**: patch, minor, or major
   - **create_release**: true (to create GitHub release)
   - **publish_packages**: false (or true if publishing to npm/PyPI)
4. Click "Run workflow"

## What Happens During a Release

When you trigger a release, the automation:

### 1. Version Bump
- Updates `spec/v1/version.txt`
- Updates all SDK package.json, pyproject.toml, etc.
- Commits changes with message: `bump: version X.Y.Z`
- Creates git tag: `vX.Y.Z`
- Pushes to GitHub

### 2. Build & Test
- Generates SDKs from specification
- Builds CLI binaries for all platforms:
  - macOS (Intel & Apple Silicon)
  - Linux (amd64 & arm64)
  - Windows (amd64)
- Runs tests on all SDKs
- Validates all components

### 3. Create Release
- Creates GitHub release with changelog
- Uploads CLI binaries as release assets
- Packages SDK source code
- Generates release notes

### 4. Deploy Website
- Regenerates documentation from spec
- Builds static site
- Deploys to Cloudflare Pages at bspec.dev

### 5. Publish Packages (Optional)
If `publish_packages` is enabled:
- TypeScript SDK → npm (@bspec/sdk)
- Python SDK → PyPI (bspec-sdk)
- Go SDK → Available via go get

## Manual Release Process

If you need to release manually:

### 1. Prepare the Release

```bash
# Ensure you're on main and up to date
git checkout main
git pull origin main

# Ensure working directory is clean
git status

# Check current version
cat spec/v1/version.txt
```

### 2. Bump Version

```bash
# For patch release
python3 scripts/bump-version.py patch

# For minor release
python3 scripts/bump-version.py minor

# For major release
python3 scripts/bump-version.py major
```

### 3. Review Changes

```bash
# See what was updated
git diff

# Check new version
cat spec/v1/version.txt
```

### 4. Commit and Tag

```bash
# Get the new version
NEW_VERSION=$(cat spec/v1/version.txt)

# Commit
git add -A
git commit -m "bump: version $NEW_VERSION"

# Tag
git tag -a "v$NEW_VERSION" -m "Release v$NEW_VERSION"

# Push
git push origin main
git push origin --tags
```

### 5. Build Release Assets

```bash
# Generate SDKs
python3 scripts/generate.py --clean

# Build CLI (macOS example)
cd sdk/cli
make build-release GOOS=darwin GOARCH=arm64
cd ../..

# Create TGZ package
python3 scripts/create_tgz.py
```

### 6. Create GitHub Release

```bash
# Using GitHub CLI
gh release create v$NEW_VERSION \
  --title "BSpec v$NEW_VERSION" \
  --generate-notes \
  ./bspec-v1-0-*.tgz \
  ./sdk/cli/bin/bspec-*
```

### 7. Deploy Website

```bash
cd apps/web

# Regenerate docs
python3 ../../scripts/generate_web_docs.py

# Build
bun install
bun run build

# Deploy
bunx wrangler pages deploy .vitepress/dist --project-name=bspec
```

## Pre-Release Checklist

Before cutting a release, ensure:

- [ ] All tests pass
- [ ] Documentation is up to date
- [ ] CHANGELOG reflects recent changes
- [ ] No known critical bugs
- [ ] All PRs for the release are merged
- [ ] Main branch is stable
- [ ] You've tested locally

## Post-Release Checklist

After cutting a release:

- [ ] Verify GitHub release was created
- [ ] Check that CLI binaries are attached
- [ ] Verify website deployed successfully
- [ ] Test downloading and using CLI
- [ ] Announce release (if major/minor)
- [ ] Update any dependent projects

## Troubleshooting

### Release workflow failed

Check the [Actions tab](https://github.com/a3tai/bspec/actions) for error details:

- **Version bump failed**: Check if version.txt exists and is valid
- **Build failed**: Check build logs for compilation errors
- **Push failed**: Ensure you have push permissions
- **Deploy failed**: Check Cloudflare secrets are set

### Version was bumped but release failed

If the version was bumped but release creation failed:

```bash
# Get the latest version
NEW_VERSION=$(cat spec/v1/version.txt)

# Manually create the release
gh release create v$NEW_VERSION --title "BSpec v$NEW_VERSION" --generate-notes
```

### Need to rollback a release

```bash
# Delete the tag
git tag -d vX.Y.Z
git push origin :refs/tags/vX.Y.Z

# Delete the release (GitHub CLI)
gh release delete vX.Y.Z

# Reset version
python3 scripts/bump-version.py --set X.Y.Z
git add -A
git commit -m "revert: rollback to version X.Y.Z"
git push origin main
```

## Release Cadence

- **Patch releases**: As needed for bug fixes (weekly/bi-weekly)
- **Minor releases**: Monthly or when new features are ready
- **Major releases**: Quarterly or when breaking changes accumulate

## Versioning Guidelines

### Patch Release (X.Y.Z)

Increment patch version for:
- Bug fixes
- Documentation updates
- Performance improvements (no API changes)
- Internal refactoring
- Dependency updates

### Minor Release (X.Y.0)

Increment minor version for:
- New document types added
- New SDK features (backward compatible)
- New CLI commands
- Deprecations (with backward compatibility)
- Significant documentation additions

### Major Release (X.0.0)

Increment major version for:
- Breaking changes to document structure
- Removed document types
- Breaking API changes in SDKs
- Major specification restructuring
- Incompatible CLI changes

## GitHub Secrets Required

For full automation, set these secrets in GitHub:

- `CLOUDFLARE_API_TOKEN`: For deploying website
- `CLOUDFLARE_ACCOUNT_ID`: Your Cloudflare account ID
- `NPM_TOKEN`: For publishing to npm (optional)
- `PYPI_TOKEN`: For publishing to PyPI (optional)

## Support

If you encounter issues with releases:

1. Check [GitHub Actions logs](https://github.com/a3tai/bspec/actions)
2. Review this documentation
3. Ask in [GitHub Discussions](https://github.com/a3tai/bspec/discussions)
4. Open an [issue](https://github.com/a3tai/bspec/issues) if needed

## Related Documentation

- [Contributing Guide](./CONTRIBUTING.md)
- [Deployment Guide](./apps/web/DEPLOY.md)
- [GitHub Actions Workflows](./.github/workflows/README.md)
