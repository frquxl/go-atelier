# Project Canvas: Web Markdown Editor

Welcome to your Project Canvas! 🖼️

This is your actual development workspace - a complete, independent software project for a web-based markdown editor.

## 🎯 Project Overview

This canvas is dedicated to building a **Progressive Web App (PWA)** that allows users to edit markdown files directly from their web browser, especially on mobile devices. It aims to provide a simple, intuitive interface for managing documentation within Git repositories.

## ✨ Latest Features (v0.3.0)

### 🔐 **Private Repository Support**
- **Full GitHub Integration**: Access both public and private repositories
- **Secure Authentication**: Server-side proxy prevents token exposure
- **SSO Support**: Works with GitHub organizations requiring SSO authorization

### 🎨 **Enhanced User Interface**
- **Compact Repository Selector**: Streamlined sidebar with scrollable repository list
- **Full-screen Editor**: Maximum editing space by removing top navigation
- **Visual Repository Indicators**: 🔒 Private / 🌐 Public repository badges
- **Mobile-first Design**: Optimized for tablets and mobile devices

### 🔧 **Technical Improvements**
- **Server-side Git Proxy**: Secure authentication without exposing tokens
- **Better Error Handling**: Clear error messages and debugging information
- **Improved Performance**: Optimized repository loading and file operations
- **Enhanced Security**: Environment-based configuration for sensitive data

## � Project Structure

```
web-app/
├── public/                  # Static assets (images, favicons)
├── src/                     # Application source code
│   ├── app/                 # Next.js App Router pages and layouts
│   ├── components/          # Reusable React components
│   └── lib/                 # Utility functions (Git service)
├── node_modules/            # Project dependencies
├── package.json             # Project metadata and dependencies (React, Next.js, isomorphic-git)
├── next.config.ts           # Next.js configuration
├── tsconfig.json            # TypeScript configuration
├── .gitignore               # Git ignore patterns
└── [other project-specific files]
```

## 🚀 Getting Started

1.  **Navigate to the canvas**: `cd user-tools/artist-gemini/canvas-web-md-editor`
2.  **Install dependencies**:
    ```bash
    npm install
    ```
3.  **Start the development server**:
    ```bash
    npm run dev
    # Or using make: make web-dev
    ```
    Open your browser to `http://localhost:3000` (or the address shown in the terminal).
4.  **Explore the codebase**: Review `src/app/page.tsx`, `src/components/`, and `src/lib/git.ts`.
5.  **Start developing**: Implement new features, refine the UI, and add tests.
6.  **Commit regularly**: `git add . && git commit -m "feat: your changes"`

## 🔧 Development Guidelines

### Code Organization
-   **Next.js App Router**: Pages and layouts are defined under `src/app/`.
-   **React Components**: Reusable UI components are in `src/components/`.
-   **Git Service**: All `isomorphic-git` related logic is encapsulated in `src/lib/git.ts`.
-   **Styling**: Uses Tailwind CSS for utility-first styling.

### Git Workflow
-   **In-browser Git**: This project uses `isomorphic-git` to perform Git operations directly within the browser's `IndexedDB`.
-   **Commit Locally**: Changes are committed to the in-browser Git repository.
-   **Push to Remote**: Use the "Push" functionality in the app to synchronize changes with the remote repository.

### Best Practices
-   ✅ **Client-Side Only**: Ensure all `isomorphic-git` and `lightningfs` operations are strictly client-side to avoid server-side rendering issues.
-   ✅ **Secure Authentication**: For production, replace PATs with a robust OAuth flow.
-   ✅ **Responsive Design**: Ensure the UI is optimized for mobile and tablet screens.
-   ✅ **PWA Features**: Implement service workers and manifest for installability and offline capabilities.

## 🎯 Project Goals

*   **Primary Goal**: Enable users to edit markdown files within Git repositories directly from a web browser, especially on mobile devices.
*   **Key Features**: Repository cloning, file navigation, markdown editing, local commits, and remote pushing.
*   **Target Audience**: Users of `atelier-cli` who need a convenient way to manage documentation on the go.
*   **Success Criteria**: A simple, reliable, and user-friendly web application that fulfills the core editing and Git synchronization needs.

## 🤝 Contributing

*   **Contribution Guidelines**: Follow standard web development practices for React/Next.js.
*   **Coding Standards**: Adhere to ESLint and TypeScript guidelines configured in the project.
*   **Collaboration**: Use Git for version control, commit messages should follow conventional commits (e.g., `feat(editor): add new feature`).

## 📚 Documentation

-   **README.md**: Human-readable project guide (this file).
-   **AGENTS.md**: AI pair programming context and patterns for this specific canvas.

Happy coding! 🚀✨
