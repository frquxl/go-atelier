#!/bin/bash

# Atelier CLI End-to-End Test Suite
# Tests the globally installed CLI with 3-level Git submodule architecture

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test configuration
TEST_DIR="/tmp/atelier-e2e-test"
EXPECTED_ATELIER="atelier-test-workspace"
EXPECTED_ARTIST="artist-test-artist"
EXPECTED_CANVAS="canvas-test-canvas"

# Helper functions
log_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

log_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

log_error() {
    echo -e "${RED}❌ $1${NC}"
}

log_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

assert_file_exists() {
    local file="$1"
    if [[ -f "$file" ]]; then
        log_success "File exists: $file"
    else
        log_error "File missing: $file"
        exit 1
    fi
}

assert_dir_exists() {
    local dir="$1"
    if [[ -d "$dir" ]]; then
        log_success "Directory exists: $dir"
    else
        log_error "Directory missing: $dir"
        exit 1
    fi
}

assert_file_contains() {
    local file="$1"
    local content="$2"
    if grep -q "$content" "$file"; then
        log_success "File contains '$content': $file"
    else
        log_error "File missing content '$content': $file"
        exit 1
    fi
}

assert_git_repo() {
    local dir="$1"
    if [[ -d "$dir/.git" ]]; then
        log_success "Git repository exists: $dir"
    else
        log_error "Git repository missing: $dir"
        exit 1
    fi
}

assert_submodule() {
    local parent_dir="$1"
    local submodule="$2"
    if grep -q "$submodule" "$parent_dir/.gitmodules"; then
        log_success "Submodule registered: $submodule in $parent_dir"
    else
        log_error "Submodule not registered: $submodule in $parent_dir"
        exit 1
    fi
}

# Setup
log_info "Setting up E2E test environment..."
rm -rf "$TEST_DIR"
mkdir -p "$TEST_DIR"
cd "$TEST_DIR"

# Install CLI globally
log_info "Installing Atelier CLI globally..."
cd /home/frquxl/code/go-atelier
go install .
cd "$TEST_DIR"

# Verify installation
if ! command -v atelier-cli &> /dev/null; then
    log_error "atelier-cli not found in PATH"
    exit 1
fi
log_success "atelier-cli installed successfully"

# Test 1: atelier init
log_info "Test 1: Running 'atelier-cli init test-workspace'..."
atelier-cli init test-workspace

# Verify atelier structure
assert_dir_exists "$EXPECTED_ATELIER"
assert_file_exists "$EXPECTED_ATELIER/.atelier"
assert_file_exists "$EXPECTED_ATELIER/README.md"
assert_file_exists "$EXPECTED_ATELIER/AGENTS.md"
assert_file_exists "$EXPECTED_ATELIER/Makefile"
assert_file_exists "$EXPECTED_ATELIER/.gitignore"
assert_file_exists "$EXPECTED_ATELIER/.geminiignore"
assert_file_exists "$EXPECTED_ATELIER/.gitmodules"
assert_git_repo "$EXPECTED_ATELIER"

# Check atelier content
assert_file_contains "$EXPECTED_ATELIER/.atelier" "atelier-test-workspace"
assert_file_contains "$EXPECTED_ATELIER/README.md" "Atelier Workspace"
assert_file_contains "$EXPECTED_ATELIER/AGENTS.md" "Atelier AI Context"

# Test 2: artist init
log_info "Test 2: Running 'atelier-cli artist init test-artist'..."
cd "$EXPECTED_ATELIER"
atelier-cli artist init test-artist

# Verify artist structure
assert_dir_exists "$EXPECTED_ARTIST"
assert_file_exists "$EXPECTED_ARTIST/.artist"
assert_file_exists "$EXPECTED_ARTIST/README.md"
assert_file_exists "$EXPECTED_ARTIST/AGENTS.md"
assert_file_exists "$EXPECTED_ARTIST/Makefile"
assert_file_exists "$EXPECTED_ARTIST/.gitignore"
assert_file_exists "$EXPECTED_ARTIST/.geminiignore"
assert_git_repo "$EXPECTED_ARTIST"

# Check artist content
assert_file_contains "$EXPECTED_ARTIST/.artist" "atelier-test-workspace"
assert_file_contains "$EXPECTED_ARTIST/.artist" "artist-test-artist"
assert_file_contains "$EXPECTED_ARTIST/README.md" "Artist Workspace"
assert_file_contains "$EXPECTED_ARTIST/AGENTS.md" "Artist AI Context"

# Verify submodule registration
assert_submodule "." "$EXPECTED_ARTIST"

# Test 3: canvas init
log_info "Test 3: Running 'atelier-cli canvas init test-canvas'..."
cd "$EXPECTED_ARTIST"
atelier-cli canvas init test-canvas

# Verify canvas structure
assert_dir_exists "$EXPECTED_CANVAS"
assert_file_exists "$EXPECTED_CANVAS/.canvas"
assert_file_exists "$EXPECTED_CANVAS/README.md"
assert_file_exists "$EXPECTED_CANVAS/AGENTS.md"
assert_file_exists "$EXPECTED_CANVAS/.gitignore"
assert_file_exists "$EXPECTED_CANVAS/.geminiignore"
assert_file_exists "$EXPECTED_CANVAS/Makefile"
assert_git_repo "$EXPECTED_CANVAS"

