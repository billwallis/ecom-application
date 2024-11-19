<div align="center">

[![Go](https://img.shields.io/badge/Go-1.22-blue.svg)](https://go.dev/)
[![GitHub last commit](https://img.shields.io/github/last-commit/Bilbottom/ecom-application)](https://shields.io/badges/git-hub-last-commit)
[![Tests](https://github.com/Bilbottom/ecom-application/actions/workflows/tests.yml/badge.svg)](https://github.com/Bilbottom/ecom-application/actions/workflows/tests.yml)

</div>

# E-commerce Application

An e-commerce application built with Go.

This is from the following YouTube tutorial:

- https://youtu.be/7VLmLOiQ3ck

The corresponding repository is:

- https://github.com/sikozonpc/ecom

## Commands

I'm on Windows (and the YouTube tutorial is not), so I can't add a Makefile. Instead, since I'm using GoLand, I am just adding run configurations.

However, the corresponding commands still need to be documented somewhere, so here they are:

```shell
# build & run
go build -o bin/ecom main.go
go run main.go

# ...alternatively, run via Docker
docker compose up --detach
docker compose down --volumes  # when you're done

# test
go test ./...

# create migration (needs the golang-migrate CLI installed)
migrate create -ext sql -dir /migrations/ <migration_name>

# migrate up/down (via Docker)
docker exec EComApp    go run cmd/migrate/main.go up
docker exec EComApp    go run cmd/migrate/main.go down
```

## Database

The database is MySQL on port `3306`. Make sure the following environment variables are set in your runtime environment:

- `PUBLIC_HOST`: `localhost` (the default is `http://localhost`)
- `DB_PORT`: `5432`
- `DB_USERNAME`: `postgres`
- `DB_PASSWORD`: `postgres`
