import { useState, useEffect } from "react";
import { fetchAllProjects } from "../api/project";
import { fetchIssuesByProject } from "../api/issue";
import { fetchCycleTimeByProject } from "../api/metrics";
import Select from "react-select";
import { ClipLoader } from "react-spinners";
import { Link } from "react-router-dom";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const CycleTime = () => {
  const [selectedProject, setSelectedProject] = useState(null);
  const [projects, setProjects] = useState([]);
  const [issues, setIssues] = useState([]);
  const [averageCycleTime, setAverageCycleTime] = useState(null);
  const [message, setMessage] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const getProjects = async () => {
      try {
        const data = await fetchAllProjects();
        setProjects(data.projects);
      } catch (error) {
        setMessage(`Failed to load projects: ${error.message}`);
      }
    };
    getProjects();
  }, []);

  useEffect(() => {
    const loadCycleTime = async () => {
      if (!selectedProject) return;

      setLoading(true);
      setMessage("");

      try {
        const project = projects.find((p) => p.id === selectedProject);
        const issueData = await fetchIssuesByProject(project.project_key);

        const resolvedIssues = (issueData.issues || []).filter(
          (issue) => issue.status.toLowerCase() === "resolved"
        );

        setIssues(resolvedIssues);

        const cycleData = await fetchCycleTimeByProject(project.project_key);
        if (cycleData.cycleTime) {
          setAverageCycleTime(cycleData.cycleTime.average_duration || null);
        }
      } catch (error) {
        setMessage(`Failed to fetch cycle time: ${error.message}`);
      } finally {
        setLoading(false);
      }
    };

    loadCycleTime();
  }, [selectedProject, projects]);

  const projectOptions = projects.map((project) => ({
    value: project.id,
    label: `${project.name} (${project.project_key})`,
  }));

  const sortedByCycle = [...issues].sort(
    (a, b) => (b.cycle_time || 0) - (a.cycle_time || 0)
  );
  const longest = sortedByCycle[0];
  const shortest = sortedByCycle[sortedByCycle.length - 1];

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 p-6">
      <div className="max-w-5xl w-full bg-white shadow-xl rounded-lg p-8 border border-gray-200">
        <h2 className="text-3xl font-semibold mb-2 text-center text-gray-800">
          Cycle Time Visualisation
        </h2>
        <p className="text-center text-sm text-gray-600 mb-6">
          Cycle Time measures the time from when work starts on a task until it is delivered (e.g., from "In Progress" to "Resolved").
        </p>

        {message && (
          <div className="mb-4 text-red-600 text-center bg-red-100 p-3 rounded-lg shadow-sm">
            {message}
          </div>
        )}

        <div className="mb-6">
          <label className="block mb-2 text-gray-700 font-medium">Select Project</label>
          <Select
            options={projectOptions}
            value={projectOptions.find((opt) => opt.value === selectedProject)}
            onChange={(option) => setSelectedProject(option?.value || null)}
            placeholder="Select a project..."
            isDisabled={loading}
            className="react-select-container"
          />
        </div>

        {loading ? (
          <div className="flex justify-center items-center">
            <ClipLoader size={50} color="#3498db" loading={loading} />
          </div>
        ) : (
          <div>
            {averageCycleTime && (
              <div className="mb-8 p-6 bg-blue-50 border border-blue-200 rounded-lg shadow-sm">
                <div className="text-2xl font-bold text-blue-800 mb-2">Average Cycle Time</div>
                <div className="text-sm text-gray-600 mb-2">
                  Time between issue start (e.g., "In Progress") and resolution.
                </div>
                <div className="text-3xl font-semibold text-blue-900">{averageCycleTime}</div>
              </div>
            )}

            {issues.length > 0 && (
              <>
                <h3 className="text-xl font-semibold mb-4 text-gray-800">Resolved Issues</h3>
                <ul className="space-y-4 mb-8">
                  {issues.map((issue) => (
                    <Link key={issue.id} to={`/issues/${issue.id}`} className="block">
                      <li
                        className={`relative p-4 rounded-lg shadow-md transition duration-300 ease-in-out transform hover:scale-105 ${
                          issue.id === longest?.id
                            ? "bg-red-50 border-l-4 border-red-500"
                            : issue.id === shortest?.id
                            ? "bg-green-50 border-l-4 border-green-500"
                            : "bg-gray-50 hover:bg-gray-200"
                        }`}
                      >
                        <div className="font-semibold text-gray-800">
                          {issue.key}: {issue.summary}
                        </div>
                        <div className="text-gray-600">Status: {issue.status}</div>
                        <div className="text-gray-700">
                          Cycle Time: {issue.cycle_time || "N/A"}{" "}
                          <span className="text-sm text-gray-500">(start → resolved)</span>
                        </div>

                        {issue.id === longest?.id && (
                          <span className="absolute top-2 right-2 text-xs font-bold text-red-700 bg-red-200 px-2 py-1 rounded-full shadow">
                            Longest Cycle Time
                          </span>
                        )}
                        {issue.id === shortest?.id && (
                          <span className="absolute top-2 right-2 text-xs font-bold text-green-700 bg-green-200 px-2 py-1 rounded-full shadow">
                            Shortest Cycle Time
                          </span>
                        )}
                      </li>
                    </Link>
                  ))}
                </ul>

                <div className="bg-white p-6 rounded-lg shadow border">
                  <h4 className="text-lg font-semibold mb-4 text-gray-800">
                    Cycle Time per Issue (Chart)
                  </h4>
                  <ResponsiveContainer width="100%" height={300}>
                    <BarChart
                      data={issues.map((issue) => ({
                        name: issue.key,
                        cycle_time: parseFloat(issue.cycle_time) || 0,
                      }))}
                    >
                      <CartesianGrid strokeDasharray="3 3" />
                      <XAxis dataKey="name" />
                      <YAxis label={{ value: "Days", angle: -90, position: "insideLeft" }} />
                      <Tooltip />
                      <Bar dataKey="cycle_time" fill="#3b82f6" />
                    </BarChart>
                  </ResponsiveContainer>
                </div>
              </>
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default CycleTime;
