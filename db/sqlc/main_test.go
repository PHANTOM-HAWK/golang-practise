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

func TestMain(m *testing.M) {
	con, err := sql.Open(driver, conn)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	testQueries = New(con)
	os.Exit(m.Run())
}
