.PHONY: build clean test fmt vet run help

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

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"
	@echo "  fmt      - Format code"
	@echo "  vet      - Run go vet"
	@echo "  run      - Build and run the binary"
	@echo "  deps     - Download and tidy dependencies"
	@echo "  all      - Format, vet, test, and build"
	@echo "  help     - Show this help"