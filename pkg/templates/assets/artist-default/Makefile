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

# Git push operations
push: ## Major recursive commit: push this artist and all canvases
	@atelier-cli artist push

# Maintenance
clean: ## Clean build artifacts and temporary files
	@echo "Cleaning up artist workspace..."
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete
	find . -type f -name "*.pyo" -delete
	find . -type d -name "node_modules" -exec rm -rf {} +
	find . -type d -name "build" -exec rm -rf {} +
	find . -type d -name "dist" -exec rm -rf {} +

# Artist-level canvas management
canvas-init: ## Initialize a new canvas in this artist (NAME=&lt;name&gt;)
	@atelier-cli canvas init $(NAME)

canvas-delete: ## Delete a canvas in this artist (FULL=&lt;canvas-full-name&gt;)
	@atelier-cli canvas delete $(FULL)
	
# Artist-specific targets (uncomment and customize as needed)
# add your special make commands from here