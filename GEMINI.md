
## Vision
- Create an intuitive, metaphor-driven CLI that helps users plan, scaffold, and evolve software projects within any user-specified directory.
- Embrace the "atelier" metaphor to make professional workflows approachable.
- Deliver a fast, dependable experience implemented in Go.

## Features
- **Metaphor-Driven Interface**: Uses atelier/artist/canvas metaphors to make CLI interactions intuitive.
- **3-Level Git Submodule Architecture**: Atelier → Artists (submodules) → Canvases (submodules) for clean version control separation.
- **Nested Repository Management**: Each canvas is an independent Git repository while maintaining atelier structure.
- **Automatic Submodule Setup**: CLI handles complex Git submodule relationships automatically.
- **Embedded Template System**: Generates README.md and GEMINI.md files from templates embedded in the binary.

## Commands

### atelier init
- **Purpose**: Initializes a new atelier workspace with 3-level Git submodule architecture.
- **Usage**: `atelier-cli init <atelier-name> [<artist-name> <canvas-name>]`
- **Functionality**:
  - Creates directory `atelier-<atelier-name>` as main Git repository.
  - If no artist/canvas provided, defaults to `van-gogh` and `sunflowers`.
  - Creates `artist-<artist-name>` as Git submodule of atelier.
  - Creates `canvas-<canvas-name>` as Git submodule of artist.
  - Sets up `.gitmodules` files to track submodule relationships.
  - Creates marker files: `.atelier`, `.artist`, and `.canvas` in respective directories.
  - Generates contextual README.md and GEMINI.md files at each level.
  - Commits all changes with proper Git submodule setup.

### artist init
- **Purpose**: Creates a new artist studio as a Git submodule within an existing atelier.
- **Usage**: `atelier-cli artist init <artist-name>`
- **Functionality**:
  - Must be run from within an atelier directory (detects `.atelier` marker).
  - Creates `artist-<artist-name>` as Git repository with initial commit.
  - Adds artist as submodule to parent atelier repository.
  - Creates default `canvas-example` as submodule of the artist.
  - Sets up `.gitmodules` file in artist to track canvas submodules.
  - Creates marker files: `.artist` and `.canvas` in respective directories.
  - Generates contextual README.md and GEMINI.md files.
  - Commits submodule relationships to both atelier and artist repositories.

### canvas init
- **Purpose**: Creates a new canvas as a Git submodule within an existing artist studio.
- **Usage**: `atelier-cli canvas init <canvas-name>`
- **Functionality**:
  - Must be run from within an artist directory (detects `.artist` marker).
  - Creates `canvas-<canvas-name>` as Git repository with initial commit.
  - Adds canvas as submodule to parent artist repository.
  - Updates artist's `.gitmodules` file to track the new canvas.
  - Creates marker file `.canvas` in the canvas directory.
  - Generates contextual README.md and GEMINI.md files.
  - Commits submodule relationship to artist repository.
  
  ## Project Structure
  The 3-level Git submodule architecture created by `atelier-cli init`:
  
  - `atelier-<name>/`: **Main Git Repository** (atelier root)
    - `.git/`: Atelier's Git repository
    - `.gitmodules`: Tracks artist submodules
    - `.atelier`: Marker file identifying this as an atelier root
    - `README.md`: Atelier overview and project documentation
    - `GEMINI.md`: AI context file for the atelier
    - `artist-<name>/`: **Git Submodule** (artist workspace)
      - `.git/`: Artist's Git repository (submodule)
      - `.gitmodules`: Tracks canvas submodules
      - `.artist`: Marker file identifying this as an artist workspace
      - `README.md`: Artist-specific documentation
      - `GEMINI.md`: AI context file for the artist
      - `canvas-<name>/`: **Git Submodule** (project workspace)
        - `.git/`: Canvas's Git repository (submodule)
        - `.canvas`: Marker file identifying this as a canvas/project area
        - `README.md`: Project-specific documentation
        - `GEMINI.md`: AI context file for the canvas
        - `src/`, `docs/`, etc.: Your actual project files
  
  ### Directory Markers
  - **`.atelier`**: Identifies atelier root directories
  - **`.artist`**: Identifies artist workspace directories
  - **`.canvas`**: Identifies canvas/project directories
  
  These marker files enable context-aware command execution and help prevent commands from running in incorrect directories.
  
  ## Git Submodule Management
  
  ### Understanding the 3-Level Architecture
  - **Atelier Level**: Main repository tracking overall project structure
  - **Artist Level**: Submodules of atelier, can contain multiple canvases
  - **Canvas Level**: Submodules of artists, independent project repositories
  
  ### Working with Submodules
  ```bash
  # Clone entire atelier with all submodules
  git clone --recursive https://github.com/user/atelier-myproject.git
  
  # Update all submodules to latest
  git submodule update --init --recursive
  
  # Work on a specific canvas
  cd artist-picasso/canvas-guernica
  git checkout -b feature/new-feature
  # Make changes...
  git commit -m "feat: add new feature"
  
  # Update parent repositories
  cd ../..  # Back to atelier
  git add artist-picasso/canvas-guernica
  git commit -m "feat: update guernica canvas"
  ```
  
  ### Best Practices
  - **Always commit submodule changes**: When you update a canvas, commit the new reference in the parent repository
  - **Use recursive operations**: Clone with `--recursive` and update with `--recursive`
  - **Keep submodules clean**: Each level should only track its own files
  - **Version pinning**: Atelier can pin artists/canvases to specific versions

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
- **Cobra CLI Framework**: Use idiomatic Go patterns with Cobra for command structure.
- **Command Separation**: Keep commands in separate files under the `cmd/` directory.
- **Internal Packages**: Shared logic is abstracted into internal packages under `pkg/`.
  - `pkg/fs`: Handles all filesystem operations, like creating directories and files.
  - `pkg/gitutil`: Wraps all `git` command executions for consistent error handling and execution.
- **Embedded Templates**: Templates are embedded directly into the binary using `go:embed` for reliable content generation at runtime.
- **Error Handling**: Commands return errors up to the root, ensuring proper exit codes and enabling deferred cleanup functions for failed initializations.
- **Marker Files**: Use `.atelier`, `.artist`, `.canvas` files for directory context detection.

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
- **3-Level Submodule Architecture**: Automatic Git submodule setup and management
- **Helpful Errors**: Error messages include available options and next steps
- **Consistent Naming**: `atelier-<name>`, `artist-<name>`, `canvas-<name>` format
- **Default Values**: Sensible defaults (van-gogh/sunflowers) for quick starts
- **Template Content**: Auto-generated contextual content for README/GEMINI files
- **Submodule Workflow**: CLI handles complex Git submodule relationships transparently

