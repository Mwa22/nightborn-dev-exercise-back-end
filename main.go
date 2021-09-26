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
			// swagger:operation GET /api/v1/users users GetAll
			// Get all users.
			// Return all users
			// ---
			//     Produces:
			//       - application/json
			//     Responses:
			//       200:
			//         description: All the users
			//         schema:
			//           type: array
			//           items:
			//             "$ref": "#/definitions/User"
			//       400:
			//         description: Bad Request
			users.GET("/", controllers.GetUsers)

			// swagger:operation POST /api/v1/users users AddUser
			// Add an user.
			// Create an user and return this user
			// ---
			//     Consumes:
			//       - application/json
			//     Produces:
			//       - application/json
			//     Parameters:
			//       - name: data
			//         in: body
			//         required: true
			//         schema:
			//           type: object
			//           "$ref": "#/definitions/User"
			//     Responses:
			//       201:
			//         description: The user created
			//         schema:
			//           type: object
			//           "$ref": "#/definitions/User"
			//       400:
			//         description: Bad Request
			users.POST("/", controllers.AddUser)

			// swagger:operation GET /api/v1/users/{id} users GetUserById
			// Get an user by id.
			// Return the user with id
			// ---
			//     Produces:
			//       - application/json
			//     Parameters:
			//       - name: id
			//         in: path
			//         type: string
			//         required: true
			//     Responses:
			//       200:
			//         description: The user
			//         schema:
			//           type: object
			//           "$ref": "#/definitions/User"
			//       404:
			//         description: User not found
			users.GET("/:id", controllers.GetUserById)

			// swagger:operation PATCH /api/v1/users/{id} users UpdateUser
			// Update an user.
			// Return the user updated
			// ---
			//     Consumes:
			//       - application/json
			//     Produces:
			//       - application/json
			//     Parameters:
			//       - name: id
			//         in: path
			//         type: string
			//         required: true
			//       - name: data
			//         in: body
			//         schema:
			//           type: object
			//           "$ref": "#/definitions/UpdateUserInput"
			//     Responses:
			//       200:
			//         description: The user updated
			//         schema:
			//           type: object
			//           "$ref": "#/definitions/User"
			//       400:
			//         description: Bad Request
			//       404:
			//         description: User not found
			users.PATCH("/:id", controllers.UpdateUser)
		}
	}

	router.Run("localhost:8080")
}
