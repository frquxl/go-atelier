import Sidebar from "@/components/Sidebar";

export default function Home() {
  return (
    <div className="flex h-full">
      <Sidebar />
      <div className="flex-1 p-6">
        <h1 className="text-2xl font-bold mb-4">Markdown Editor</h1>
        <textarea
          className="w-full h-5/6 p-2 border border-gray-300 rounded-md resize-none"
          placeholder="Select a file from the sidebar to begin editing..."
          readOnly
        />
      </div>
    </div>
  );
}