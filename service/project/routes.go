package project

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
)

type Handler struct {
	store types.ProjectStore
}

func NewHandler(store types.ProjectStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/projects", h.handleGetProjects).Methods("GET")
	router.HandleFunc("/projects", h.handleCreateProject).Methods("POST")
	router.HandleFunc("/projects/{id}", h.handleGetProject).Methods("GET")
}

func (h *Handler) handleGetProjects(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleCreateProject(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleGetProject(w http.ResponseWriter, r *http.Request) {
}
