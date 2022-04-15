package controllers

import (
	"fmt"
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

func (receiver UserController) SumBalance(c *gin.Context) gin.H {

	var sum int
	userId := c.Query("user_id")
	if len(userId) == 0 {
		return gin.H{
			"success": false,
			"message": "user_id doesnt exits",
		}
	}

	id := functions.StrToInt(userId)
	db := mysql.InitDb()
	row := db.Table("balances").Where("user_id = ?", id).Select("sum(amount)").Row()
	err := row.Scan(&sum)
	if err != nil {
		return gin.H{
			"data":    nil,
			"success": false,
			"message": fmt.Sprintf("data not found for user_id=%d", id),
		}
	}
	dataSum := map[string]int{}
	dataSum["sum"] = sum
	h := gin.H{
		"success": true,
		"data":    dataSum,
	}
	return h
}
