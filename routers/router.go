package routers

import (
	"github.com/gin-gonic/gin"
	"super_pay/controllers"
)

func Route(route *gin.Engine) {

	route.GET("/api", func(context *gin.Context) {
		context.JSON(200, new(controllers.HelloController).SayHello(context))
	})
}
