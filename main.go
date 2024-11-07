package main

import "github.com/gin-gonic/gin"

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H({
			"message": "Hello World",
		}))
	})

}
