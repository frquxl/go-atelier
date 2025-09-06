#!/bin/bash
# Git Push Engine - Main Orchestrator
# Automatically detects level and delegates to appropriate push script

set -uo pipefail

# Source dependencies
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/config.sh"
source "$SCRIPT_DIR/git-helpers.sh"

# Script variables
# Make engine non-interactive by default unless user overrides.
export ENGINE_ASSUME_YES=${ENGINE_ASSUME_YES:-true}

DRY_RUN=false
VERBOSE=true
FORCE=false
LEVEL_ARGS=()

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
        --force)
            FORCE=true
            shift
            ;;
        --help)
            echo "Git Push Engine - Atelier/Artist/Canvas Push Tool"
            echo ""
            echo "Usage: $0 [OPTIONS] [LEVEL_ARGS...]"
            echo ""
            echo "Automatically detects current level (atelier/artist/canvas)"
            echo "and delegates to the appropriate push script."
            echo ""
            echo "Options:"
            echo "  --dry-run    Show what would be pushed without pushing"
            echo "  --quiet      Suppress verbose output"
            echo "  --force      Force push (use with caution)"
            echo "  --help       Show this help message"
            echo ""
            echo "Level-specific arguments are passed through:"
            echo "  Atelier: --artists, --all"
            echo "  Artist:  --canvases"
            echo "  Canvas:  (none)"
            echo ""
            echo "Examples:"
            echo "  $0                          # Push current level"
            echo "  $0 --artists               # Push all artists (atelier only)"
            echo "  $0 --canvases              # Push all canvases (artist only)"
            echo "  $0 --all                   # Push everything (atelier only)"
            echo "  $0 --dry-run               # Preview what would be pushed"
            exit 0
            ;;
        *)
            # Pass through level-specific arguments
            LEVEL_ARGS+=("$1")
            shift
            ;;
    esac
done

# Detect current level
CURRENT_LEVEL=$(detect_level)
echo "DEBUG: Detected level: $CURRENT_LEVEL" >&2

case $CURRENT_LEVEL in
    "atelier")
        log_info "Atelier detected - delegating to atelier push engine"
        SCRIPT="$SCRIPT_DIR/atelier-push.sh"
        ;;
    "artist")
        log_info "Artist detected - delegating to artist push engine"
        SCRIPT="$SCRIPT_DIR/artist-push.sh"
        ;;
    "canvas")
        log_info "Canvas detected - delegating to canvas push engine"
        SCRIPT="$SCRIPT_DIR/canvas-push.sh"
        ;;
    "unknown")
        handle_error "Unable to detect atelier/artist/canvas level. Are you in the right directory?" $EXIT_INVALID_LEVEL
        ;;
    *)
        handle_error "Invalid level detected: $CURRENT_LEVEL" $EXIT_INVALID_LEVEL
        ;;
esac

# Build command arguments
CMD_ARGS=("$SCRIPT")
if [ "$DRY_RUN" = true ]; then
    CMD_ARGS+=("--dry-run")
fi
if [ "$VERBOSE" = false ]; then
    CMD_ARGS+=("--quiet")
fi
if [ "$FORCE" = true ]; then
    CMD_ARGS+=("--force")
fi

# Add level-specific arguments
if [ ${#LEVEL_ARGS[@]} -gt 0 ]; then
    CMD_ARGS+=("${LEVEL_ARGS[@]}")
fi

# Sanity check delegate path
if [ ! -x "$SCRIPT" ]; then
    log_error "Delegate not executable or not found: $SCRIPT"
    ls -la "$SCRIPT_DIR" >&2 || true
    exit $EXIT_ERROR
fi

# Execute the appropriate script and propagate exit status (normalize no-op)
log_info "Delegate: $SCRIPT | DRY_RUN=${DRY_RUN:-false} | VERBOSE=${VERBOSE:-true}"
log_debug "Executing: ${CMD_ARGS[*]}"

# Temporarily disable 'exit on error' to capture delegate status safely
set +e
"${CMD_ARGS[@]}"
status=$?
set -e

log_info "Delegate exit status: $status"

# In dry-run mode we always return success to enable orchestration previews
if [ "${DRY_RUN:-false}" = true ]; then
    log_info "Dry-run mode: returning success."
    exit 0
fi

# Treat EXIT_NO_CHANGES (3) as success (0) so higher levels don't fail on no-ops
if [ "$status" -eq "${EXIT_NO_CHANGES:-3}" ]; then
    log_info "No changes to push at this level."
    exit 0
fi

exit "$status"