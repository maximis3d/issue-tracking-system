package projectassignment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
	"github.com/maximis3d/issue-tracking-system/utils"
)

type Handler struct {
	store types.ProjectAssignmentStore
}

func NewHandler(store types.ProjectAssignmentStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registers all routes related to project assignments.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/projects-assignment/{projectID}/assign/{userID}", h.handleAssignUser).Methods("POST")
	router.HandleFunc("/projects-assignment/{projectID}/remove/{userID}", h.handleRemoveUser).Methods("DELETE")
	router.HandleFunc("/projects-assignment/{projectID}/assigned-users", h.handleGetAssignedUsers).Methods("GET")
	router.HandleFunc("/users", h.handleGetAllUsers).Methods("GET")
}

// handleAssignUser assigns a user to a project.
func (h *Handler) handleAssignUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Convert projectID and userID string to int
	projectIDStr := vars["projectID"]
	userIDStr := vars["userID"]

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid projectID: %v", err))
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid userID: %v", err))
		return
	}

	// Optional: Get the role from query params or request body
	role := r.URL.Query().Get("role")
	if role == "" {
		role = "member" // Default to "member"
	}

	// Attempt to assign user to the project
	if err := h.store.AssignUserToProject(projectID, userID, role); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to assign user to project: %v", err))
		return
	}

	// Success response
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "User assigned to project successfully"})
}

// handleRemoveUser removes a user from a project.
func (h *Handler) handleRemoveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Convert projectID and userID string to int
	projectIDStr := vars["projectID"]
	userIDStr := vars["userID"]

	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid projectID: %v", err))
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid userID: %v", err))
		return
	}

	// Attempt to remove the user from the project
	if err := h.store.RemoveUserFromProject(projectID, userID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to remove user from project: %v", err))
		return
	}

	// Success response
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "User removed from project successfully"})
}

// handleGetAssignedUsers retrieves all users assigned to a project.
func (h *Handler) handleGetAssignedUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Convert projectID string to int
	projectIDStr := vars["projectID"]
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid projectID: %v", err))
		return
	}

	// Get all users assigned to the project
	users, err := h.store.GetUsersForProject(projectID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch assigned users: %v", err))
		return
	}

	// Return list of users
	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to retrive users: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "Users retrieved succesfully",
		"users":   users,
	})

}
