package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:pass@localhost:5432/db?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
