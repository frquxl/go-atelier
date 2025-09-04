"use client";

import { useState } from "react";
import Sidebar from "@/components/Sidebar";
import { readFileContent } from "@/lib/git";

export default function Home() {
  const [selectedFile, setSelectedFile] = useState<string | null>(null);
  const [fileContent, setFileContent] = useState<string>("");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleFileSelect = async (filepath: string) => {
    setIsLoading(true);
    setError(null);
    setSelectedFile(filepath);
    try {
      const content = await readFileContent(filepath);
      setFileContent(content);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to read file.");
      setFileContent("");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="flex h-full">
      <Sidebar onFileSelect={handleFileSelect} />
      <main className="flex-1 p-6 flex flex-col">
        {selectedFile && (
          <h1 className="text-2xl font-bold mb-4">
            Editing: <span className="font-mono">{selectedFile}</span>
          </h1>
        )}
        {isLoading ? (
          <div className="flex-1 flex items-center justify-center">
            <p>Loading...</p>
          </div>
        ) : (
          <textarea
            key={selectedFile} // Re-mount textarea when file changes
            className="w-full flex-1 p-2 border border-gray-300 rounded-md resize-none font-mono"
            placeholder={selectedFile ? "" : "Select a file from the sidebar to begin editing..."}
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
