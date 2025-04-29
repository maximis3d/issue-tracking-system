import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { fetchAllScopeDetails } from "../api/scope";

const Scopes = () => {
  const [scopes, setScopes] = useState([]);
  const [search, setSearch] = useState("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetchAllScopeDetails()
      .then((data) => {
        setScopes(data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  const filteredScopes = scopes.filter((scope) =>
    scope.name.toLowerCase().includes(search.toLowerCase())
  );

  if (loading) return <div className="text-center p-8">Loading scopes...</div>;
  if (error) return <div className="text-center text-red-500 p-8">Error: {error}</div>;

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-600 to-pink-700 text-white flex flex-col">
      <header className="text-center py-12">
        <h1 className="text-4xl font-bold mb-4">Scopes</h1>
        <p className="text-lg max-w-2xl mx-auto mb-8">
          View all current scopes and associated project tracking details.
        </p>
      </header>

      <div className="max-w-4xl mx-auto p-4 w-full">
        <input
          type="text"
          placeholder="Search scopes..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="border rounded-lg p-2 w-full mb-8"
        />

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          {filteredScopes.length === 0 ? (
            <p>No scopes found.</p>
          ) : (
            filteredScopes.map((scope) => (
              <div
                key={scope.id}
                onClick={() => navigate(`/scope/${scope.id}`)}
                className="bg-white text-gray-900 shadow-md rounded-lg p-4 cursor-pointer hover:shadow-lg transition"
              >
                <h2 className="text-xl font-semibold mb-2">{scope.name}</h2>
                <p className="text-sm text-gray-600 mb-2">{scope.description}</p>
                <p className="text-sm text-gray-500">
                  Created: {new Date(scope.created_at).toLocaleDateString()}
                </p>
              </div>
            ))
          )}
        </div>
      </div>

      <footer className="text-center py-6 text-gray-300 mt-auto">
        © {new Date().getFullYear()} StandUpTracker
      </footer>
    </div>
  );
};

export default Scopes;
