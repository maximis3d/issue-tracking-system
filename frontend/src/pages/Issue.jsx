import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import { fetchIssue } from "../api/issue"; // Import the fetchIssue function

const Issue = () => {
  const { id } = useParams();
  const [issue, setIssue] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const getIssueData = async () => {
      try {
        setLoading(true);
        const issueData = await fetchIssue(id); // Fetch issue using the provided fetchIssue function
        setIssue(issueData);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    getIssueData();
  }, [id]);

  if (loading) {
    return <div className="min-h-screen flex items-center justify-center">Loading issue...</div>;
  }

  if (error) {
    return <div className="min-h-screen flex items-center justify-center text-red-500">Error: {error}</div>;
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-green-500 to-blue-600 text-white flex flex-col">
      <header className="text-center py-12">
        <h1 className="text-4xl font-bold mb-4">{issue.summary}</h1>
        <p className="text-lg max-w-2xl mx-auto">{issue.description}</p>
      </header>

      <div className="max-w-4xl mx-auto p-4 flex justify-end">
        <Link to={`/issues/${id}/edit`}>
          <button className="py-2 px-6 bg-yellow-500 text-black rounded-lg hover:bg-yellow-600 focus:outline-none">
            Edit Issue
          </button>
        </Link>
      </div>

      <div className="max-w-4xl mx-auto p-4 bg-white text-gray-900 rounded-lg shadow-lg">
        <p><strong>Issue Key:</strong> {issue.key}</p>
        <p><strong>Status:</strong> {issue.status}</p>
        <p><strong>Assignee:</strong> {issue.assignee || "Unassigned"}</p>
        <p><strong>Reporter:</strong> {issue.reporter}</p>
        <p><strong>Created At:</strong> {new Date(issue.createdAt).toLocaleDateString()}</p>
        <p><strong>Last Updated:</strong> {new Date(issue.updatedAt).toLocaleDateString()}</p>
      </div>

      <footer className="text-center py-6 text-gray-300 mt-auto">
        © {new Date().getFullYear()} StandUpTracker
      </footer>
    </div>
  );
};

export default Issue;
