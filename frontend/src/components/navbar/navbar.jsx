import { useNavigate } from "react-router-dom";
import { useState } from "react";
import Dropdown from "../dropDown/dropDown";

export default function Navbar() {
  const navigate = useNavigate();
  
  const [openDropdown, setOpenDropdown] = useState(null);

  const handleDropdownToggle = (index) => {
    setOpenDropdown(prev => (prev === index ? null : index));
  };

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
  ];

  const scopeOptions = [
    { label: "All Scopes", onClick: () => navigate("/scopes") },
    { label: "Create Scope", onClick: () => navigate("/create-scope") },
    { label: "Edit Scope", onClick: () => navigate("/edit-scope") }
  ];
  
  const metricOptions = [
    {label: "Cycle Time", onClick: () => navigate("/cycle-time")},
    {label: "Throughput", onClick: () => navigate("/throughput")}
  ]

  return (
    <nav className="flex items-center justify-between px-6 py-4 shadow-sm bg-white">
      <div className="flex items-center space-x-6">
        <div className="text-xl font-bold">Agile Tracker</div>

        <Dropdown
          label="Projects"
          options={projectOptions}
          isOpen={openDropdown === 0}
          onToggle={() => handleDropdownToggle(0)}
        />
        <Dropdown
          label="Issues"
          options={issueOptions}
          isOpen={openDropdown === 1}
          onToggle={() => handleDropdownToggle(1)}
        />
        <Dropdown
          label="Users"
          options={userOptions}
          isOpen={openDropdown === 2}
          onToggle={() => handleDropdownToggle(2)}
        />
        <Dropdown
          label="Scopes"
          options={scopeOptions}
          isOpen={openDropdown === 3}
          onToggle={() => handleDropdownToggle(3)}
        />
        <Dropdown
          label="Metrics"
          options={metricOptions}
          isOpen={openDropdown === 4}
          onToggle={() => handleDropdownToggle(4)}
        />
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
