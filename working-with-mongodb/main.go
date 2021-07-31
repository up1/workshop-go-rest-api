package main

import (
	"demo/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	dbClient *db.MongoClient
}

func main() {
	r := gin.New()
	dbClient := db.NewMongoClient()
    server := Server{dbClient: dbClient}
	r.GET("/account/:id", server.GetAccountById)
	r.Run(":8080")
}

func (s *Server) GetAccountById(c *gin.Context) {
	accountID := c.Param("id")
	account, err := s.dbClient.InquiryAccount(c, accountID)
	if err != nil {
		logrus.Errorf("Error reading accountID '%v' from DB: %v", accountID, err.Error())
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		    return
		}
	}
	c.JSON(http.StatusOK, account)
}