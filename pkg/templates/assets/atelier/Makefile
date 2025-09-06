# Atelier Makefile
# Customize this file for your specific atelier environment and workflows

.PHONY: help setup clean update submodules status

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

# Setup and initialization
setup: ## Initial setup for the atelier
	@echo "Setting up atelier environment..."
	# Add your setup commands here

# Git submodule management
submodules: ## Initialize and update all git submodules
	git submodule init
	git submodule update --recursive

update: ## Update all submodules to latest
	git submodule update --remote --recursive

status: ## Show status of all submodules
	git submodule status --recursive

# Git push operations
push: ## Major recursive commit: push atelier, all artists, and all canvases
	@atelier-cli push

# Development workflow
build: ## Build all projects in the atelier
	@echo "Building all projects..."
	# Add build commands for each canvas/project

test: ## Run tests for all projects
	@echo "Running tests..."
	# Add test commands for each canvas/project

run: ## Run the main application
	@echo "Running main application..."
	# Add run command for your primary canvas

# Maintenance
clean: ## Clean build artifacts and temporary files
	@echo "Cleaning up..."
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete
	find . -type f -name "*.pyo" -delete
	find . -type f -name "*.pyd" -delete
	find . -type d -name ".pytest_cache" -exec rm -rf {} +
	find . -type d -name ".coverage" -exec rm -rf {} +
	find . -type d -name "node_modules" -exec rm -rf {} +
	find . -type d -name "build" -exec rm -rf {} +
	find . -type d -name "dist" -exec rm -rf {} +

# Atelier-level artist management
artist-init: ## Initialize a new artist (NAME=&lt;name&gt;)
	@atelier-cli artist init $(NAME)

artist-delete: ## Delete an artist (FULL=&lt;artist-full-name&gt;)
	@atelier-cli artist delete $(FULL)

# Atelier-specific targets (uncomment and customize as needed)
# add your special make commands from here