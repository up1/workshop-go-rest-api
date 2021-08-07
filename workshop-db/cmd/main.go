package main

import (
	"database/sql"
	"demo"
	"demo/login"
	"demo/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func main() {
	// Database connection
	db := demo.CreateDatabaseConnection()

	router := gin.New()
	router.Use(gin.Recovery())
	router.HandleMethodNotAllowed = true
	router.NoMethod(middlewares.NoMethodHandler()) // Bug !!
	router.NoRoute(middlewares.NoRouteHandler())

	// Metrics :: Prometheus format
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(router)

	route := router.Group("/api/v1")
	route.GET("/user", 
			middlewares.AuthRequired(), 
			middlewares.AuthenticateWithJwt(), 
			handleGetUsers(db))

	// Login API (public api)
	login.NewLoginAPI(route)

	router.Run(":8080")
}

// Closure
func handleGetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := demo.GetAllUsers(db)
		c.JSON(http.StatusOK, users)
	}
}
