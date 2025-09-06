# Project Canvas

Welcome to your Project Canvas! 🖼️

This is your actual development workspace - a complete, independent software project with its own Git repository.

## 🎯 Project Overview

**Project canvases** are self-contained software projects within the atelier/artist/canvas architecture. Each canvas:
- Has its own independent Git repository
- Can be developed, tested, and deployed separately
- Maintains its own dependencies and configurations
- Is organized thematically within an artist workspace

## 📁 Project Structure

```
canvas/
├── .git/                    # Independent Git repository
├── README.md               # This project documentation
├── GEMINI.md               # AI pair programming context
├── src/                    # Source code directory
├── tests/                  # Test files and test suites
├── docs/                   # Project documentation
├── .gitignore              # Git ignore patterns
└── [project-specific files]
```

## 🚀 Getting Started

1. **Set up your environment**: Install dependencies, configure tools
2. **Explore the codebase**: Review existing code and documentation
3. **Start developing**: Add features, fix bugs, write tests
4. **Commit regularly**: `git add . && git commit -m "feat: your changes"`

## 🔧 Development Guidelines

### Code Organization
- Keep source code in `src/` directory
- Place tests in `tests/` directory
- Use `docs/` for project documentation
- Follow your team's coding standards

### Git Workflow

- Day-to-day: use regular Git in this canvas repo as you normally would:
  - Example: git add -A && git commit -m "feat: changes" && git push
- Major commit (this canvas only):
  - Run: make push
  - This calls the engine [util/git/push-engine.sh](util/git/push-engine.sh:1) to:
    - Auto-stage and commit any working tree changes in this canvas (single commit)
    - Push the canvas to its remote
- Notes:
  - This canvas has no submodules beneath it; make push is non-recursive here.
  - AUTO_COMMIT_DEFAULT=true enables auto-staging and committing.

### Best Practices
- ✅ Write tests for new features
- ✅ Update documentation as you go
- ✅ Keep dependencies up to date
- ✅ Follow security best practices
- ✅ Review code before committing

## 🎯 Project Goals

*What is this project trying to achieve?*

*What technologies and frameworks are you using?*

*Who is the target audience?*

*What are the success criteria?*

## 🤝 Contributing

*How should others contribute to this project?*

*What are the coding standards and conventions?*

*How do you want to collaborate?*

## 📚 Documentation

- **README.md**: Human-readable project guide (this file)
- **GEMINI.md**: AI pair programming context and patterns

Happy coding! 🚀✨