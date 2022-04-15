package routers

import (
	"github.com/gin-gonic/gin"
	"super_pay/controllers"
)

func Hello() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(200, new(controllers.HelloController).SayHello(context))
	}
}
func User() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(200, new(controllers.HelloController).SayHello(context))
	}
}


func a() func(context *gin.Context) {
		return func(context *gin.Context) {
		context.JSON(200, new(controllers.HelloController).SayHello(context))
	}
}


func Balance() func(context *gin.Context) {
	return func(context *gin.Context) {
		context.JSON(200, new(controllers.Balance).GetAll(context))
	}
}
