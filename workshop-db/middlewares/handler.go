package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("+++++")
		c.JSON(http.StatusNotFound, 
		gin.H{"message": "The processing function of the request route was not found"})
	}
}

func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("=======")
		c.JSON(http.StatusMethodNotAllowed, 
		gin.H{"message": "Method not found"})
	}
}