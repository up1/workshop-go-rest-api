package demo

import (
	"demo/db"
	"demo/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func StartServer() {
	// Create database connection
	resource := db.NewResource()

	router := gin.New()
	// ===== Public APIs
	publicRoute := router.Group("/api/v1")
	user.NewUserAPI(publicRoute, resource)

	// ===== Start server
	router.Run(":8080")
}

