package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/api/data", apitDataHandler)

	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page")
}

func apitDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data from the api"

	fmt.Fprintln(w, data)
}
