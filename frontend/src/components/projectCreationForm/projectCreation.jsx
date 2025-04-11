import { useState } from "react";

const createProject = async (project) => {
  try {
    const response = await fetch("http://localhost:8080/api/v1/projects", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(project),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || "Failed to create project");
    }

    return await response.json();
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

const CreateProject = () => {
  const [form, setForm] = useState({
    project_key: "",
    name: "",
    description: "",
    projectLead: "",
    wip_limit: "",
  });
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);

    try {
      await createProject({
        project_key: form.project_key,
        name: form.name,
        description: form.description,
        projectLead: parseInt(form.projectLead),
        wip_limit: parseInt(form.wip_limit),
      });

      setSuccess("Project created successfully!");
      setForm({ project_key: "", name: "", description: "", projectLead: "", wip_limit: "" });
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="flex min-h-screen items-center justify-center flex-col px-6 py-12 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="text-center text-2xl font-bold text-gray-900">Create New Project</h2>
      </div>

      <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-md">
        <form className="space-y-4" onSubmit={handleSubmit}>
          {error && <p className="text-red-500 text-sm">{error}</p>}
          {success && <p className="text-green-600 text-sm">{success}</p>}

          {[
            { id: "project_key", label: "Project Key" },
            { id: "name", label: "Project Name" },
            { id: "description", label: "Description" },
            { id: "projectLead", label: "Project Lead ID", type: "number" },
            { id: "wip_limit", label: "WIP Limit", type: "number" },
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
              Create Project
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateProject;
