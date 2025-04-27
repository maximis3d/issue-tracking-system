package standups

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
	"github.com/maximis3d/issue-tracking-system/utils"
)

type Handler struct {
	store types.StandupStore
}

func NewHandler(store types.StandupStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/standups/start", h.handleCreateStandup).Methods("POST")
	router.HandleFunc("/standups/end", h.handleEndStandUp).Methods("POST")
}

func (h *Handler) handleCreateStandup(w http.ResponseWriter, r *http.Request) {
	var standup types.Standup

	if err := utils.ParseJSON(r, &standup); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get the last finished standup's end_time
	lastEndTime, err := h.store.GetLastStandupEndTime(standup.ProjectKey)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to retrieve last standup's end time: %v", err))
		return
	}

	// Filter issues that were updated after the last standup ended
	issues, err := h.store.FilterTicketsByEndTime(standup.ProjectKey, lastEndTime)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to filter issues: %v", err))
		return
	}

	// Create the new standup (this will only happen after filtering the issues)
	if err := h.store.CreateStandup(standup); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with filtered issues and the creation success message
	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"message": "Standup Created Successfully",
		"issues":  issues,
	})
}

func (h *Handler) handleEndStandUp(w http.ResponseWriter, r *http.Request) {
	var standup types.Standup

	if err := utils.ParseJSON(r, &standup); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(standup); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid palyad: %v", errors))
	}

	if err := h.store.EndCurrentStandUp(standup); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Standup ended succesfully",
	})

}

func (h *Handler) handleIssueFiltering(w http.ResponseWriter, r *http.Request) {
	var project types.Project

	if err := utils.ParseJSON(r, &project); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if project.ProjectKey == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing project_key from request body"))
		return

	}

	issues, err := h.store.FilterTickets(project)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, issues)
}
