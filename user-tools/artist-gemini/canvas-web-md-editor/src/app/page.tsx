"use client";

import { useState, useEffect } from "react";
import Sidebar from "@/components/Sidebar";
import { readFileContent, saveFile } from "@/lib/git";

export default function Home() {
  const [selectedFile, setSelectedFile] = useState<string | null>(null);
  const [originalContent, setOriginalContent] = useState<string>("");
  const [fileContent, setFileContent] = useState<string>("");
  const [isLoading, setIsLoading] = useState(false);
  const [isSaving, setIsSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const isModified = fileContent !== originalContent;

  const handleFileSelect = async (filepath: string) => {
    setIsLoading(true);
    setError(null);
    setSelectedFile(filepath);
    try {
      const content = await readFileContent(filepath);
      setFileContent(content);
      setOriginalContent(content);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to read file.");
      setFileContent("");
      setOriginalContent("");
    } finally {
      setIsLoading(false);
    }
  };

  const handleSave = async () => {
    if (!selectedFile || !isModified) return;

    setIsSaving(true);
    setError(null);
    try {
      await saveFile(selectedFile, fileContent);
      setOriginalContent(fileContent); // Update original content after save
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to save file.");
    } finally {
      setIsSaving(false);
    }
  };

  return (
    <div className="flex h-full">
      <Sidebar onFileSelect={handleFileSelect} />
      <main className="flex-1 p-6 flex flex-col">
        <div className="flex justify-between items-center mb-4">
          {selectedFile ? (
            <h1 className="text-2xl font-bold">
              Editing: <span className="font-mono">{selectedFile}</span>
            </h1>
          ) : (
            <h1 className="text-2xl font-bold">Markdown Editor</h1>
          )}
          {selectedFile && (
            <button
              onClick={handleSave}
              disabled={!isModified || isSaving}
              className="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-600 disabled:bg-gray-400"
            >
              {isSaving ? "Saving..." : "Save Changes"}
            </button>
          )}
        </div>

        {isLoading ? (
          <div className="flex-1 flex items-center justify-center">
            <p>Loading...</p>
          </div>
        ) : (
          <textarea
            key={selectedFile}
            className="w-full flex-1 p-2 border border-gray-300 rounded-md resize-none font-mono text-gray-900" // Added text-gray-900
            placeholder="Select a file from the sidebar to begin editing..."
            value={fileContent}
            onChange={(e) => setFileContent(e.target.value)}
            readOnly={!selectedFile}
          />
        )}
        {error && <p className="text-red-500 text-sm mt-2">{error}</p>}
      </main>
    </div>
  );
}
