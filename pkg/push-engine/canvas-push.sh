#!/bin/bash
# Canvas Push Engine
# Handles pushing canvas-level changes

set -euo pipefail

# Source dependencies
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/config.sh"
source "$SCRIPT_DIR/git-helpers.sh"

# Script variables
LEVEL="canvas"
DRY_RUN=false
VERBOSE=true
PUSHED_ITEMS=()

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        --quiet)
            VERBOSE=false
            shift
            ;;
        --help)
            echo "Usage: $0 [OPTIONS]"
            echo "Push canvas-level changes"
            echo ""
            echo "Options:"
            echo "  --dry-run    Show what would be pushed without pushing"
            echo "  --quiet      Suppress verbose output"
            echo "  --help       Show this help message"
            exit 0
            ;;
        *)
            log_error "Unknown option: $1"
            exit $EXIT_ERROR
            ;;
    esac
done

# Validate we're in a canvas
current_level=$(detect_level)
if [ "$current_level" != "$LEVEL" ]; then
    handle_error "Not in a canvas directory (detected: $current_level)" $EXIT_INVALID_LEVEL
fi

# Validate git repository
validate_git_repo
validate_remote

log_info "Starting canvas push operation"

# Handle uncommitted changes
if has_uncommitted_changes; then
    git status --short
    if [ "${DRY_RUN:-false}" = true ]; then
        log_warn "Canvas has uncommitted changes (dry-run; not committing)."
        # In dry-run, continue to show what would happen next
    else
        if [ "${AUTO_COMMIT_DEFAULT:-false}" = true ]; then
            log_warn "Canvas has uncommitted changes; auto-committing."
            git add -A
            git commit -m "${AUTO_COMMIT_MESSAGE:-engine: auto-commit uncommitted changes}"
        else
            log_warn "You have uncommitted changes. Please commit or stash them first."
            exit $EXIT_ERROR
        fi
    fi
fi

# Check for unpushed commits
if ! has_unpushed_commits; then
    log_info "No unpushed commits found"
    exit $EXIT_NO_CHANGES
fi

# Get repository information
get_repo_info
repo_name=$(basename "$(git rev-parse --show-toplevel)")

if [ "$DRY_RUN" = true ]; then
    log_info "[DRY RUN] Would push canvas: $repo_name"
    log_info "[DRY RUN] Remote: $REMOTE_URL"
    log_info "[DRY RUN] Branch: $CURRENT_BRANCH"
    exit $EXIT_SUCCESS
fi

# Confirm push
if ! confirm_action "Push canvas '$repo_name' to $REMOTE_URL?"; then
    exit $EXIT_SUCCESS
fi

# Perform the push
log_info "Pushing canvas: $repo_name"
if safe_git_command push "$REMOTE_NAME" "$CURRENT_BRANCH"; then
    log_success "Successfully pushed canvas: $repo_name"
    PUSHED_ITEMS+=("$repo_name (canvas)")
else
    handle_error "Failed to push canvas: $repo_name"
fi

# Show summary
show_summary "${PUSHED_ITEMS[@]}"

log_success "Canvas push operation completed"
exit $EXIT_SUCCESS