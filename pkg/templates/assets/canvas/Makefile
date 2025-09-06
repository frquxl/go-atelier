# Canvas Makefile
# Generic boilerplate for any development project
# Customize these targets for your specific language/framework

.PHONY: help setup build test run clean install deps lint format docs

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

# Environment setup
setup: deps ## Setup development environment
	@echo "Setting up development environment..."
	# Add your setup commands here

install: ## Install the project
	@echo "Installing project..."
	# Add installation commands here

deps: ## Install dependencies
	@echo "Installing dependencies..."
	# Add dependency installation commands here
	# Examples:
	# pip install -r requirements.txt
	# npm install
	# go mod download
	# cargo build

# Development workflow
build: ## Build the project
	@echo "Building project..."
	# Add build commands here
	# Examples:
	# python setup.py build
	# npm run build
	# go build
	# cargo build --release

test: ## Run tests
	@echo "Running tests..."
	# Add test commands here
	# Examples:
	# python -m pytest
	# npm test
	# go test ./...
	# cargo test

run: ## Run the application
	@echo "Running application..."
	# Add run commands here
	# Examples:
	# python main.py
	# npm start
	# go run main.go
	# cargo run

# Code quality
lint: ## Run linting tools
	@echo "Running linters..."
	# Add linting commands here
	# Examples:
	# flake8 .
	# eslint .
	# golint ./...
	# cargo clippy

format: ## Format code
	@echo "Formatting code..."
	# Add formatting commands here
	# Examples:
	# black .
	# prettier --write .
	# gofmt -w .
	# cargo fmt

# Documentation
docs: ## Generate documentation
	@echo "Generating documentation..."
	# Add documentation commands here
	# Examples:
	# sphinx-build docs docs/_build
	# jsdoc -r . -d docs
	# godoc -http=:6060
	# cargo doc

# Git operations
push: ## Major commit: push this canvas (no recursion below)
	@atelier-cli canvas push

# Maintenance
clean: ## Clean build artifacts and temporary files
	@echo "Cleaning up..."
	# Add cleanup commands here
	# Examples:
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete
	find . -type f -name "*.pyo" -delete
	find . -type d -name "node_modules" -exec rm -rf {} +
	find . -type d -name "target" -exec rm -rf {} +
	find . -type d -name "build" -exec rm -rf {} +
	find . -type d -name "dist" -exec rm -rf {} +


# Canvas-specific targets (uncomment and customize as needed)
# add your special make commands from here

# Development helpers (uncomment and customize as needed)
# dev: ## Start development server
# dev-server: ## Run development server with hot reload
# docker-build: ## Build Docker image
# docker-run: ## Run in Docker container
# docker-stop: ## Stop Docker containers
# deploy: ## Deploy to production
# release: ## Create a release
# coverage: ## Generate test coverage report

# Environment-specific (uncomment and customize)
# dev-setup: deps ## Development environment setup
# prod-setup: deps ## Production environment setup
# ci: test lint ## Continuous integration pipeline