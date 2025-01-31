package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/maximis3d/issue-tracking-system/pkg/user"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newUser.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	userCache := user.NewUserCache()
	userCache.AddUser(newUser)

	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCache := user.NewUserCache()
	user, err := userCache.GetUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new UserCache instance
	userCache := user.NewUserCache()
	if err := userCache.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
