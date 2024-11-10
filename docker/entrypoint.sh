#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

wait_for_mysql() {
    DB_HOST=mysql

    while ! mysqladmin ping -h "$DB_HOST" --silent; do
        sleep 5
    done
}

main() {
    echo "Waiting for MySQL..."
    wait_for_mysql

    echo "Running e-commerce application"
    /usr/local/bin/app
}


main "$@"
