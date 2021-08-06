package demo

import (
	"demo/middleware"
	"demo/user"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func StartServer() {
	router := gin.New()

	// ===== Middlewares
	router.Use(gin.Recovery())

	// ===== Custom middlewares
	router.Use(middleware.AuthRequired())
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

	// ===== Prefix of all routes
	publicRoute := router.Group("/api/v1")

	// ===== Initial resource from MongoDB
	// resource, err := db.CreateResource()
	// if err != nil {
	// 	logrus.Error(err)
	// }
	// defer resource.Close()

	// ===== Add routes of users
	user.NewUserAPI(publicRoute)

	// ===== Start server
	router.Run(":8080")
}