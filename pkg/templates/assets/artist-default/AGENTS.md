# Artist AI Context

## 👨‍🎨 Artist Overview
This is an artist workspace within the atelier, containing multiple project canvases that are thematically related.

## 🎯 Artist Purpose
**Artist workspaces** represent thematic or organizational groupings of related software projects. This could be:
- Projects using the same technology stack
- Projects for the same client or organization
- Projects sharing similar goals or patterns
- Personal projects with common themes

## 📁 Workspace Structure
```
artist/
├── .git/                    # Artist's Git repository (submodule)
├── .gitmodules            # Tracks canvas submodules
├── README.md              # Human documentation
├── AGENTS.md              # AI context (this file)
└── canvas-*/              # Project canvases (Git submodules)
    ├── .git/              # Canvas repository (submodule)
    ├── README.md          # Project documentation
    └── AGENTS.md          # Project AI context
```

## 🤖 AI Pair Programming Guidelines

### Understanding This Artist
- **Identity**: What theme, technology, or purpose groups these canvases?
- **Scope**: What types of projects belong in this artist workspace?
- **Relationships**: How do the canvases in this artist relate to each other?

### Canvas Management
- Each canvas is an independent Git repository.
- Canvases can be added, removed, or moved between artists.
- Artist tracks specific versions of each canvas.
- Changes in canvases don't automatically affect the artist.

### Communication Patterns
- Reference the artist's theme when suggesting new canvases.
- Consider canvas relationships when proposing architectural changes.
- Use the artist metaphor to explain workspace organization.

### Development Patterns
- Each canvas is a complete, independent Git repository
- Artists provide organizational grouping (e.g., by technology, client, or theme)

## 🔄 Workflow Patterns
- **Adding Canvases**: Create new projects that fit the artist's theme.
- **Organizing Work**: Group related projects together.
- **Version Control**: Manage canvas versions within the artist.
- **Collaboration**: Coordinate across related projects.

### 📚 Available Documentation
- **Artist README**: Overview of the artist's purpose and canvases.
- **Canvas READMEs**: Individual project documentation.
- **Artist AGENTS**: AI context for the artist's workspace (this file).
- **Canvas AGENTSs**: AI context for individual projects.
- **.gitmodules files**: Track submodule relationships for canvases

### Git Workflows and info
- Update after canvas changes: `git add canvas-name && git commit -m "Update canvas"`
- Push artist changes: `git push` to update the artist repository
- Check canvas status: `git submodule status` shows current commit hashes
- Initialize new canvases: `git submodule update --init canvas-name`
- **'atelier-cli artist push' is available**: Use `make push` or `atelier-cli artist push` from the artist directory to recursively push this artist and all canvases.
- Recursive order: canvases are committed and pushed first, then the artist (with updated submodule pointers).
- For a full workspace roll-up across all artists from the atelier root, use `make push` or `atelier-cli push`.
- Useful flags: `--dry-run`, `--quiet`, `--force`.

Example:
```bash
# Preview recursive artist-level push (no changes pushed)
atelier-cli artist push --dry-run

# Execute recursive artist-level push
make push
# equivalent to:
atelier-cli artist push
```

#### Atelier Canvas Commands at this level
- Initialize a new canvas in this artist (run from the artist directory):
  - `atelier-cli canvas init &lt;canvas-name&gt;`
- Delete a canvas (run from the artist directory, requires full directory name, e.g., canvas-example):
  - `atelier-cli canvas delete &lt;canvas-full-name&gt;`

Make targets:
```bash
# Initialize/delete canvases from this artist's Makefile
make canvas-init NAME=example
make canvas-delete FULL=canvas-example
```

Note:
- To delete an entire artist, run from the atelier root:
  - `atelier-cli artist delete &lt;artist-full-name&gt;`

