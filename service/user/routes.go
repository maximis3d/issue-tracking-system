package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/types"
	"github.com/maximis3d/issue-tracking-system/utils"
)

type Handler struct {
	stre types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{ store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	return nil
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload

	var payload types.RegisterUserPayload

	if err := utils.ParseJson(r.Body, payload), err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if user exists
	h.store.GetUserByEmail(payload.Email)



	
	// if not, create new user

}
