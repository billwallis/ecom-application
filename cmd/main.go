package main

import (
	"database/sql"
	"github.com/Bilbottom/ecom-application/cmd/api"
	"github.com/Bilbottom/ecom-application/config"
	"github.com/Bilbottom/ecom-application/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	store, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(store)

	server := api.NewWebServer(":8080", store)
	if err := server.Run(); err != nil {
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
