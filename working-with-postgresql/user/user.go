package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewUserAPI to create the router of user
func NewUserAPI(app *gin.RouterGroup) {
	app.GET("/user", handleGetUsers)
}

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Using closure
func handleGetUsers(c *gin.Context) {
	// Get data from database
	users, _ := GetAllUsers(c.Keys["db"].(*sql.DB))
	c.JSON(http.StatusOK, users)
}
