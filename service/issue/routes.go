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
	router.HandleFunc("/issues/{projectKey}", h.handleGetIssuesByProject).Methods("GET")

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
	projectKey := vars["projectKey"]

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
