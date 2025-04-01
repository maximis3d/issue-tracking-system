package issue

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	var payload types.IssueUpdatePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	vars := mux.Vars(r)
	issueID, err := strconv.Atoi(vars["id"])
	fmt.Println("Extracted issue ID:", issueID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid issue ID"))
		return
	}

	existingIssue, err := h.store.GetIssueByID(issueID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("issue not found"))
		return
	}

	if payload.Summary != nil {
		existingIssue.Summary = *payload.Summary
	}
	if payload.Description != nil {
		existingIssue.Description = *payload.Description
	}
	if payload.ProjectKey != nil {
		existingIssue.ProjectKey = *payload.ProjectKey
	}
	if payload.Reporter != nil {
		existingIssue.Reporter = *payload.Reporter
	}
	if payload.Assignee != nil {
		existingIssue.Assignee = *payload.Assignee
	}
	if payload.Status != nil {
		existingIssue.Status = *payload.Status
	}
	if payload.IssueType != nil {
		existingIssue.IssueType = *payload.IssueType
	}

	existingIssue.UpdatedAt = time.Now()

	if err := h.store.UpdateIssue(*existingIssue); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to update issue"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Issue updated successfully",
	})
}
