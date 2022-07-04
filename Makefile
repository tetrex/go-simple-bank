compose-up:
		docker-compose up
compose-down:
		docker-compose down

# useage: make migrate args="<args>"
migrate:
		go run ./cmd/migrate $(args)

.PHONT: migrate compose-up compose-down