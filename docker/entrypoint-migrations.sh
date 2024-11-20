#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

wait_for_database() {
    while ! pg_isready --host "$DB_HOST" --quiet; do
        sleep 5
    done
}

main() {
    echo "Waiting for database..."
    wait_for_database

    echo "Running migrations"
    go run cmd/migrate/main.go up
    echo "Migrations finished"
}


main "$@"
