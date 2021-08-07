package login

import (
	"demo"
	"demo/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Creational pattern (Builder)
func NewLoginAPI(app *gin.RouterGroup) {
	app.POST("/login", handleLogin)
}

func handleLogin(c *gin.Context) {
	var payload LoginPayload
	var user demo.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		c.Abort()
		return
	}

	// Check email and password
	if payload.Email != "demo@demo.com" || payload.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}
	// Success
	user.Email = payload.Email

	// TODO :: Move to config or environment variables
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "xxxsecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	c.JSON(http.StatusOK, tokenResponse)
}
