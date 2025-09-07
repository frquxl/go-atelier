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
- Each canvas is an independent Git repository
- Canvases can be added, removed, or moved between artists
- Artist tracks specific versions of each canvas
- Changes in canvases don't automatically affect the artist

### Development Workflow
- **Canvas Development**: Work independently in each canvas
- **Artist Coordination**: Manage relationships between canvases
- **Version Management**: Decide which canvas versions to include

### Context Awareness
- **Artist Level**: Understand the grouping theme and purpose
- **Canvas Level**: Respect each canvas's independence
- **Cross-Canvas**: Identify opportunities for code sharing or patterns

### Communication Patterns
- Reference the artist's theme when suggesting new canvases
- Consider canvas relationships when proposing architectural changes
- Use the artist metaphor to explain workspace organization

## 🎨 Artist Metaphor
Think of this artist as:
- **A painter's studio**: Containing multiple works in progress
- **A portfolio**: Showcasing related projects and skills
- **A workspace**: Organized around a particular theme or technology

## 🔄 Workflow Patterns
- **Adding Canvases**: Create new projects that fit the artist's theme
- **Organizing Work**: Group related projects together
- **Version Control**: Manage canvas versions within the artist
- **Collaboration**: Coordinate across related projects

## 📚 Documentation Hierarchy
- **Artist README**: Overview of the artist's purpose and canvases
- **Canvas READMEs**: Individual project documentation
- **Artist AGENTS**: AI context for the artist's workspace (this file)
- **Canvas AGENTSs**: AI context for individual projects

This artist workspace provides a focused environment for developing related software projects within the broader atelier context.