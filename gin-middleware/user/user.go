package user

import (
	"demo/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewUserAPI to create the router of user
func NewUserAPI(app *gin.RouterGroup) {
	app.GET("/user", middleware.AuthRequired(), handleGetUsers())
	app.GET("/user/:id", middleware.AuthRequired(), handleGetUserByID())
}

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleGetUsers() func(c *gin.Context) {
	return func(c *gin.Context) {
		code := http.StatusOK
		c.JSON(code, Users{})
	}
}

func handleGetUserByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		code := http.StatusOK
		id := c.Param("id")
		c.JSON(code, User{Id: id})
	}
}

func getErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}