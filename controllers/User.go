package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_crud_api/models"
)

func (res Resources) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	usr, err := user.Create(res.db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usr})
}

func (res Resources) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	usr, err := user.Get(res.db, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usr})

}

func (res Resources) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	c.BindJSON(&user)

	usr, err := user.Edit(res.db, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usr})
}
func (res Resources) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	_, err := user.Delete(res.db, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
