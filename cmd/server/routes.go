package main

import "net/http"

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/users", createUser)
	mux.HandleFunc("/users/", getUser)
	mux.HandleFunc("/users/", deleteUser)

	return mux
}
