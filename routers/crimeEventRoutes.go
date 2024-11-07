package routers

import (
	"github.com/capstone-be/controllers"
	"github.com/gin-gonic/gin"
)

func CrimeEventRoutes(router *gin.Engine) {
	groupRoute := router.Group("/crime-detection")
	groupRoute.POST("/send-event", controllers.SendCrimeEventV2)
	groupRoute.POST("/send-file", controllers.SendFile)
}
