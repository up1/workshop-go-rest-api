package demo

import (
	"demo/user"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.New()
	// ===== Public APIs
	publicRoute := router.Group("/api/v1")
	user.NewUserAPI(publicRoute)

	// ===== Start server
	router.Run(":8080")
}
