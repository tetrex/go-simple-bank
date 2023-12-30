package main

import (
	"database/sql"
	"log"
	"runtime/debug"

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
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host		localhost:8080
// @basePath	/
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}

	DB, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to DB")
		log.Fatal(err)
	}

	store := db.NewStore(DB)
	s, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create new server")
		log.Fatal(err)
	}
	s.Start()

}
