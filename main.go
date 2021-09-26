package main

import (
	"github.com/Mwa22/nightborn-dev-exercise-back-end/src/controllers"
	"github.com/Mwa22/nightborn-dev-exercise-back-end/src/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDataBase()

	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserById)
	router.POST("/users", controllers.AddUser)

	router.Run("localhost:8080")
}
