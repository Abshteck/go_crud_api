package routes

import (
	"github.com/go_crud_api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	res := controllers.NewConfig(db)
	router := gin.Default()

	router.POST("/user", res.CreateUser)
	router.GET("/user/:id", res.GetUser)
	router.PUT("/user/:id", res.UpdateUser)
	router.DELETE("/user/:id", res.DeleteUser)

	router.Run()
	return router
}
