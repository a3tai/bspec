#!/usr/bin/env bash
#
# Setup GitHub Secrets for BSpec
#

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_header() {
    echo ""
    echo "╔═══════════════════════════════════════════════════════╗"
    echo "║        BSpec GitHub Secrets Setup                     ║"
    echo "╚═══════════════════════════════════════════════════════╝"
    echo ""
}

# Check if gh CLI is installed
check_gh_cli() {
    if ! command -v gh &> /dev/null; then
        print_error "GitHub CLI (gh) is not installed"
        echo ""
        echo "Install it with:"
        echo "  macOS: brew install gh"
        echo "  Linux: See https://github.com/cli/cli/blob/trunk/docs/install_linux.md"
        exit 1
    fi
}

# Check if authenticated
check_auth() {
    if ! gh auth status &> /dev/null; then
        print_error "Not authenticated with GitHub"
        echo ""
        echo "Run: gh auth login"
        exit 1
    fi
}

# Get repository name
get_repo() {
    gh repo view --json nameWithOwner -q .nameWithOwner 2>/dev/null || echo "a3tai/bspec"
}

# List existing secrets
list_secrets() {
    print_info "Current secrets:"
    gh secret list || echo "No secrets set"
    echo ""
}

# Set a secret
set_secret() {
    local secret_name=$1
    local secret_value=$2
    local description=$3
    
    if [ -z "$secret_value" ]; then
        print_warning "Skipping $secret_name (no value provided)"
        return
    fi
    
    echo "$secret_value" | gh secret set "$secret_name"
    
    if [ $? -eq 0 ]; then
        print_success "Set $secret_name"
    else
        print_error "Failed to set $secret_name"
    fi
}

# Set a secret interactively
set_secret_interactive() {
    local secret_name=$1
    local description=$2
    local example=$3
    local required=$4
    
    echo ""
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${GREEN}$secret_name${NC}"
    echo "$description"
    
    if [ -n "$example" ]; then
        echo -e "${YELLOW}Example:${NC} $example"
    fi
    
    echo ""
    
    # Check if secret already exists
    if gh secret list | grep -q "^$secret_name"; then
        print_warning "Secret already exists"
        read -p "Overwrite? (y/N): " overwrite
        if [[ ! "$overwrite" =~ ^[Yy]$ ]]; then
            print_info "Skipped $secret_name"
            return
        fi
    fi
    
    # Try to get value from environment variable
    local env_value="${!secret_name}"
    
    if [ -n "$env_value" ]; then
        print_info "Found value in environment variable \$${secret_name}"
        read -p "Use this value? (Y/n): " use_env
        if [[ ! "$use_env" =~ ^[Nn]$ ]]; then
            set_secret "$secret_name" "$env_value" "$description"
            return
        fi
    fi
    
    if [[ "$required" == "required" ]]; then
        echo -e "${RED}REQUIRED${NC}"
    else
        echo -e "${YELLOW}OPTIONAL${NC} - Press Enter to skip"
    fi
    
    echo ""
    read -s -p "Enter value (hidden): " secret_value
    echo ""
    
    if [ -z "$secret_value" ]; then
        if [[ "$required" == "required" ]]; then
            print_error "This secret is required!"
            return 1
        else
            print_info "Skipped $secret_name"
            return 0
        fi
    fi
    
    set_secret "$secret_name" "$secret_value" "$description"
}

# Main script
main() {
    print_header
    
    check_gh_cli
    check_auth
    
    local repo=$(get_repo)
    print_info "Setting secrets for repository: $repo"
    echo ""
    
    list_secrets
    
    echo ""
    echo "This script will help you set up the required GitHub secrets."
    echo "Secrets are needed for:"
    echo "  • Deploying the website to Cloudflare Pages"
    echo "  • Publishing packages to npm and PyPI (optional)"
    echo ""
    read -p "Continue? (Y/n): " continue
    
    if [[ "$continue" =~ ^[Nn]$ ]]; then
        print_info "Setup cancelled"
        exit 0
    fi
    
    # Cloudflare API Token (Required)
    set_secret_interactive \
        "CLOUDFLARE_API_TOKEN" \
        "Required for deploying the website to Cloudflare Pages.
Get this from: https://dash.cloudflare.com/profile/api-tokens
Create a token with 'Cloudflare Pages:Edit' permissions." \
        "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" \
        "required"
    
    # Cloudflare Account ID (Required)
    set_secret_interactive \
        "CLOUDFLARE_ACCOUNT_ID" \
        "Your Cloudflare Account ID.
Find this in Cloudflare dashboard → Account → Account ID" \
        "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" \
        "required"
    
    # NPM Token (Optional)
    set_secret_interactive \
        "NPM_TOKEN" \
        "Optional: For publishing TypeScript SDK to npm.
Create at: https://www.npmjs.com/settings/YOUR_USERNAME/tokens
Token type: Automation" \
        "npm_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" \
        "optional"
    
    # PyPI Token (Optional)
    set_secret_interactive \
        "PYPI_TOKEN" \
        "Optional: For publishing Python SDK to PyPI.
Create at: https://pypi.org/manage/account/token/
Scope: Entire account or specific project" \
        "pypi-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" \
        "optional"
    
    echo ""
    echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo ""
    print_success "Setup complete!"
    echo ""
    print_info "Current secrets:"
    gh secret list
    echo ""
    print_info "You can update secrets anytime by running this script again"
    echo ""
    print_info "Next steps:"
    echo "  1. Test deployment: gh workflow run deploy.yml"
    echo "  2. Cut a release: ./scripts/release.sh"
}

# Run main
main
