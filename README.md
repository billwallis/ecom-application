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
go build -o bin/ecom cmd/main.go
go run cmd/main.go

# ...alternatively, run via Docker
docker compose up --detach

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
- `DB_PORT`: `3306`
- `DB_USER`: `root`
- `DB_PASSWORD`: `password`

## REST API

Docs are made available by the following package:

- https://github.com/swaggo/swag

Check out the docs at:

- [docs/swagger.json](docs/swagger.json)

## Data and service models

### Data model

```mermaid
---
title: E-com Data Model
---
erDiagram
    User {
        int    id
        string first_name
        string last_name
        string email
        string password
    }
    
    Address {
        int    id
        int    user_id
        bool   is_default
        string line_1
        string line_2
        string city
        string country
        string postcode
    }
    
    Product {
        int     id
        string  name
        string  description
        string  image
        decimal price
        int     quantity
    }
    
    Order {
        int     id
        int     user_id
        decimal total
        string  status
        %% TODO: create foreign key relationship to Address
        string  address
    }
    
    OrderItem {
        int     id
        int     order_id
        int     product_id
        int     quantity
        decimal price
    }
    
    User ||--o{ Address : "has"
    User ||--o{ Order : "places"
    Order ||--o{ OrderItem : "contains"
    OrderItem }o--|| Product : "ordered"
    Order }o--o{ Address : "ships to"
```

### Service model

Each service generally has a handler and a store:

- The handler is responsible for handling the request and response. This is the interface between the API and the service.
- The store is responsible for the database operations. This is the interface between the service and the database.

```mermaid
---
title: E-com Service Model
---
graph LR
    User <--> UI <--API Calls--> Router <--> Service <--> Database
```
