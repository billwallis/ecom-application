package api

import (
	"database/sql"
	"github.com/Bilbottom/ecom-application/service/address"
	_ "github.com/Bilbottom/ecom-application/docs"
	"github.com/Bilbottom/ecom-application/service/cart"
	"github.com/Bilbottom/ecom-application/service/order"
	"github.com/Bilbottom/ecom-application/service/product"
	"github.com/Bilbottom/ecom-application/service/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

// Run godoc
//
// @title           E-Commerce Application
// @version         1.0
// @description     An e-commerce application built with Go.
//
// @contact.name    Bilbottom
// @contact.url     https://github.com/Bilbottom
//
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host            localhost:8080
// @BasePath        /api/v1
func (s *WebServer) Run() error {
	router := mux.NewRouter()
	router.Use(requestLogger)
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	addressStore := address.NewStore(s.db)
	addressHandler := address.NewHandler(addressStore, userStore)
	addressHandler.RegisterRoutes(subRouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subRouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore, addressStore, productStore, userStore)
	cartHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

// requestLogger prints the incoming request method and URL
func requestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
