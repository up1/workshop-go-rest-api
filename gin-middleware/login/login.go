package login

import (
	"demo/auth"
	"demo/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLoginAPI(app *gin.RouterGroup) {
	app.POST("/login", handleLogin())
	app.POST("/signup", handleSignUp())
}

func handleLogin() func(c *gin.Context) {
	return func(c *gin.Context) {
		var payload LoginPayload
		var user user.User

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
}

func handleSignUp() func(c *gin.Context) {
	return func(c *gin.Context) {
		var user user.User

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "invalid request",
			})
			c.Abort()
			return
		}

		err = user.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "error hashing password",
			})
			c.Abort()
			return
		}
		// TODO :: Create user in database

		c.JSON(http.StatusOK, user)
	}
}
