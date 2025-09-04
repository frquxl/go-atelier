
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
- **Usage**: `atelier init [<artist-name> <canvas-name>]`
- **Functionality**:
  - If no arguments are provided, defaults to creating `van-gogh` as the artist and `sunflowers` as the canvas.
  - Creates the directory `atelier` if it doesn't exist.
  - Within it, generates the subdirectories: `atelier/<artist-name>/<canvas-name>`.
  - Initializes a Git repository in the `atelier` directory.
  - Creates boilerplate files: `README.md` and `GEMINI.md` in each directory (`atelier`, `<artist-name>`, `<canvas-name>`), drawn from templates.
  
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

- Initialize a new atelier with the default example: `atelier init`
- Initialize a new atelier with custom artist and canvas: `atelier init picasso guernica`
- Add a new artist to the existing atelier: `artist init monet`

