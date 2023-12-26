package db

import (
	"database/sql"
	"log"
	"os"
	"runtime/debug"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/tetrex/backend-masterclass-go/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}

	testDB, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to DB")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
