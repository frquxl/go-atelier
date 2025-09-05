## Vision
- Create an intuitive, metaphor-driven CLI that helps users plan, scaffold, and evolve software projects within any user-specified directory.
- Embrace the "atelier" metaphor to make professional workflows approachable.
- Deliver a fast, dependable experience implemented in Go.

## Features
- **Metaphor-Driven Interface**: Uses atelier/artist/canvas metaphors to make CLI interactions intuitive.
- **3-Level Git Submodule Architecture**: Atelier → Artists (submodules) → Canvases (submodules) for clean version control separation.
- **Nested Repository Management**: Each canvas is an independent Git repository while maintaining atelier structure.
- **Automatic Submodule Setup**: CLI handles complex Git submodule relationships automatically.
- **Embedded Template System**: Generates README.md, GEMINI.md, Makefile, .gitignore, and .geminiignore files from templates embedded in the binary.

## Commands

### atelier init
- **Purpose**: Initializes a new atelier workspace with 3-level Git submodule architecture.
- **Usage**: `atelier-cli init <atelier-name> [<artist-name> <canvas-name>] [--sketch] [--gallery]`
- **Functionality**:
  - Creates directory `atelier-<atelier-name>` as main Git repository.
  - If no artist/canvas provided, defaults to `van-gogh` and `sunflowers`.
  - If `--sketch` flag is used, also creates `artist-sketch` with a default `canvas-example`.
  - If `--gallery` flag is used, also creates `artist-gallery` with a default `canvas-example`.
  - Sets up `.gitmodules` files to track submodule relationships.
  - Creates marker files: `.atelier`, `.artist`, and `.canvas` in respective directories.
  - Generates contextual README.md, GEMINI.md, .gitignore, and .geminiignore files at each level.
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
  - Generates contextual README.md, GEMINI.md, .gitignore, and .geminiignore files.
  - Commits submodule relationships to both atelier and artist repositories.

### artist delete
- **Purpose**: Deletes an artist studio and removes it from Git tracking.
- **Usage**: `atelier-cli artist delete <artist-full-name>`
- **Functionality**:
  - Must be run from within an atelier directory.
  - Prompts for confirmation before deletion.
  - Deinitializes the artist submodule, removes it from Git, and deletes the directory.
  - Commits changes to the parent atelier repository.

### canvas init
- **Purpose**: Creates a new canvas as a Git submodule within an existing artist studio.
- **Usage**: `atelier-cli canvas init <canvas-name>`
- **Functionality**:
  - Must be run from within an artist directory (detects `.artist` marker).
  - Creates `canvas-<canvas-name>` as Git repository with initial commit.
  - Adds canvas as submodule to parent artist repository.
  - Updates artist's `.gitmodules` file to track the new canvas.
  - Creates marker file `.canvas` in the canvas directory.
  - Generates contextual README.md, GEMINI.md, .gitignore, and .geminiignore files.
  - Commits submodule relationship to artist repository.

### canvas delete
- **Purpose**: Deletes a canvas and removes it from Git tracking.
- **Usage**: `atelier-cli canvas delete <canvas-full-name>`
- **Functionality**:
  - Must be run from within an artist directory.
  - Prompts for confirmation before deletion.
  - Deinitializes the canvas submodule, removes it from Git, and deletes the directory.
  - Commits changes to the parent artist repository.
  
## Project Structure
- **`main.go`**: CLI entry point.
- **`cmd/`**: Cobra command definitions. These are thin wrappers that parse flags and arguments and call the `engine` package.
- **`pkg/`**: Internal packages containing all shared and core logic.
  - **`pkg/engine`**: Contains the core application logic for creating ateliers, artists, and canvases.
  - **`pkg/fs`**: Low-level filesystem utilities.
  - **`pkg/gitutil`**: Wrappers for executing Git commands.
  - **`pkg/templates`**: Manages the embedded boilerplate files.
- **`test/e2e`**: Contains the end-to-end test suite for the application.

## AI Context Patterns

### Development Workflow
- **Study Code First**: Always analyze existing code to understand flow and logic before implementing changes.
- **Engine-First Development**: New core functionality should be added to the `pkg/engine` first. The `cmd` package should only contain presentation logic.
- **Iterative Implementation**: Build up features incrementally, testing at each step.
- **Update E2E Tests**: For any change in functionality, update or add to the E2E test suite in `test/e2e` to validate the behavior.
- **Iterate Until Pass**: Fix issues and re-test until `make test` and `make e2e-test` both pass.
- **Update Documentation**: Keep GEMINI.md and README.md current with implementation details.

### Testing Approach
- **Unit Tests**: Standard Go tests located alongside the code they test (primarily in the `pkg/` directory). They should be fast and focused. Run with `make test`.
- **End-to-End (E2E) Tests**: A comprehensive test suite located in `test/e2e`. These tests build the final CLI binary and execute it to verify the application's behavior from a user's perspective. They are slower but more thorough. Run with `make e2e-test`.

### Git Workflow
- **Feature Branches**: Create branches for new features when needed.
- **Conventional Commits**: Use `feat:`, `fix:`, `docs:`, `refactor:`, `test:` prefixes for commit messages.
- **Version Tagging**: Tag releases with semantic versioning (v0.1.0, v0.2.0, v1.0.0, etc.).
- **Release Documentation**: Update GEMINI.md and README.md before tagging releases.
- **GEMINI.md Refresh**: Always review and update GEMINI.md with current implementation details before any commit, especially before releases.