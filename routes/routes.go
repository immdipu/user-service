package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/immdipu/user-service/controllers"
)

func UserRoutes(route *gin.Engine) {

	route.POST("/user", controllers.CreateUser)

}
