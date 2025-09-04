"use client";

import { useState, useEffect } from "react";
import { cloneRepo, listFiles, pushChanges } from "@/lib/git";

const HARDCODED_REPO_URL = "https://github.com/frquxl/go-atelier.git"; // Hardcoded URL

export default function Sidebar({ onFileSelect }: { onFileSelect: (filepath: string) => void }) {
  const [pat, setPat] = useState(""); // Personal Access Token
  const [files, setFiles] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState(true); // Set to true initially to show loading
  const [isPushing, setIsPushing] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Auto-clone on component mount
  useEffect(() => {
    const autoClone = async () => {
      setError(null);
      setFiles([]);
      try {
        console.log("Attempting to clone repository...");
        await cloneRepo(HARDCODED_REPO_URL);
        console.log("Repository cloned. Attempting to list files...");
        const fileList = await listFiles();
        console.log("Files listed:", fileList);
        setFiles(fileList);
        if (fileList.length === 0) {
          setError("No .md files found in the repository.");
        }
      } catch (err) {
        console.error("Error during auto-clone or listFiles:", err);
        setError(err instanceof Error ? err.message : "An unknown error occurred during auto-clone.");
      } finally {
        setIsLoading(false);
      }
    };
    autoClone();
  }, []); // Empty dependency array means this runs once on mount

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
      {/* Removed repo URL input and Clone button */}
      <p className="text-sm text-gray-700">Cloned: <span className="font-mono">{HARDCODED_REPO_URL}</span></p>
      {isLoading && <p className="text-sm text-gray-500 mt-2">Cloning repository...</p>}

      <hr className="my-4" />

      <div className="mb-4">
        <h2 className="text-lg font-semibold mb-2">Push Changes</h2>
        <input
          type="password"
          value={pat}
          onChange={(e) => setPat(e.target.value)}
          placeholder="GitHub Personal Access Token"
          className="w-full p-2 border border-gray-300 rounded-md mb-2 text-sm text-gray-900 dark:text-gray-100" // Added dark:text-gray-100
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
        {isLoading ? (
          <p className="text-sm text-gray-500">Loading files...</p>
        ) : files.length > 0 ? (
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
            {error ? "" : "No files found. Check console for errors or ensure repository has .md files."}
          </p>
        )}
      </div>
    </aside>
  );
}
