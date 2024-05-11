package api

import (
	"database/sql"
	"github.com/Bilbottom/ecom-application/service/cart"
	"github.com/Bilbottom/ecom-application/service/order"
	"github.com/Bilbottom/ecom-application/service/product"
	"github.com/Bilbottom/ecom-application/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WebServer struct {
	addr string
	db   *sql.DB
}

func NewWebServer(addr string, db *sql.DB) *WebServer {
	return &WebServer{
		addr: addr,
		db:   db,
	}
}

func (s *WebServer) Run() error {
	router := mux.NewRouter()
	// A sub-router allows us to version our API
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter) // prefix routes with `/api/v1`

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subRouter) // prefix routes with `/api/v1`

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subRouter) // prefix routes with `/api/v1`

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
