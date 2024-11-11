package config

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(config mysql.Config) (*sql.DB, error) {
	dsn := config.FormatDSN()
	log.Println("Connecting to the database at:", dsn)

	if err := createDatabase(config); err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func createDatabase(config mysql.Config) error {
	dbName := config.DBName
	config.DBName = ""

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return err
	}

	_, err = db.Exec("create database if not exists " + dbName)
	if err != nil {
		return err
	}

	return nil
}
