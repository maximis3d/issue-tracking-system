import { useNavigate } from "react-router-dom";
import Dropdown from "../dropDown/dropDown";

export default function Navbar() {
  const navigate = useNavigate();

  const projectOptions = [
    { label: "All Projects", onClick: () => navigate("/projects") },
    { label: "Create New Project", onClick: () => navigate("/create-project") },
  ];

  const issueOptions = [
    { label: "All Issues", onClick: () => navigate("/issues") },
    { label: "Create Issue", onClick: () => navigate("/create-issue") },
  ];

  const userOptions = [
    { label: "Assign Users", onClick: () => navigate("/assign-user") },
  ]
  const scopeOptions = [
    {label: "All Scopes", onClick: () => navigate("/scopes")},
    {label: "Create Scope", onClick: () => navigate("/create-scope")},
  ]

  return (
    <nav className="flex items-center justify-between px-6 py-4 shadow-sm bg-white">
      <div className="flex items-center space-x-6">
        <div className="text-xl font-bold">Agile Tracker</div>

        <Dropdown label="Projects" options={projectOptions} />
        <Dropdown label="Issues" options={issueOptions} />
        <Dropdown label="Users" options={userOptions} />
        <Dropdown label="Scopes" options={scopeOptions} />

      </div>

      {/* Right Section: Auth buttons */}
      <div className="space-x-4 flex items-center">
        <button
          className="text-gray-700 hover:text-blue-600"
          onClick={() => navigate("/login")}
        >
          Login
        </button>
        <button
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
          onClick={() => navigate("/register")}
        >
          Get Started
        </button>
      </div>
    </nav>
  );
}
