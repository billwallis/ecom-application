FROM golang:1.24-bullseye


# use `datastore` for Docker (see docker-compose.yaml) and `localhost` for local development
ENV DB_HOST=datastore
ENV DB_PORT=5432
#ENV DB_NAME=ecom
ENV DB_USERNAME=postgres
ENV DB_PASSWORD=postgres

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download \
 && go mod verify \
 && apt-get update \
 && apt-get install -y postgresql-client

# TODO: Don't copy everything; build the binary and expose that
COPY . .
RUN go build -v -o /usr/local/bin/app /usr/src/app/main.go \
 && chmod +x /usr/src/app/docker/entrypoint.sh \
 && chmod +x /usr/src/app/docker/entrypoint-migrations.sh

# RUN sleep 10
# COPY cmd/migrate/migrations/*.sql cmd/migrate/migrations/
# RUN go run /usr/src/app/cmd/migrate/main.go up

#ENTRYPOINT ["/usr/src/app/docker/entrypoint.sh"]
CMD ["/usr/src/app/docker/entrypoint.sh"]
