#!/bin/bash
# Atelier Push Engine
# Handles pushing atelier-level changes and orchestrating artist pushes

set -euo pipefail

# Source dependencies
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/config.sh"
source "$SCRIPT_DIR/git-helpers.sh"

# Script variables
LEVEL="atelier"
DRY_RUN=false
VERBOSE=true
PUSH_ARTISTS=false
PUSH_ALL=false
PUSHED_ITEMS=()
ARTISTS_PROCESSED=()

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --artists)
            PUSH_ARTISTS=true
            shift
            ;;
        --all)
            PUSH_ALL=true
            PUSH_ARTISTS=true
            shift
            ;;
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
            echo "Push atelier-level changes and optionally artists/canvases"
            echo ""
            echo "Options:"
            echo "  --artists    Push all artists and their canvases"
            echo "  --all        Push everything in the hierarchy"
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

# Default to recursing into artists and their canvases
PUSH_ARTISTS=${PUSH_ARTISTS:-true}
PUSH_ALL=${PUSH_ALL:-true}

# Validate we're in the atelier
current_level=$(detect_level)
if [ "$current_level" != "$LEVEL" ]; then
    handle_error "Not in atelier directory (detected: $current_level)" $EXIT_INVALID_LEVEL
fi

# Validate git repository
validate_git_repo
validate_remote

log_info "Starting atelier push operation"

# Get repository information
get_repo_info
atelier_name=$(basename "$(git rev-parse --show-toplevel)")

# Detect modified artist submodule pointers
mapfile -t sub_paths < <(git submodule status 2>/dev/null | awk '{print $2}')
any_sub_modified=false
for p in "${sub_paths[@]}"; do
    if [ -n "$p" ] && is_submodule_modified "$p"; then
        any_sub_modified=true
        break
    fi
done

# Handle artists (default recursive behavior)
log_info "Processing artists in atelier: $atelier_name"

# Find all artists
mapfile -t artists < <(find_artists)

if [ ${#artists[@]} -eq 0 ]; then
    log_warn "No artists found in atelier: $atelier_name"
else
    artist_count=0
    for artist_dir in "${artists[@]}"; do
        artist_count=$((artist_count + 1))
        artist_name=$(basename "$artist_dir")

        show_progress $artist_count ${#artists[@]} "Processing artist: $artist_name"

        # Always delegate to artist engine; it will no-op if nothing to do
        artist_args=()
        if [ "$DRY_RUN" = true ]; then
            artist_args+=("--dry-run")
        fi
        if [ "$VERBOSE" = false ]; then
            artist_args+=("--quiet")
        fi

        if (cd "$artist_dir" && "$SCRIPT_DIR/artist-push.sh" "${artist_args[@]}"); then
            ARTISTS_PROCESSED+=("$artist_name")
            PUSHED_ITEMS+=("$artist_name (artist)")
        else
            log_warn "Failed to push artist: $artist_name"
        fi
    done

    echo # New line after progress
fi

# Check atelier-level changes
has_atelier_changes=false

# Check for uncommitted changes
if has_uncommitted_changes; then
    log_warn "Atelier has uncommitted changes."
    git status --short
    if [ "$DRY_RUN" = true ]; then
        log_warn "Continuing due to --dry-run; uncommitted changes ignored."
    else
        if [ "${AUTO_COMMIT_DEFAULT:-false}" = true ]; then
            log_warn "Staging atelier working tree changes."
            git add -A
        else
            exit $EXIT_ERROR
        fi
    fi
fi

# Re-detect modified artist submodule pointers after processing artists
mapfile -t sub_paths < <(git submodule status 2>/dev/null | awk '{print $2}')
any_sub_modified=false
for p in "${sub_paths[@]}"; do
    if [ -n "$p" ] && is_submodule_modified "$p"; then
        any_sub_modified=true
        break
    fi
done

# If artist submodule pointers changed, stage them for a single combined commit
if [ "$any_sub_modified" = true ] && [ "${DRY_RUN:-false}" != true ] && [ "${AUTO_COMMIT_DEFAULT:-false}" = true ]; then
    log_info "Staging updated artist submodule pointers."
    for p in "${sub_paths[@]}"; do
        if [ -n "$p" ] && is_submodule_modified "$p"; then
            git add "$p"
        fi
    done
fi

# Single combined commit for any staged changes
if [ "${DRY_RUN:-false}" != true ] && [ "${AUTO_COMMIT_DEFAULT:-false}" = true ]; then
    if ! git diff --cached --quiet; then
        git commit -m "${AUTO_COMMIT_MESSAGE:-engine: auto-commit atelier changes}"
    fi
fi

# Check for unpushed commits
if has_unpushed_commits; then
    has_atelier_changes=true
fi

# Check if any artists were updated (which would require atelier commit)
if [ ${#ARTISTS_PROCESSED[@]} -gt 0 ]; then
    has_atelier_changes=true
    log_info "Artists processed, will update atelier references"
fi

if [ "$has_atelier_changes" = false ]; then
    log_info "No atelier-level changes found"
    if [ ${#PUSHED_ITEMS[@]} -gt 0 ]; then
        show_summary "${PUSHED_ITEMS[@]}"
        log_success "Atelier submodules pushed successfully"
    fi
    exit $EXIT_SUCCESS
fi

if [ "$DRY_RUN" = true ]; then
    log_info "[DRY RUN] Would push atelier: $atelier_name"
    if [ ${#ARTISTS_PROCESSED[@]} -gt 0 ]; then
        log_info "[DRY RUN] Would update references for artists: ${ARTISTS_PROCESSED[*]}"
    fi
    exit $EXIT_SUCCESS
fi

# Confirm push
if ! confirm_action "Push atelier '$atelier_name' to $REMOTE_URL?"; then
    exit $EXIT_SUCCESS
fi

# Perform the push
log_info "Pushing atelier: $atelier_name"
if safe_git_command push "$REMOTE_NAME" "$CURRENT_BRANCH"; then
    log_success "Successfully pushed atelier: $atelier_name"
    PUSHED_ITEMS+=("$atelier_name (atelier)")
else
    handle_error "Failed to push atelier: $atelier_name"
fi

# Show summary
show_summary "${PUSHED_ITEMS[@]}"

log_success "Atelier push operation completed"
exit $EXIT_SUCCESS