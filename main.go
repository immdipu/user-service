package main

import (
	"log"

	"github.com/immdipu/user-service/db"
	"github.com/immdipu/user-service/models"
)

func main() {

	err := db.InitDB("postgres://postgres:dipu@localhost:5432/user-service")

	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	log.Println("Connected to the database")

	db.DB.Create(&models.User{
		FullName:   "Dipu",
		Email:      "dipu@gmail.com",
		Username:   "immdipu",
		Bio:        "I am a software engineer",
		ProfilePic: "https://www.google.com",
		Verified:   true,
	})

	// var router *gin.Engine = gin.Default()

}
