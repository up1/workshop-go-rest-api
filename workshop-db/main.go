package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Resource struct {
	db *sql.DB
}

func main() {
	// Database connection
	r := Resource{ db : createDatabaseConnection() }

	router := gin.New()
	route := router.Group("/api/v1")
	route.GET("/user", r.handleGetUsers)
	router.Run(":8080")
}

func (r *Resource)handleGetUsers(c *gin.Context) {
	users, _ := getAllUsers(r.db)
	c.JSON(http.StatusOK, users)
}

func createDatabaseConnection() *sql.DB {
	var db *sql.DB
	var err error

	db, err = sql.Open("postgres", "postgres://user:pass@localhost/demo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}
	return db
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func getAllUsers(db *sql.DB) (Users, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
