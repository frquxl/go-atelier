"use client";

import { useState } from "react";
import { cloneRepo, listFiles, pushChanges } from "@/lib/git";

export default function Sidebar({ onFileSelect }: { onFileSelect: (filepath: string) => void }) {
  const [repoUrl, setRepoUrl] = useState("");
  const [pat, setPat] = useState(""); // Personal Access Token
  const [files, setFiles] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [isPushing, setIsPushing] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleClone = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);
    setFiles([]);
    try {
      await cloneRepo(repoUrl);
      const fileList = await listFiles();
      setFiles(fileList);
    } catch (err) {
      setError(err instanceof Error ? err.message : "An unknown error occurred during clone.");
    } finally {
      setIsLoading(false);
    }
  };

  const handlePush = async () => {
    setIsPushing(true);
    setError(null);
    try {
      await pushChanges(pat);
      alert("Changes pushed successfully!");
    } catch (err) {
      setError(err instanceof Error ? err.message : "An unknown error occurred during push.");
    } finally {
      setIsPushing(false);
    }
  };

  return (
    <aside className="w-72 bg-gray-100 p-4 border-r border-gray-200 flex flex-col">
      <h2 className="text-lg font-semibold mb-4">Repository</h2>
      <form onSubmit={handleClone} className="mb-4">
        <input
          type="text"
          value={repoUrl}
          onChange={(e) => setRepoUrl(e.target.value)}
          placeholder="Enter public repo URL"
          className="w-full p-2 border border-gray-300 rounded-md mb-2 text-sm text-gray-900" // Added text-gray-900
          disabled={isLoading}
        />
        <button
          type="submit"
          className="w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 disabled:bg-gray-400"
          disabled={isLoading}
        >
          {isLoading ? "Cloning..." : "Clone"}
        </button>
      </form>

      <hr className="my-4" />

      <div className="mb-4">
        <h2 className="text-lg font-semibold mb-2">Push Changes</h2>
        <input
          type="password"
          value={pat}
          onChange={(e) => setPat(e.target.value)}
          placeholder="GitHub Personal Access Token"
          className="w-full p-2 border border-gray-300 rounded-md mb-2 text-sm text-gray-900" // Added text-gray-900
          disabled={isPushing}
        />
        <button
          onClick={handlePush}
          className="w-full bg-green-500 text-white p-2 rounded-md hover:bg-green-600 disabled:bg-gray-400"
          disabled={isPushing || files.length === 0}
        >
          {isPushing ? "Pushing..." : "Push"}
        </button>
      </div>

      {error && <p className="text-red-500 text-sm mt-2">{error}</p>}
      
      <hr className="my-4" />

      <h2 className="text-lg font-semibold mb-4">Files</h2>
      <div className="flex-1 overflow-y-auto">
        {files.length > 0 ? (
          <ul>
            {files.map((file) => (
              <li key={file} className="mb-1">
                <button
                  onClick={() => onFileSelect(file)}
                  className="text-left w-full text-sm text-gray-700 hover:bg-gray-200 p-1 rounded"
                >
                  {file}
                </button>
              </li>
            ))}
          </ul>
        ) : (
          <p className="text-sm text-gray-500">
            {isLoading ? "" : "Clone a repository to see files."}
          </p>
        )}
      </div>
    </aside>
  );
}