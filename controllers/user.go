package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"super_pay/database/mysql"
	functions "super_pay/libs"
	customHttp "super_pay/libs/http"
	"super_pay/models"
)

type UserController struct {
}

func (receiver UserController) FindUserBalance(c *gin.Context) customHttp.Response {

	var response customHttp.Response
	userId := c.Query("user_id")
	id := functions.StrToInt(userId)
	var balances []models.BalanceStructure
	db := mysql.InitDb()
	db.Find(&balances, "user_id=?", id)

	h := gin.H{
		"balances": balances,
	}
	response.Status = http.StatusOK
	response.Data = h
	return response
}

func (receiver UserController) SumBalance(c *gin.Context) customHttp.Response {

	var sum int
	var response customHttp.Response
	userId := c.Query("user_id")
	if len(userId) == 0 {
		response.Status = http.StatusNotFound
		response.Message = "user_id doesnt exits"
	}

	id := functions.StrToInt(userId)
	db := mysql.InitDb()
	row := db.Table("balances").Where("user_id = ?", id).Select("sum(amount)").Row()
	err := row.Scan(&sum)
	if err != nil {
		response.Data = nil
		response.Message = fmt.Sprintf("sql not found for user_id=%d", id)
		response.Status = http.StatusNotFound
		return response
	}
	dataSum := map[string]int{}
	dataSum["value"] = sum
	h := gin.H{
		"success": true,
		"sum":     dataSum,
	}
	response.Status = http.StatusOK
	response.Message = "success"
	response.Data = h
	return response
}
