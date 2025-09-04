"use client";

import {
  Bold,
  Italic,
  Strikethrough,
  Heading1,
  Heading2,
  Heading3,
  List,
  ListOrdered,
  Link,
  Image,
  Code,
  Quote,
  Table,
  Undo,
  Redo,
  Eye,
  EyeOff
} from "lucide-react";

interface MarkdownToolbarProps {
  onFormat: (format: string) => void;
  showPreview: boolean;
  onTogglePreview: () => void;
}

export default function MarkdownToolbar({
  onFormat,
  showPreview,
  onTogglePreview
}: MarkdownToolbarProps) {
  const formatButtons = [
    { icon: Bold, label: "Bold", format: "bold", shortcut: "Ctrl+B" },
    { icon: Italic, label: "Italic", format: "italic", shortcut: "Ctrl+I" },
    { icon: Strikethrough, label: "Strikethrough", format: "strikethrough", shortcut: "Ctrl+Shift+X" },
    { icon: Heading1, label: "Heading 1", format: "h1", shortcut: "Ctrl+1" },
    { icon: Heading2, label: "Heading 2", format: "h2", shortcut: "Ctrl+2" },
    { icon: Heading3, label: "Heading 3", format: "h3", shortcut: "Ctrl+3" },
    { icon: List, label: "Bullet List", format: "ul", shortcut: "Ctrl+Shift+8" },
    { icon: ListOrdered, label: "Numbered List", format: "ol", shortcut: "Ctrl+Shift+7" },
    { icon: Link, label: "Link", format: "link", shortcut: "Ctrl+K" },
    { icon: Image, label: "Image", format: "image", shortcut: "Ctrl+Shift+I" },
    { icon: Code, label: "Code Block", format: "code", shortcut: "Ctrl+Shift+C" },
    { icon: Quote, label: "Quote", format: "quote", shortcut: "Ctrl+Shift+Q" },
    { icon: Table, label: "Table", format: "table", shortcut: "Ctrl+T" },
  ];

  return (
    <div className="border-b border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-800 p-2">
      <div className="flex flex-wrap items-center gap-1">
        {/* Format buttons */}
        {formatButtons.map(({ icon: Icon, label, format, shortcut }) => (
          <button
            key={format}
            onClick={() => onFormat(format)}
            className="flex items-center gap-1 px-3 py-1.5 text-sm bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors group relative"
            title={`${label} (${shortcut})`}
          >
            <Icon size={16} />
            <span className="hidden sm:inline">{label}</span>
          </button>
        ))}

        {/* Separator */}
        <div className="w-px h-6 bg-gray-300 dark:bg-gray-600 mx-2" />

        {/* Preview toggle */}
        <button
          onClick={onTogglePreview}
          className={`flex items-center gap-1 px-3 py-1.5 text-sm border rounded-md transition-colors ${
            showPreview
              ? "bg-blue-500 text-white border-blue-500"
              : "bg-white dark:bg-gray-700 border-gray-300 dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-600"
          }`}
          title={showPreview ? "Hide Preview" : "Show Preview"}
        >
          {showPreview ? <EyeOff size={16} /> : <Eye size={16} />}
          <span className="hidden sm:inline">{showPreview ? "Hide Preview" : "Show Preview"}</span>
        </button>

        {/* Undo/Redo */}
        <div className="w-px h-6 bg-gray-300 dark:bg-gray-600 mx-2" />
        <button
          onClick={() => onFormat("undo")}
          className="flex items-center gap-1 px-3 py-1.5 text-sm bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
          title="Undo (Ctrl+Z)"
        >
          <Undo size={16} />
          <span className="hidden sm:inline">Undo</span>
        </button>
        <button
          onClick={() => onFormat("redo")}
          className="flex items-center gap-1 px-3 py-1.5 text-sm bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-md hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors"
          title="Redo (Ctrl+Y)"
        >
          <Redo size={16} />
          <span className="hidden sm:inline">Redo</span>
        </button>
      </div>

      {/* Keyboard shortcuts hint */}
      <div className="mt-2 text-xs text-gray-500 dark:text-gray-400">
        💡 Tip: Use keyboard shortcuts for quick formatting
      </div>
    </div>
  );
}