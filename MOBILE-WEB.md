# Vision

This cli generates a scaffold for a set of software projects. It would be ideal if users could also be provided a tool to edit the generated markdown files on a mobile phone. 

A web app approach simplifies development and distribution, avoiding the need for app stores.

I envisage it to be very simple, and only to allow them to update and push the files back to the remote repo.

## High-Level Plan for a Web Application

### 1. Core Concept: A Progressive Web App (PWA)

Instead of a native mobile app, we'll build a **Progressive Web App (PWA)**. This is a modern website that can be "installed" on a user's phone home screen, can work offline, and feels almost identical to a native app. This approach gives us the best of both worlds: mobile accessibility without the app stores.

### 2. Technology Choice (Web Stack)

*   **Recommendation:** **React** (using the **Next.js** framework). This is a powerful, industry-standard choice for building modern, fast web applications. We can pair it with a simple styling library like **Tailwind CSS** to create a clean and responsive mobile-first interface.

### 3. The Git Engine: In-Browser Git

This is the biggest and most important change from the native mobile plan. Since a web browser cannot directly use a device's `git` command, we must use a JavaScript-based Git implementation that runs entirely within the browser.

*   **Recommendation:** **`isomorphic-git`**. This is a remarkable library that can clone, edit, commit, and push Git repositories using the browser's internal storage (`IndexedDB`). The entire Git repository essentially lives inside a browser tab.

### 4. The User's Journey (Web Version)

The user experience remains simple and focused:

1.  **Visit the Web App**: The user navigates to the app's URL on their phone's browser.
2.  **Connect via GitHub**: To interact with repositories, the user will log in using their GitHub account (via OAuth). This is a secure, standard way for web apps to get permission to act on a user's behalf without ever seeing their password.
3.  **Clone an Atelier**: After logging in, the user can provide the URL of their atelier repository. The app uses `isomorphic-git` and the user's GitHub token to clone the repo directly into the browser's storage.
4.  **Navigate and Edit Offline**: Once cloned, the user can navigate their files, edit markdown, and commit changes, **even if they go offline**. The PWA and the in-browser Git repo make this possible.
5.  **Push Changes**: When the user has an internet connection, they can simply press a "Push" button to send their saved commits back to the GitHub repository.

### 5. Architectural Blueprint (Web Version)

*   **Frontend**: A Next.js (React) application.
*   **Git Engine**: The `isomorphic-git` library for all Git operations.
*   **Authentication**: GitHub OAuth for login and repository permissions.
*   **Storage**: The browser's `IndexedDB` for storing the cloned Git repository data.
*   **Deployment**: The final web app can be easily deployed globally on modern hosting platforms like Vercel or Netlify.

### 6. A Phased Approach (Web MVP)

This approach also lends itself to a simple MVP to prove the concept quickly.

1.  **Phase 1 (MVP)**:
    *   Build a React app that uses `isomorphic-git`.
    *   Focus *only* on **public** repositories. This lets us bypass the complexity of user login (OAuth) for the first version.
    *   The user can clone a public repo, edit markdown, and commit changes to the in-browser copy.
    *   Since pushing requires authentication, the MVP could allow the user to **download their modified files** or a **patch file**. This delivers the core editing value immediately.

2.  **Phase 2 (Full Vision)**:
    *   Implement the full GitHub OAuth login flow.
    *   Enable pushing changes directly to the remote repository.
    *   Add the PWA features (service worker, manifest file) to make the app installable and enhance its offline capabilities.