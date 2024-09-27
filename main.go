package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/Bilbottom/ecom-application/config"
	"github.com/Bilbottom/ecom-application/db"
	"github.com/Bilbottom/ecom-application/domain"
	"github.com/Bilbottom/ecom-application/inbound"
	"github.com/Bilbottom/ecom-application/outbound/datastore"
)

const (
	ExternalPort    = "8080"
	NetworkProtocol = "tcp"
)

func main() {
	mySQLStore, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  NetworkProtocol,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(mySQLStore)

	store := datastore.NewStore(mySQLStore)

	healthChecker := domain.NewHealthChecker()
	userService := domain.NewUserService(store)
	authService := domain.NewAuthService(*userService)
	addressService := domain.NewAddressService(store)
	productService := domain.NewProductService(store)
	orderService := domain.NewOrderService(store)
	cartService := domain.NewCartModifier(
		*addressService,
		*productService,
		*orderService,
	)

	server := inbound.NewServer(
		ExternalPort,
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

func initStorage(database *sql.DB) {
	err := database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully connected!")
}
