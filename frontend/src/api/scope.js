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

export const fetchAllScopeDetails = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/v1/scopes`)
    if (!response.ok) {
      throw new Error("Failed to fetch all scope details")
    }
    const data = await response.json()
    return data.scopes
  } catch (error) {
    throw new Error(`Error fetching scopes: ${error.message}`);
  }
}

export const createScope = async (name, description, projects) => {
  try {
    const response = await fetch("http://localhost:8080/api/v1/scopes", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name, description, projects }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || "Failed to create scope");
    }

    return await response.json();
  } catch (error) {
    throw new Error(`Error creating scope: ${error.message}`);
  }
};


// Function to fetch all scopes
export const fetchAllScopes = async () => {
  try {
    const response = await fetch("http://localhost:8080/api/v1/scopes");
    if (!response.ok) {
      throw new Error("Failed to fetch scopes");
    }
    const data = await response.json();
    return data.scopes;
  } catch (error) {
    throw new Error(`Error fetching scopes: ${error.message}`);
  }
};


// Function to remove projects from a scope
export const removeProjectsFromScope = async (scopeId, projectKeys) => {
  try {
    const response = await fetch(`http://localhost:8080/api/v1/scopes/${scopeId}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        project_keys: projectKeys,
      }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || "Failed to remove projects");
    }

    return await response.json(); // Return the response on successful deletion
  } catch (error) {
    throw new Error(`Error removing projects from scope: ${error.message}`);
  }
};
