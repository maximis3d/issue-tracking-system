import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { createScope } from "../api/scope";// Adjust the path to your API file

const RegisterScope = () => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [projectKeys, setProjectKeys] = useState("");
  const [_, setError] = useState(null);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setLoading(true);

    try {
      const projects = projectKeys
        .split(",")
        .map((p) => p.trim())
        .filter((p) => p.length > 0);

      await createScope(name, description, projects);
      setMessage("Scope successfully created.");
      setTimeout(() => navigate("/scopes"), 1200);
    } catch (error) {
      setMessage(`Failed to create scope: ${error.message}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="max-w-4xl mx-auto bg-white shadow-lg rounded-lg p-8 border border-gray-200 w-full">
        <h2 className="text-3xl font-semibold mb-6 text-center text-gray-800">
          Register New Scope
        </h2>

        {message && (
          <div
            className={`p-4 mb-4 text-center rounded-lg ${
              message.includes("successfully")
                ? "bg-green-100 text-green-600"
                : "bg-red-100 text-red-600"
            }`}
          >
            {message}
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-6">
          <div className="flex flex-col">
            <label htmlFor="name" className="mb-2 font-medium text-gray-700">
              Scope Name
            </label>
            <input
              id="name"
              type="text"
              required
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="e.g. Sprint Alpha"
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              disabled={loading}
            />
          </div>

          <div className="flex flex-col">
            <label htmlFor="description" className="mb-2 font-medium text-gray-700">
              Description
            </label>
            <textarea
              id="description"
              rows={4}
              required
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="Describe the purpose or coverage of this scope..."
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              disabled={loading}
            />
          </div>

          <div className="flex flex-col">
            <label htmlFor="projectKeys" className="mb-2 font-medium text-gray-700">
              Project Keys (comma-separated)
            </label>
            <input
              id="projectKeys"
              type="text"
              value={projectKeys}
              onChange={(e) => setProjectKeys(e.target.value)}
              placeholder="e.g. PROJ1, PROJ2, PROJ3"
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              disabled={loading}
            />
            <span className="text-xs text-gray-400 mt-1">
              Enter the project keys separated by commas.
            </span>
          </div>

          <button
            type="submit"
            disabled={loading}
            className={`w-full py-3 rounded-lg text-white cursor-pointer focus:outline-none transition-colors ${
              loading
                ? "bg-gray-400 cursor-not-allowed"
                : "bg-blue-600 hover:bg-blue-700"
            }`}
          >
            {loading ? "Creating..." : "Create Scope"}
          </button>
        </form>
      </div>
    </div>
  );
};

export default RegisterScope;
