# Canvas AI Context

## 🖼️ Canvas Overview
This is a project canvas - an independent software development project within the atelier/artist/canvas architecture.

## 🎯 Project Identity
**Canvas projects** are complete, self-contained software projects with their own Git repository. Each canvas:
- Has its own independent Git repository
- Can be developed, tested, and deployed separately
- Maintains its own dependencies and configurations
- Is organized thematically within an artist workspace

## 🤖 AI Pair Programming Guidelines

### Understanding This Canvas
- **Purpose**: What problem does this project solve?
- **Scope**: What functionality is included/excluded?
- **Technology**: What languages, frameworks, and tools are used?
- **Audience**: Who will use this software?

### Development Context
- **Independent Repository**: This canvas has its own Git history
- **Isolated Environment**: Own dependencies and configurations
- **Version Control**: Complete Git workflow within this directory
- **Deployment**: Can be built, tested, and deployed independently

### Code Patterns
- **Architecture**: What design patterns and architectural decisions?
- **Coding Standards**: What conventions and style guidelines?
- **Testing Strategy**: How is code quality ensured?
- **Documentation**: How is the codebase documented?

### Git Workflow
- **Branching**: Feature branches, release branches, etc.
- **Commits**: Atomic, descriptive commit messages
- **Pull Requests**: Code review and collaboration process
- **Releases**: Versioning and release management

# 🔧 Development Best Practices

### Code Quality
- Write clean, readable, maintainable code
- Follow established patterns and conventions
- Include comprehensive tests
- Document complex logic and decisions

### Collaboration
- Use clear commit messages
- Write helpful PR descriptions
- Review code thoroughly
- Communicate design decisions

### Project Management
- Break work into manageable tasks
- Track progress and blockers
- Plan releases and milestones
- Maintain project documentation

## 🚀 Development Workflow

1. **Understand the codebase**: Review existing code and documentation
2. **Plan your changes**: Break down tasks and create a plan
3. **Implement features**: Write clean, tested code
4. **Commit regularly**: Use descriptive commit messages
5. **Collaborate**: Share progress and get feedback
6. **Deploy**: Test and release your changes

## 📚 Documentation
- **README.md**: Human-readable project guide
- **AGENTS.md**: AI pair programming context (this file)
- **Code comments**: Inline documentation

#### Atelier Commands at this level
- Use `make push` or `atelier-cli canvas push` from this canvas directory to push only this canvas.
- Non-recursive: only this canvas is auto-staged, committed (when configured), and pushed.
- For recursive roll-ups:
  - From the artist directory use `make push` or `atelier-cli artist push` to recurse through all canvases and then the artist pointer.
  - From the atelier root use `make push` or `atelier-cli push` to recurse canvases → artists → atelier.
- Useful flags: `--dry-run`, `--quiet`, `--force`.

Example:
```bash
# Preview canvas-only push (no changes pushed)
atelier-cli canvas push --dry-run

# Execute canvas-only push
make push
# equivalent to:
atelier-cli canvas push
```
##### Init/Delete scope
- Canvas initialization and deletion are managed from the artist level (not from inside a canvas):
  - Init: `atelier-cli canvas init &lt;canvas-name&gt;` (run from the artist directory)
  - Delete: `atelier-cli canvas delete &lt;canvas-full-name&gt;` (run from the artist directory, e.g., canvas-example)

# Canvas specific context from here

## 🎯 Canvas Vision and Goals 

## Canvas Requirements

## Canvas MVP

## 📁 Canvas Structure
```
canvas/
├── .git/                    # Independent Git repository
├── README.md               # Human project documentation
├── AGENTS.md               # AI context (this file)
├── Makefile                # Build and development tasks
├── src/                    # Source code
├── tests/                  # Test files
├── docs/                   # Documentation
├── .gitignore              # Git ignore patterns
└── [project-specific files]
```

- and add more below as you paint this new amaxing canvas!
