package controllers

import (
	"net/http"

	"github.com/Mwa22/nightborn-dev-exercise-back-end/src/database"
	"github.com/Mwa22/nightborn-dev-exercise-back-end/src/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func AddUser(c *gin.Context) {
	// Get data from request
	var newUser models.User
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

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Get existing user
	var user models.User
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	// Get data from request
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Update user
	if err := database.DB.Model(&user).Updates(input).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
