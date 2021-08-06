package demo

import (
	"demo/login"
	"demo/middleware"
	"demo/user"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func StartServer() {
	router := gin.New()

	// ===== Global Middlewares
	router.Use(gin.Recovery())

	// ===== Custom middlewares
	router.NoRoute(middleware.NoRouteHandler())
	router.NoMethod(middleware.NoMethodHandler())

	// ===== Prometheus middleware ========
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(router)

	// ===== Protected APIs
	protectedRoute := router.Group("/api/v1")
	protectedRoute.Use(middleware.AuthRequired())

	// ===== Add routes of users
	user.NewUserAPI(protectedRoute)

	// ===== Public APIs
	publicRoute := router.Group("/api/v1")
	login.NewLoginAPI(publicRoute)

	// ===== Start server
	router.Run(":8080")
}
