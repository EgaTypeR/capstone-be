package routers

import (
	"github.com/capstone-be/controllers"
	"github.com/gin-gonic/gin"
)

func CrimeEventRoutes(router *gin.Engine) {
	groupRoute := router.Group("/crime-detection")
	groupRoute.POST("/send-event", controllers.SendCrimeEvent)
}
