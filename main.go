package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/immdipu/user-service/db"
	"github.com/immdipu/user-service/routes"
)

func main() {

	err := db.InitDB()

	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	log.Println("Connected to the database")

	// db.DB.Create(&models.User{
	// 	FullName:   "Dipu",
	// 	Email:      "dipu@gmail.com",
	// 	Username:   "immdipu",
	// 	Bio:        "I am a software engineer",
	// 	ProfilePic: "https://www.google.com",
	// 	Verified:   true,
	// })

	var router *gin.Engine = gin.Default()

	routes.UserRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
	log.Println("Server started at port 8080")

}
