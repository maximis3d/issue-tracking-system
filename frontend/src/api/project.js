export const fetchProjectDetails = async (key) => {
    try {
      const res = await fetch(`http://localhost:8080/api/v1/projects/${key}`);
      if (!res.ok) throw new Error("Project not found");
      return res.json();
    } catch (err) {
      throw new Error(err.message);
    }
  };
  