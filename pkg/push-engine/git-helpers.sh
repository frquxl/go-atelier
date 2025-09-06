#!/bin/bash
# Git Push Engine - Shared Helper Functions
# Provides common git operations and dynamic discovery functions

# Source configuration
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/config.sh"

# Initialize variables
LOG_LEVEL="${LOG_LEVEL:-$LOG_LEVEL_DEFAULT}"

# Logging functions
log_error() { echo -e "\033[31m[ERROR]\033[0m $*" >&2; }
log_warn() { echo -e "\033[33m[WARN]\033[0m $*" >&2; }
log_info() { echo -e "\033[36m[INFO]\033[0m $*" >&2; }
# Use an if to avoid set -e exiting on a non-zero [ ... ] return status
log_debug() {
  if [ "${LOG_LEVEL:-$LOG_LEVEL_DEFAULT}" = "debug" ]; then
    echo -e "\033[35m[DEBUG]\033[0m $*" >&2
  fi
}
log_success() { echo -e "\033[32m[SUCCESS]\033[0m $*" >&2; }

# Context detection functions
detect_level() {
    if [ -f "$ATELIER_MARKER" ]; then
        echo "atelier"
    elif [ -f "$ARTIST_MARKER" ]; then
        echo "artist"
    elif [ -f "$CANVAS_MARKER" ]; then
        echo "canvas"
    else
        echo "unknown"
    fi
}

# Repository information functions
get_repo_info() {
    REMOTE_URL=$(git remote get-url "$REMOTE_NAME" 2>/dev/null || echo "")
    CURRENT_BRANCH=$(git branch --show-current 2>/dev/null || echo "")
    if [ -n "$REMOTE_URL" ]; then
        GITHUB_ORG=$(echo "$REMOTE_URL" | sed 's/.*github.com[:/]\([^/]*\).*/\1/' || echo "")
    fi
}

# Submodule discovery functions
get_submodules() {
    git config --file .gitmodules --list 2>/dev/null |
    grep '^submodule\..*\.path=' |
    sed 's/submodule\.\([^.]*\)\.path=.*/\1/' ||
    echo ""
}

get_submodule_url() {
    local submodule=$1
    git config --file .gitmodules submodule."$submodule".url 2>/dev/null || echo ""
}

# Directory discovery functions
find_artists() {
    find . -maxdepth 1 -name "$ARTIST_PATTERN" -type d 2>/dev/null | sort
}

find_canvases() {
    find . -maxdepth 1 -name "$CANVAS_PATTERN" -type d 2>/dev/null | sort
}

# Git status functions
has_unpushed_commits() {
    local ahead_behind
    ahead_behind=$(git rev-list --count --left-right "@{upstream}"...HEAD 2>/dev/null || echo "0 0")
    local ahead=$(echo "$ahead_behind" | awk '{print $2}' | tr -d '\t')
    [ "${ahead:-0}" -gt 0 ] 2>/dev/null
}

has_uncommitted_changes() {
    [ -n "$(git status --porcelain 2>/dev/null)" ]
}

# Submodule status functions
get_submodule_status() {
    local submodule=$1
    git submodule status "$submodule" 2>/dev/null | head -1 || echo ""
}

is_submodule_modified() {
    local submodule=$1
    local status
    status=$(get_submodule_status "$submodule")
    [ -n "$status" ] && [[ "$status" == +* ]]
}

# Git operation functions
safe_git_command() {
    local cmd=$*
    log_debug "Executing: git $cmd"
    if git "$@" 2>&1; then
        return 0
    else
        local exit_code=$?
        log_error "Git command failed: git $cmd"
        return $exit_code
    fi
}

# Validation functions
validate_git_repo() {
    if ! git rev-parse --git-dir >/dev/null 2>&1; then
        log_error "Not a git repository"
        return $EXIT_GIT_ERROR
    fi
    return $EXIT_SUCCESS
}

validate_remote() {
    if ! git remote get-url "$REMOTE_NAME" >/dev/null 2>&1; then
        log_error "No remote '$REMOTE_NAME' configured"
        return $EXIT_GIT_ERROR
    fi
    return $EXIT_SUCCESS
}

# Progress and output functions
show_progress() {
    local current=$1
    local total=$2
    local item=$3
    if [ "$PROGRESS_BAR" = true ]; then
        printf "\r[%d/%d] %s" "$current" "$total" "$item"
    else
        log_info "Processing: $item"
    fi
}

show_summary() {
    local pushed=("$@")
    if [ "$SUMMARY_REPORT" = true ] && [ ${#pushed[@]} -gt 0 ]; then
        log_success "Push Summary:"
        printf '%s\n' "${pushed[@]}"
    fi
}

# Confirmation functions
confirm_action() {
    local message=$1
    local default=${2:-$CONFIRM_PUSH_DEFAULT}

    # Non-interactive acceptance:
    # - If default says no prompt (false)
    # - If terminal is not interactive (no TTY)
    # - If ENGINE_ASSUME_YES=true is set
    if [ "$default" = false ] || [ ! -t 0 ] || [ "${ENGINE_ASSUME_YES:-false}" = true ]; then
        return $EXIT_SUCCESS
    fi

    echo -n "$message [y/N]: "
    read -r response
    case $response in
        [Yy]|[Yy][Ee][Ss])
            return $EXIT_SUCCESS
            ;;
        *)
            log_info "Operation cancelled"
            return $EXIT_ERROR
            ;;
    esac
}

# Error handling
handle_error() {
    local message=$1
    local exit_code=${2:-$EXIT_ERROR}
    log_error "$message"
    exit $exit_code
}

# Cleanup function
cleanup() {
    # Reset terminal if needed
    if [ "$USE_COLORS" = true ]; then
        echo -e "\033[0m" >&2
    fi
}

# Set up error handling
trap cleanup EXIT