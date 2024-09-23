package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouters(db *mongo.Database, route *gin.Engine) {
	CrimeEventRoutes(db, route)
}
