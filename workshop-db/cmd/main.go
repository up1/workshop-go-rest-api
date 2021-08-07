package main

import (
	"database/sql"
	"demo"
	"demo/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	db := demo.CreateDatabaseConnection()

	router := gin.New()
	router.Use(gin.Recovery())
	router.NoMethod(middlewares.NoMethodHandler()) // Bug !!
	router.NoRoute(middlewares.NoRouteHandler())

	route := router.Group("/api/v1")
	route.GET("/user", middlewares.AuthRequired(), handleGetUsers(db))
	router.Run(":8080")
}
// Closure
func handleGetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := demo.GetAllUsers(db)
		c.JSON(http.StatusOK, users)
	}
}