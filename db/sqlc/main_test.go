package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/uti"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	uti.LoadConfig("../..")
	config, err := uti.GetConfig()
	if err != nil {
		log.Fatal("cannot connect to config: ", err)
		return
	}
	testDB, err = sql.Open(config.DB_DRIVER, config.URI_DB)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	m.Run()
}
