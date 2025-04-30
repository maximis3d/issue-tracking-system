import React, { useState, useEffect } from "react";
import Select from "react-select";
import { fetchAllScopes, fetchScopeDetails, removeProjectsFromScope } from "../api/scope";
const EditScope = () => {
  const [scopes, setScopes] = useState([]);
  const [selectedScope, setSelectedScope] = useState(null);
  const [projectsInScope, setProjectsInScope] = useState([]);
  const [selectedProjects, setSelectedProjects] = useState([]);
  const [loading, setLoading] = useState(false);
  const [successMessage, setSuccessMessage] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const loadScopes = async () => {
      try {
        const data = await fetchAllScopes();
        setScopes(data); 
      } catch (err) {
        setError(err.message);
      }
    };
    loadScopes();
  }, []);

  useEffect(() => {
    if (!selectedScope) return;

    const loadScopeProjects = async () => {
      try {
        const scopeDetails = await fetchScopeDetails(selectedScope.value);
        setProjectsInScope(scopeDetails.projects || []);
        setSelectedProjects([]);
      } catch (err) {
        setError("Failed to fetch scope details: " + err.message);
      }
    };

    loadScopeProjects();
  }, [selectedScope]);

  const handleProjectChange = (selectedOptions) => {
    setSelectedProjects(selectedOptions.map(option => option.value));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!selectedProjects.length) return; 

    setLoading(true);
    setError(null);
    setSuccessMessage(null);

    try {
      const _ = await removeProjectsFromScope(selectedScope.value, selectedProjects);
      setSuccessMessage("Projects removed successfully.");
      setProjectsInScope(prev => prev.filter(project => !selectedProjects.includes(project)));
      setSelectedProjects([]); // Clear selected projects
    } catch (err) {
      setError("Error: " + err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-8">
      <div className="max-w-4xl mx-auto bg-white shadow-xl rounded-xl p-10 border border-gray-200 w-full">
        <h2 className="text-3xl font-semibold mb-8 text-center text-gray-800">Remove Projects from Scope</h2>

        {error && (
          <div className="p-4 mb-4 text-center rounded-lg bg-red-100 text-red-600">
            {error}
          </div>
        )}

        {successMessage && (
          <div className="p-4 mb-4 text-center rounded-lg bg-green-100 text-green-600">
            {successMessage}
          </div>
        )}

        <div className="mb-6">
          <label className="block text-lg font-medium text-gray-700 mb-2">Select Scope</label>
          <Select
            options={scopes.map(scope => ({ value: scope.id, label: scope.name }))}
            onChange={setSelectedScope}
            value={selectedScope}
            placeholder="Choose a scope to edit"
            className="w-full"
          />
        </div>

        {selectedScope && projectsInScope.length > 0 ? (
          <form onSubmit={handleSubmit}>
            <div className="mb-6">
              <label className="block text-lg font-medium text-gray-700 mb-2">Select Projects to Remove</label>
              <Select
                isMulti
                options={projectsInScope.map(project => ({ value: project, label: project }))}
                onChange={handleProjectChange}
                value={selectedProjects.map(project => ({ value: project, label: project }))}
                placeholder="Select projects..."
                className="w-full"
              />
            </div>

            <button
              type="submit"
              disabled={loading}
              className={`w-full py-3 rounded-lg text-white cursor-pointer focus:outline-none transition-colors ${
                loading
                  ? "bg-gray-400 cursor-not-allowed"
                  : "bg-red-600 hover:bg-red-700"
              }`}
            >
              {loading ? "Removing..." : "Remove Selected Projects"}
            </button>
          </form>
        ) : (
          selectedScope && (
            <div className="p-4 text-center rounded-lg bg-yellow-100 text-yellow-600 mt-6">
              No projects available to remove from this scope.
            </div>
          )
        )}
      </div>
    </div>
  );
};

export default EditScope;
