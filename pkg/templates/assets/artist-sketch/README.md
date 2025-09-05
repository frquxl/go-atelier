# Sketch Artist Workspace

Welcome to your Sketch Artist workspace! ✏️

This artist workspace contains multiple project canvases, each representing a different software development project. Think of this as your personal sketchbook for quick ideas and rough drafts.

## 🎯 Artist Purpose

**Sketch Artist workspaces** are designed for:
- Rapid prototyping and experimentation.
- Exploring new ideas or technologies quickly.
- Creating minimal viable products (MVPs) or proof-of-concepts.
- Projects where speed and iteration are prioritized over polish.

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
    ├── src/               # Source code
    ├── tests/             # Test files
    └── docs/              # Project documentation
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

Think of yourself as a sketch artist:
- **Atelier**: The entire workspace/studio
- **Artist**: Your personal sketchbook for quick ideas (this level)
- **Canvas**: Individual sketches or rough drafts you're working on

Each canvas represents a complete, independent project that you can develop, deploy, and maintain separately while being organized thematically within this artist workspace.

## 📚 Documentation

- **README.md**: Human-readable artist and project overview (this file)
- **GEMINI.md**: AI pair programming context for this artist's projects

Keep sketching! ✏️
