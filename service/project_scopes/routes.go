package projectscopes

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
	store types.ScopeStore
}

func NewHandler(store types.ScopeStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/scopes", h.handleCreateScope).Methods("POST")
	router.HandleFunc("/scopes/{id}", h.handleAddProject).Methods("POST")
	router.HandleFunc("/scopes/issues/{id}", h.handleGetIssuesByScope).Methods("GET")

}

func (h *Handler) handleCreateScope(w http.ResponseWriter, r *http.Request) {
	var scope types.Scope

	if err := utils.ParseJSON(r, &scope); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(scope); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	// Call the Store's CreateScope method to insert the new scope
	err := h.store.CreateScope(scope)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create scope: %v", err))
		return
	}

	// Return the newly created scope in the response
	utils.WriteJSON(w, http.StatusCreated, scope)
}

func (h *Handler) handleAddProject(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ProjectKey string `json:"project_key" validate:"required"`
	}

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		fmt.Println("Validation errors:", errors)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	scopeID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error converting ID:", err)
		http.Error(w, "Invalid scope ID", http.StatusBadRequest)
		return
	}

	err = h.store.AddProjectToScope(scopeID, payload.ProjectKey)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add project to scope: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "project added to scope successfully",
	})
}

func (h *Handler) handleGetIssuesByScope(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	scopeId, err := strconv.Atoi(id)

	issues, err := h.store.GetIssuesByScope(scopeId)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("cannot retrive issues %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, issues)

}
