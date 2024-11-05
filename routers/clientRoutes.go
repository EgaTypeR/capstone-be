package routers

import (
	"github.com/capstone-be/controllers"
	"github.com/gin-gonic/gin"
)

func ClientRoute(router *gin.Engine) {
	groupRouter := router.Group("/client")
	groupRouter.GET("/crime-event", nil)
	groupRouter.GET("/notification", controllers.GetNotification)

	// groupRouter.POST("/send-notif", controllers.SendNotification)
	groupRouter.GET("/alert-history", controllers.GetHistoryAlerts)
	groupRouter.PATCH("/update-alert/:id", controllers.UpdateAlert)

	// Alert Page
	groupRouter.GET("/count-unread-notif", controllers.CountUnreadNotification)
	groupRouter.PATCH("/read-notif", controllers.ReadNotification)
}
