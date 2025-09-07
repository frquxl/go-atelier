# Atelier AI Context

## 🎨 Atelier Overview
This is the root level of an atelier workspace using the atelier/artist/canvas metaphor for software development with a 3-level Git submodule architecture.

## 🏗️ Architecture Pattern
**3-Level Git Submodule System:**
- **Atelier** (main repository): Project container and structure (this level)
- **Artists** (submodules): Individual workspaces containing project groups
- **Canvases** (submodules): Actual project repositories for development

## 📊 Repository Structure
```
atelier/
├── .git/                    # Main atelier repository
├── .gitmodules            # Tracks artist submodules
├── README.md              # Human documentation
├── AGENTS.md              # AI context (this file)
├── artist-*/              # Artist workspaces (Git submodules)
│   ├── .git/              # Artist submodule repository
│   ├── .gitmodules        # Tracks canvas submodules
│   ├── README.md          # Artist documentation
│   ├── AGENTS.md          # Artist AI context
│   └── canvas-*/          # Project canvases (Git submodules)
│       ├── .git/          # Canvas submodule repository
│       ├── README.md      # Project documentation
│       └── AGENTS.md      # Project AI context
```

## 🤖 AI Pair Programming Guidelines

### Understanding the Structure
- **Atelier Level**: High-level project organization and documentation
- **Artist Level**: Thematic grouping of related projects
- **Canvas Level**: Individual, independent software projects
- Always identify which level (atelier/artist/canvas) you're working in
- Respect the independence of each canvas's Git repository
- Understand submodule relationships for proper version control

### Communication Style
- Use atelier/artist/canvas metaphors in discussions
- Reference the hierarchical structure when explaining concepts
- Consider the independence of each canvas when suggesting changes

### Development Patterns
- Each canvas is a complete, independent Git repository
- Artists provide organizational grouping (e.g., by technology, client, or theme)
- Atelier maintains overall project structure and relationships

### 📚 Available Documentation
- **README.md files**: Human-readable guides at each level
- **AGENTS.md files**: AI context for each workspace level
- **.gitmodules files**: Track submodule relationships for artists

### Git Workflows and info
- **Atelier repo updates**: Track artist versions in main repository
- **Artist repo updates**: Track canvas versions within artists
- **Canvas repo commits**: Independent development in each canvas
- **Check submodule status**: `git submodule status` to see current commit hashes
- **Atelier level repo**: yes Atlelier is repo too, mainly docs
- **'atelier-cli push' is available**: `make push` or `atelier-cli push` works at this level to perform a recursive roll-up push from the atelier root.
- Recursive order: canvases are committed and pushed first, then artists (with updated submodule pointers), then the atelier (with updated artist pointers).
- Useful flags: `--dry-run`, `--quiet`, `--force`.

Example:
```bash
# Safe preview of the recursive push (no changes pushed)
atelier-cli push --dry-run

# Execute the recursive push from the atelier root
make push
# equivalent to:
atelier-cli push
```
#### Atelier Artist Commands at this level
- Initialize a new artist from the atelier root:
  - `atelier-cli artist init &lt;artist-name&gt;`
- Delete an artist (run from the atelier root, requires full directory name, e.g., artist-van-gogh):
  - `atelier-cli artist delete &lt;artist-full-name&gt;`

Make targets:
```bash
# Initialize/delete from the atelier's Makefile
make artist-init NAME=van-gogh
make artist-delete FULL=artist-van-gogh
```

This atelier provides a structured yet flexible environment for software development using Git submodules and the atelier metaphor.