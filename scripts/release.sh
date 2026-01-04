#!/usr/bin/env bash
#
# BSpec Release Script
# Automates the process of cutting a new release
#

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print colored output
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

# Print header
print_header() {
    echo ""
    echo "╔═══════════════════════════════════════════════════════╗"
    echo "║           BSpec Release Manager                        ║"
    echo "╚═══════════════════════════════════════════════════════╝"
    echo ""
}

# Check if we're in the correct directory
check_directory() {
    if [ ! -f "spec/v1/version.txt" ]; then
        print_error "This script must be run from the BSpec root directory"
        exit 1
    fi
}

# Get current version
get_current_version() {
    if [ -f "spec/v1/version.txt" ]; then
        cat spec/v1/version.txt
    else
        echo "0.0.0"
    fi
}

# Show release options
show_menu() {
    local current_version=$1
    
    echo "Current Version: ${GREEN}v${current_version}${NC}"
    echo ""
    echo "Select release type:"
    echo "  1) ${BLUE}Patch${NC} - Bug fixes and minor updates (v1.0.0 → v1.0.1)"
    echo "  2) ${YELLOW}Minor${NC} - New features, backward compatible (v1.0.0 → v1.1.0)"
    echo "  3) ${RED}Major${NC} - Breaking changes (v1.0.0 → v2.0.0)"
    echo "  4) ${GREEN}Custom${NC} - Specify exact version"
    echo "  q) Quit"
    echo ""
    read -p "Enter choice: " choice
}

# Preview changes
preview_changes() {
    print_info "Recent changes since last release:"
    echo ""
    
    local last_tag=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
    
    if [ -n "$last_tag" ]; then
        git log $last_tag..HEAD --oneline --pretty=format:"  • %s" | head -20
    else
        git log --oneline --pretty=format:"  • %s" | head -20
    fi
    
    echo ""
    echo ""
}

# Check git status
check_git_status() {
    print_info "Checking git status..."
    
    # Check for uncommitted changes
    if [ -n "$(git status --porcelain)" ]; then
        print_warning "You have uncommitted changes:"
        git status --short
        echo ""
        read -p "Continue anyway? (y/N): " continue
        if [[ ! "$continue" =~ ^[Yy]$ ]]; then
            print_error "Release cancelled"
            exit 1
        fi
    else
        print_success "Working directory is clean"
    fi
    
    # Check if we're on main branch
    local current_branch=$(git branch --show-current)
    if [ "$current_branch" != "main" ]; then
        print_warning "You're on branch '$current_branch', not 'main'"
        read -p "Continue anyway? (y/N): " continue
        if [[ ! "$continue" =~ ^[Yy]$ ]]; then
            print_error "Release cancelled"
            exit 1
        fi
    else
        print_success "On main branch"
    fi
    
    # Check if we're up to date
    git fetch origin main --quiet
    local local_commit=$(git rev-parse HEAD)
    local remote_commit=$(git rev-parse origin/main)
    
    if [ "$local_commit" != "$remote_commit" ]; then
        print_warning "Local branch is not up to date with remote"
        read -p "Pull latest changes? (Y/n): " pull
        if [[ ! "$pull" =~ ^[Nn]$ ]]; then
            git pull origin main
            print_success "Pulled latest changes"
        fi
    else
        print_success "Branch is up to date"
    fi
}

# Run pre-release checks
run_checks() {
    print_info "Running pre-release checks..."
    
    # Check if required files exist
    local required_files=(
        "spec/v1/version.txt"
        "scripts/bump-version.py"
        ".github/workflows/release.yml"
    )
    
    for file in "${required_files[@]}"; do
        if [ ! -f "$file" ]; then
            print_error "Required file not found: $file"
            exit 1
        fi
    done
    
    print_success "All required files present"
}

