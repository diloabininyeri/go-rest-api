package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"super_pay/models"
)

type BalanceController struct {
}

func (receiver BalanceController) GetAll() gin.H {
	data := models.GetAll()
	h := gin.H{
		"success": true,
		"data":    data,
	}
	return h
}
