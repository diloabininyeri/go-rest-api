package main

import (
	"github.com/gin-gonic/gin"
	"super_pay/helpers"
	"super_pay/routers"
)

func main() {

	engine := gin.Default()
	routers.Route(engine)
	_ = engine.Run(helpers.AppPort())
}
