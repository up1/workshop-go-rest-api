package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewUserAPI to create the router of user
func NewUserAPI(app *gin.RouterGroup) {
	app.GET("/user", handleGetUsers)
	app.GET("/user/:id", handleGetUserByID)
}

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleGetUsers(c *gin.Context) {
	// Get data from database
	c.JSON(http.StatusOK, Users{})
}

func handleGetUserByID(c *gin.Context) {
	id := c.Param("id")
	// Get data from database
	c.JSON(http.StatusOK, User{Id: id})
}

func getErrorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
