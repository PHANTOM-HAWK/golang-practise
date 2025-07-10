package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	driver = "postgres"
	conn   = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driver, conn)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
