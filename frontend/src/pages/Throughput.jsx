import { useState, useEffect } from "react";
import { fetchAllProjects } from "../api/project";
import { fetchThroughputByProject } from "../api/metrics";
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
  LineChart,
  Line,
  ReferenceLine
} from "recharts";

const Throughput = () => {
  const [selectedProject, setSelectedProject] = useState(null);
  const [projects, setProjects] = useState([]);
  const [throughputData, setThroughputData] = useState(null);
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
    const loadThroughput = async () => {
      if (!selectedProject) return;

      setLoading(true);
      setMessage("");

      try {
        const project = projects.find((p) => p.id === selectedProject);
        const throughputData = await fetchThroughputByProject(project.project_key);
        setThroughputData(throughputData.throughput);
        console.log('Fetched Throughput Data:', throughputData);
      } catch (error) {
        setMessage(`Failed to fetch throughput: ${error.message}`);
      } finally {
        setLoading(false);
      }
    };

    loadThroughput();
  }, [selectedProject, projects]);

  const projectOptions = projects.map((project) => ({
    value: project.id,
    label: `${project.name} (${project.project_key})`,
  }));

  const formattedThroughputData = throughputData
    ? Object.entries(throughputData).map(([week, count]) => ({
        week,
        count,
      }))
    : [];

  console.log('Formatted Throughput Data:', formattedThroughputData);

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 p-6">
      <div className="max-w-5xl w-full bg-white shadow-xl rounded-lg p-8 border border-gray-200">
        <h2 className="text-3xl font-semibold mb-2 text-center text-gray-800">
          Weekly Throughput
        </h2>
        <p className="text-center text-sm text-gray-600 mb-6">
          Throughput measures the number of issues completed within a specific time period (e.g., weekly).
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
            {formattedThroughputData.length > 0 ? (
              <>
                {/* Bar Chart */}
                <div className="mb-8 p-6 bg-white border border-gray-200 rounded-lg shadow-sm">
                  <h3 className="text-xl font-semibold mb-4 text-gray-800">Throughput per Week (Bar Chart)</h3>
                  <ResponsiveContainer width="100%" height={300}>
                    <BarChart data={formattedThroughputData}>
                      <CartesianGrid strokeDasharray="3 3" />
                      <XAxis dataKey="week" />
                      <YAxis />
                      <Tooltip />
                      <Bar dataKey="count" fill="#3b82f6" />
                    </BarChart>
                  </ResponsiveContainer>
                </div>

                {/* Line Chart */}
                <div className="mb-8 p-6 bg-white border border-gray-200 rounded-lg shadow-sm">
                  <h3 className="text-xl font-semibold mb-4 text-gray-800">Throughput Trend (Line Chart)</h3>
                  <ResponsiveContainer width="100%" height={300}>
                    <LineChart data={formattedThroughputData}>
                      <CartesianGrid strokeDasharray="3 3" />
                      <XAxis dataKey="week" />
                      <YAxis />
                      <Tooltip />
                      <Line
                        type="monotone"
                        dataKey="count"
                        stroke="#8884d8"
                        dot={false}
                      />
                      <ReferenceLine y={0} stroke="#000" />
                    </LineChart>
                  </ResponsiveContainer>
                </div>
              </>
            ) : (
              <p>No data available for throughput</p>
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default Throughput;
