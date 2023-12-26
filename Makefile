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
	clear && go test -v -cover -short ./...

doc:
	go run ./cmd/docs/main.go init -g cmd/api/main.go -o docs

start:
	go run ./cmd/api/main.go

clean:
	clear

run: doc clean start
	

.PHONT: migrate composeup composedown createdb dropdb dbinit test sqlc doc run start clean