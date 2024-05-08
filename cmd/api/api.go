package api

import (
	"database/sql"
	"github.com/Bilbottom/ecom-application/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	// A sub-router allows us to version our API
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter) // prefix routes with `/api/v1`

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
