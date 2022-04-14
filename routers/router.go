package routers

import (
	"github.com/gin-gonic/gin"
	routers "super_pay/routers/dispatches"
)

func Route(route *gin.Engine) {
	route.GET("/api", routers.Hello())
	route.GET("/user", routers.User())
}

