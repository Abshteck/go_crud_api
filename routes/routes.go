package routes

import (
	"go_crud_api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/user", controllers.CreateUser)
	route.GET("/user", controllers.GetAllUsers)
	route.PUT("/user/:id", controllers.UpdateUser)
	route.DELETE("/user/:id", controllers.DeleteUser)

	route.Run()
}
