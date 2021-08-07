package main

import (
	"database/sql"
	"log"
)

func main() {
	createDatabaseConnection()
}

func createDatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/demo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// Ping to database
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}
	return db
}