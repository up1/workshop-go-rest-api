package demo

import (
	"demo/db"
	"demo/middleware"
	"demo/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	// Create database connection
	resource := db.NewResource()

	router := gin.New()
	// ===== Public APIs
	publicRoute := router.Group("/api/v1")
	publicRoute.Use(middleware.InjectDBToRequest(resource.DB))
	user.NewUserAPI(publicRoute)

	// ===== Start server
	router.Run(":8080")
}
