#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

wait_for_database() {
#    while ! pg_isready --host "$DB_HOST" --dbname "$DB_NAME" --quiet; do
    while ! pg_isready --host "$DB_HOST" --quiet; do
        sleep 5
    done
}

main() {
    echo "Waiting for database..."
    wait_for_database

    echo "Running e-commerce application"
    /usr/local/bin/app
}


main "$@"
