package routers

import (
	"github.com/gin-gonic/gin"
	routers "super_pay/routers/dispatches"
)

func Route(route *gin.Engine) {
	route.GET("/api", routers.Hello())
	route.GET("/", routers.Hello())
	route.GET("/user", routers.User())
	route.GET("/balance", routers.BalanceGetAll())
	route.POST("/balance", routers.BalanceStore())
	route.GET("/user/balance", routers.UserBalance())
	route.GET("/user/balance/sum", routers.UserSumBalance())
}
