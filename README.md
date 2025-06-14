<div align="center">

[![Go](https://img.shields.io/badge/Go-1.22-blue.svg)](https://go.dev/)
[![GitHub last commit](https://img.shields.io/github/last-commit/billwallis/ecom-application)](https://shields.io/badges/git-hub-last-commit)
[![Tests](https://github.com/billwallis/ecom-application/actions/workflows/tests.yml/badge.svg)](https://github.com/billwallis/ecom-application/actions/workflows/tests.yml)

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
```

## Database

The database is PostgreSQL on port `5432`. After spinning up the Docker containers, you can connect to it using the following credentials:

- `DB_HOST`: `postgres`
- `DB_PORT`: `5432`
- `DB_USERNAME`: `postgres`
- `DB_PASSWORD`: `postgres`

### Analytics

This project uses DuckDB for analytics. Using version `>=1.1` we can hook DuckDB up to the PostgreSQL instance with:

```
install postgres;
load postgres;
attach 'host=localhost port=5432 dbname=postgres user=postgres password=postgres' as backend (
    type postgres,
    schema 'ecom',
    read_only
);
use backend.ecom;

/* Check that it works */
show tables;
```

Copies of the tables can easily be made into a local DuckDB instance while the app is up so that the data can be queried without needing the Docker containers running.
