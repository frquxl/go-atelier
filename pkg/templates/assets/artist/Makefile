# Artist Makefile
# Git submodule management for artist workspaces
# Customize this file for your specific artist-level workflows

.PHONY: help submodules update status clean list-canvases

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

# Git submodule management (primary focus for artists)
submodules: ## Initialize and update all canvas submodules
	git submodule init
	git submodule update --recursive

update: ## Update all canvas submodules to latest
	git submodule update --remote --recursive

status: ## Show status of all canvas submodules
	git submodule status --recursive

list-canvases: ## List all canvas submodules in this artist
	@echo "Canvas submodules in this artist:"
	@git submodule status | awk '{print "  - " $$2}'

# Maintenance
clean: ## Clean build artifacts and temporary files
	@echo "Cleaning up artist workspace..."
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete
	find . -type f -name "*.pyo" -delete
	find . -type d -name "node_modules" -exec rm -rf {} +
	find . -type d -name "build" -exec rm -rf {} +
	find . -type d -name "dist" -exec rm -rf {} +

# Canvas management helpers (uncomment and customize as needed)
# add-canvas: ## Add a new canvas submodule
#	@echo "To add a new canvas:"
#	@echo "1. Create the canvas: atelier-cli canvas init <canvas-name>"
#	@echo "2. Or manually: git submodule add <repo-url> canvas-<canvas-name>"
#	@echo "3. Then commit: git commit -m 'feat: add canvas <canvas-name>'"

# Workflow helpers (uncomment and customize)
# build-all: ## Build all canvases in this artist
#	@echo "Building all canvases..."
#	@for dir in canvas-*; do \
#		if [ -d "$$dir" ]; then \
#			echo "Building $$dir..."; \
#			(cd "$$dir" && make build 2>/dev/null || echo "No build target in $$dir"); \
#		fi \
#	done

# test-all: ## Run tests for all canvases in this artist
#	@echo "Testing all canvases..."
#	@for dir in canvas-*; do \
#		if [ -d "$$dir" ]; then \
#			echo "Testing $$dir..."; \
#			(cd "$$dir" && make test 2>/dev/null || echo "No test target in $$dir"); \
#		fi \
#	done