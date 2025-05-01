import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { createScope } from "../api/scope"; 
import { fetchAllProjects } from "../api/project"; 
import Select from "react-select"; 

const CreateScope = () => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [selectedProjects, setSelectedProjects] = useState([]);
  const [projects, setProjects] = useState([]); 
  const [_, setError] = useState(null);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const navigate = useNavigate();

  useEffect(() => {
    const getProjects = async () => {
      try {
        const projectsData = await fetchAllProjects();
        setProjects(projectsData.projects);
      } catch (error) {
        setMessage(`Failed to load projects: ${error}`);
      }
    };
    getProjects();
  }, []);

  const projectOptions = projects.map((project) => ({
    value: project.id,
    label: `${project.name} (${project.project_key})`,
  }));

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setLoading(true);

    try {
      const projectKeys = selectedProjects
        .map((id) => {
          const project = projects.find((p) => p.id === id);
          return project?.project_key;
        })
        .filter(Boolean);

      await createScope(name, description, projectKeys);
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
        <h2 className="text-3xl font-semibold mb-6 text-center text-gray-800">Register New Scope</h2>

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
          {/* Scope Name */}
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

          {/* Description */}
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

          {/* Searchable Multi-select Project Dropdown */}
          <div className="flex flex-col">
            <label htmlFor="projects" className="mb-2 font-medium text-gray-700">
              Select Projects
            </label>
            <Select
              id="projects"
              isMulti
              options={projectOptions}
              value={projectOptions.filter((option) =>
                selectedProjects.includes(option.value)
              )}
              onChange={(selectedOptions) => {
                setSelectedProjects(selectedOptions.map((option) => option.value));
              }}
              className="w-full"
              isDisabled={loading}
              placeholder="Select projects..."
            />
            <span className="text-xs text-gray-400 mt-1">
              You can select multiple projects.
            </span>
          </div>

          {/* Submit Button */}
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

export default CreateScope;
