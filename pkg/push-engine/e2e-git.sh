#!/usr/bin/env bash
# End-to-end tests for the Git Push Engine
# This suite targets only the guaranteed repos in a new atelier:
#   - canvas: artist-van-gogh/canvas-sunflowers
#   - artist: artist-van-gogh
#   - atelier: root repository
#
# Intent and expectations:
# - Engine defaults:
#   * Artist-level pushes recurse into all canvases by default.
#   * Atelier-level pushes recurse into all artists and their canvases by default.
# - Commit expectations after the full run (relative to initial baselines):
#   * Canvas repo: 3 commits (canvas test, artist test, atelier test)
#   * Artist repo: 2 commits (artist test, atelier test)
#   * Atelier repo: 1 commit (atelier test)
#
# Non-interactive mode:
#   We disable confirmation prompts and enable auto-commit so this script can run unattended.
#   You can override by exporting CONFIRM_PUSH_DEFAULT=true or AUTO_COMMIT_DEFAULT=false before running.
export CONFIRM_PUSH_DEFAULT=${CONFIRM_PUSH_DEFAULT:-false}
export AUTO_COMMIT_DEFAULT=${AUTO_COMMIT_DEFAULT:-true}
export AUTO_COMMIT_MESSAGE=${AUTO_COMMIT_MESSAGE:-"engine(e2e): auto-commit uncommitted changes"}

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
ENGINE_DIR="$ROOT_DIR/util/git"

CANVAS_DIR="$ROOT_DIR/artist-van-gogh/canvas-sunflowers"
ARTIST_DIR="$ROOT_DIR/artist-van-gogh"

timestamp() { date -Iseconds; }

log() { echo -e "\033[36m[INFO]\033[0m $*"; }
warn() { echo -e "\033[33m[WARN]\033[0m $*" >&2; }
err() { echo -e "\033[31m[ERROR]\033[0m $*" >&2; }

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || { err "Missing required command: $1"; exit 1; }
}

remote_head() {
  git -C "$1" ls-remote origin main | awk '{print $1}'
}

local_head() {
  git -C "$1" rev-parse HEAD
}

assert_remote_synced() {
  local repo_dir="$1"
  local lhead rhead
  lhead="$(local_head "$repo_dir")"
  rhead="$(remote_head "$repo_dir")"
  if [[ -z "$rhead" ]]; then
    err "Remote HEAD not found for $repo_dir (is the remote configured and accessible?)"
    exit 1
  fi
  if [[ "$lhead" != "$rhead" ]]; then
    err "Remote not synced for $repo_dir. Local=$lhead Remote=$rhead"
    exit 1
  fi
  log "Remote synced for $repo_dir (HEAD=$lhead)"
}

append_readme_only() {
  local repo_dir="$1"
  local msg="$2"
  local readme="$repo_dir/README.md"
  echo "" >> "$readme" || true
  echo "e2e: $msg @ $(timestamp)" >> "$readme"
  # Do NOT stage or commit here; the engine must handle all git operations.
}

push_canvas_with_engine() {
  local canvas_dir="$1"
  ( cd "$canvas_dir" && "$ENGINE_DIR/push-engine.sh" )
}

push_artist_with_engine() {
  local artist_dir="$1"
  ( cd "$artist_dir" && "$ENGINE_DIR/push-engine.sh" )
}

push_atelier_with_engine() {
  # Defaults now recurse into all artists and canvases
  ( cd "$ROOT_DIR" && "$ENGINE_DIR/push-engine.sh" )
}

count_delta() {
  local repo_dir="$1"
  local base_sha="$2"
  git -C "$repo_dir" rev-list --count "${base_sha}..HEAD"
}

# Capture baselines for delta checks
capture_baselines() {
  BASE_CANVAS="$(local_head "$CANVAS_DIR")"
  BASE_ARTIST="$(local_head "$ARTIST_DIR")"
  BASE_ROOT="$(local_head "$ROOT_DIR")"
}

