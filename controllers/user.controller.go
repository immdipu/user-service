package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/immdipu/user-service/models"
)

func CreateUser(c *gin.Context) {

	var user models.User

	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide valid data",
		})
		return
	}

	err = user.Validate()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
