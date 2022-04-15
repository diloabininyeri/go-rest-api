package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"super_pay/database/mysql"
	"super_pay/models"
)

type BalanceController struct {
}

func (receiver BalanceController) GetAll() gin.H {
	var balanceModel models.BalanceModel
	data := balanceModel.GetAll()
	h := gin.H{
		"success": true,
		"data":    data,
	}
	return h
}

func (receiver BalanceController) SaveBalance(c *gin.Context) gin.H {

	var balanceApi models.BalanceApi
	err := c.BindJSON(&balanceApi)
	if err != nil {
		panic(err)
	}
	db := mysql.InitDb()
	db.Model(balanceApi).Create(&balanceApi)
	return gin.H{
		"data": balanceApi,
	}

}
