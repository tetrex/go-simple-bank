package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/tetrex/backend-masterclass-go/api"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
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
// @basePath	/v1
func main() {
	conn, err := sql.Open(dbSource, dbSource)
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot connect to DB :: main")
	}
	store := db.NewStore(conn)
	api.NewServer(store)

}
