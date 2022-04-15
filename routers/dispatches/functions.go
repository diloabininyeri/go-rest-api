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
func BalanceGetAll() func(context *gin.Context) {
	return func(context *gin.Context) {
		res := new(controllers.BalanceController).GetAll()
		context.JSON(res.Status, res.Json())
	}
}
func BalanceStore() func(context *gin.Context) {
	return func(context *gin.Context) {
		response := new(controllers.BalanceController).SaveBalance(context)
		context.JSON(response.Status, response.Json())
	}
}
func UserBalance() func(context *gin.Context) {
	return func(context *gin.Context) {
		response := new(controllers.UserController).FindUserBalance(context)
		context.JSON(response.Status, response.Json())
	}
}
func UserSumBalance() func(context *gin.Context) {
	return func(context *gin.Context) {

		response := new(controllers.UserController).SumBalance(context)
		context.JSON(response.Status, response.Json())
	}
}