require_paths() {
  if [[ ! -d "$CANVAS_DIR" ]]; then
    err "Required canvas not found: $CANVAS_DIR"
    exit 1
  fi
  if [[ ! -d "$ARTIST_DIR" ]]; then
    err "Required artist not found: $ARTIST_DIR"
    exit 1
  fi
  if [[ ! -f "$ROOT_DIR/README.md" ]]; then
    warn "Root README.md not found, creating"
    touch "$ROOT_DIR/README.md"
  fi
}

# Tests

test_canvas_sunflowers() {
  log "TEST: Canvas push (sunflowers)"
  append_readme_only "$CANVAS_DIR" "e2e(canvas-sunflowers): canvas stage"
  push_canvas_with_engine "$CANVAS_DIR"
  assert_remote_synced "$CANVAS_DIR"

  local canvas_delta
  canvas_delta="$(count_delta "$CANVAS_DIR" "$BASE_CANVAS")"
  if [[ "$canvas_delta" -ne 1 ]]; then
    err "Expected canvas commit delta 1 after canvas test, got $canvas_delta"
    exit 1
  fi
}

test_artist_van_gogh() {
  log "TEST: Artist push (artist-van-gogh) after canvas & artist edits"
  append_readme_only "$CANVAS_DIR" "e2e(canvas-sunflowers): artist stage"
  append_readme_only "$ARTIST_DIR" "e2e(artist-van-gogh): artist stage"
  # Rely on default recursion at artist level
  push_artist_with_engine "$ARTIST_DIR"
  assert_remote_synced "$CANVAS_DIR"
  assert_remote_synced "$ARTIST_DIR"

  local canvas_delta artist_delta
  canvas_delta="$(count_delta "$CANVAS_DIR" "$BASE_CANVAS")"
  artist_delta="$(count_delta "$ARTIST_DIR" "$BASE_ARTIST")"
  if [[ "$canvas_delta" -ne 2 ]]; then
    err "Expected canvas commit delta 2 after artist test, got $canvas_delta"
    exit 1
  fi
  if [[ "$artist_delta" -ne 1 ]]; then
    err "Expected artist commit delta 1 after artist test, got $artist_delta"
    exit 1
  fi
}

test_atelier_root() {
  log "TEST: Atelier push after updating canvas, artist, and root"
  append_readme_only "$CANVAS_DIR" "e2e(canvas-sunflowers): atelier stage"
  append_readme_only "$ARTIST_DIR" "e2e(artist-van-gogh): atelier stage"
  append_readme_only "$ROOT_DIR" "e2e(atelier): atelier stage"
  # Rely on default recursion at atelier level
  push_atelier_with_engine
  assert_remote_synced "$ROOT_DIR"

  local canvas_delta artist_delta root_delta
  canvas_delta="$(count_delta "$CANVAS_DIR" "$BASE_CANVAS")"
  artist_delta="$(count_delta "$ARTIST_DIR" "$BASE_ARTIST")"
  root_delta="$(count_delta "$ROOT_DIR" "$BASE_ROOT")"
  if [[ "$canvas_delta" -ne 3 ]]; then
    err "Expected canvas commit delta 3 after atelier test, got $canvas_delta"
    exit 1
  fi
  if [[ "$artist_delta" -ne 2 ]]; then
    err "Expected artist commit delta 2 after atelier test, got $artist_delta"
    exit 1
  fi
  if [[ "$root_delta" -ne 1 ]]; then
    err "Expected root commit delta 1 after atelier test, got $root_delta"
    exit 1
  fi
}

main() {
  require_cmd git
  require_cmd awk
  log "Starting E2E Git Engine tests at $(timestamp)"
  log "Root directory: $ROOT_DIR"
  require_paths
  capture_baselines

  test_canvas_sunflowers
  test_artist_van_gogh
  test_atelier_root

  log "All tests completed successfully."
}

main "$@"