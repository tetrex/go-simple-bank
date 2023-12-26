package main

import (
	"database/sql"
	"log"

	"github.com/tetrex/backend-masterclass-go/api"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/util"

	_ "github.com/jackc/pgx/v5/stdlib"
)

//	@title			API
//	@version		1.0
//	@description	This is a backend api for simple bank

//	@contact.name	Tetrex

//	@license.name	MIT License

// @host		localhost:8080
// @basePath	/
func main() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	DB, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to DB")
		log.Fatal(err)
	}

	store := db.NewStore(DB)
	s := api.NewServer(store)
	s.Start()
}
