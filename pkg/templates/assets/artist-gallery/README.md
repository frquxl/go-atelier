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

Think of yourself as a gallery artist:
- **Atelier**: The entire workspace/studio
- **Artist**: Your curated exhibition space (this level)
- **Canvas**: Individual masterpieces ready for display

Each canvas represents a complete, independent project that you can develop, deploy, and maintain separately while being organized thematically within this artist workspace.

## 📚 Documentation

- **README.md**: Human-readable artist and project overview (this file)
- **GEMINI.md**: AI pair programming context for this artist's projects

Keep curating! 🖼️
