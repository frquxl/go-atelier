## Vision
- Create an intuitive, metaphor-driven CLI that helps users plan, scaffold, and evolve software projects within any user-specified directory.
- Embrace the "atelier" metaphor to make professional workflows approachable.
- Deliver a fast, dependable experience implemented in Go.

## Features
- **Metaphor-Driven Interface**: Uses atelier/artist/canvas metaphors to make CLI interactions intuitive.
- **3-Level Git Submodule Architecture**: Atelier → Artists (submodules) → Canvases (submodules) for clean version control separation.
- **Nested Repository Management**: Each canvas is an independent Git repository while maintaining atelier structure.
- **Canvas Movement**: Move canvases between artists with automatic Git submodule relationship updates.
- **Canvas Cloning**: Clone canvases to different artists while preserving Git history and relationships.
- **Automatic Submodule Setup**: CLI handles complex Git submodule relationships automatically.
- **Embedded Template System**: Generates README.md, AGENTS.md, Makefile, .gitignore, and .geminiignore files from templates embedded in the binary.

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
  - Generates contextual README.md, AGENTS.md, .gitignore, and .geminiignore files at each level.
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
  - Generates contextual README.md, AGENTS.md, .gitignore, and .geminiignore files.
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
  - Generates contextual README.md, AGENTS.md, .gitignore, and .geminiignore files.
  - Commits submodule relationship to artist repository.

### canvas delete
- **Purpose**: Deletes a canvas and removes it from Git tracking.
- **Usage**: `atelier-cli canvas delete <canvas-full-name>`
- **Functionality**:
  - Must be run from within an artist directory.
  - Prompts for confirmation before deletion.
  - Deinitializes the canvas submodule, removes it from Git, and deletes the directory.
  - Commits changes to the parent artist repository.

### canvas move
- **Purpose**: Moves a canvas from one artist to another, updating Git submodules and internal paths accordingly.
- **Usage**: `atelier-cli canvas move <canvas-full-name> <new-artist-full-name>`
- **Functionality**:
  - Can be run from any directory within an atelier (automatically finds the atelier root).
  - Automatically discovers which artist currently contains the target canvas.
  - Validates that the destination artist exists and prevents naming conflicts.
  - Removes canvas from current artist's Git submodule tracking while preserving the directory.
  - Physically moves the canvas directory to the new artist.
  - Adds canvas as submodule to the new artist with proper Git tracking.
  - Updates the canvas's `.canvas` file with new artist context information.
  - Creates appropriate commit messages in both source and destination artists.
  - Provides comprehensive error handling with cleanup on failure.

### canvas clone
- **Purpose**: Clones a canvas to another artist, creating a copy with proper Git submodule relationships.
- **Usage**: `atelier-cli canvas clone <canvas-full-name> <target-artist-full-name> [new-canvas-name]`
- **Functionality**:
  - Can be run from any directory within an atelier (automatically finds the atelier root).
  - Automatically discovers which artist currently contains the source canvas.
  - Validates that the target artist exists.
  - If no new name is provided and a naming conflict exists, prompts the user for a new canvas name.
  - If a new name is provided as the third parameter, uses that name for the cloned canvas.
  - Copies the entire canvas directory to the target artist (preserving all files and Git history).
  - Updates the cloned canvas's `.canvas` file with new artist context and name information.
  - Adds the cloned canvas as submodule to the target artist with proper Git tracking.
  - Creates appropriate commit messages in the target artist.
  - Leaves the original canvas unchanged in its source artist.
  - Provides comprehensive error handling with cleanup on failure.

### atelier push
- **Purpose**: Pushes changes across the atelier hierarchy with automatic recursion.
- **Usage**: `atelier-cli push [--dry-run] [--quiet] [--force]`
- **Functionality**:
  - Must be run from within an atelier directory (detects `.atelier` marker).
  - Recursively pushes all artists and their canvases in correct order (canvases → artists → atelier).
  - Auto-commits uncommitted changes and creates single combined commits per level.
  - Supports dry-run mode to preview actions without making changes.
  - Non-interactive by default; uses the Git Push Engine for orchestration.

### artist push
- **Purpose**: Pushes changes within an artist studio with canvas recursion.
- **Usage**: `atelier-cli artist push [--dry-run] [--quiet] [--force]`
- **Functionality**:
  - Must be run from within an artist directory (detects `.artist` marker).
  - Recursively pushes all canvases first, then creates combined artist commit.
  - Auto-commits uncommitted changes and updated submodule pointers.
  - Supports dry-run mode to preview actions without making changes.
  - Uses the Git Push Engine for proper submodule handling.

### canvas push
- **Purpose**: Pushes changes for a single canvas.
- **Usage**: `atelier-cli canvas push [--dry-run] [--quiet] [--force]`
- **Functionality**:
  - Must be run from within a canvas directory (detects `.canvas` marker).
  - Pushes the canvas repository with auto-commit for uncommitted changes.
  - Supports dry-run mode to preview actions without making changes.
  - Uses the Git Push Engine for consistent behavior.
  
## Project Structure
- **`main.go`**: CLI entry point.
- **`cmd/`**: Cobra command definitions. These are thin wrappers that parse flags and arguments and call the `engine` package.
- **`pkg/`**: Internal packages containing all shared and core logic.
  - **`pkg/engine`**: Contains the core application logic for creating ateliers, artists, and canvases.
  - **`pkg/fs`**: Low-level filesystem utilities.
  - **`pkg/gitutil`**: Wrappers for executing Git commands.
  - **`pkg/templates`**: Manages the embedded boilerplate files.
  - **`pkg/push-engine`**: Git Push Engine for hierarchical repository management. See [pkg/push-engine/README.md](pkg/push-engine/README.md) and [pkg/push-engine/ENGINE-MANUAL.md](pkg/push-engine/ENGINE-MANUAL.md) for detailed documentation.
- **`test/e2e`**: Contains the end-to-end test suite for the application.

## AI Context Patterns

### Development Workflow
- **Study Code First**: Always analyze existing code to understand flow and logic before implementing changes.
- **Engine-First Development**: New core functionality should be added to the `pkg/engine` first. The `cmd` package should only contain presentation logic.
- **Iterative Implementation**: Build up features incrementally, testing at each step.
- **Update E2E Tests**: For any change in functionality, update or add to the E2E test suite in `test/e2e` to validate the behavior.
- **Iterate Until Pass**: Fix issues and re-test until `make test` and `make e2e-test` both pass.
- **Update Documentation**: Keep AGENTS.md and README.md current with implementation details.

### Testing Approach
- **Unit Tests**: Standard Go tests located alongside the code they test (primarily in the `pkg/` directory). They should be fast and focused. Run with `make test`.
- **End-to-End (E2E) Tests**: A comprehensive test suite located in `test/e2e`. These tests build the final CLI binary and execute it to verify the application's behavior from a user's perspective. They are slower but more thorough. Run with `make e2e-test`.

### Git Workflow
- **Feature Branches**: Create branches for new features when needed.
- **Conventional Commits**: Use `feat:`, `fix:`, `docs:`, `refactor:`, `test:` prefixes for commit messages.
- **Version Tagging**: Tag releases with semantic versioning (v0.1.0, v0.2.0, v1.0.0, etc.).
- **Release Documentation**: Update AGENTS.md and README.md before tagging releases.
- **AGENTS.md Refresh**: Always review and update AGENTS.md with current implementation details before any commit, especially before releases.