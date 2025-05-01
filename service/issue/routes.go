package issue

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
	"github.com/maximis3d/issue-tracking-system/utils"
)

type Handler struct {
	store types.IssueStore
}

func NewHandler(store types.IssueStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/createIssue", h.handleCreateIssue).Methods("POST")
	router.HandleFunc("/issues/{id}", h.handleUpdateIssue).Methods("PUT")
	router.HandleFunc("/issue/{id}", h.handleGetIssueById).Methods("GET")

	router.HandleFunc("/issues/{key}", h.handleGetIssuesByProject).Methods("GET")
	router.HandleFunc("/cycle-time/{project_key}", h.handleGetAverageCycleTime).Methods("GET")
	router.HandleFunc("/throughput/{project_key}", h.handleGetWeeklyThroughput).Methods("GET")
}

func (h *Handler) handleCreateIssue(w http.ResponseWriter, r *http.Request) {
	var issue types.IssuePayload
	if err := utils.ParseJSON(r, &issue); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(issue); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	newIssue := types.Issue{
		Summary:     issue.Summary,
		Description: issue.Description,
		ProjectKey:  issue.ProjectKey,
		Reporter:    issue.Reporter,
		Assignee:    issue.Assignee,
		Status:      issue.Status,
		IssueType:   issue.IssueType,
	}

	err := h.store.CreateIssue(newIssue)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "Issue created successfully",
	})
}

func (h *Handler) handleUpdateIssue(w http.ResponseWriter, r *http.Request) {
	var issue types.Issue

	if err := utils.ParseJSON(r, &issue); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if err := utils.Validate.Struct(issue); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		fmt.Println("Validation errors:", errors)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	issueID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting ID:", err)
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	existingIssue, err := h.store.GetIssueByID(issueID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("issue not found"))
		fmt.Println(err)
		return
	}

	issue.ID = existingIssue.ID

	if err := h.store.UpdateIssue(issue); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to update issue"))
		fmt.Println(err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Issue updated successfully",
		"issue":   issue,
	})
}

func (h *Handler) handleGetIssuesByProject(w http.ResponseWriter, r *http.Request) {
	// Extract the project key from URL parameters
	vars := mux.Vars(r)
	projectKey := vars["key"]

	// Fetch the issues for the given project from the store
	issues, err := h.store.GetIssuesByProject(projectKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No issues found for the provided project key
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no issues found for project %s", projectKey))
			fmt.Println("No issues found for project:", projectKey)
		} else {
			// An error occurred while fetching the issues
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch issues for project %s: %v", projectKey, err))
			fmt.Println("Error fetching issues for project:", projectKey, err)
		}
		return
	}

	// Return the issues in a JSON response
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message":    "Issues fetched successfully",
		"projectKey": projectKey,
		"issues":     issues,
	})
}

func (h *Handler) handleGetIssueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]
	issueID, err := strconv.Atoi(Id)

	if err != nil {
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	issue, err := h.store.GetIssueByID(issueID)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("issue not found"))
		return
	}

	// Adding cycle time to the response
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message":   "Issue fetched successfully",
		"issue":     issue,
		"cycleTime": issue.CycleTime,
	})
}

func (h *Handler) handleGetAverageCycleTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectKey := vars["project_key"]
	if projectKey == "" {
		http.Error(w, "Missing project_key", http.StatusBadRequest)
		return
	}

	avgCycleTime, err := h.store.GetAverageCycleTime(projectKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get average cycle time: %v", err), http.StatusInternalServerError)
		return
	}

	type response struct {
		ProjectKey      string  `json:"project_key"`
		AverageSeconds  float64 `json:"average_seconds"`
		AverageDuration string  `json:"average_duration"`
	}

	resp := response{
		ProjectKey:      projectKey,
		AverageSeconds:  avgCycleTime.Seconds(),
		AverageDuration: avgCycleTime.String(),
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message":   "Average cycle time fetched successfully",
		"cycleTime": resp,
	})
}

func (h *Handler) handleGetWeeklyThroughput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectKey := vars["project_key"]
	if projectKey == "" {
		http.Error(w, "Missing project_key", http.StatusBadRequest)
		return
	}

	data, err := h.store.GetWeeklyThroughput(projectKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get throughput: %v", err), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"message":    "Weekly throughput fetched successfully",
		"throughput": data,
	})
}
