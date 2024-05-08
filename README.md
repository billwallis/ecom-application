# E-commerce Application

An e-commerce application built with Go.

This is from the following YouTube tutorial:

- https://youtu.be/7VLmLOiQ3ck

## Commands

I'm on Windows (and the YouTube tutorial is not), so I can't add a Makefile. Instead, since I'm using GoLand, I am just adding run configurations.

However, the corresponding commands still need to be documented somewhere, so here they are:

```shell
# build & run
go build -o bin/ecom cmd/main.go
go run cmd/main.go

# test
go test ./...

# create migration (needs the golang-migrate CLI installed)
migrate create -ext sql -dir /migrations/ <migration_name>

# migrate up/down
go run cmd/migrate/main.go up
go run cmd/migrate/main.go down
```

## Database

The database is MySQL on port `3306`. Make sure the following environment variables are set in your runtime environment:

- `PORT`: `3306`
- `DB_USER`: `root`
- `DB_PASSWORD`: `password`
