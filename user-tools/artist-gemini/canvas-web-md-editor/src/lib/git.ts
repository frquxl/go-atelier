// src/lib/git.ts
import git from 'isomorphic-git';
import http from 'isomorphic-git/http/web';
import FS from '@isomorphic-git/lightning-fs';

let fs: FS;
let pfs: any;
const dir = '/';

// Initialize FS only on the client side
function initFs() {
  if (typeof window !== 'undefined' && !fs) {
    fs = new FS('fs');
    pfs = fs.promises;
  }
  if (!fs) {
    throw new Error('Filesystem is not initialized. This code should only run on the client side.');
  }
}

// Get authentication credentials based on repository URL
function getAuthFromUrl(url: string) {
  if (url.includes('github.com')) {
    const token = process.env.NEXT_PUBLIC_GITHUB_TOKEN || process.env.GITHUB_TOKEN;
    if (token) {
      // For GitHub, use token as username for better CORS proxy compatibility
      return { username: token, password: 'x-oauth-basic' };
    }
  } else if (url.includes('gitlab.com')) {
    const token = process.env.NEXT_PUBLIC_GITLAB_TOKEN || process.env.GITLAB_TOKEN;
    if (token) {
      return { username: 'oauth2', password: token };
    }
  } else if (url.includes('bitbucket.org')) {
    const username = process.env.NEXT_PUBLIC_BITBUCKET_USERNAME || process.env.BITBUCKET_USERNAME;
    const password = process.env.NEXT_PUBLIC_BITBUCKET_PASSWORD || process.env.BITBUCKET_PASSWORD;
    if (username && password) {
      return { username, password };
    }
  }
  return null;
}

// NOTE: Cloning from GitHub will require a CORS proxy.
// For MVP, we can use a repo hosted on a server with permissive CORS.
export const cloneRepo = async (url: string) => {
  initFs(); // Ensure FS is initialized

  // Get authentication credentials
  const auth = getAuthFromUrl(url);

  // Before cloning, wipe the filesystem to ensure a clean slate
  const allFiles = await pfs.readdir('/');
  for (const file of allFiles) {
    if (file !== '.' && file !== '..') {
      // This is a bit of a hack to clean the root. A better way would be to use a dedicated directory for each repo.
      try {
        await pfs.unlink('/' + file);
      } catch {
        await pfs.rmdir('/' + file, { recursive: true });
      }
    }
  }

  // For GitHub, try embedding credentials in URL for CORS proxy
  let finalUrl = url;
  if (url.includes('github.com') && auth) {
    const token = process.env.NEXT_PUBLIC_GITHUB_TOKEN || process.env.GITHUB_TOKEN;
    if (token) {
      finalUrl = url.replace('https://', `https://${token}:x-oauth-basic@`);
    }
  }

  await git.clone({
    fs,
    http,
    dir,
    url,
    corsProxy: '/api/git-proxy',
    singleBranch: true,
    depth: 1,
    ...(auth && {
      onAuth: () => auth,
      onAuthFailure: (url: string, auth: any) => {
        throw new Error(`Authentication failed for ${url}`);
      }
    }),
  });

  // After cloning, verify that there's at least one commit
  // If not, the clone might have failed or the repo is empty
  try {
    const commits = await git.log({ fs, dir, depth: 1 });
    if (commits.length === 0) {
      throw new Error('Cloned repository has no commits. It might be empty or the clone failed.');
    }
    console.log("Cloned repository has commits:", commits); // Log commits
  } catch (e: any) {
    throw new Error(`Failed to verify cloned repository: ${e.message}`);
  }

  // --- NEW DEBUGGING STEP --- 
  // List contents of the root directory in LightningFS after clone
  try {
    const rootContents = await pfs.readdir('/');
    console.log("Contents of LightningFS root after clone:", rootContents);
  } catch (e: any) {
    console.error("Error reading LightningFS root after clone:", e);
  }
  // --- END NEW DEBUGGING STEP --- 
};

export const listFiles = async () => {
  initFs(); // Ensure FS is initialized

  const files = await git.listFiles({
    fs,
    dir,
  });

  console.log("Files found by git.listFiles (before .md filter):", files); // Log files before filtering

  // Filter for .md files only
  return files.filter(filepath => filepath.endsWith('.md'));
};

export const readFileContent = async (filepath: string) => {
  initFs(); // Ensure FS is initialized
  const content = await git.readBlob({
    fs,
    dir,
    oid: await git.resolveRef({ fs, dir, ref: 'HEAD' }),
    filepath,
  });
  return new TextDecoder().decode(content.blob);
};

export const saveFile = async (filepath: string, content: string) => {
  initFs(); // Ensure FS is initialized
  await pfs.writeFile(`${dir}${filepath}`, content, 'utf8');
  
  await git.add({ fs, dir, filepath });

  await git.commit({
    fs,
    dir,
    message: `docs: update ${filepath}`,
    author: {
      name: 'Atelier Web Editor',
      email: 'editor@atelier.dev',
    },
  });
};

export const pushChanges = async (token?: string) => {
  initFs(); // Ensure FS is initialized

  let auth: any = null;

  if (token) {
    auth = { username: token, password: 'x-oauth-basic' };
  } else {
    // Try to get auth from current repo URL (assuming it's set somewhere)
    // For now, we'll use a generic approach
    const githubToken = process.env.NEXT_PUBLIC_GITHUB_TOKEN || process.env.GITHUB_TOKEN;
    if (githubToken) {
      auth = { username: githubToken, password: 'x-oauth-basic' };
    }
  }

  await git.push({
    fs,
    http,
    dir,
    corsProxy: '/api/git-proxy',
    ...(auth && { onAuth: () => auth }),
  });
};

// Optional: A function to check the status, useful for UI feedback
export const getStatus = async (filepath: string) => {
  initFs(); // Ensure FS is initialized
  return git.status({ fs, dir, filepath });
};