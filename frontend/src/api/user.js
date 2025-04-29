export const fetchAllUsers = async () => {
    try {
        const res = await fetch(`http://localhost:8080/api/v1/users`);
        if (!res.ok) throw new Error("Cant retrievee users");
        return res.json();
    } catch (err) {
        throw new Error(err.message);
    }
};
export const assignUserToProject = async ({ projectId, userId, role = "member" }) => {
    try {
        const response = await fetch(
            `http://localhost:8080/api/v1/projects-assignment/${projectId}/assign/${userId}?role=${role}`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
            }
        );

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || "Failed to assign user to project.");
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error assigning user:", error);
        throw error;
    }
};