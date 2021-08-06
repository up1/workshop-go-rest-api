package user

import (
	"demo/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewUserAPI to create the router of user
func NewUserAPI(app *gin.RouterGroup, resource *db.Resource) {
	app.GET("/user", handleGetUsers(resource))
}

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Using closure
func handleGetUsers(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get data from database
		users, _ := GetAllUsers(resource.DB)
		c.JSON(http.StatusOK, users)
	}
}
