# BSpec Automation Quick Reference

## Release a New Version

```bash
./scripts/release.sh
```

Or via GitHub UI: [Actions → Release BSpec](https://github.com/a3tai/bspec/actions/workflows/release.yml)

## Deploy Website

Website deploys automatically on push to `main` when files in these paths change:
- `apps/web/**`
- `spec/v1/**`
- `scripts/generate_web_docs.py`

Manual deployment:
```bash
cd apps/web
python3 ../../scripts/generate_web_docs.py  # Regenerate docs
bun run build                                 # Build site
bunx wrangler pages deploy .vitepress/dist --project-name=bspec
```

Or trigger via GitHub: [Actions → Deploy](https://github.com/a3tai/bspec/actions/workflows/deploy.yml)

## Generate SDKs

```bash
python3 scripts/generate.py --clean
```

This generates:
- TypeScript SDK → `sdk/v1/typescript/`
- Python SDK → `sdk/v1/python/`
- Go SDK → `sdk/v1/go/`
- JSON Schemas → `sdk/v1/json/`

## Build CLI

```bash
cd sdk/cli
make build          # Build for current platform
make build-all      # Build for all platforms
```

Binaries go to `sdk/cli/bin/`

## Run Tests

```bash
# All tests
make test

# Specific SDK tests
cd sdk/v1/typescript && bun test
cd sdk/v1/python && pytest
cd sdk/cli && make test
```

## CI/CD Workflows

### Automatic Triggers

| Workflow | Trigger | What It Does |
|----------|---------|--------------|
| **Deploy** | Push to `main` (web/spec changes) | Regenerates docs, builds site, deploys to Cloudflare |
| **CI** | Pull requests, push to `main` | Runs tests, linters, builds |
| **Release** | Manual trigger | Bumps version, builds assets, creates GitHub release |

### Manual Triggers

All workflows can be manually triggered from [GitHub Actions](https://github.com/a3tai/bspec/actions):

1. **Release BSpec** - Cut a new release
2. **Deploy to Cloudflare Pages** - Deploy website
3. **CI** - Run full test suite

## Version Management

### Check Current Version
```bash
cat spec/v1/version.txt
```

### Bump Version Manually
```bash
python3 scripts/bump-version.py patch   # 1.0.0 → 1.0.1
python3 scripts/bump-version.py minor   # 1.0.0 → 1.1.0
python3 scripts/bump-version.py major   # 1.0.0 → 2.0.0
```

### Show Current Version Info
```bash
python3 scripts/bump-version.py --show
```

## Common Tasks

### Add a New Document Type

1. Create spec file: `spec/v1/{domain}/{TYPE}-spec.md`
2. Generate SDKs: `python3 scripts/generate.py`
3. Regenerate web docs: `python3 scripts/generate_web_docs.py`
4. Test locally: `cd apps/web && bun run dev`
5. Commit and push

The website will auto-deploy!

### Update Website Content

1. Edit files in `apps/web/`
2. Or edit spec files in `spec/v1/` (will regenerate docs)
3. Test locally: `cd apps/web && bun run dev`
4. Commit and push to `main`

Deployment happens automatically!

### Update SDK Logic

1. Edit generator: `scripts/generators/{language}_generator.py`
2. Regenerate: `python3 scripts/generate.py --clean`
3. Test changes in the SDK
4. Commit both generator and generated code

## GitHub Secrets

Required secrets in repository settings:

- `CLOUDFLARE_API_TOKEN` - For website deployment
- `CLOUDFLARE_ACCOUNT_ID` - Your Cloudflare account ID
- `NPM_TOKEN` - For npm publishing (optional)
- `PYPI_TOKEN` - For PyPI publishing (optional)

## Useful Commands

```bash
# See what's changed since last release
git log $(git describe --tags --abbrev=0)..HEAD --oneline

# List all releases
gh release list

# View latest release
gh release view

# Check workflow status
gh run list --workflow=release.yml

# Watch a workflow run
gh run watch

# Build and preview website locally
cd apps/web
bun install
bun run dev     # http://localhost:5173

# Build for production
bun run build
bun run preview # http://localhost:4173
```

## Troubleshooting

### Website not updating
- Check [deploy workflow](https://github.com/a3tai/bspec/actions/workflows/deploy.yml)
- Verify Cloudflare secrets are set
- Try manual deployment: `cd apps/web && bunx wrangler pages deploy .vitepress/dist --project-name=bspec`

### Release failed
- Check [release workflow logs](https://github.com/a3tai/bspec/actions/workflows/release.yml)
- Verify version.txt exists: `cat spec/v1/version.txt`
- Check git status: `git status`

### SDK generation fails
- Check parser: `python3 scripts/test_generator.py`
- Verify spec files are valid YAML + Markdown
- Check generator logs

### Build errors
- Clear node_modules: `cd apps/web && rm -rf node_modules && bun install`
- Clear CLI cache: `cd sdk/cli && make clean && make build`
- Update dependencies: `bun update` or `pip install --upgrade -r requirements.txt`

## Quick Links

- **Repository**: https://github.com/a3tai/bspec
- **Website**: https://bspec.dev
- **Actions**: https://github.com/a3tai/bspec/actions
- **Releases**: https://github.com/a3tai/bspec/releases
- **Issues**: https://github.com/a3tai/bspec/issues

## Related Docs

- [RELEASING.md](./RELEASING.md) - Detailed release process
- [CONTRIBUTING.md](./CONTRIBUTING.md) - Contribution guidelines
- [apps/web/DEPLOY.md](./apps/web/DEPLOY.md) - Website deployment
- [.github/workflows/README.md](./.github/workflows/README.md) - CI/CD details
