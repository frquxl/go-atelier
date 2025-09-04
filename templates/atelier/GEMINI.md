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
├── GEMINI.md              # AI context (this file)
├── artist-*/              # Artist workspaces (Git submodules)
│   ├── .git/              # Artist submodule repository
│   ├── .gitmodules        # Tracks canvas submodules
│   ├── README.md          # Artist documentation
│   ├── GEMINI.md          # Artist AI context
│   └── canvas-*/          # Project canvases (Git submodules)
│       ├── .git/          # Canvas submodule repository
│       ├── README.md      # Project documentation
│       └── GEMINI.md      # Project AI context
```

## 🤖 AI Pair Programming Guidelines

### Understanding the Structure
- **Atelier Level**: High-level project organization and documentation
- **Artist Level**: Thematic grouping of related projects
- **Canvas Level**: Individual, independent software projects

### Development Patterns
- Each canvas is a complete, independent Git repository
- Artists provide organizational grouping (e.g., by technology, client, or theme)
- Atelier maintains overall project structure and relationships

### Git Workflow
- **Canvas commits**: Independent development in each canvas
- **Artist updates**: Track canvas versions within artists
- **Atelier updates**: Track artist versions in main repository

### Context Awareness
- Always identify which level (atelier/artist/canvas) you're working in
- Respect the independence of each canvas's Git repository
- Understand submodule relationships for proper version control

### Communication Style
- Use atelier/artist/canvas metaphors in discussions
- Reference the hierarchical structure when explaining concepts
- Consider the independence of each canvas when suggesting changes

## 🎯 Working Effectively
- **Navigation**: Always know which level you're operating at
- **Changes**: Consider impact across the 3-level hierarchy
- **Collaboration**: Respect each canvas's independent nature
- **Documentation**: Maintain context at each level

## 📚 Available Documentation
- README.md files: Human-readable guides at each level
- GEMINI.md files: AI context for each workspace level
- .gitmodules files: Track submodule relationships

This atelier provides a structured yet flexible environment for software development using Git submodules and the atelier metaphor.