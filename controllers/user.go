package controllers

import (
	"github.com/gin-gonic/gin"
	"super_pay/database/mysql"
	functions "super_pay/libs"
	"super_pay/models"
)

type UserController struct {
}

func (receiver UserController) FindUserBalance(c *gin.Context) gin.H {

	userId := c.Query("user_id")
	id := functions.StrToInt(userId)
	var balances []models.BalanceStructure
	db := mysql.InitDb()
	db.Find(&balances, "user_id=?", id)

	h := gin.H{
		"success": true,
		"data":    balances,
	}
	return h
}
