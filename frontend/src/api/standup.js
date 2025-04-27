export const startStandup = async (projectKey) => {
  const response = await fetch('http://localhost:8080/api/v1/standups/start', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ project_key: projectKey })
  });

  if (!response.ok) {
    console.error(projectKey)
    throw new Error('Failed to start standup');
   
  }

  return response.json();
};

export const endStandup = async (projectKey) => {
  const response = await fetch('http://localhost:8080/api/v1/standups/end', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ project_key: projectKey })
  });

  if (!response.ok) {
    throw new Error('Failed to end standup');
  }

  return response.json();
};

