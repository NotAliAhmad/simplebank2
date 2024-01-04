package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbSource = "postgres"
	dbDriver = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	// Replace the connection string with your PostgreSQL database details
	testDB, err = sql.Open(dbSource, dbDriver)
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	os.Exit(m.Run())
}
