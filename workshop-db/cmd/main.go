package main

import (
	"database/sql"
	"demo"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Resource struct {
	db *sql.DB
}

func main() {
	// Database connection
	db := createDatabaseConnection()

	router := gin.New()
	route := router.Group("/api/v1")
	route.GET("/user", handleGetUsers(db))
	router.Run(":8080")
}
// Closure
func handleGetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := demo.GetAllUsers(db)
		c.JSON(http.StatusOK, users)
	}
}

func createDatabaseConnection() *sql.DB {
	var db *sql.DB
	var err error

	db, err = sql.Open("postgres", "postgres://user:pass@localhost/demo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}
	return db
}
