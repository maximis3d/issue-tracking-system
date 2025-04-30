import React, { useState, useEffect } from "react";
import Select from "react-select"; // Import react-select
import { fetchAllProjects } from "../api/project"; // Assume this API call will fetch the list of projects

const CycleTime = () => {
  const [projects, setProjects] = useState([]);
  const [selectedProjects, setSelectedProjects] = useState([]); // Track selected projects
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    const getProjects = async () => {
      setLoading(true);
      setError(null);

      try {
        const projectsData = await fetchAllProjects(); // Fetch projects from the API
        setProjects(projectsData); // Assuming the response is an array of projects
      } catch  {
        setError("Failed to load projects");
      } finally {
        setLoading(false);
      }
    };

    getProjects();
  }, []);

  // Format the projects for react-select
  const projectOptions = projects.map((project) => ({
    value: project.id, // Use project id as the value
    label: `${project.name} (${project.project_key})`, // Combine name and key for display
  }));

  return (
    <div className="flex flex-col">
      <label htmlFor="projects" className="mb-2 font-medium text-gray-700">
        Select Projects
      </label>
      {error && <p className="text-red-500 text-sm">{error}</p>}
      <Select
        id="projects"
        isMulti
        options={projectOptions} // Set project options here
        value={projectOptions.filter((option) =>
          selectedProjects.includes(option.value)
        )}
        onChange={(selectedOptions) => {
          setSelectedProjects(selectedOptions.map((option) => option.value)); // Update selected projects
        }}
        className="w-full"
        isDisabled={loading}
        placeholder={loading ? "Loading projects..." : "Select projects..."}
      />
      <span className="text-xs text-gray-400 mt-1">
        You can select multiple projects.
      </span>
    </div>
  );
};

export default CycleTime;
