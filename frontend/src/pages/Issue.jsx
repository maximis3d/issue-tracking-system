import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import { fetchIssue } from "../api/issue";

const Issue = () => {
  const { id } = useParams();
  const [issue, setIssue] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const getIssueData = async () => {
      try {
        setLoading(true);
        const issueData = await fetchIssue(id);
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
    return (
      <div className="min-h-screen flex items-center justify-center text-gray-600">
        Loading issue...
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen flex items-center justify-center text-red-500">
        Error: {error}
      </div>
    );
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 py-8">
      <div className="w-full max-w-5xl bg-white rounded-lg shadow-lg p-8">
        <div className="flex justify-between items-start mb-6">
          <div>
            <h1 className="text-3xl font-bold text-gray-800 mb-2">{issue.summary}</h1>
            <div className="flex items-center space-x-4">
              <span className="text-sm text-gray-500">Key: {issue.key}</span>
              <span className="px-3 py-1 bg-blue-100 text-blue-700 rounded-full text-xs font-semibold">
                {issue.status.replace("_", " ").toUpperCase()}
              </span>
            </div>
          </div>
          <Link to={`/issues/${id}/edit`}>
            <button className="py-2 px-6 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition">
              Edit Issue
            </button>
          </Link>
        </div>

        <div className="mb-8">
          <h2 className="text-xl font-semibold text-gray-700 mb-2">Description</h2>
          <p className="text-gray-600 whitespace-pre-wrap">{issue.description || "No description provided."}</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="bg-gray-50 p-4 rounded-lg">
            <h3 className="text-sm font-semibold text-gray-500 mb-1">Assignee</h3>
            <p className="text-gray-800">{issue.assignee || "Unassigned"}</p>
          </div>

          <div className="bg-gray-50 p-4 rounded-lg">
            <h3 className="text-sm font-semibold text-gray-500 mb-1">Reporter</h3>
            <p className="text-gray-800">{issue.reporter}</p>
          </div>

          <div className="bg-gray-50 p-4 rounded-lg">
            <h3 className="text-sm font-semibold text-gray-500 mb-1">Created At</h3>
            <p className="text-gray-800">{new Date(issue.createdAt).toLocaleDateString()}</p>
          </div>

          <div className="bg-gray-50 p-4 rounded-lg">
            <h3 className="text-sm font-semibold text-gray-500 mb-1">Last Updated</h3>
            <p className="text-gray-800">{new Date(issue.updatedAt).toLocaleDateString()}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Issue;
