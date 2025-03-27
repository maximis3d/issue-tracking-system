package issue

import (
	"fmt"
	"net/http"

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

	err := h.store.CreateIssue(types.Issue{
		Summary:     issue.Summary,
		Description: issue.Description,
		Project:     issue.Project,
		Reporter:    issue.Reporter,
		Assignee:    issue.Assignee,
		Status:      issue.Status,
		IssueType:   issue.IssueType,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)

}
