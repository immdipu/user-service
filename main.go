package main

import (
	"log"

	"github.com/immdipu/user-service/db"
)

func main() {

	err := db.InitDB("postgres://postgres:dipu@localhost:5432/user-service")

	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	log.Println("Connected to the database")

	defer db.CloseDB()

	// var router *gin.Engine = gin.Default()

}
