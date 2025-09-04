.PHONY: build clean test fmt vet run help web-dev web-build web-start

# Binary name
BINARY_NAME=atelier-cli

# Build the binary
build:
	go build -o $(BINARY_NAME) .

# Clean build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Run the binary
run: build
	./$(BINARY_NAME)

# Install dependencies
deps:
	go mod download
	go mod tidy

# Build and test
all: fmt vet test build

# E2E Testing
e2e-test:
	@echo "Running end-to-end tests..."
	./test-e2e.sh

# Web App Development
WEB_APP_DIR=user-tools/artist-gemini/canvas-web-md-editor

web-dev:
	@echo "Starting web app development server..."
	cd $(WEB_APP_DIR) && npm run dev

web-build:
	@echo "Building web app for production..."
	cd $(WEB_APP_DIR) && npm run build

web-start:
	@echo "Starting web app in production mode..."
	cd $(WEB_APP_DIR) && npm run start

# Show help
help:
	@echo "Available targets:"
	@echo "  build       - Build the CLI binary"
	@echo "  clean       - Clean CLI build artifacts"
	@echo "  test        - Run CLI tests"
	@echo "  fmt         - Format CLI code"
	@echo "  vet         - Run Go vet on CLI code"
	@echo "  run         - Build and run the CLI binary"
	@echo "  deps        - Download and tidy CLI dependencies"
	@echo "  all         - Format, vet, test, and build CLI"
	@echo "  web-dev     - Start web app development server"
	@echo "  web-build   - Build web app for production"
	@echo "  web-start   - Start web app in production mode"
	@echo "  help        - Show this help"
