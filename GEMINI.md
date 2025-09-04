
## Vision
- Create an intuitive, metaphor-driven CLI that helps users plan, scaffold, and evolve software projects within any user-specified directory.
- Embrace the "atelier" metaphor to make professional workflows approachable.
- Deliver a fast, dependable experience implemented in Go.

## Features
- **Metaphor-Driven Interface**: Uses atelier/artist/canvas metaphors to make CLI interactions intuitive.
- **Basic Project Scaffolding**: Creates a vanilla project structure with essential directories and boilerplate files.
- **Version Control Integration**: Initializes a Git repository for basic version tracking.
- **Template-Based Boilerplate Generation**: Generates README.md and GEMINI.md files in each directory from predefined templates.

## Commands

### atelier init
- **Purpose**: Initializes a new atelier workspace with the complete directory structure.
- **Usage**: `atelier-cli init <atelier-name> [<artist-name> <canvas-name>]`
- **Functionality**:
  - Creates directory `atelier-<atelier-name>` (required first argument).
  - If no artist/canvas provided, defaults to `van-gogh` and `sunflowers`.
  - Within it, generates the subdirectories: `artist-<artist-name>/canvas-example`.
  - Initializes a Git repository in the atelier directory.
  - Creates marker file `.atelier` in the root.
  - Creates boilerplate files: `README.md` and `GEMINI.md` in each directory with contextual content.
  - Creates marker files: `.artist` and `.canvas` in respective directories.

### artist init
- **Purpose**: Creates a new artist studio within an existing atelier.
- **Usage**: `atelier-cli artist init <artist-name>`
- **Functionality**:
  - Must be run from within an atelier directory (detects `.atelier` marker).
  - Creates the subdirectory `artist-<artist-name>` if it doesn't exist.
  - Within it, generates the subdirectory `canvas-example`.
  - Creates marker files: `.artist` and `.canvas` in respective directories.
  - Creates boilerplate files: `README.md` and `GEMINI.md` in each directory with contextual content.

### canvas init
- **Purpose**: Creates a new canvas within an existing artist studio.
- **Usage**: `atelier-cli canvas init <canvas-name>`
- **Functionality**:
  - Must be run from within an artist directory (detects `.artist` marker).
  - Creates the subdirectory `canvas-<canvas-name>` if it doesn't exist.
  - Creates marker file `.canvas` in the canvas directory.
  - Creates boilerplate files: `README.md` and `GEMINI.md` with contextual content.
  
  ## Project Structure
  The skeleton structure created by `atelier-cli init`:
  
  - `atelier-<name>/`: Root workspace directory (represents the artist's studio).
    - `.atelier`: Marker file identifying this as an atelier root.
    - `README.md`: Auto-generated readme with atelier overview.
    - `GEMINI.md`: AI context file for the atelier.
    - `artist-<name>/`: Artist workspace directory.
      - `.artist`: Marker file identifying this as an artist workspace.
      - `README.md`: Auto-generated readme for the artist.
      - `GEMINI.md`: AI context file for the artist.
      - `canvas-example/`: Default canvas directory (represents the project workspace).
        - `.canvas`: Marker file identifying this as a canvas/project area.
        - `README.md`: Auto-generated readme for the canvas.
        - `GEMINI.md`: AI context file for the canvas.
  
  ### Directory Markers
  - **`.atelier`**: Identifies atelier root directories
  - **`.artist`**: Identifies artist workspace directories
  - **`.canvas`**: Identifies canvas/project directories
  
  These marker files enable context-aware command execution and help prevent commands from running in incorrect directories.

## Installation

### Global Installation (Recommended)
```bash
# Install globally using Go
go install .

# Or use the installation script
./install.sh

# Verify installation
atelier-cli --version
```

### Local Development
```bash
# Build locally
go build -o ateliercli .

# Use directly
./ateliercli --help
```

## Usage Examples

### Basic Workflow
```bash
# 1. Initialize a new atelier with defaults
atelier-cli init myproject

# 2. Navigate to the created atelier
cd atelier-myproject

# 3. Add a new artist
atelier-cli artist init picasso

# 4. Navigate to the artist
cd artist-picasso

# 5. Add a new canvas
atelier-cli canvas init guernica
```

### Advanced Usage
- Initialize with custom names: `atelier-cli init workshop dali persistence`
- Add multiple artists: `atelier-cli artist init monet && atelier-cli artist init vangogh`
- Add multiple canvases: `atelier-cli canvas init waterlilies && atelier-cli canvas init sunflowers`

### Context-Aware Commands
Commands automatically detect their execution context:
- `atelier-cli init` works anywhere
- `atelier-cli artist init` requires being in an atelier directory
- `atelier-cli canvas init` requires being in an artist directory

## AI Context Patterns

### Development Workflow
- **Iterative Development**: Build up features incrementally, testing at each step
- **Waypoint Commits**: Use `git add .` and `git commit` at major milestones, not every change
- **Simple Makefile**: Use `make build`, `make test`, `make run` for common tasks
- **Keep it Simple**: Focus on MVP functionality, avoid over-engineering
- **Global Installation**: Design for `go install` compatibility from day one

### Code Organization
- **Cobra CLI Framework**: Use idiomatic Go patterns with Cobra for command structure
- **Command Separation**: Keep commands in separate files under `cmd/` directory
- **Template Generation**: Built-in default content generation (no external files needed)
- **Error Handling**: Context-aware error messages with helpful suggestions
- **Marker Files**: Use `.atelier`, `.artist`, `.canvas` files for directory context detection

### Testing Approach
- **Manual Testing**: Test CLI commands directly with `atelier-cli <command>`
- **Build Verification**: Use `go build` to ensure code compiles
- **Functional Testing**: Create test directories and verify output structures
- **Global Testing**: Test after `go install` to ensure global functionality works
- **Cross-Directory Testing**: Test commands from different directory contexts

### Git Workflow
- **Feature Branches**: Create branches for new features when needed
- **Conventional Commits**: Use `feat:`, `fix:`, `docs:` prefixes for commit messages
- **Version Tagging**: Tag releases with semantic versioning (v0.1.0, v0.2.0, v1.0.0, etc.)
- **Release Documentation**: Update GEMINI.md and README.md before tagging releases
- **GEMINI.md Refresh**: Always review and update GEMINI.md with current implementation details before any commit, especially before releases

### CLI Design Patterns
- **Context Awareness**: Commands validate execution context using marker files
- **Helpful Errors**: Error messages include available options and next steps
- **Consistent Naming**: `atelier-<name>`, `artist-<name>`, `canvas-<name>` format
- **Default Values**: Sensible defaults (van-gogh/sunflowers) for quick starts
- **Template Content**: Auto-generated contextual content for README/GEMINI files

