# Canvas AI Context

## 🖼️ Canvas Overview
This is a project canvas - an independent software development project within the atelier/artist/canvas architecture.

## 🎯 Project Identity
**Canvas projects** are complete, self-contained software projects with their own Git repository. Each canvas:
- Has its own independent Git repository
- Can be developed, tested, and deployed separately
- Maintains its own dependencies and configurations
- Is organized thematically within an artist workspace

## 📁 Project Structure
```
canvas/
├── .git/                    # Independent Git repository
├── README.md               # Human project documentation
├── GEMINI.md               # AI context (this file)
├── src/                    # Source code
├── tests/                  # Test files
├── docs/                   # Documentation
├── .gitignore              # Git ignore patterns
└── [project-specific files]
```

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

### Python Projects
- **Virtual Environment**: All Python projects must use `.venv` for virtual environments
- **Dependencies**: Use `requirements.txt` or `pyproject.toml` for dependency management
- **Environment Activation**: Always activate the virtual environment before development

### Git Workflow
- **Branching**: Feature branches, release branches, etc.
- **Commits**: Atomic, descriptive commit messages
- **Pull Requests**: Code review and collaboration process
- **Releases**: Versioning and release management

## 🔧 Development Best Practices

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

## 🎯 Project Goals & Success Criteria

*What are the measurable outcomes for this project?*

*How will success be evaluated?*

*What are the key milestones and deliverables?*

## 🚀 Development Workflow

1. **Understand the codebase**: Review existing code and documentation
2. **Plan your changes**: Break down tasks and create a plan
3. **Implement features**: Write clean, tested code
4. **Commit regularly**: Use descriptive commit messages
5. **Collaborate**: Share progress and get feedback
6. **Deploy**: Test and release your changes

## 📚 Documentation
- **README.md**: Human-readable project guide
- **GEMINI.md**: AI pair programming context (this file)
- **Code comments**: Inline documentation
- **API docs**: Interface and usage documentation

This canvas provides a focused, independent environment for developing high-quality software within the atelier ecosystem.