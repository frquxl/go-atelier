"use client";

import { useState, useEffect } from "react";
import { cloneRepo, listFiles, pushChanges } from "@/lib/git";

interface Repository {
  id: number;
  name: string;
  full_name: string;
  html_url: string;
  clone_url: string;
  private: boolean;
  description: string | null;
}

export default function Sidebar({ onFileSelect }: { onFileSelect: (filepath: string) => void }) {
  const [repositories, setRepositories] = useState<Repository[]>([]);
  const [selectedRepo, setSelectedRepo] = useState<Repository | null>(null);
  const [files, setFiles] = useState<string[]>([]);
  const [isLoadingRepos, setIsLoadingRepos] = useState(true);
  const [isCloning, setIsCloning] = useState(false);
  const [isPushing, setIsPushing] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Fetch repositories on component mount
  useEffect(() => {
    const fetchRepositories = async () => {
      setError(null);
      try {
        const token = process.env.NEXT_PUBLIC_GITHUB_TOKEN || process.env.GITHUB_TOKEN;
        if (!token) {
          setError("GitHub token not found. Please set GITHUB_TOKEN in .env.local");
          setIsLoadingRepos(false);
          return;
        }

        const response = await fetch('https://api.github.com/user/repos?sort=updated&per_page=100&type=all', {
          headers: {
            'Authorization': `token ${token}`,
            'Accept': 'application/vnd.github.v3+json'
          }
        });

        if (!response.ok) {
          throw new Error(`Failed to fetch repositories: ${response.statusText}`);
        }

        if (!response.ok) {
          if (response.status === 401) {
            setError("Authentication failed. Please check your GitHub token has 'repo' scope.");
          } else if (response.status === 403) {
            setError("Access forbidden. Your token may not have permission to access private repositories.");
          } else {
            setError(`Failed to fetch repositories: ${response.statusText}`);
          }
          setIsLoadingRepos(false);
          return;
        }

        const repos: Repository[] = await response.json();
        setRepositories(repos);

        // Log repository visibility for debugging
        const privateRepos = repos.filter(repo => repo.private);
        const publicRepos = repos.filter(repo => !repo.private);
        console.log(`Loaded ${repos.length} repositories: ${privateRepos.length} private, ${publicRepos.length} public`);

        if (privateRepos.length === 0 && repos.length > 0) {
          console.warn("No private repositories found. Check token permissions.");
        }
      } catch (err) {
        console.error("Error fetching repositories:", err);
        setError(err instanceof Error ? err.message : "Failed to fetch repositories");
      } finally {
        setIsLoadingRepos(false);
      }
    };
    fetchRepositories();
  }, []);

  const handleRepoSelect = async (repo: Repository) => {
    setIsCloning(true);
    setError(null);
    setFiles([]);
    setSelectedRepo(repo);
    try {
      // Use the canonical HTTPS clone URL (includes .git)
      await cloneRepo(repo.clone_url);
      const fileList = await listFiles();
      setFiles(fileList);
      if (fileList.length === 0) {
        setError("No .md files found in the repository.");
      }
    } catch (err) {
      console.error("Error cloning repository:", err);
      setError(err instanceof Error ? err.message : "Failed to clone repository");
      setSelectedRepo(null);
    } finally {
      setIsCloning(false);
    }
  };


  const handlePush = async () => {
    setIsPushing(true);
    setError(null);
    try {
      await pushChanges();
      alert("Changes pushed successfully!");
    } catch (err) {
      setError(err instanceof Error ? err.message : "An unknown error occurred during push.");
    } finally {
      setIsPushing(false);
    }
  };

  return (
    <aside className="w-72 bg-gray-900 p-4 border-r border-gray-700 flex flex-col text-gray-100">
      <h2 className="text-lg font-semibold mb-4 text-gray-100">Atelier Markdown Editor</h2>


      <hr className="my-4 border-gray-700" />

      {/* Repository selector - compact at top */}
      <div className="mb-4">
        <div className="max-h-32 overflow-y-auto border border-gray-700 rounded">
          {isLoadingRepos ? (
            <p className="text-sm text-gray-300 p-2">Loading repositories...</p>
          ) : repositories.length > 0 ? (
            <div className="space-y-1 p-1">
              {repositories.map((repo) => (
                <button
                  key={repo.id}
                  onClick={() => handleRepoSelect(repo)}
                  disabled={isCloning}
                  className="text-left w-full text-sm text-gray-300 hover:bg-gray-700 p-2 rounded flex items-center justify-between"
                >
                  <span className="truncate font-medium">{repo.name}</span>
                  <span className="text-xs ml-2 flex-shrink-0">
                    {repo.private ? '🔒' : '🌐'}
                  </span>
                </button>
              ))}
            </div>
          ) : (
            <p className="text-sm text-gray-300 p-2">
              {error ? "Failed to load repositories" : "No repositories found"}
            </p>
          )}
        </div>
      </div>

      {selectedRepo && (
        <>
          <hr className="my-2 border-gray-700" />
          <div className="mb-2">
            <div className="flex items-center justify-between mb-2">
              <h3 className="text-sm font-semibold text-gray-100">Current</h3>
              <button
                onClick={handlePush}
                disabled={isPushing || files.length === 0}
                className="bg-green-600 text-white px-2 py-1 rounded text-xs hover:bg-green-700 disabled:bg-gray-600"
              >
                {isPushing ? "Pushing..." : "Push"}
              </button>
            </div>
            <p className="text-xs text-gray-300 truncate">
              {selectedRepo.full_name}
            </p>
          </div>
        </>
      )}

      {error && <p className="text-red-400 text-xs mt-2">{error}</p>}

      {/* File tree - takes remaining space */}
      {files.length > 0 && (
        <>
          <hr className="my-2 border-gray-700" />
          <h3 className="text-sm font-semibold mb-2 text-gray-100">Files</h3>
          <div className="flex-1 overflow-y-auto">
            <ul className="space-y-1">
              {files.map((file) => (
                <li key={file}>
                  <button
                    onClick={() => onFileSelect(file)}
                    className="text-left w-full text-sm text-gray-300 hover:bg-gray-700 p-2 rounded"
                  >
                    {file}
                  </button>
                </li>
              ))}
            </ul>
          </div>
        </>
      )}
    </aside>
  );
}