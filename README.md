# Atelier CLI

A metaphor-driven CLI tool for software project management using the atelier/artist/canvas concept.

## Installation

### For Development/Testing (Recommended)

Install globally using Go's built-in installation:

```bash
# Clone the repository
git clone <repository-url>
cd go-atelier

# Install globally
go install .

# Verify installation
atelier-cli --version
```

This will install the `atelier-cli` binary to your `$GOPATH/bin` or `$HOME/go/bin` directory, which should be in your PATH.

### Alternative: Local Build

```bash
# Build locally
go build -o atelier-cli .

# Use directly
./ateliercli --help
```

## Usage

### Initialize a New Atelier

```bash
# Create a new atelier with default artist/canvas
atelier-cli init myproject

# This creates:
# atelier-myproject/
# ├── .atelier
# ├── README.md
# ├── GEMINI.md
# ├── artist-van-gogh/
# │   ├── .artist
# │   ├── README.md
# │   ├── GEMINI.md
# │   └── canvas-sunflowers/
# │       ├── .canvas
# │       ├── README.md
# │       └── GEMINI.md
```

### Add Artists to Your Atelier

```bash
# Navigate to your atelier
cd atelier-myproject

# Add a new artist
atelier-cli artist init picasso

# This creates:
# artist-picasso/
# ├── .artist
# ├── README.md
# ├── GEMINI.md
# └── canvas-example/
#     ├── .canvas
#     ├── README.md
#     └── GEMINI.md
```

### Add Canvases to Artists

```bash
# Navigate to an artist directory
cd artist-picasso

# Add a new canvas
atelier-cli canvas init guernica

# This creates:
# canvas-guernica/
# ├── .canvas
# ├── README.md
# └── GEMINI.md
```

## Directory Structure

The CLI creates a hierarchical structure:

```
atelier-<name>/
├── .atelier           # Marks the atelier root
├── README.md         # Atelier documentation
├── GEMINI.md         # Atelier AI context
├── artist-<name>/
│   ├── .artist       # Marks an artist workspace
│   ├── README.md     # Artist documentation
│   ├── GEMINI.md     # Artist AI context
│   └── canvas-example/
│       ├── .canvas   # Marks a canvas/project
│       ├── README.md # Canvas documentation
│       └── GEMINI.md # Canvas AI context
```

## Context-Aware Commands

Commands are context-aware and only work in appropriate directories:

- `atelier-cli init` - Works anywhere
- `atelier-cli artist init` - Only works in atelier directories
- `atelier-cli canvas init` - Only works in artist directories

If you try to run a command in the wrong context, you'll get helpful error messages with suggestions.

## Features

- **Metaphor-Driven Interface**: Uses atelier/artist/canvas metaphors for intuitive project organization
- **Basic Project Scaffolding**: Creates structured directories with essential files
- **Version Control Integration**: Initializes Git repositories automatically
- **Embedded Template System**: Generates README.md and GEMINI.md files from templates embedded in the binary
- **Context-Aware Commands**: Commands validate their execution context
- **Hierarchical Organization**: atelier → artist → canvas structure

## Development

### Prerequisites

- Go 1.19 or later
- Git

### Building

```bash
# Clone the repository
git clone <repository-url>
cd go-atelier

# Install dependencies
go mod tidy

# Build
go build -o ateliercli .

# Run tests
go test ./...
```

### Project Structure

```
.
├── main.go              # CLI entry point
├── cmd/                 # Cobra command definitions
│   ├── root.go
│   ├── init.go
│   ├── artist.go
│   └── canvas.go
├── pkg/                 # Internal packages
│   ├── fs/              # Filesystem utilities (file creation)
│   └── gitutil/         # Git command utilities
├── go.mod
├── go.sum
└── README.md
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.