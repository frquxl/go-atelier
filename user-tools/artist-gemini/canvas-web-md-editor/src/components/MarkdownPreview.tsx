"use client";

import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { oneDark, oneLight } from "react-syntax-highlighter/dist/esm/styles/prism";

interface MarkdownPreviewProps {
  content: string;
  theme?: "light" | "dark";
}

export default function MarkdownPreview({ content, theme = "light" }: MarkdownPreviewProps) {
  return (
    <div className="w-full h-full p-4 overflow-auto bg-white dark:bg-gray-900 text-gray-900 dark:text-gray-100">
      <div className="prose prose-sm sm:prose lg:prose-lg xl:prose-xl dark:prose-invert max-w-none">
        <ReactMarkdown
          remarkPlugins={[remarkGfm]}
          components={{
            code({ node, inline, className, children, ...props }: any) {
              const match = /language-(\w+)/.exec(className || "");
              return !inline && match ? (
                <SyntaxHighlighter
                  style={theme === "dark" ? oneDark : oneLight}
                  language={match[1]}
                  PreTag="div"
                  className="rounded-md"
                  {...props}
                >
                  {String(children).replace(/\n$/, "")}
                </SyntaxHighlighter>
              ) : (
                <code
                  className="px-1.5 py-0.5 bg-gray-100 dark:bg-gray-800 rounded text-sm font-mono"
                  {...props}
                >
                  {children}
                </code>
              );
            },
            h1: ({ children }: any) => (
              <h1 className="text-3xl font-bold mb-4 mt-6 first:mt-0 border-b border-gray-200 dark:border-gray-700 pb-2">
                {children}
              </h1>
            ),
            h2: ({ children }: any) => (
              <h2 className="text-2xl font-bold mb-3 mt-5 border-b border-gray-200 dark:border-gray-700 pb-1">
                {children}
              </h2>
            ),
            h3: ({ children }: any) => (
              <h3 className="text-xl font-semibold mb-2 mt-4">{children}</h3>
            ),
            blockquote: ({ children }: any) => (
              <blockquote className="border-l-4 border-blue-500 pl-4 italic text-gray-700 dark:text-gray-300 my-4">
                {children}
              </blockquote>
            ),
            ul: ({ children }: any) => (
              <ul className="list-disc list-inside space-y-1 my-4">{children}</ul>
            ),
            ol: ({ children }: any) => (
              <ol className="list-decimal list-inside space-y-1 my-4">{children}</ol>
            ),
            li: ({ children }: any) => (
              <li className="leading-relaxed">{children}</li>
            ),
            table: ({ children }: any) => (
              <div className="overflow-x-auto my-4">
                <table className="min-w-full border-collapse border border-gray-300 dark:border-gray-600">
                  {children}
                </table>
              </div>
            ),
            th: ({ children }: any) => (
              <th className="border border-gray-300 dark:border-gray-600 px-4 py-2 bg-gray-100 dark:bg-gray-800 font-semibold text-left">
                {children}
              </th>
            ),
            td: ({ children }: any) => (
              <td className="border border-gray-300 dark:border-gray-600 px-4 py-2">
                {children}
              </td>
            ),
            a: ({ children, href }: any) => (
              <a
                href={href}
                className="text-blue-600 dark:text-blue-400 hover:underline"
                target="_blank"
                rel="noopener noreferrer"
              >
                {children}
              </a>
            ),
            img: ({ src, alt }: any) => (
              <img
                src={src}
                alt={alt}
                className="max-w-full h-auto rounded-lg shadow-md my-4"
              />
            ),
            hr: () => (
              <hr className="border-gray-300 dark:border-gray-600 my-8" />
            ),
            p: ({ children }: any) => (
              <p className="leading-relaxed mb-4 last:mb-0">{children}</p>
            ),
          }}
        >
          {content}
        </ReactMarkdown>
      </div>
    </div>
  );
}