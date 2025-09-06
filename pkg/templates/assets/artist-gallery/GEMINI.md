# Gallery Artist AI Context

## 👨‍🎨 Artist Overview
This is a gallery artist workspace within the atelier, containing polished, production-ready project canvases.

## 🎯 Artist Purpose
**Gallery Artist workspaces** are for:
- **Showcasing**: Presenting finished and high-quality projects.
- **Production**: Hosting applications ready for deployment or long-term maintenance.
- **Refinement**: Focusing on optimization, robust testing, and comprehensive documentation.
- **Collaboration**: Working on projects intended for public release or significant user bases.

## 📁 Workspace Structure
```
artist/
├── .git/                    # Artist's Git repository (submodule)
├── .gitmodules            # Tracks canvas submodules
├── README.md              # Human documentation
├── GEMINI.md              # AI context (this file)
└── canvas-*/              # Project canvases (Git submodules)
    ├── .git/              # Canvas repository (submodule)
    ├── README.md          # Project documentation
    └── GEMINI.md          # Project AI context
```

## 🤖 AI Pair Programming Guidelines

### Understanding This Artist
- **Identity**: Focus on quality, stability, and long-term maintainability.
- **Scope**: Projects here are typically mature, well-tested, and intended for production use.
- **Relationships**: Canvases are often part of a larger, cohesive product or portfolio.

### Canvas Management
- Each canvas is an independent Git repository.
- Canvases can be added, removed, or moved between artists.
- Artist tracks specific versions of each canvas.
- Changes in canvases don't automatically affect the artist.

### Development Workflow
- **Canvas Development**: Work independently in each canvas.
- **Artist Coordination**: Manage relationships between canvases.
- **Version Management**: Decide which canvas versions to include.

### Git Workflow
**AI Pair Programming Context:**
- Fix .gitmodules URLs immediately: Change `./canvas-name` to `git@github.com:user/canvas-name.git`
- Create GitHub repos first: Use `gh repo create user/repo --public --source canvas-dir --push`
- Update after canvas changes: `git add canvas-name && git commit -m "Update canvas"`
- Push artist changes: `git push` to update the artist repository
- Check canvas status: `git submodule status` shows current commit hashes
- Initialize new canvases: `git submodule update --init canvas-name`

### Context Awareness
- **Artist Level**: Understand the grouping theme and purpose.
- **Canvas Level**: Respect each canvas's independence.
- **Cross-Canvas**: Identify opportunities for code sharing or patterns.

### Communication Patterns
- Reference the artist's theme when suggesting new canvases.
- Consider canvas relationships when proposing architectural changes.
- Use the artist metaphor to explain workspace organization.

## 🎨 Artist Metaphor
Think of this gallery artist as:
- **A curated exhibition**: For polished, production-ready works.
- **A museum**: Housing valuable and well-preserved creations.
- **A showcase**: Presenting the best of your development efforts.

## 🔄 Workflow Patterns
- **Adding Canvases**: Create new projects that fit the artist's theme.
- **Organizing Work**: Group related projects together.
- **Version Control**: Manage canvas versions within the artist.
- **Collaboration**: Coordinate across related projects.

## 📚 Documentation Hierarchy
- **Artist README**: Overview of the artist's purpose and canvases.
- **Canvas READMEs**: Individual project documentation.
- **Artist GEMINI**: AI context for the artist's workspace (this file).
- **Canvas GEMINIs**: AI context for individual projects.

This gallery artist workspace provides a focused environment for developing related software projects within the broader atelier context.
