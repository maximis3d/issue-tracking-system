import React, { useState } from "react";
import { useParams } from "react-router-dom";
import { startStandup, endStandup } from "../api/standup";

const Standup = () => {
  const { key } = useParams();
  const [issues, setIssues] = useState([]);
  const [standupActive, setStandupActive] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleStartStandup = async () => {
    try {
      setLoading(true);
      const data = await startStandup(key);
      setIssues(data.issues || []);
      setStandupActive(true);
      setLoading(false);
    } catch (err) {
      console.error(err);
      setError("Failed to start standup");
      setLoading(false);
    }
  };

  const handleEndStandup = async () => {
    try {
      setLoading(true);
      await endStandup(key);
      setIssues([]);
      setStandupActive(false);
      setLoading(false);
    } catch (err) {
      console.error(err);
      setError("Failed to end standup");
      setLoading(false);
    }
  };

  const statusMapping = {
    open: "To Do",
    in_progress: "In Progress",
    closed: "Done",
  };

  // Group issues by mapped status
  const groupedIssues = issues.reduce((acc, issue) => {
    const status = statusMapping[issue.status] || "To Do";
    if (!acc[status]) {
      acc[status] = [];
    }
    acc[status].push(issue);
    return acc;
  }, {});

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-500 to-indigo-600 text-white flex flex-col">
      <header className="text-center py-12">
        <h1 className="text-4xl font-bold mb-4">Standup - {key}</h1>
        <div className="flex justify-center space-x-4">
          <button
            onClick={handleStartStandup}
            className="py-2 px-6 bg-green-400 text-black rounded-lg hover:bg-green-500 focus:outline-none"
            disabled={standupActive || loading}
          >
            Start Standup
          </button>
          <button
            onClick={handleEndStandup}
            className="py-2 px-6 bg-red-400 text-black rounded-lg hover:bg-red-500 focus:outline-none"
            disabled={!standupActive || loading}
          >
            End Standup
          </button>
        </div>
      </header>

      {error && (
        <div className="text-center text-red-400 mb-4">{error}</div>
      )}

      <section className="w-full p-4 mt-8">
        <h2 className="text-2xl font-bold text-white mb-4 text-center">Kanban Board</h2>

        {loading ? (
          <p className="text-center text-gray-200">Loading...</p>
        ) : !standupActive ? (
          <p className="text-center text-gray-300">Start a standup to view issues.</p>
        ) : Object.keys(groupedIssues).length === 0 ? (
          <p className="text-center text-gray-300">No issues for standup.</p>
        ) : (
          <div className="flex w-full space-x-4 overflow-x-auto text-black">
            {["To Do", "In Progress", "Done"].map((status) => (
              <div key={status} className="bg-gray-200 rounded-lg flex-1 p-4 min-w-[300px]">
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

export default Standup;
