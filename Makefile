.PHONY: postgres database migrate lint test start

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

database:
	createdb tracking

migrate:
	migrate -path db/migrations \
					-database postgres://postgres:secret@localhost:5432/tracking?sslmode=disable up

lint:
	golangci-lint run

test:
	go test ./tests/...

start:
	go run main.go
