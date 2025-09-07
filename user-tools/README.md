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
2. **Work on projects**: `cd artist-name/canvas-name` to enter a project
3. **Develop independently**: Each canvas is its own Git repository

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

Happy creating! 🎨✨