# Canvas AI Context: Web Markdown Editor

## 🖼️ Canvas Overview
This is a project canvas for the **Web Markdown Editor**, an independent software development project within the atelier/artist/canvas architecture. It is being developed as a Progressive Web App (PWA) to allow users to edit markdown files directly from a web browser, especially on mobile devices.

## 🎯 Project Identity
**Canvas projects** are complete, self-contained software projects with their own Git repository. This canvas:
-   Is a **Progressive Web App (PWA)** built with Next.js (React).
-   Utilizes `isomorphic-git` for in-browser Git operations.
-   Focuses on providing a simple interface for editing markdown files within Git repositories.
-   Aims to be deployable as a web application, avoiding app store distribution.

## 📁 Project Structure
```
web-app/
├── .git/                    # Independent Git repository (managed by isomorphic-git in IndexedDB)
├── README.md               # Human project documentation
├── GEMINI.md               # AI context (this file)
├── public/                  # Static assets (images, favicons)
├── src/                     # Application source code
│   ├── app/                 # Next.js App Router pages and layouts
│   ├── components/          # Reusable React components (Header, Sidebar)
│   └── lib/                 # Utility functions (Git service)
├── node_modules/            # Project dependencies
├── package.json            # Project metadata and dependencies (React, Next.js, isomorphic-git)
├── next.config.ts          # Next.js configuration
├── tsconfig.json           # TypeScript configuration
└── [other project-specific files]
```

## 🤖 AI Pair Programming Guidelines

### Understanding This Canvas
-   **Purpose**: To provide a web-based, mobile-friendly markdown editor for Git-managed documentation.
-   **Scope**: Focus on core editing, cloning, and pushing functionalities.
-   **Technology**: Next.js (React), TypeScript, Tailwind CSS, `isomorphic-git`, `lightningfs`.
-   **Audience**: `atelier-cli` users who need on-the-go documentation management.

### Development Context
-   **Web Application**: Runs in a browser sandbox; server-side rendering (SSR) considerations for browser-specific APIs are crucial.
-   **In-browser Git**: All Git operations are handled by `isomorphic-git` within the browser's `IndexedDB`.
-   **Authentication**: Currently uses PAT for push; future goal is GitHub OAuth.
-   **PWA Focus**: Design for installability and offline capabilities.

### Code Patterns
-   **Next.js App Router**: Standard Next.js application structure.
-   **React Components**: Modular and reusable UI components.
-   **Git Service Module**: `src/lib/git.ts` centralizes all Git-related logic.
-   **Client-Side Execution**: Browser-specific code (e.g., `isomorphic-git` initialization) must be guarded to run only on the client.

### Git Workflow
-   **Local Commits**: Changes are committed to the in-browser Git repository.
-   **Remote Push**: Use the app's UI to push changes to the remote.
-   **No Direct `git` CLI**: All Git interactions are abstracted through `isomorphic-git`.

## 🔧 Development Best Practices

### Code Quality
-   Write clean, readable, maintainable React/TypeScript code.
-   Follow Next.js and Tailwind CSS best practices.
-   Ensure robust error handling for all Git operations.

### Collaboration
-   Use clear and descriptive commit messages (e.g., `feat(editor):`, `fix(ui):`).
-   Adhere to project's ESLint and Prettier configurations.

### Project Management
-   Iterative development, focusing on the MVP first.
-   Prioritize responsive design for mobile usability.

## 🎯 Project Goals & Success Criteria

*   **Primary Goal**: Deliver a functional PWA for editing Git-managed markdown files on mobile.
*   **Success Criteria**: 
    *   Ability to clone public repositories.
    *   Seamless file navigation and editing experience.
    *   Reliable local saving and remote pushing of changes.
    *   Intuitive and responsive user interface.
    *   Achieve PWA installability and basic offline functionality.

## 🚀 Development Workflow

1.  **Set up environment**: `npm install`
2.  **Start dev server**: `npm run dev` (or `make web-dev`)
3.  **Implement features**: Build UI components, integrate Git logic.
4.  **Test**: Manually test features in the browser.
5.  **Commit regularly**: Save progress to the main repository.
6.  **Refine**: Improve UI/UX, add advanced features.

## 📚 Documentation
-   **README.md**: Human-readable project guide.
-   **GEMINI.md**: AI pair programming context (this file).
-   **Code comments**: Inline documentation for complex logic.

Happy coding! 🚀✨
