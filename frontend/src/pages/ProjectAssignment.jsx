import React, { useEffect, useState } from "react";
import { assignUserToProject, removeUserFromProject, fetchAllUsers } from "../api/user";
import { fetchAllProjects } from "../api/project";

const ProjectAssignment = () => {
  const [projects, setProjects] = useState([]);
  const [users, setUsers] = useState([]);
  const [selectedUser, setSelectedUser] = useState("");
  const [selectedProject, setSelectedProject] = useState("");
  const [selectedRole, setSelectedRole] = useState("member");
  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);
  const [isAssigning, setIsAssigning] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const projectsData = await fetchAllProjects();
        const usersData = await fetchAllUsers();

        if (usersData && Array.isArray(usersData.users)) {
          setUsers(usersData.users);
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

  // Handle Assign User functionality
  const handleAssignSubmit = async (e) => {
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
        role: selectedRole,
      });
      setMessage("User successfully assigned to project.");
    } catch (error) {
      setMessage(`Failed to assign user to project: ${error}`);
    } finally {
      setLoading(false);
    }
  };

  // Handle Remove User functionality
  const handleRemoveUser = async () => {
    if (!selectedUser || !selectedProject) {
      setMessage("Please select both a user and a project to remove.");
      return;
    }

    const confirmRemoval = window.confirm(
      "Are you sure you want to remove this user from the project?"
    );
    if (!confirmRemoval) {
      return; // If the user cancels, do nothing
    }

    setLoading(true);
    try {
      await removeUserFromProject({
        userId: parseInt(selectedUser),
        projectId: parseInt(selectedProject),
      });
      setMessage("User successfully removed from project.");
    } catch (error) {
      setMessage(`Failed to remove user from project: ${error}`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="max-w-4xl mx-auto bg-white shadow-lg rounded-lg p-8 border border-gray-200 w-full">
        <h2 className="text-3xl font-semibold mb-6 text-center text-gray-800">
          {isAssigning ? "Assign/Remove User from Project" : "Remove User from Project"}
        </h2>

        {/* Display Message */}
        {message && (
          <div
            className={`p-4 mb-4 text-center rounded-lg ${
              message.includes("successfully")
                ? "bg-green-100 text-green-600"
                : "bg-red-100 text-red-600"
            }`}
          >
            {message}
          </div>
        )}

        {/* Tabs for Toggle between Assign and Remove */}
        <div className="flex justify-center mb-6">
          <button
            className={`px-4 py-2 font-medium text-lg ${
              isAssigning ? "bg-blue-600 text-white" : "bg-gray-200 text-gray-700"
            } rounded-l-lg`}
            onClick={() => setIsAssigning(true)}
            disabled={loading}
          >
            Assign User
          </button>
          <button
            className={`px-4 py-2 font-medium text-lg ${
              !isAssigning ? "bg-red-600 text-white" : "bg-gray-200 text-gray-700"
            } rounded-r-lg`}
            onClick={() => setIsAssigning(false)}
            disabled={loading}
          >
            Remove User
          </button>
        </div>

        {/* Form for Assigning or Removing */}
        <form
          onSubmit={isAssigning ? handleAssignSubmit : (e) => e.preventDefault()}
          className="space-y-6"
        >
          <div className="flex flex-col">
            <label htmlFor="user" className="mb-2 font-medium text-gray-700">
              Select User
            </label>
            <select
              id="user"
              className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              value={selectedUser}
              onChange={(e) => setSelectedUser(e.target.value)}
              disabled={loading}
            >
              <option value="">-- Choose a user --</option>
              {users.map((user) => (
                <option key={user.id} value={user.id}>
                  {user.firstName} {user.lastName} ({user.email})
                </option>
              ))}
            </select>
          </div>

          <div className="flex flex-col">
            <label htmlFor="project" className="mb-2 font-medium text-gray-700">
              Select Project
            </label>
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

          {isAssigning && (
            <div className="flex flex-col">
              <label htmlFor="role" className="mb-2 font-medium text-gray-700">
                Select Role
              </label>
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
          )}

          <button
            type="submit"
            className={`w-full py-3 rounded-lg text-white cursor-pointer focus:outline-none transition-colors ${
              loading ? "bg-gray-400 cursor-not-allowed" : "bg-blue-600 hover:bg-blue-700"
            }`}
            disabled={loading}
            onClick={isAssigning ? handleAssignSubmit : handleRemoveUser}
          >
            {loading ? (isAssigning ? "Assigning..." : "Removing...") : (isAssigning ? "Assign User" : "Remove User")}
          </button>
        </form>
      </div>
    </div>
  );
};

export default ProjectAssignment;
