export const fetchScopeDetails = async (scopeId) => {
    try {
      const response = await fetch(`http://localhost:8080/api/v1/scopes/details/${scopeId}`);
      if (!response.ok) {
        throw new Error("Failed to fetch scope details");
      }
      return await response.json();
    } catch (error) {
      throw new Error(`Error fetching scope details: ${error.message}`);
    }
  };
  
  export const fetchScopeIssues = async (scopeId) => {
    try {
      const response = await fetch(`http://localhost:8080/api/v1/scopes/issues/${scopeId}`);
      if (!response.ok) {
        throw new Error("Failed to fetch scope issues");
      }
      const data = await response.json();
      return data.issues;
    } catch (error) {
      throw new Error(`Error fetching scope issues: ${error.message}`);
    }
  };
  