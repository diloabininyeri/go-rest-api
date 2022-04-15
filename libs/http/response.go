package http

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int    `json:"status"`
	Data    gin.H  `json:"data"`
	Message string `json:"message"`
}

func (r Response) Json() gin.H {
	return gin.H{
		"data":    r.Data,
		"message": r.Message,
	}
}
