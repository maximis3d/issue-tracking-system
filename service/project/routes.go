package project

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
	"github.com/maximis3d/issue-tracking-system/utils"
)

type Handler struct {
	store types.ProjectStore
}

func NewHandler(store types.ProjectStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/projects", h.handleGetProjects).Methods("GET")
	router.HandleFunc("/projects/{name}", h.handleGetProjectByName).Methods("GET")
	router.HandleFunc("/projects", h.handleCreateProject).Methods("POST")
}

func (h *Handler) handleGetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.store.GetProjects()

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("no projects found"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, projects)
}

func (h *Handler) handleGetProjectByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	project, err := h.store.GetProjectByName(name)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("project not found"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, project)

}

func (h *Handler) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	var project types.ProjectPayload
	if err := utils.ParseJSON(r, &project); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(project); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	//check if project exits
	_, err := h.store.GetProjectByName(project.Name)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("project already exists"))
		return
	}
	err = h.store.CreateProject(types.Project{
		Name:        project.Name,
		Description: project.Description,
		ProjectLead: project.ProjectLead,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "Project created successfully"})

}
