package http

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status      int    `json:"status"`
	Data        gin.H  `json:"sql"`
	Message     string `json:"message"`
	ReferenceId int    `json:"reference_id"`
}

func (r Response) Json() gin.H {

	h := gin.H{
		"sql":          r.Data,
		"message":      r.Message,
		"reference_id": r.ReferenceId,
	}

	delete(h, "reference_id")

	if r.ReferenceId != 0 {
		h["reference_id"] = r.ReferenceId
	}

	return h
}
