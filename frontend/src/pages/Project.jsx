import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { fetchProjectDetails } from "../api/project";
import { fetchIssues } from "../api/issue";

const ProjectDetails = () => {
  const { key } = useParams();
  const [project, setProject] = useState(null);
  const [issues, setIssues] = useState([]);
  const [loading, setLoading] = useState(true);
  const [issuesLoading, setIssuesLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const getProjectData = async () => {
      try {
        setLoading(true);
        const projectData = await fetchProjectDetails(key);
        setProject(projectData);
        setLoading(false);

        const issuesData = await fetchIssues(key);
        setIssues(issuesData);
        setIssuesLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
        setIssuesLoading(false);
      }
    };

    getProjectData();
  }, [key]);

  if (loading) {
    return <div className="min-h-screen flex items-center justify-center">Loading project...</div>;
  }

  if (error) {
    return <div className="min-h-screen flex items-center justify-center text-red-500">Error: {error}</div>;
  }

  const statusMapping = {
    open: "To Do",
    in_progress: "In Progress",
    closed: "Done",
  };

  // Group issues by their mapped status
  const groupedIssues = issues.reduce((acc, issue) => {
    const status = statusMapping[issue.status] || "To Do";
    if (!acc[status]) {
      acc[status] = [];
    }
    acc[status].push(issue);
    return acc;
  }, {});

  return (
    <div className="min-h-screen bg-gradient-to-br from-green-500 to-blue-600 text-white flex flex-col">
      <header className="text-center py-12">
        <h1 className="text-4xl font-bold mb-4">{project.name}</h1>
        <p className="text-lg max-w-2xl mx-auto">{project.description}</p>
      </header>

      <div className="max-w-4xl mx-auto p-4 bg-white text-gray-900 rounded-lg shadow-lg">
        <p><strong>Project Key:</strong> {project.project_key}</p>
        <p><strong>Lead:</strong> {project.projectLead}</p>
        <p><strong>Issues:</strong> {project.issueCount}</p>
        <p><strong>Created At:</strong> {new Date(project.createdAt).toLocaleDateString()}</p>
      </div>

      <section className="w-full p-4 mt-8">
        <h2 className="text-2xl font-bold text-white mb-4 text-center">Kanban Board</h2>

        {/* Kanban Board with Columns */}
        {issuesLoading ? (
          <p className="text-center text-gray-200">Loading issues...</p>
        ) : Object.keys(groupedIssues).length === 0 ? (
          <p className="text-center text-gray-300">No issues found for this project.</p>
        ) : (
          <div className="flex w-full space-x-4 overflow-x-auto text-black">
            {["To Do", "In Progress", "Done"].map((status) => (
              <div
                key={status}
                className="bg-gray-200 rounded-lg flex-1 p-4"
              >
                <h3 className="text-xl font-semibold mb-4 text-center">{status}</h3>
                <div className="space-y-4">
                  {groupedIssues[status]?.map((issue) => (
                    <div key={issue.id} className="bg-white p-4 rounded-lg shadow-md text-black">
                      <h4 className="text-lg font-semibold mb-2">{issue.summary}</h4>
                      <p className="text-sm mb-1">
                        <strong>Key:</strong> {issue.key}
                      </p>
                      <p className="text-sm mb-1">
                        <strong>Status:</strong> {status} | <strong>Type:</strong> {issue.issueType}
                      </p>
                      <p className="text-sm mb-1">
                        <strong>Reporter:</strong> {issue.reporter} <strong>Assignee:</strong> {issue.assignee || "Unassigned"}
                      </p>
                      <p className="text-xs text-gray-600">
                        Last updated: {new Date(issue.updatedAt).toLocaleDateString()}
                      </p>
                    </div>
                  ))}
                </div>
              </div>
            ))}
          </div>
        )}
      </section>

      <footer className="text-center py-6 text-gray-300 mt-auto">
        © {new Date().getFullYear()} StandUpTracker
      </footer>
    </div>
  );
};

export default ProjectDetails;
