package demo

import (
	"database/sql"
	"demo/user"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func StartServer() {
	// Create database connection
	createDatabaseConnection()

	router := gin.New()
	// ===== Public APIs
	publicRoute := router.Group("/api/v1")
	user.NewUserAPI(publicRoute)

	// ===== Start server
	router.Run(":8080")
}

func createDatabaseConnection() {
	var err error
	// Move to config file or environment variables
	user.DB, err = sql.Open("postgres", "postgres://user:pass@localhost/demo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// Ping to database
	err = user.DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}
}
