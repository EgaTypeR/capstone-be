package routers

import (
	"github.com/capstone-be/controllers"
	"github.com/gin-gonic/gin"
)

func ClientRoute(router *gin.Engine) {
	groupRouter := router.Group("/client")
	groupRouter.GET("/crime-event", nil)
	groupRouter.GET("/notification", nil)
	groupRouter.PATCH("/read-notif", nil)

	groupRouter.POST("/send-notif", controllers.SendNotification)
}