# Check canvas content
assert_file_contains "$EXPECTED_CANVAS/.canvas" "atelier-test-workspace"
assert_file_contains "$EXPECTED_CANVAS/.canvas" "artist-test-artist"
assert_file_contains "$EXPECTED_CANVAS/.canvas" "canvas-test-canvas"
assert_file_contains "$EXPECTED_CANVAS/README.md" "Project Canvas"
assert_file_contains "$EXPECTED_CANVAS/AGENTS.md" "Canvas AI Context"

# Verify submodule registration
assert_submodule "." "$EXPECTED_CANVAS"

# Test 4: Git submodule relationships
log_info "Test 4: Verifying Git submodule relationships..."

# Check that atelier tracks artist
cd "$TEST_DIR/$EXPECTED_ATELIER"
if git submodule status | grep -q "$EXPECTED_ARTIST"; then
    log_success "Atelier properly tracks artist submodule"
else
    log_error "Atelier does not track artist submodule"
    exit 1
fi

# Check that artist tracks canvas
cd "$EXPECTED_ARTIST"
if git submodule status | grep -q "$EXPECTED_CANVAS"; then
    log_success "Artist properly tracks canvas submodule"
else
    log_error "Artist does not track canvas submodule"
    exit 1
fi

# Test 5: Context awareness
log_info "Test 5: Testing context-aware commands..."

# Should fail when not in atelier
cd "$TEST_DIR"
if atelier-cli artist init should-fail 2>&1 | grep -q "atelier directory"; then
    log_success "Context awareness works: artist init fails outside atelier"
else
    log_error "Context awareness broken: artist init should fail outside atelier"
    exit 1
fi

# Should fail when not in artist
cd "$EXPECTED_ATELIER"
if atelier-cli canvas init should-fail 2>&1 | grep -q "artist directory"; then
    log_success "Context awareness works: canvas init fails outside artist"
else
    log_error "Context awareness broken: canvas init should fail outside artist"
    exit 1
fi

# Ensure we're back at the test root for path-based assertions
cd "$TEST_DIR"

# Test 6: Template content verification
log_info "Test 6: Verifying template content..."

# Check atelier README has expected sections
assert_file_contains "$EXPECTED_ATELIER/README.md" "Atelier Workspace"
# Optional extended checks (uncomment if needed)
# assert_file_contains "$EXPECTED_ATELIER/README.md" "🏗️ Architecture Overview"
# assert_file_contains "$EXPECTED_ATELIER/README.md" "🚀 Getting Started"

# Check artist README has expected sections (paths relative to TEST_DIR)
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/README.md" "🎯 Artist Purpose"
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/README.md" "👨‍🎨"

# Check canvas README has expected sections (paths relative to TEST_DIR)
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/$EXPECTED_CANVAS/README.md" "🎯 Project Overview"
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/$EXPECTED_CANVAS/README.md" "🖼️"

# Test 7: Hierarchical context in marker files
log_info "Test 7: Verifying hierarchical context in marker files..."

# Verify .atelier contains full directory name
assert_file_contains "$EXPECTED_ATELIER/.atelier" "atelier-test-workspace"

# Verify .artist contains full hierarchy
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/.artist" "atelier-test-workspace"
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/.artist" "artist-test-artist"

# Verify .canvas contains complete hierarchy
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/$EXPECTED_CANVAS/.canvas" "atelier-test-workspace"
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/$EXPECTED_CANVAS/.canvas" "artist-test-artist"
assert_file_contains "$EXPECTED_ATELIER/$EXPECTED_ARTIST/$EXPECTED_CANVAS/.canvas" "canvas-test-canvas"

log_success "Hierarchical context properly stored in all marker files"

# Test 8: Directory structure completeness
log_info "Test 8: Verifying complete directory structure..."

# Expected structure:
# atelier-test-workspace/
# ├── .atelier
# ├── .git/
# ├── .gitmodules
# ├── README.md
# ├── AGENTS.md
# ├── artist-test-artist/
# │   ├── .artist
# │   ├── .git/
# │   ├── .gitmodules
# │   ├── README.md
# │   ├── AGENTS.md
# │   └── canvas-test-canvas/
# │       ├── .canvas
# │       ├── .git/
# │       ├── README.md
# │       └── AGENTS.md

# Count files and directories
atelier_files=$(find "$EXPECTED_ATELIER" -type f | wc -l)
atelier_dirs=$(find "$EXPECTED_ATELIER" -type d | wc -l)

if [[ $atelier_files -ge 17 ]]; then
    log_success "Atelier has expected number of files ($atelier_files)"
else
    log_error "Atelier has too few files ($atelier_files)"
    exit 1
fi

if [[ $atelier_dirs -ge 6 ]]; then
    log_success "Atelier has expected number of directories ($atelier_dirs)"
else
    log_error "Atelier has too few directories ($atelier_dirs)"
    exit 1
fi

# Cleanup
log_info "Cleaning up test environment..."
cd /
rm -rf "$TEST_DIR"

# Final success
echo ""
log_success "🎉 ALL E2E TESTS PASSED!"
log_success "✅ Atelier CLI with 3-level Git submodule architecture is working perfectly!"
echo ""
log_info "Test Summary:"
echo "  ✅ Global installation works"
echo "  ✅ atelier init creates proper structure"
echo "  ✅ artist init creates submodule"
echo "  ✅ canvas init creates submodule"
echo "  ✅ Git repositories properly initialized"
echo "  ✅ Submodule relationships correct"
echo "  ✅ Context-aware commands work"
echo "  ✅ Templates generate correct content"
echo "  ✅ Hierarchical context in marker files"
echo "  ✅ All files and directories created"
echo ""
log_info "The Atelier CLI tested just fine! 🚀"