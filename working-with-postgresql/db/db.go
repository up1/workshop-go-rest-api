package db

import (
	"database/sql"
	"log"
	"time"
)

type Resource struct {
	DB *sql.DB
}

func NewResource() *Resource {
	db := createDatabaseConnection()
	return &Resource{DB: db}
}

func createDatabaseConnection() *sql.DB {
	// Move to config file or environment variables
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/demo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	// Ping to database
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}
	return db
}
