import React, { useEffect, useState } from "react";

const Projects = () => {
  const [projects, setProjects] = useState([]);
  const [search, setSearch] = useState("");
  const [loading, setLoading] = useState(true); 
  const [error, setError] = useState(null); 

  useEffect(() => {
    fetch("http://localhost:8080/api/v1/projects")
      .then((res) => {
        if (!res.ok) {
          throw new Error("Failed to fetch projects");
        }
        return res.json();
      })
      .then((data) => {
        setProjects(data); 
        setLoading(false); 
      })
      .catch((err) => {
        setError(err.message); 
        setLoading(false); 
      });
  }, []);

  const filteredProjects = projects.filter((project) =>
    project.name.toLowerCase().includes(search.toLowerCase())
  );

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-600 to-indigo-700 text-white flex flex-col">
      {/* Header Section */}
      <header className="text-center py-12">
        <h1 className="text-4xl font-bold mb-4">Projects</h1>
        <p className="text-lg max-w-2xl mx-auto mb-8">
          Manage and track all of your team's projects seamlessly.
        </p>
      </header>

      {/* Projects Section */}
      <div className="max-w-4xl mx-auto p-4">
        {/* Search input */}
        <input
          type="text"
          placeholder="Search projects..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="border rounded-lg p-2 w-full mb-8"
        />

        {/* Project cards */}
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          {filteredProjects.length === 0 ? (
            <p>No projects found.</p>
          ) : (
            filteredProjects.map((project) => (
              <div key={project.id} className="bg-white text-gray-900 shadow-md rounded-lg p-4">
                <h2 className="text-xl font-semibold mb-2">{project.name}</h2>
                <p className="text-sm text-gray-600 mb-1">
                  Key: <span className="font-mono">{project.project_key}</span>
                </p>
                <p className="text-sm text-gray-800 mb-2">{project.description}</p>
                <p className="text-sm text-gray-500 mb-2">
                  Lead: {project.projectLead} | Issues: {project.issueCount}
                </p>
              </div>
            ))
          )}
        </div>
      </div>

      {/* Footer Section */}
      <footer className="text-center py-6 text-gray-300 mt-auto">
        © {new Date().getFullYear()} StandUpTracker
      </footer>
    </div>
  );
};

export default Projects;
