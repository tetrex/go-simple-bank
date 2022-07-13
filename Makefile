# to start docker compose
composeup:
		docker-compose up

# to stop socker compose
composedown:
		docker-compose down

# to migrate db , and helper commands
# usage: make migrate args="<args>"
migrate:
		go run ./cmd/migrate $(args)

# to migrate db to latest bersion
createdb:
		go run ./cmd/migrate up

# to drop db completely
dropdb:
		go run ./cmd/migrate reset

# to re-setup a clean db
dbinit:
		make dropdb && make createdb

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

.PHONT: migrate composeup composedown createdb dropdb dbinit test sqlc