export const fetchIssues = async (key) => {
  try {
    const res = await fetch(`http://localhost:8080/api/v1/issues/${key}`);
    if (!res.ok) throw new Error("Issues not found");
    const data = await res.json();
    return data.issues;
  } catch (err) {
    throw new Error(err.message);
  }
};

export const fetchIssue = async (id) => {
  try {
    const res = await fetch(`http://localhost:8080/api/v1/issue/${id}`)
    const data = await res.json();
    return data.issue
  } catch (err) {
    throw new Error(err.message)
  }
}