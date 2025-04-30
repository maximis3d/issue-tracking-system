import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { fetchIssue } from "../api/issue";

const EditIssue = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [issue, setIssue] = useState(null);
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    const loadIssue = async () => {
      try {
        const issueData = await fetchIssue(id);
        setIssue(issueData);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    loadIssue();
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setIssue((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Map human-readable status to backend format
    const statusMapping = {
      'To Do': 'open',
      'In Progress': 'in_progress',
      'Resolved': 'resolved'
    };

    // Ensure that the status is mapped to the correct backend value
    const mappedStatus = statusMapping[issue.status] || issue.status; // Fallback to issue.status if mapping is not found

    // Update the issue object with the correct status format
    const updatedIssue = {
      ...issue,
      status: mappedStatus, // Send the correct backend value for status
    };

    try {
      setSaving(true);
      const res = await fetch(`http://localhost:8080/api/v1/issues/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(updatedIssue),
      });

      if (!res.ok) {
        const errorData = await res.json();
        throw new Error(errorData.message || "Failed to update issue");
      }

      navigate(`/issues/${id}`);
    } catch (err) {
      setError(err.message);
    } finally {
      setSaving(false);
    }
  };

  if (loading) {
    return <div className="min-h-screen flex items-center justify-center">Loading issue...</div>;
  }

  if (error) {
    return <div className="min-h-screen flex items-center justify-center text-red-500">Error: {error}</div>;
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-green-500 to-blue-600 p-4">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-8 rounded-lg shadow-lg w-full max-w-3xl space-y-6"
      >
        <h1 className="text-2xl font-bold text-center text-gray-800">Edit Issue</h1>

        <div className="space-y-4">
          <div>
            <label className="block mb-2 text-gray-700">Summary</label>
            <input
              type="text"
              name="summary"
              value={issue.summary}
              onChange={handleChange}
              className="w-full p-3 border rounded-lg"
              required
            />
          </div>

          <div>
            <label className="block mb-2 text-gray-700">Description</label>
            <textarea
              name="description"
              value={issue.description}
              onChange={handleChange}
              className="w-full p-3 border rounded-lg h-32 resize-none"
              required
            />
          </div>

          {/* Status field - updated for proper display and submission */}
          <div>
            <label className="block mb-2 text-gray-700">Status</label>
            <select
              name="status"
              value={issue.status}
              onChange={handleChange}
              className="w-full p-3 border rounded-lg"
            >
              <option value="open">To Do</option>
              <option value="in_progress">In Progress</option>
              <option value="resolved">Done</option>
            </select>
          </div>

          <div>
            <label className="block mb-2 text-gray-700">Assignee</label>
            <input
              type="text"
              name="assignee"
              value={issue.assignee || ""}
              onChange={handleChange}
              className="w-full p-3 border rounded-lg"
            />
          </div>

          <div>
            <label className="block mb-2 text-gray-700">Reporter</label>
            <input
              type="text"
              name="reporter"
              value={issue.reporter}
              onChange={handleChange}
              className="w-full p-3 border rounded-lg"
              required
            />
          </div>

          <div>
            <label className="block mb-2 text-gray-700">Issue Type</label>
            <input
              type="text"
              name="issueType"
              value={issue.issueType}
              onChange={handleChange}
              className="w-full p-3 border rounded-lg"
              required
            />
          </div>

          <button
            type="submit"
            disabled={saving}
            className="w-full py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-semibold"
          >
            {saving ? "Saving..." : "Update Issue"}
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditIssue;
