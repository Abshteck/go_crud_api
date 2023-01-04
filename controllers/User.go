package controllers

import (
	"net/http"

	"github.com/go_crud_api/models"
	"github.com/go_crud_api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = utils.ConnectDB()

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "Record not found!"})
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "Record not found!"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "Record deleted successfully!"})
}
