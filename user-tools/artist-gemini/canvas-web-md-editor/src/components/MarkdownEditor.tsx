"use client";

import { useRef, useEffect } from "react";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { oneDark, oneLight } from "react-syntax-highlighter/dist/esm/styles/prism";

interface MarkdownEditorProps {
  value: string;
  onChange: (value: string) => void;
  placeholder?: string;
  readOnly?: boolean;
  showLineNumbers?: boolean;
  theme?: "light" | "dark";
}

export default function MarkdownEditor({
  value,
  onChange,
  placeholder = "Start writing your markdown...",
  readOnly = false,
  showLineNumbers = true,
  theme = "light"
}: MarkdownEditorProps) {
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  // Auto-resize textarea
  useEffect(() => {
    const textarea = textareaRef.current;
    if (textarea) {
      textarea.style.height = "auto";
      textarea.style.height = `${textarea.scrollHeight}px`;
    }
  }, [value]);

  // Handle keyboard shortcuts
  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.ctrlKey || e.metaKey) {
      switch (e.key) {
        case "b":
          e.preventDefault();
          insertFormatting("**", "**");
          break;
        case "i":
          e.preventDefault();
          insertFormatting("*", "*");
          break;
        case "k":
          e.preventDefault();
          insertLink();
          break;
        case "1":
        case "2":
        case "3":
        case "4":
        case "5":
        case "6":
          e.preventDefault();
          insertHeading(parseInt(e.key));
          break;
      }

      if (e.shiftKey) {
        switch (e.key) {
          case "7":
            e.preventDefault();
            insertList("1. ");
            break;
          case "8":
            e.preventDefault();
            insertList("- ");
            break;
          case "X":
            e.preventDefault();
            insertFormatting("~~", "~~");
            break;
          case "C":
            e.preventDefault();
            insertCodeBlock();
            break;
          case "Q":
            e.preventDefault();
            insertQuote();
            break;
          case "I":
            e.preventDefault();
            insertImage();
            break;
        }
      }
    }
  };

  const insertFormatting = (before: string, after: string = before) => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    const selectedText = value.substring(start, end);

    const newText = value.substring(0, start) + before + selectedText + after + value.substring(end);
    onChange(newText);

    // Reset cursor position
    setTimeout(() => {
      textarea.focus();
      textarea.setSelectionRange(start + before.length, end + before.length);
    }, 0);
  };

  const insertHeading = (level: number) => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const lineStart = value.lastIndexOf("\n", start - 1) + 1;
    const lineEnd = value.indexOf("\n", start);
    const currentLine = value.substring(lineStart, lineEnd === -1 ? value.length : lineEnd);

    const hashes = "#".repeat(level);
    const newLine = currentLine.startsWith("#")
      ? currentLine.replace(/^#+\s*/, `${hashes} `)
      : `${hashes} ${currentLine}`;

    const newText = value.substring(0, lineStart) + newLine + value.substring(lineEnd === -1 ? value.length : lineEnd);
    onChange(newText);
  };

  const insertList = (prefix: string) => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const lineStart = value.lastIndexOf("\n", start - 1) + 1;
    const lineEnd = value.indexOf("\n", start);
    const currentLine = value.substring(lineStart, lineEnd === -1 ? value.length : lineEnd);

    const newLine = currentLine.startsWith("- ") || /^\d+\.\s/.test(currentLine)
      ? currentLine.replace(/^[-*]\s|^(\d+)\.\s/, prefix)
      : `${prefix}${currentLine}`;

    const newText = value.substring(0, lineStart) + newLine + value.substring(lineEnd === -1 ? value.length : lineEnd);
    onChange(newText);
  };

  const insertLink = () => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    const selectedText = value.substring(start, end);

    const linkText = selectedText || "link text";
    const newText = value.substring(0, start) + `[${linkText}](url)` + value.substring(end);
    onChange(newText);

    // Position cursor at URL
    setTimeout(() => {
      textarea.focus();
      const urlStart = start + linkText.length + 3;
      textarea.setSelectionRange(urlStart, urlStart + 3);
    }, 0);
  };

  const insertImage = () => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    const selectedText = value.substring(start, end);

    const altText = selectedText || "alt text";
    const newText = value.substring(0, start) + `![${altText}](image-url)` + value.substring(end);
    onChange(newText);

    // Position cursor at URL
    setTimeout(() => {
      textarea.focus();
      const urlStart = start + altText.length + 4;
      textarea.setSelectionRange(urlStart, urlStart + 9);
    }, 0);
  };

  const insertCodeBlock = () => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    const selectedText = value.substring(start, end);

    const codeBlock = `\`\`\`\n${selectedText}\n\`\`\``;
    const newText = value.substring(0, start) + codeBlock + value.substring(end);
    onChange(newText);

    // Position cursor inside code block
    setTimeout(() => {
      textarea.focus();
      const cursorPos = start + 4;
      textarea.setSelectionRange(cursorPos, cursorPos + selectedText.length);
    }, 0);
  };

  const insertQuote = () => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const lineStart = value.lastIndexOf("\n", start - 1) + 1;
    const lineEnd = value.indexOf("\n", start);
    const currentLine = value.substring(lineStart, lineEnd === -1 ? value.length : lineEnd);

    const newLine = currentLine.startsWith("> ")
      ? currentLine.replace(/^>\s*/, "")
      : `> ${currentLine}`;

    const newText = value.substring(0, lineStart) + newLine + value.substring(lineEnd === -1 ? value.length : lineEnd);
    onChange(newText);
  };

  return (
    <div className="relative">
      <textarea
        ref={textareaRef}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        onKeyDown={handleKeyDown}
        placeholder={placeholder}
        readOnly={readOnly}
        className="w-full min-h-[400px] p-4 border border-gray-300 dark:border-gray-600 rounded-md resize-none font-mono text-sm leading-relaxed bg-white dark:bg-gray-900 text-gray-900 dark:text-gray-100 placeholder-gray-500 dark:placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        style={{
          tabSize: 2,
          whiteSpace: "pre-wrap",
          wordWrap: "break-word",
        }}
      />

    </div>
  );
}