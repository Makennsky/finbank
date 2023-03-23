postgres: 
	docker run --name mypostgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it mypostgres createdb --username=root --owner=root finbank

dropdb:
	docker exec -it mypostgres dropdb finbank

migrateup:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/finbank?sslmode=disable" -verbose up

migrateupstep:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/finbank?sslmode=disable" -verbose up 1

migratedown:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/finbank?sslmode=disable" -verbose down

migratedownstep:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/finbank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

exec:
	docker exec -it mypostgres psql -U root -d finbank

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Makennsky/finbank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup, migratedown sqlc test exec server mock migrateupstep migratedownstep