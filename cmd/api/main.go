package main

import (
	"database/sql"
	"log"

	"github.com/tetrex/backend-masterclass-go/api"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:pass@localhost:5432/db?sslmode=disable"
)

//	@title			API
//	@version		1.0
//	@description	This is a backend api for simple bank

//	@contact.name	Tetrex

//	@license.name	MIT License

// @host		localhost:8080
// @basePath	/
func main() {
	DB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB")
		log.Fatal(err)
	}

	store := db.NewStore(DB)
	api.NewServer(store)
}
