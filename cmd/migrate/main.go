package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Bilbottom/ecom-application/config"
)

func main() {
	db, err := sql.Open(
		"pgx",
		config.NewAppConfig().DBConfig.GetDSN(),
	)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"pgx",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	switch cmd {
	case "up":
		if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	case "down":
		if err = m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
}
