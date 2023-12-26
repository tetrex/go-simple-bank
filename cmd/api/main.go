package main

import (
	"database/sql"
	"log"

	"github.com/tetrex/backend-masterclass-go/api"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:pass@localhost:5432/db?sslmode=disable"
)

//	@title			MYAPP API
//	@version		1.0
//	@description	This is a sample RESTful API with a CRUD

//	@contact.name	Dumindu Madunuwan
//	@contact.url	https://learning-cloud-native-go.github.io

//	@license.name	MIT License
//	@license.url	https://github.com/learning-cloud-native-go/myapp/blob/master/LICENSE

// @host		localhost:8080
// @basePath	/v1
func main() {
	conn, err := sql.Open(dbSource, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB")
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start()
}
