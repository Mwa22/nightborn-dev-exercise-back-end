package main

import (
	"github.com/Mwa22/nightborn-dev-exercise-back-end/src/controllers"
	"github.com/Mwa22/nightborn-dev-exercise-back-end/src/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDataBase()

	v1 := router.Group("api/v1")
	{
		users := v1.Group("users")
		{
			users.GET("/", controllers.GetUsers)
			users.POST("/", controllers.AddUser)
			users.GET("/:id", controllers.GetUserById)
			users.PATCH("/:id", controllers.UpdateUser)
		}

	}

	router.Run("localhost:8080")
}
