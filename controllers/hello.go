package controllers

import "github.com/gin-gonic/gin"

type HelloController struct{}

func (controller HelloController) SayHello(c *gin.Context) gin.H {
	h := gin.H{
		"version": "1.0.0",
		"about":   "golang rest api task for interview ",
		"author":  "Dılo sürücü",
		"email":   "berxudar@gmail.com",
	}
	return h
}
