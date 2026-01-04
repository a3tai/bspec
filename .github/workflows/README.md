# BSpec GitHub Actions Workflows

This directory contains GitHub Actions workflows for automating BSpec development, testing, and deployment.

## Workflows

| Workflow | Trigger | Purpose |
|----------|---------|---------|
| `ci.yml` | Push/PR to main | Continuous integration - builds and tests all components |
| `deploy.yml` | Push to main (apps/web changes) | Deploys website to Cloudflare Pages |
| `deploy-mcp.yml` | Push to main (apps/mcp changes) | Deploys MCP server to Cloudflare Workers |
| `release.yml` | Manual dispatch | Full release process with version bump, CLI builds, and deployments |

## Required Secrets

Before the workflows can run, you must configure these secrets in your GitHub repository:

### Setting Up Secrets

1. Go to your repository on GitHub
2. Navigate to **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret** for each secret below

### Required Secrets

| Secret | Description | How to Get |
|--------|-------------|------------|
| `CLOUDFLARE_API_TOKEN` | Cloudflare API token for deployments | [Cloudflare Dashboard](https://dash.cloudflare.com/profile/api-tokens) → Create Token → "Edit Cloudflare Workers" template |
| `CLOUDFLARE_ACCOUNT_ID` | Your Cloudflare account ID | Cloudflare Dashboard → Any domain → Overview → Right sidebar "Account ID" |

### Creating the Cloudflare API Token

1. Go to [Cloudflare API Tokens](https://dash.cloudflare.com/profile/api-tokens)
2. Click **Create Token**
3. Use the **"Edit Cloudflare Workers"** template, or create a custom token with:
   - **Account** → Cloudflare Pages: Edit
   - **Account** → Workers Scripts: Edit
   - **Zone** → Zone: Read (for custom domains)
4. Set appropriate IP/TTL restrictions for security
5. Copy the token and add it as `CLOUDFLARE_API_TOKEN` secret

### Optional Secrets (for package publishing)

| Secret | Description | How to Get |
|--------|-------------|------------|
| `NPM_TOKEN` | npm access token for publishing TypeScript SDK | [npmjs.com](https://www.npmjs.com/) → Access Tokens |
| `PYPI_TOKEN` | PyPI API token for publishing Python SDK | [pypi.org](https://pypi.org/) → Account Settings → API Tokens |

## Workflow Details

### CI Workflow (`ci.yml`)

Runs on every push and pull request to `main` branch:

- ✅ Validates repository structure
- ✅ Tests version bump script
- ✅ Builds all SDKs (TypeScript, Python, Go, Rust)
- ✅ Builds CLI tool
- ✅ Builds web application
- ✅ Security checks for leaked secrets

### Deploy Website (`deploy.yml`)

Automatically deploys when changes are pushed to:
- `apps/web/**`
- `spec/v1/**`
- `scripts/generate_web_docs.py`

Deploys to: **https://bspec.dev**

### Deploy MCP Server (`deploy-mcp.yml`)

Automatically deploys when changes are pushed to:
- `apps/mcp/**`
- `spec/v1/**`

Deploys to: **https://mcp.bspec.dev**

### Release Workflow (`release.yml`)

Manual workflow for creating releases. Options:

- **version_bump**: `patch`, `minor`, or `major`
- **create_release**: Create GitHub release with CLI binaries
- **deploy_web**: Deploy website after release
- **deploy_mcp**: Deploy MCP server after release

This workflow:
1. Bumps version across all components
2. Creates git tag
3. Builds CLI for all platforms (Linux, macOS, Windows × amd64/arm64)
4. Creates GitHub release with downloadable binaries
5. Optionally deploys website and MCP server

## Local Development

For local development, use [direnv](https://direnv.net/) with 1Password:

```bash
# Install prerequisites
brew install direnv 1password-cli

# Add to ~/.zshrc
eval "$(direnv hook zsh)"

# In the repo root
cp .envrc.example .envrc
# Edit .envrc to point to your 1Password vault
direnv allow
```

## Troubleshooting

### "Authentication error" during deployment

1. Verify `CLOUDFLARE_API_TOKEN` is set correctly
2. Verify `CLOUDFLARE_ACCOUNT_ID` matches your Cloudflare account
3. Ensure the API token has the required permissions

### CLI builds fail on specific platform

The CLI builds use cross-compilation via `CGO_ENABLED=0`. If builds fail:
1. Check Go version compatibility
2. Ensure no CGO dependencies in the CLI code

### Website deployment doesn't update

Cloudflare Pages may cache content. Try:
1. Clear Cloudflare cache in dashboard
2. Wait a few minutes for propagation
3. Check the deployment logs in GitHub Actions

## Security

- Never commit secrets to the repository
- Use GitHub Secrets for all sensitive values
- The `.envrc` file uses 1Password's `op read` for local development
- Cloudflare Account IDs are not secret (visible in dashboard URLs)
