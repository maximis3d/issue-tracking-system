import { useState, useEffect } from "react";
import { fetchAllProjects } from "../api/project";
import { fetchIssuesByProject } from "../api/issue";
import { fetchCycleTimeByProject } from "../api/metrics";
import Select from "react-select";
import { ClipLoader } from "react-spinners"; 
import { Link } from "react-router-dom";

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

        // Filter resolved issues and add cycle time
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

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 p-6">
      <div className="max-w-4xl w-full bg-white shadow-xl rounded-lg p-8 border border-gray-200">
        <h2 className="text-3xl font-semibold mb-6 text-center text-gray-800">
          Cycle Time Visualisation
        </h2>

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
            {issues.length > 0 && (
              <div>
                <h3 className="text-xl font-semibold mb-4 text-gray-800">
                  Resolved Issues
                </h3>
                <ul className="space-y-4">
                  {issues.map((issue) => (
                    <Link
                      key={issue.id}
                      to={`/issues/${issue.id}`}
                      className="block"
                    >
                      <li
                        className="p-4 bg-gray-50 rounded-lg shadow-md hover:bg-gray-200 transition duration-300 ease-in-out transform hover:scale-105"
                      >
                        <div className="font-semibold text-gray-800">
                          {issue.key}: {issue.summary}
                        </div>
                        <div className="text-gray-600">Status: {issue.status}</div>
                        <div className="text-gray-700">Cycle Time: {issue.cycle_time || "N/A"}</div>
                      </li>
                    </Link>
                  ))}
                </ul>

              </div>
            )}

            {averageCycleTime && (
              <div className="mt-6 text-lg text-gray-800 font-semibold">
                <div className="text-xl font-bold mb-2">Average Cycle Time</div>
                <div className="text-lg text-gray-600">{averageCycleTime}</div>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default CycleTime;
