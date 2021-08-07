package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	createDatabaseConnection()
}

func createDatabaseConnection() *sql.DB {
	var db *sql.DB
	var err error

	db, err = sql.Open("postgres", "postgres://user:pass@localhost/demo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}
	return db
}