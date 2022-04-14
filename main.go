package main

import (
	"github.com/gin-gonic/gin"
	"super_pay/helpers"
	"super_pay/routers"
)

func main() {
	r := gin.Default()
	routers.Route(r)
	_ = r.Run(helpers.AppPort())
}
