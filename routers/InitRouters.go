package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(route *gin.Engine) {
	CrimeEventRoutes(route)
	ClientRoute(route)
	WebsocketRouter(route)
}
