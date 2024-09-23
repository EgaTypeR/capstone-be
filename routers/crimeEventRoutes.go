package routers

import (
	"github.com/capstone-be/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CrimeEventRoutes(db *mongo.Database, router *gin.Engine) {
	groupRoute := router.Group("/crime-event")
	groupRoute.POST("/send-event", controllers.SendCrimeEvent)

}
