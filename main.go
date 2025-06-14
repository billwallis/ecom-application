package main

import (
	"log"

	"github.com/billwallis/ecom-application/config"
	"github.com/billwallis/ecom-application/domain"
	"github.com/billwallis/ecom-application/inbound"
	"github.com/billwallis/ecom-application/outbound/datastore"
)

const (
	NetworkProtocol = "tcp"
)

func main() {
	appConfig := config.NewAppConfig()
	dbConfig := appConfig.DBConfig
	dbConn, err := config.GetDatabaseConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	store := datastore.NewPostgresStore(dbConn)

	healthChecker := domain.NewHealthChecker()
	userService := domain.NewUserService(store)
	authService := domain.NewAuthService(appConfig.AuthConfig, *userService)
	addressService := domain.NewAddressService(store)
	productService := domain.NewProductService(store)
	orderService := domain.NewOrderService(store)
	cartService := domain.NewCartService(
		*addressService,
		*productService,
		*orderService,
	)

	server := inbound.NewServer(
		appConfig,
		*authService,
		healthChecker,
		authService,
		addressService,
		addressService,
		userService,
		userService,
		productService,
		productService,
		cartService,
	)
	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
