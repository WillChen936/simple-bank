package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	testDB, err := sql.Open("postgres", "postgres://root:secret@localhost:5432/simple-bank?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