# Trigger GitHub Actions release
trigger_release() {
    local bump_type=$1
    
    print_info "Triggering GitHub Actions release workflow..."
    print_info "Bump type: $bump_type"
    
    # Check if gh CLI is installed
    if ! command -v gh &> /dev/null; then
        print_error "GitHub CLI (gh) is not installed"
        echo ""
        echo "Install it with:"
        echo "  macOS: brew install gh"
        echo "  Linux: See https://github.com/cli/cli/blob/trunk/docs/install_linux.md"
        echo ""
        echo "Or manually trigger the release workflow from:"
        echo "https://github.com/a3tai/bspec/actions/workflows/release.yml"
        exit 1
    fi
    
    # Check if authenticated
    if ! gh auth status &> /dev/null; then
        print_error "Not authenticated with GitHub"
        echo ""
        echo "Run: gh auth login"
        exit 1
    fi
    
    # Trigger the workflow
    gh workflow run release.yml \
        -f version_bump=$bump_type \
        -f create_release=true \
        -f publish_packages=false
    
    print_success "Release workflow triggered!"
    echo ""
    print_info "Monitor the release at:"
    echo "  https://github.com/a3tai/bspec/actions/workflows/release.yml"
    echo ""
    
    # Wait a moment and show the run
    sleep 2
    print_info "Latest workflow runs:"
    gh run list --workflow=release.yml --limit=3
}

# Manual local release (for testing)
manual_release() {
    local bump_type=$1
    local current_version=$(get_current_version)
    
    print_info "Performing local release (for testing only)..."
    
    # Bump version
    python3 scripts/bump-version.py $bump_type
    local new_version=$(get_current_version)
    
    print_success "Version bumped: $current_version → $new_version"
    
    # Show what would be committed
    echo ""
    print_info "Changes to be committed:"
    git diff --stat
    
    echo ""
    read -p "Commit these changes? (y/N): " commit
    if [[ "$commit" =~ ^[Yy]$ ]]; then
        git add -A
        git commit -m "bump: version $new_version"
        git tag -a "v$new_version" -m "Release v$new_version"
        
        print_success "Committed and tagged v$new_version"
        echo ""
        read -p "Push to remote? (y/N): " push
        if [[ "$push" =~ ^[Yy]$ ]]; then
            git push origin main
            git push origin --tags
            print_success "Pushed to remote"
        fi
    else
        print_warning "Changes not committed. Reset with: git reset --hard HEAD"
    fi
}

# Main script
main() {
    print_header
    check_directory
    
    local current_version=$(get_current_version)
    
    # Show menu
    show_menu "$current_version"
    
    case $choice in
        1)
            bump_type="patch"
            ;;
        2)
            bump_type="minor"
            ;;
        3)
            bump_type="major"
            ;;
        4)
            read -p "Enter exact version (e.g., 1.2.3): " custom_version
            # TODO: Implement custom version
            print_error "Custom version not yet implemented"
            exit 1
            ;;
        q|Q)
            print_info "Release cancelled"
            exit 0
            ;;
        *)
            print_error "Invalid choice"
            exit 1
            ;;
    esac
    
    echo ""
    print_header
    
    # Preview changes
    preview_changes
    
    # Run checks
    check_git_status
    run_checks
    
    echo ""
    print_info "Ready to release!"
    print_warning "This will trigger a $bump_type version bump"
    echo ""
    
    read -p "Continue? (y/N): " confirm
    if [[ ! "$confirm" =~ ^[Yy]$ ]]; then
        print_error "Release cancelled"
        exit 1
    fi
    
    echo ""
    print_info "Choose release method:"
    echo "  1) ${GREEN}GitHub Actions${NC} (Recommended - full automation)"
    echo "  2) ${YELLOW}Manual${NC} (Local testing only)"
    echo ""
    read -p "Enter choice: " method
    
    case $method in
        1)
            trigger_release "$bump_type"
            ;;
        2)
            manual_release "$bump_type"
            ;;
        *)
            print_error "Invalid choice"
            exit 1
            ;;
    esac
    
    echo ""
    print_success "Release process complete!"
}

# Run main
main
