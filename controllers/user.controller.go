package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/immdipu/user-service/db"
	"github.com/immdipu/user-service/models"
)

func CreateUser(c *gin.Context) {

	var user models.User

	err := c.ShouldBindBodyWithJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Please provide valid data",
			"error":   err.Error(),
		})
		return
	}

	err = user.Validate()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	err = user.CheckIfUnique()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "User creation failed",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User created successfully",
		"user":    user,
	})
}
