import { useState } from "react";

const createIssue = async (issue) => {
  try {
    const response = await fetch("http://localhost:8080/api/v1/createIssue", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(issue),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || "Failed to create Issue");
    }

    return await response.json();
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

const CreateIssue = () => {
  const [form, setForm] = useState({
    summary: "",
    key: "",
    description: "",
    project_key: "",
    reporter: "",
    assignee: "",
    status: "",
    issueType: ""
  });
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);

    try {
      await createIssue({
        summary: form.summary,
        description: form.description,
        project_key: form.project_key,
        reporter: form.reporter,
        assignee: form.assignee,
        status: form.status,
        issueType: form.issueType
      });

      setSuccess("Issue created successfully!");
      setForm({ summary: "", description: "", project_key: "", reporter: "", assignee: "", status: "", issueType: "" });
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="flex min-h-screen items-center justify-center flex-col px-6 py-12 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="text-center text-2xl font-bold text-gray-900">Create New Issue</h2>
      </div>

      <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-md">
        <form className="space-y-4" onSubmit={handleSubmit}>
          {error && <p className="text-red-500 text-sm">{error}</p>}
          {success && <p className="text-green-600 text-sm">{success}</p>}

          {[
            { id: "summary", label: "Issue Summary" },
            { id: "description", label: "Description" },
            { id: "project_key", label: "Project" },
            { id: "reporter", label: "Reporter" },
            { id: "assignee", label: "Assignee" },
            { id: "status", label: "status" },
            { id: "issueType", label: "Issue Type" },


          ].map(({ id, label, type = "text" }) => (
            <div key={id}>
              <label htmlFor={id} className="block text-sm font-medium text-gray-900">
                {label}
              </label>
              <div className="mt-1">
                <input
                  type={type}
                  id={id}
                  name={id}
                  required={id !== "description"}
                  value={form[id]}
                  onChange={(e) => setForm({ ...form, [id]: e.target.value })}
                  className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline outline-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:outline-indigo-600 sm:text-sm"
                />
              </div>
            </div>
          ))}

          <div>
            <button
              type="submit"
              className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white hover:bg-indigo-500 focus:outline-indigo-600"
            >
              Create Issue
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateIssue;
