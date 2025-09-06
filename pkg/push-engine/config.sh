#!/bin/bash
# Git Push Engine Configuration
# Static constants and defaults

# Marker files for level detection
ATELIER_MARKER=".atelier"
ARTIST_MARKER=".artist"
CANVAS_MARKER=".canvas"

# Git defaults
DEFAULT_BRANCH="main"
REMOTE_NAME="origin"
MAX_RETRIES=3
TIMEOUT_SECONDS=300

# Behavior defaults (allow environment to override)
DRY_RUN_DEFAULT=${DRY_RUN_DEFAULT:-false}
VERBOSE_DEFAULT=${VERBOSE_DEFAULT:-false}
CONFIRM_PUSH_DEFAULT=${CONFIRM_PUSH_DEFAULT:-true}
CONFIRM_FORCE_DEFAULT=${CONFIRM_FORCE_DEFAULT:-false}

# Logging
LOG_LEVEL_DEFAULT="info"
LOG_LEVELS=("error" "warn" "info" "debug")

# Output formatting
USE_COLORS=true
PROGRESS_BAR=true
SUMMARY_REPORT=true

# Push settings
PUSH_TIMEOUT=60
FORCE_PUSH_CONFIRM=true

# Auto-commit behavior (allow environment override)
# Default to true so 'make push' performs major roll-up commits without extra env vars.
AUTO_COMMIT_DEFAULT=${AUTO_COMMIT_DEFAULT:-true}
AUTO_COMMIT_MESSAGE=${AUTO_COMMIT_MESSAGE:-"engine: auto-commit uncommitted changes"}

# Directory patterns
ARTIST_PATTERN="artist-*"
CANVAS_PATTERN="canvas-*"

# Required files for validation
CANVAS_REQUIRED_FILES=("README.md" "Makefile")
ARTIST_REQUIRED_FILES=("README.md")
ATELIER_REQUIRED_FILES=("README.md")

# Exit codes
EXIT_SUCCESS=0
EXIT_ERROR=1
EXIT_INVALID_LEVEL=2
EXIT_NO_CHANGES=3
EXIT_GIT_ERROR=4