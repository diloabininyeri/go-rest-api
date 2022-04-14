package main

import (
	"github.com/gin-gonic/gin"
	"super_pay/routers"
)

func main() {
	r := gin.Default()
	routers.Route(r)
	_ = r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
