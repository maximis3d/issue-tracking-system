import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import { fetchScopeDetails, fetchScopeIssues } from "../api/scope";

const Scope = () => {
  const { id } = useParams(); // scope ID from URL
  const [scope, setScope] = useState(null);
  const [issues, setIssues] = useState([]);
  const [loading, setLoading] = useState(true);
  const [issuesLoading, setIssuesLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const loadScopeData = async () => {
      try {
        const scopeData = await fetchScopeDetails(id);
        const issuesData = await fetchScopeIssues(id);
        setScope(scopeData);
        setIssues(issuesData);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
        setIssuesLoading(false);
      }
    };

    loadScopeData();
  }, [id]);

  const statusMapping = {
    open: "To Do",
    in_progress: "In Progress",
    closed: "Done",
  };

  const groupedIssues = issues.reduce((acc, issue) => {
    const status = statusMapping[issue.status] || "To Do";
    if (!acc[status]) acc[status] = [];
    acc[status].push(issue);
    return acc;
  }, {});

  if (loading) return <div className="min-h-screen flex items-center justify-center">Loading scope...</div>;
  if (error) return <div className="min-h-screen flex items-center justify-center text-red-500">Error: {error}</div>;

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-500 to-blue-600 text-white flex flex-col">
      <header className="text-center py-10">
        <h1 className="text-4xl font-bold mb-2">{scope.name}</h1>
        <p className="text-lg max-w-2xl mx-auto">{scope.description}</p>
      </header>

      <div className="max-w-4xl mx-auto p-4 bg-white text-gray-900 rounded-lg shadow-lg">
        <h2 className="text-xl font-semibold mb-2">Scope Details</h2>
        <p><strong>ID:</strong> {scope.id}</p>
        <p><strong>Created At:</strong> {new Date(scope.createdAt).toLocaleDateString()}</p>
        {scope.projects && scope.projects.length > 0 && (
          <p><strong>Projects:</strong> {scope.projects.join(", ")}</p>
        )}
      </div>

      <section className="w-full p-6 mt-8">
        <h2 className="text-2xl font-bold text-white mb-4 text-center">Kanban Board</h2>

        {issuesLoading ? (
          <p className="text-center text-gray-200">Loading issues...</p>
        ) : Object.keys(groupedIssues).length === 0 ? (
          <p className="text-center text-gray-300">No issues found for this scope.</p>
        ) : (
          <div className="flex space-x-4 overflow-x-auto text-black">
            {["To Do", "In Progress", "Done"].map((status) => (
              <div key={status} className="bg-gray-100 rounded-lg flex-1 p-4 min-w-[300px]">
                <h3 className="text-xl font-semibold mb-4 text-center">{status}</h3>
                <div className="space-y-4">
                  {groupedIssues[status]?.map((issue) => (
                    <Link
                      key={issue.id}
                      to={`/issues/${issue.id}`}
                      className="block bg-white p-4 rounded-lg shadow text-black"
                    >
                      <h4 className="text-lg font-semibold">{issue.summary}</h4>
                      <p className="text-sm"><strong>Key:</strong> {issue.key}</p>
                      <p className="text-sm"><strong>Type:</strong> {issue.issueType}</p>
                      <p className="text-sm">
                        <strong>Reporter:</strong> {issue.reporter} | <strong>Assignee:</strong> {issue.assignee || "Unassigned"}
                      </p>
                      <p className="text-xs text-gray-500 mt-1">
                        Last updated: {new Date(issue.updatedAt).toLocaleDateString()}
                      </p>
                    </Link>
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

export default Scope;
