package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/Bilbottom/ecom-application/config"
	"github.com/Bilbottom/ecom-application/domain"
	"github.com/Bilbottom/ecom-application/inbound"
	"github.com/Bilbottom/ecom-application/outbound/datastore"
)

const (
	NetworkProtocol = "tcp"
)

func main() {
	appConfig := config.NewAppConfig()
	dbConfig := appConfig.DBConfig
	mySQLStore, err := config.NewMySQLStorage(mysql.Config{
		User:                 appConfig.DBConfig.User,
		Passwd:               appConfig.DBConfig.Password,
		Addr:                 fmt.Sprintf("%s:%s", dbConfig.Host, dbConfig.Port),
		DBName:               appConfig.DBConfig.Name,
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

func initStorage(database *sql.DB) {
	err := database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully connected!")
}
