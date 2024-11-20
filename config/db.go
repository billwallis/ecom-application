package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func GetDatabaseConnection(config DBConfig) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), config.GetDSN())
	if err != nil {
		log.Fatalf("Database error: %s", err)
		return nil, err
	}

	return conn, err
}
