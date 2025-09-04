// src/lib/git.ts
import git from 'isomorphic-git';
import http from 'isomorphic-git/http/web';
import FS from '@isomorphic-git/lightning-fs';

let fs: FS;
let pfs: typeof FS.promises;
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

// NOTE: Cloning from GitHub will require a CORS proxy.
// For MVP, we can use a repo hosted on a server with permissive CORS.
export const cloneRepo = async (url: string) => {
  initFs(); // Ensure FS is initialized

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

  await git.clone({
    fs,
    http,
    dir,
    url,
    corsProxy: 'https://cors.isomorphic-git.org',
    singleBranch: true,
    depth: 1,
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

export const pushChanges = async (token: string) => {
  initFs(); // Ensure FS is initialized
  await git.push({
    fs,
    http,
    dir,
    corsProxy: 'https://cors.isomorphic-git.org',
    onAuth: () => ({ username: token }), // For GitHub, the token is used as the username
  });
};

// Optional: A function to check the status, useful for UI feedback
export const getStatus = async (filepath: string) => {
  initFs(); // Ensure FS is initialized
  return git.status({ fs, dir, filepath });
};