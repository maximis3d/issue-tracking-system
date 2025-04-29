import React, { useEffect, useState } from "react";
import { assignUserToProject } from "../api/user";
import { fetchAllProjects } from "../api/project";
import { fetchAllUsers } from "../api/user";

const ProjectAssignment = () => {
  const [projects, setProjects] = useState([]);
  const [users, setUsers] = useState([]);
  const [selectedUser, setSelectedUser] = useState("");
  const [selectedProject, setSelectedProject] = useState("");
  const [selectedRole, setSelectedRole] = useState("member"); // Default role is 'member'
  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const projectsData = await fetchAllProjects();
        const usersData = await fetchAllUsers();

        // Access the users array from the response
        if (usersData && Array.isArray(usersData.users)) {
          setUsers(usersData.users); // Set the users array correctly
        } else {
          setMessage("Failed to load users. Invalid data format.");
        }

        setProjects(projectsData);
      } catch (error) {
        setMessage(`Failed to load users or projects: ${error}`);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!selectedUser || !selectedProject) {
      setMessage("Please select both a user and a project.");
      return;
    }

    setLoading(true);
    try {
      await assignUserToProject({
        userId: parseInt(selectedUser),
        projectId: parseInt(selectedProject),
        role: selectedRole, // Include role in the request
      });
      setMessage("User successfully assigned to project.");
    } catch (error) {
      setMessage(`Failed to assign user to project: ${error}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="max-w-4xl mx-auto bg-white shadow-lg rounded-lg p-8 border border-gray-200 w-full">
        <h2 className="text-3xl font-semibold mb-6 text-center text-gray-800">Assign User to Project</h2>

        {message && (
          <div
            className={`p-4 mb-4 text-center rounded-lg ${
              message.includes("successfully") ? "bg-green-100 text-green-600" : "bg-red-100 text-red-600"
            }`}
          >
            {message}
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-6">
          <div className="flex flex-col">
            <label htmlFor="user" className="mb-2 font-medium text-gray-700">Select User</label>
            <select
              id="user"
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={selectedUser}
              onChange={(e) => setSelectedUser(e.target.value)}
              disabled={loading}
            >
              <option value="">-- Choose a user --</option>
              {Array.isArray(users) && users.length > 0 ? (
                users.map((user) => (
                  <option key={user.id} value={user.id}>
                    {user.firstName} {user.lastName} ({user.email})
                  </option>
                ))
              ) : (
                <option value="">No users available</option>
              )}
            </select>
          </div>

          <div className="flex flex-col">
            <label htmlFor="project" className="mb-2 font-medium text-gray-700">Select Project</label>
            <select
              id="project"
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={selectedProject}
              onChange={(e) => setSelectedProject(e.target.value)}
              disabled={loading}
            >
              <option value="">-- Choose a project --</option>
              {projects.map((project) => (
                <option key={project.id} value={project.id}>
                  {project.name} ({project.project_key})
                </option>
              ))}
            </select>
          </div>

          <div className="flex flex-col">
            <label htmlFor="role" className="mb-2 font-medium text-gray-700">Select Role</label>
            <select
              id="role"
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={selectedRole}
              onChange={(e) => setSelectedRole(e.target.value)}
              disabled={loading}
            >
              <option value="member">Member</option>
              <option value="developer">Developer</option>
              <option value="scrum_master">Scrum Master</option>
            </select>
          </div>

          <button
            type="submit"
            className={`w-full py-3 rounded-lg text-white cursor-pointer focus:outline-none transition-colors ${
              loading ? "bg-gray-400 cursor-not-allowed" : "bg-blue-600 hover:bg-blue-700"
            }`}
            disabled={loading}
            
          >
            {loading ? "Assigning..." : "Assign User"}
          </button>
        </form>
      </div>
    </div>
  );
};

export default ProjectAssignment;
