
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
- **Purpose**: Initializes a new atelier workspace by creating the basic skeleton directory structure in the specified directory.
- **Usage**: `atelier init <atelier-name> [<artist-name> <canvas-name>]`
- **Functionality**:
  - Creates directory `atelier-<atelier-name>` (required first argument).
  - If no artist/canvas provided, defaults to `van-gogh` and `sunflowers`.
  - Within it, generates the subdirectories: `atelier-<name>/<artist-name>/<canvas-name>`.
  - Initializes a Git repository in the atelier directory.
  - Creates boilerplate files: `README.md` and `GEMINI.md` in each directory, drawn from templates.
  
  ### artist init
  - **Purpose**: Creates a new artist studio within the existing atelier.
  - **Usage**: `artist init <artist-name>`
  - **Functionality**:
    - Creates the subdirectory `atelier/<artist-name>` if it doesn't exist.
    - Within it, generates the subdirectory `atelier/<artist-name>/canvas`.
    - Creates boilerplate files: `README.md` and `GEMINI.md` in each directory (`<artist-name>`, `canvas`), drawn from templates.
  
  ## Project Structure
The skeleton structure created by `atelier init`:

- `atelier/`: Represents the artist's studio - the root workspace for the project.
  - `README.md`: Template-based readme for the atelier.
  - `GEMINI.md`: Template-based AI context file for the atelier.
  - `van-gogh/`: The artist's personal area - contains tools, configurations, and personal workspace.
    - `README.md`: Template-based readme for the artist area.
    - `GEMINI.md`: Template-based AI context file for the artist area.
    - `sunflowers/`: The canvas - the main project area where the actual development work (code, files) takes place.
      - `README.md`: Template-based readme for the canvas.
      - `GEMINI.md`: Template-based AI context file for the canvas.

This structure embraces the atelier metaphor to organize software projects intuitively.

## Usage Examples

- Initialize a new atelier with default artist/canvas: `atelier init myproject`
- Initialize with custom artist and canvas: `atelier init myproject picasso guernica`
- Add a new artist to existing atelier: `artist init monet`

## AI Context Patterns

### Development Workflow
- **Iterative Development**: Build up features incrementally, testing at each step
- **Waypoint Commits**: Use `git add .` and `git commit` at major milestones, not every change
- **Simple Makefile**: Use `make build`, `make test`, `make run` for common tasks
- **Keep it Simple**: Focus on MVP functionality, avoid over-engineering

### Code Organization
- **Cobra CLI Framework**: Use idiomatic Go patterns with Cobra for command structure
- **Command Separation**: Keep commands in separate files under `cmd/` directory
- **Template Generation**: Simple string-based templates for boilerplate files
- **Error Handling**: Basic error checking with descriptive messages

### Testing Approach
- **Manual Testing**: Test CLI commands directly with `./atelier-cli <command>`
- **Build Verification**: Use `make build` to ensure code compiles
- **Functional Testing**: Create test directories and verify output structures

### Git Workflow
- **Feature Branches**: Create branches for new features when needed
- **Conventional Commits**: Use `feat:`, `fix:`, `docs:` prefixes for commit messages
- **Version Tagging**: Tag releases with semantic versioning (v0.1.0, v1.0.0, etc.)

