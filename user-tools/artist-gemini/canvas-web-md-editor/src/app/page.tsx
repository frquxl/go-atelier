"use client";

import { useState, useEffect } from "react";
import Sidebar from "@/components/Sidebar";
import MarkdownToolbar from "@/components/MarkdownToolbar";
import MarkdownEditor from "@/components/MarkdownEditor";
import MarkdownPreview from "@/components/MarkdownPreview";
import { readFileContent, saveFile } from "@/lib/git";

export default function Home() {
  const [selectedFile, setSelectedFile] = useState<string | null>(null);
  const [originalContent, setOriginalContent] = useState<string>("");
  const [fileContent, setFileContent] = useState<string>("");
  const [isLoading, setIsLoading] = useState(false);
  const [isSaving, setIsSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [showPreview, setShowPreview] = useState(false);
  const [theme, setTheme] = useState<"light" | "dark">("light");

  const isModified = fileContent !== originalContent;

  // Handle markdown formatting
  const handleFormat = (format: string) => {
    if (!selectedFile) return;

    let newContent = fileContent;
    const textarea = document.querySelector('textarea') as HTMLTextAreaElement;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    const selectedText = fileContent.substring(start, end);

    switch (format) {
      case "bold":
        newContent = fileContent.substring(0, start) + `**${selectedText}**` + fileContent.substring(end);
        break;
      case "italic":
        newContent = fileContent.substring(0, start) + `*${selectedText}*` + fileContent.substring(end);
        break;
      case "strikethrough":
        newContent = fileContent.substring(0, start) + `~~${selectedText}~~` + fileContent.substring(end);
        break;
      case "h1":
        newContent = insertHeading(fileContent, start, 1);
        break;
      case "h2":
        newContent = insertHeading(fileContent, start, 2);
        break;
      case "h3":
        newContent = insertHeading(fileContent, start, 3);
        break;
      case "ul":
        newContent = insertList(fileContent, start, "- ");
        break;
      case "ol":
        newContent = insertList(fileContent, start, "1. ");
        break;
      case "link":
        newContent = fileContent.substring(0, start) + `[${selectedText || "link text"}](url)` + fileContent.substring(end);
        break;
      case "image":
        newContent = fileContent.substring(0, start) + `![${selectedText || "alt text"}](image-url)` + fileContent.substring(end);
        break;
      case "code":
        newContent = fileContent.substring(0, start) + `\`\`\`\n${selectedText}\n\`\`\`` + fileContent.substring(end);
        break;
      case "quote":
        newContent = insertQuote(fileContent, start);
        break;
      case "table":
        newContent = fileContent.substring(0, start) + `| Header 1 | Header 2 |\n|----------|----------|\n| Cell 1   | Cell 2   |` + fileContent.substring(end);
        break;
    }

    setFileContent(newContent);
  };

  const insertHeading = (content: string, position: number, level: number): string => {
    const lineStart = content.lastIndexOf("\n", position - 1) + 1;
    const lineEnd = content.indexOf("\n", position);
    const currentLine = content.substring(lineStart, lineEnd === -1 ? content.length : lineEnd);
    const hashes = "#".repeat(level);
    const newLine = currentLine.startsWith("#")
      ? currentLine.replace(/^#+\s*/, `${hashes} `)
      : `${hashes} ${currentLine}`;
    return content.substring(0, lineStart) + newLine + content.substring(lineEnd === -1 ? content.length : lineEnd);
  };

  const insertList = (content: string, position: number, prefix: string): string => {
    const lineStart = content.lastIndexOf("\n", position - 1) + 1;
    const lineEnd = content.indexOf("\n", position);
    const currentLine = content.substring(lineStart, lineEnd === -1 ? content.length : lineEnd);
    const newLine = currentLine.startsWith("- ") || /^\d+\.\s/.test(currentLine)
      ? currentLine.replace(/^[-*]\s|^(\d+)\.\s/, prefix)
      : `${prefix}${currentLine}`;
    return content.substring(0, lineStart) + newLine + content.substring(lineEnd === -1 ? content.length : lineEnd);
  };

  const insertQuote = (content: string, position: number): string => {
    const lineStart = content.lastIndexOf("\n", position - 1) + 1;
    const lineEnd = content.indexOf("\n", position);
    const currentLine = content.substring(lineStart, lineEnd === -1 ? content.length : lineEnd);
    const newLine = currentLine.startsWith("> ")
      ? currentLine.replace(/^>\s*/, "")
      : `> ${currentLine}`;
    return content.substring(0, lineStart) + newLine + content.substring(lineEnd === -1 ? content.length : lineEnd);
  };

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
      <main className="flex-1 flex flex-col">
        <div className="flex justify-between items-center p-4 border-b border-gray-200 dark:border-gray-700">
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

        {selectedFile && (
          <MarkdownToolbar
            onFormat={handleFormat}
            showPreview={showPreview}
            onTogglePreview={() => setShowPreview(!showPreview)}
          />
        )}

        {isLoading ? (
          <div className="flex-1 flex items-center justify-center">
            <p>Loading...</p>
          </div>
        ) : showPreview && selectedFile ? (
          <div className="flex flex-1">
            <div className="flex-1 border-r border-gray-200 dark:border-gray-700">
              <MarkdownEditor
                value={fileContent}
                onChange={setFileContent}
                placeholder="Select a file from the sidebar to begin editing..."
                readOnly={!selectedFile}
                theme={theme}
              />
            </div>
            <div className="flex-1">
              <MarkdownPreview content={fileContent} theme={theme} />
            </div>
          </div>
        ) : (
          <div className="flex-1 p-4">
            <MarkdownEditor
              value={fileContent}
              onChange={setFileContent}
              placeholder="Select a file from the sidebar to begin editing..."
              readOnly={!selectedFile}
              theme={theme}
            />
          </div>
        )}
        {error && <p className="text-red-500 text-sm mt-2 p-4">{error}</p>}
      </main>
    </div>
  );
}