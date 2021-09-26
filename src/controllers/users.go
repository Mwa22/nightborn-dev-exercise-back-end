package controllers

import (
	"net/http"
	"strconv"

	"github.com/Mwa22/nightborn-dev-exercise-back-end/src/database"
	"github.com/Mwa22/nightborn-dev-exercise-back-end/src/models"
	"github.com/gin-gonic/gin"
)

// TODO: Add database
// var users = []models.User{
// 	{ID: 1, FIRSTNAME: "David", LASTNAME: "Leclerq", EMAIL: "david.leclerq@gmail.com", ROLE: models.Administrator, PASSWORD: "Jean123!"},
// 	{ID: 2, FIRSTNAME: "Matthieu", LASTNAME: "Bocquet", EMAIL: "matthieu.bocquet@gmail.com", ROLE: models.Administrator, PASSWORD: "Jean123!"},
// 	{ID: 3, FIRSTNAME: "Sharon", LASTNAME: "Dupont", EMAIL: "sharon@live.be", ROLE: models.RegularUser, PASSWORD: "Jean123!"},
// 	{ID: 4, FIRSTNAME: "Lisa", LASTNAME: "De Groof", EMAIL: "lisa@luminecapital.com", ROLE: models.RegularUser, PASSWORD: "Jean123!"},
// }

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid user id"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func AddUser(c *gin.Context) {
	var newUser models.User

	// Get data from request
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Create new user
	if err := database.DB.Create(&newUser).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}
