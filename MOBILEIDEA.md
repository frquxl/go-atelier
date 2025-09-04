# Vision

This cli generates a scaffold for a set of software projects. It would be ideal if users could also be provided a tool to edit the generated markdown files on a mobile phone. 

Perhaps one approach is for us to plan a mobile app that clones the git subomduled atteliar then presents the user with a navigation system and a way to edit the markdown files. 

I envisage it to be very simple, and only to allow them to update and push the files back to the remote repo.

## High-Level Plan

### 1. Core Concept: A Companion App

The mobile app should be treated as a **companion** to the `atelier-cli`. It doesn't need to generate new ateliers, but it should be an expert at consuming and editing the structures the CLI creates. Its primary purpose is to provide a simple and intuitive interface for a very specific task: editing markdown files within a Git repository.

### 2. Technology Choice

To build a "very simple" app for both iOS and Android, a cross-platform framework is the most efficient path.

*   **Recommendation:** **Flutter**. It's a modern framework by Google that allows us to build a beautiful, high-performance app for both platforms from a single Dart codebase. It has strong libraries for handling Git operations and rendering Markdown.

### 3. The User's Journey (Key Features)

We can imagine the user's experience as a series of simple steps:

1.  **Connect to a Repository**: The first time a user opens the app, they are prompted to clone their "atelier" by providing a Git URL (e.g., from GitHub). For private repos, they would also need to provide credentials (which we would need to store securely).
2.  **Navigate the Atelier**: Once cloned, the app would present a clean, hierarchical view of the repository, understanding the `atelier -> artist -> canvas` structure. Instead of a raw file list, it could be a more curated navigation experience.
3.  **Edit Markdown**: The user can tap on any `.md` file to open it in a simple, mobile-friendly editor. The editor would show basic Markdown syntax highlighting and perhaps a small toolbar for common actions like bold or italics.
4.  **Commit & Push Changes**: After editing, the user can save the file. The app would then allow them to write a short commit message and, with a single tap, push the changes back to the remote repository. The app could automatically pull the latest changes before each editing session to prevent conflicts.

### 4. Architectural Blueprint

*   **Git Engine**: The heart of the app will be a Go library compiled for mobile using `gomobile` or a native Dart/Flutter Git library. This engine will handle all the `clone`, `pull`, `commit`, and `push` operations in the background. Using a library is better than trying to run `git` commands on the phone's shell.
*   **Simple UI**: The interface should be minimal. A list of repositories, a file navigator, and an editor screen. That's it. We avoid complex features like branching, merging, or detailed commit history to stick to the "very simple" vision.
*   **Local Storage**: The app will store the cloned repositories on the phone's local storage, allowing for offline editing. Changes can be pushed whenever the user has an internet connection.

### 5. A Phased Approach (MVP)

To get started, we could aim for a Minimum Viable Product (MVP) that proves the core concept:

1.  **Phase 1 (MVP)**:
    *   Clone a *public* Git repository.
    *   Navigate the file structure.
    *   View and edit Markdown files.
    *   Commit and push changes back to the public repository.
    *   (This avoids the complexity of handling private repository authentication initially).

2.  **Phase 2 (Full Vision)**:
    *   Add support for private repositories (e.g., with SSH keys or personal access tokens).
    *   Implement a secure way to store user credentials on the device.
    *   Add a simple "diff" view so the user can see what they've changed before committing.
