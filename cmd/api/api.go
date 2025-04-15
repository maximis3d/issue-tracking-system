package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/maximis3d/issue-tracking-system/service/issue"
	"github.com/maximis3d/issue-tracking-system/service/project"
	projectscopes "github.com/maximis3d/issue-tracking-system/service/project_scopes"
	"github.com/maximis3d/issue-tracking-system/service/standups"
	"github.com/maximis3d/issue-tracking-system/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	projectStore := project.NewStore(s.db)
	projectHandler := project.NewHandler(projectStore)
	projectHandler.RegisterRoutes(subrouter)

	issueStore := issue.NewStore(s.db)
	issueHandler := issue.NewHandler(issueStore)
	issueHandler.RegisterRoutes(subrouter)

	standupStore := standups.NewStore(s.db)
	standupHandler := standups.NewHandler(standupStore)
	standupHandler.RegisterRoutes(subrouter)

	scopeStore := projectscopes.NewStore(s.db)
	scopeHandler := projectscopes.NewHandler(scopeStore)
	scopeHandler.RegisterRoutes(subrouter)

	// Enable CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, corsHandler(router))
}
