# Gallery Artist Workspace

Welcome to your Gallery Artist workspace! 🖼️

This artist workspace contains multiple project canvases, each representing a polished, production-ready software project. Think of this as your curated exhibition of finished works.

## 🎯 Artist Purpose

**Gallery Artist workspaces** are for:
- Showcasing polished and production-ready applications.
- Projects that require high levels of quality, testing, and documentation.
- Maintaining stable versions of deployed software.
- Collaborating on projects intended for public release or long-term support.

## 📁 Canvas Projects

This artist contains the following project canvases:
```
artist/
├── .git/                    # Artist's Git repository (submodule)
├── .gitmodules            # Tracks canvas submodules
├── README.md              # This artist documentation
├── GEMINI.md              # AI context for this artist
└── canvas-*/              # Project canvases (Git submodules)
    ├── .git/              # Canvas's Git repository
    ├── README.md          # Project documentation
    ├── GEMINI.md          # Project AI context
    ├── tests/             # Test files
```

## 🚀 Working with Canvases

Each canvas in this artist workspace is:
- ✅ **Independent Git repository** - develop without affecting other projects
- ✅ **Isolated environment** - own dependencies and configurations
- ✅ **Version controlled** - track changes and collaborate
- ✅ **Self-contained** - complete project with its own documentation

## 🔄 Development Workflow

1. **Choose a canvas**: `cd canvas-project-name`
2. **Work independently**: Each canvas has its own Git history
3. **Commit changes**: `git add . && git commit -m "feat: your changes"`
4. **Push updates**: `git push origin main`

## 🎨 Artist Philosophy

Think of yourself as a gallery artist:
- **Atelier**: The entire workspace/studio
- **Artist**: Your curated exhibition space (this level)
- **Canvas**: Individual masterpieces ready for display

Each canvas represents a complete, independent project that you can develop, deploy, and maintain separately while being organized thematically within this artist workspace.

## 📚 Documentation

- **README.md**: Human-readable artist and project overview (this file)
- **GEMINI.md**: AI pair programming context for this artist's projects

### Git Workflow

- Day-to-day: use regular Git in this artist repo as you normally would:
  - Example: git add -A && git commit -m "feat: changes" && git push
- Major recursive commit (this artist and all canvases beneath it):
  - From this directory, run: make push
  - This invokes the CLI atelier-cli artist push to:
    - Recurse through all canvases in this artist
    - Commit and push canvases first (if changes)
    - Stage updated canvas pointers and any artist working tree changes
    - Create a single combined artist commit and push it
- Notes:
  - This artist is a submodule of the atelier (root). To roll up multiple artists and the root in one go, run make push at the atelier root.
  - AUTO_COMMIT_DEFAULT=true enables auto-staging and auto-commit for working tree and pointer updates.

  ### Atelier Commands at this level
- Push this artist recursively (canvases → artist):
  - CLI: `atelier-cli artist push [--dry-run] [--quiet] [--force]`
  - Make: `make push`
- Manage canvases from this artist directory:
  - Init a new canvas: `atelier-cli canvas init &lt;canvas-name&gt;`
  - Delete a canvas (requires full directory name, e.g., canvas-example): `atelier-cli canvas delete &lt;canvas-full-name&gt;`
  - Make equivalents:
    - `make canvas-init NAME=example`
    - `make canvas-delete FULL=canvas-example`

Notes:
- To delete an entire artist, run from the atelier root: `atelier-cli artist delete &lt;artist-full-name&gt;`
- Commands are scope-aware: they must be run at the correct level (atelier, artist, canvas) per the CLI’s cobra validation.

Keep curating! 🖼️
