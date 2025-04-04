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

	if err := utils.Validate.Struct(standup); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	newStandup := types.Standup{
		ProjectKey: standup.ProjectKey,
	}

	if err := h.store.CreateStandup(newStandup); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "Standup Created Successfully",
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
