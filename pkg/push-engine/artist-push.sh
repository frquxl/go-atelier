#!/bin/bash
# Artist Push Engine
# Handles pushing artist-level changes and updating canvases

set -euo pipefail

# Source dependencies
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/config.sh"
source "$SCRIPT_DIR/git-helpers.sh"

# Script variables
LEVEL="artist"
DRY_RUN=false
VERBOSE=true
PUSHED_ITEMS=()
CANVASES_UPDATED=()

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --canvases)
            PUSH_CANVASES=true
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
            echo "Push artist-level changes and optionally canvases"
            echo ""
            echo "Options:"
            echo "  --canvases   Also push all canvases in this artist"
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

# Default to recursing into canvases
PUSH_CANVASES=${PUSH_CANVASES:-true}

# In dry-run, disable exit-on-error to allow full flow and reporting
if [ "${DRY_RUN:-false}" = true ]; then
    set +e
fi

# Validate we're in an artist
current_level=$(detect_level)
if [ "$current_level" != "$LEVEL" ]; then
    handle_error "Not in an artist directory (detected: $current_level)" $EXIT_INVALID_LEVEL
fi

# Validate git repository
validate_git_repo
validate_remote

log_info "Starting artist push operation"

# Get repository information
get_repo_info
artist_name=$(basename "$(git rev-parse --show-toplevel)")

# Detect modified submodule pointers (canvases)
# Be tolerant if 'git submodule status' returns non-zero (e.g., no submodules yet)
mapfile -t sub_paths < <((git submodule status 2>/dev/null || true) | awk '{print $2}')
any_sub_modified=false
for p in "${sub_paths[@]}"; do
    if [ -n "$p" ] && is_submodule_modified "$p"; then
        any_sub_modified=true
        break
    fi
done

# Handle canvases first (default recursive behavior)
log_info "Processing canvases in artist: $artist_name"

# Find all canvases
mapfile -t canvases < <(find_canvases)

if [ ${#canvases[@]} -eq 0 ]; then
    log_info "No canvases found in artist: $artist_name"
else
    canvas_count=0
    for canvas_dir in "${canvases[@]}"; do
        canvas_count=$((canvas_count + 1))
        canvas_name=$(basename "$canvas_dir")

        show_progress $canvas_count ${#canvases[@]} "Processing canvas: $canvas_name"

        # Always delegate to canvas engine; it will no-op if nothing to do
        canvas_args=()
        if [ "$DRY_RUN" = true ]; then
            canvas_args+=("--dry-run")
        fi
        if [ "$VERBOSE" = false ]; then
            canvas_args+=("--quiet")
        fi

        if (cd "$canvas_dir" && "$SCRIPT_DIR/canvas-push.sh" "${canvas_args[@]}"); then
            :
        else
            log_warn "Failed to push canvas: $canvas_name"
        fi
    done

    echo # New line after progress
fi

# Re-detect modified canvas submodule pointers after processing canvases
mapfile -t sub_paths < <((git submodule status 2>/dev/null || true) | awk '{print $2}')
any_sub_modified=false
for p in "${sub_paths[@]}"; do
    if [ -n "$p" ] && is_submodule_modified "$p"; then
        any_sub_modified=true
        break
    fi
done

# Check artist-level changes
has_artist_changes=false

# Check for uncommitted changes
if has_uncommitted_changes; then
    log_warn "Artist has uncommitted changes."
    git status --short
    if [ "${DRY_RUN:-false}" = true ]; then
        log_warn "Continuing due to --dry-run; uncommitted changes ignored."
    else
        if [ "${AUTO_COMMIT_DEFAULT:-false}" = true ]; then
            log_warn "Staging artist working tree changes."
            git add -A
        else
            exit $EXIT_ERROR
        fi
    fi
fi

# If submodule pointers changed, stage them for a single combined commit
if [ "$any_sub_modified" = true ] && [ "${DRY_RUN:-false}" != true ] && [ "${AUTO_COMMIT_DEFAULT:-false}" = true ]; then
    log_info "Staging updated canvas submodule pointers."
    for p in "${sub_paths[@]}"; do
        if [ -n "$p" ] && is_submodule_modified "$p"; then
            git add "$p"
        fi
    done
fi

# Single combined commit for any staged changes
if [ "${DRY_RUN:-false}" != true ] && [ "${AUTO_COMMIT_DEFAULT:-false}" = true ]; then
    if ! git diff --cached --quiet; then
        git commit -m "${AUTO_COMMIT_MESSAGE:-engine: auto-commit artist changes}"
        has_artist_changes=true
    fi
fi

# Check for unpushed commits
if has_unpushed_commits; then
    has_artist_changes=true
    log_debug "Artist has unpushed commits -> will push."
fi

# Consider detected submodule pointer changes as a reason to push
if [ "$any_sub_modified" = true ]; then
    has_artist_changes=true
    log_debug "Artist has modified submodule pointers -> will push."
fi

# Check if any canvases were updated (which would require artist commit)
if [ ${#CANVASES_UPDATED[@]} -gt 0 ]; then
    has_artist_changes=true
    log_info "Canvases updated, will update artist references"
fi

if [ "$has_artist_changes" = false ]; then
    log_info "No artist-level changes detected; attempting push to ensure remote sync"
fi

if [ "$DRY_RUN" = true ]; then
    log_info "[DRY RUN] Would push artist: $artist_name"
    if [ ${#CANVASES_UPDATED[@]} -gt 0 ]; then
        log_info "[DRY RUN] Would update references for canvases: ${CANVASES_UPDATED[*]}"
    fi
    exit $EXIT_SUCCESS
fi

# Confirm push
if ! confirm_action "Push artist '$artist_name' to $REMOTE_URL?"; then
    exit $EXIT_SUCCESS
fi

# Perform the push
log_info "Pushing artist: $artist_name"
if safe_git_command push "$REMOTE_NAME" "$CURRENT_BRANCH"; then
    log_success "Successfully pushed artist: $artist_name"
    PUSHED_ITEMS+=("$artist_name (artist)")
else
    handle_error "Failed to push artist: $artist_name"
fi

# Show summary
show_summary "${PUSHED_ITEMS[@]}"

log_success "Artist push operation completed"
exit $EXIT_SUCCESS