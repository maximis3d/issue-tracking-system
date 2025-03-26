package issue

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
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

}
