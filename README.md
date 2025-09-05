# Atelier CLI

A metaphor-driven CLI tool for software project management using the atelier/artist/canvas concept.

## Features

- **Metaphor-Driven Interface**: Uses atelier/artist/canvas metaphors for intuitive project organization.
- **3-Level Git Submodule Architecture**: Automatically scaffolds a nested Git repository structure (`atelier` -> `artist` -> `canvas`) for clean version control separation.
- **Context-Aware Commands**: Ensures commands like `artist init` and `canvas init` are run in the correct directory context.
- **Boilerplate Generation**: Creates useful starter files (`README.md`, `GEMINI.md`, `Makefile`, `.gitignore`) from an embedded template system.

## Prerequisites

- Go 1.19 or later
- Git

## Installation

The recommended way to install the `atelier-cli` is using the `Makefile` after cloning the repository. This handles all necessary steps.

```bash
# Clone the repository
git clone https://github.com/your-username/go-atelier.git
cd go-atelier

# Install the CLI to your GOPATH
make install

# Verify installation
atelier-cli --version
```

## Usage

### Initialize a New Atelier

```bash
# Create a new atelier with default artist (van-gogh) and canvas (sunflowers)
atelier-cli init my-project

# Create an atelier with additional default artists
atelier-cli init my-project --sketch --gallery
# This will create 'artist-van-gogh', 'artist-sketch', and 'artist-gallery'

# Create an atelier with a custom primary artist and additional default artists
atelier-cli init my-project picasso guernica --sketch
# This will create 'artist-picasso' (with 'canvas-guernica') and 'artist-sketch' (with 'canvas-example')
```

### Add a New Artist

```bash
# Navigate to your atelier directory
cd atelier-my-project

# Add a new artist (creates a default 'example' canvas inside)
atelier-cli artist init picasso
```

### Add a New Canvas

```bash
# Navigate to an artist directory
cd artist-picasso

# Add a new canvas
atelier-cli canvas init guernica
```

## Development & Testing

All common development tasks are managed through the `Makefile`.

- `make build`: Build the binary locally.
- `make test`: Run the fast unit tests.
- `make e2e-test`: Run the full Go-based end-to-end test suite.
- `make e2e-test-sh`: Run the legacy shell-based E2E tests.
- `make fmt`: Format the Go source code.

## Project Structure

```
.
├── main.go              # CLI entry point
├── cmd/                 # Cobra command definitions
├── pkg/                 # Internal packages (core logic)
│   ├── engine/          # Core application logic
│   ├── fs/              # Filesystem utilities
│   ├── gitutil/         # Git command utilities
│   └── templates/       # Embedded boilerplate files
├── test/e2e/            # End-to-end tests
├── go.mod
└── Makefile
```

## Contributing

1. Fork the repository.
2. Create a feature branch.
3. Make your changes.
4. Add or update tests.
5. Ensure `make test` and `make e2e-test` pass.
6. Submit a pull request.