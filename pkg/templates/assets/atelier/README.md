# Atelier Workspace

Welcome to your Atelier workspace! 🎨

This is the root directory of your atelier - your personal studio for software development using a 3-level Git submodule architecture.

## 🏗️ Architecture Overview

Your atelier uses a hierarchical structure:
- **Atelier** (this level): Main workspace and project container
- **Artists** (Git submodules): Thematic groupings of related projects
- **Canvases** (Git submodules): Individual software projects

## 🚀 Getting Started

1. **Explore artists**: `ls artist-*` to see available artist workspaces
2. **Work on projects**: `cd artist-name/canvas-name` to enter a canas (project)
3. **Set up remote repositories**: Work with your AI pair programmer to push the entire atelier (including all submodules) to private GitHub repositories using SSH. Example prompt: "this is a new project pre configured with submodules, can you get it all remoted private using gh cli, please use the ssh version"
4. **Develop independently**: Each canvas is its own Git repository

## 📁 Directory Structure

```
atelier/
├── .git/              # Main atelier repository
├── .gitmodules       # Tracks artist submodules
├── README.md         # This file (human guide)
├── AGENTS.md         # AI context for pair programming
└── artist-*/         # Artist workspaces (Git submodules)
    ├── .git/         # Artist's Git repository
    ├── .gitmodules   # Tracks canvas submodules
    ├── README.md     # Artist documentation
    ├── AGENTS.md     # Artist AI context
    └── canvas-*/     # Project canvases (Git submodules)
        ├── .git/     # Canvas's Git repository
        ├── README.md # Project documentation
        └── AGENTS.md # Project AI context
```

## 🔧 Development Workflow

- Each canvas is an **independent Git repository**
- Artists organize related canvases thematically
- The atelier tracks the overall project structure
- Use Git submodules for clean version control separation

## 🎨 Working with Your Atelier

- **Add artists**: Create new artist workspaces for different themes
- **Add canvases**: Create new projects within artists
- **Version control**: Each level has its own Git history
- **Independence**: Projects don't interfere with each other

## 📚 Documentation

- **README.md** (this file): Human-readable workspace guide
- **AGENTS.md**: AI pair programming context and patterns

### Git Workflow

- Day-to-day: use regular Git in each repo (canvas, artist, or atelier) as you normally would:
  - Example: cd artist-van-gogh/canvas-sunflowers && git add -A && git commit -m "feat: changes" && git push
- Major recursive commit (end-of-day, multiple canvases/artists touched):
  - From the atelier root, run: make push
  - This invokes the CLI atelier-cli push to:
    - Recurse through all artists and their canvases
    - Commit and push canvases first (if changes)
    - Commit and push artists with updated canvas pointers (single combined commit per artist)
    - Commit and push the atelier with updated artist pointers (single combined commit)
- Notes:
  - AUTO_COMMIT_DEFAULT=true will auto-stage and auto-commit working tree changes and pointer updates.
  - The CLI is non-interactive by default and designed for major roll-ups; prefer manual Git for normal incremental commits.

  ### Atelier Commands at this level
- Push the entire workspace recursively (canvases → artists → atelier):
  - CLI: `atelier-cli push [--dry-run] [--quiet] [--force]`
  - Make: `make push`
- Manage artists from the atelier root:
  - Init a new artist: `atelier-cli artist init &lt;artist-name&gt;`
  - Delete an artist (requires full directory name, e.g., artist-van-gogh): `atelier-cli artist delete &lt;artist-full-name&gt;`
  - Make equivalents:
    - `make artist-init NAME=van-gogh`
    - `make artist-delete FULL=artist-van-gogh`

Notes:
- The push command is recursive when run from the atelier and will:
  1) Commit/push canvases that changed,
  2) Commit/push artists with updated canvas pointers,
  3) Commit/push the atelier with updated artist pointers.
- Commands are scope-aware: they must be run at the correct level (atelier, artist, canvas) per the CLI’s cobra validation.

Happy creating! 🎨✨