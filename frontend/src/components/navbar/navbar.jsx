import { useNavigate } from 'react-router-dom';

export default function Navbar() {
  const navigate = useNavigate();  // Initialize the navigation hook

  return (
    <nav className="flex justify-between items-center px-6 py-4 shadow-sm">
      <div className="text-xl font-bold">Agile Tracker</div>

      {/* Middle section with "Projects" button */}
      <div className="flex-grow text-center">
        <button
          className="text-gray-700 hover:text-blue-600"
          onClick={() => navigate('/projects')} 
        >
          Projects
        </button>
      </div>

      {/* Right section with "Login" and "Register" buttons */}
      <div className="space-x-4">
        <button
          className="text-gray-700 hover:text-blue-600"
          onClick={() => navigate('/login')}  
        >
          Login
        </button>
        <button
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
          onClick={() => navigate('/register')} 
        >
          Get Started
        </button>
      </div>
    </nav>
  );
}
