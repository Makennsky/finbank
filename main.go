package main

import (
	"database/sql"
	"log"

	"github.com/Makennsky/finbank/api"
	db "github.com/Makennsky/finbank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver   = "postgres"
	dbSource   = "postgresql://root:secret@localhost:5433/finbank?sslmode=disable"
	srvAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db!", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(srvAddress)
	if err != nil {
		log.Fatal("Cannot start server!", err)
	}
}
