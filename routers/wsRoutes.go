package routers

import (
	"github.com/capstone-be/controllers"
	"github.com/gin-gonic/gin"
)

func WebsocketRouter(router *gin.Engine) {
	groupRoute := router.Group("/ws")
	groupRoute.GET("/get-notification", controllers.HandleWebsocketConnection)
}
