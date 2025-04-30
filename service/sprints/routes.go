package sprints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
	"github.com/maximis3d/issue-tracking-system/utils"
)

type Handler struct {
	store types.SprintStore
}

func NewHandler(store types.SprintStore) *Handler {
	return &Handler{store: store}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/sprints", h.handleCreateSprint).Methods("POST")
	router.HandleFunc("/sprints/{sprintID}/issues/{issueID}", h.handleAddIssueToSprint).Methods("POST")
	router.HandleFunc("/sprints/{sprintID}/issues", h.handleGetIssuesInSprint).Methods("GET")

}

func (h *Handler) handleCreateSprint(w http.ResponseWriter, r *http.Request) {
	var sprint types.Sprint
	// Parse the incoming request JSON payload into a SprintPayload struct
	if err := utils.ParseJSON(r, &sprint); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload structure (using validation library)
	if err := utils.Validate.Struct(sprint); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Create the sprint using the store
	err := h.store.CreateSprint(types.Sprint{
		Name:        sprint.Name,
		Description: sprint.Description,
		StartDate:   sprint.StartDate,
		EndDate:     sprint.EndDate,
		ProjectKey:  sprint.ProjectKey,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Return success response with a message
	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "Sprint created successfully",
	})
}

func (h *Handler) handleAddIssueToSprint(w http.ResponseWriter, r *http.Request) {
	// Extract the issue ID and sprint ID from URL parameters
	vars := mux.Vars(r)
	issueID, err := strconv.Atoi(vars["issueID"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid issue ID"))
		return
	}
	sprintID, err := strconv.Atoi(vars["sprintID"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid sprint ID"))
		return
	}

	// Call store method to add the issue to the sprint
	err = h.store.AddIssueToSprint(issueID, sprintID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add issue to sprint: %v", err))
		return
	}

	// Send success response
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Issue successfully added to sprint",
	})
}

func (h *Handler) handleGetIssuesInSprint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sprintID, err := strconv.Atoi(vars["sprintID"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid sprint ID"))
		return
	}

	issues, err := h.store.GetIssuesInSprint(sprintID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch issues for sprint: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message": "Issues fetched successfully",
		"issues":  issues,
	})
}
