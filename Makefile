postgres: 
	docker run --name mypostgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it mypostgres createdb --username=root --owner=root finbank

dropdb:
	docker exec -it mypostgres dropdb finbank

migrateup:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5433/finbank?sslmode=disable" -verbose up

migratedown:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5433/finbank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

exec:
	docker exec -it mypostgres psql -U root -d finbank

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup, migratedown sqlc test exec server