package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Role int16

const (
	Administrator Role = iota
	RegularUser
)

type User struct {
	ID        string `json:"id"`
	FIRSTNAME string `json:"first_name"`
	LASTNAME  string `json:"last_name"`
	EMAIL     string `json:"email"`
	ROLE      string `json:"role"`
	PASSWORD  string `json:"password"`
}

// TODO: Add database
var users = []User{
	{ID: "1", FIRSTNAME: "David", LASTNAME: "Leclerq", EMAIL: "david.leclerq@gmail.com", ROLE: "0", PASSWORD: "Jean123!"},
	{ID: "2", FIRSTNAME: "Matthieu", LASTNAME: "Bocquet", EMAIL: "matthieu.bocquet@gmail.com", ROLE: "0", PASSWORD: "Jean123!"},
	{ID: "3", FIRSTNAME: "Sharon", LASTNAME: "Dupont", EMAIL: "sharon@live.be", ROLE: "1", PASSWORD: "Jean123!"},
	{ID: "4", FIRSTNAME: "Lisa", LASTNAME: "De Groof", EMAIL: "lisa@luminecapital.com", ROLE: "1", PASSWORD: "Jean123!"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserById)
	router.POST("/users", addUser)

	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")

	for _, u := range users {
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func addUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	lastId, err := strconv.Atoi(users[len(users)-1].ID)
	if err == nil {
		newUser.ID = strconv.Itoa(lastId + 1)
	} else {
		newUser.ID = "1"
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
