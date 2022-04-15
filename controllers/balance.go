package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"super_pay/database/mysql"
	CustomHttp "super_pay/libs/http"
	"super_pay/models"
)

type BalanceController struct {
}

func (receiver BalanceController) GetAll() CustomHttp.Response {
	var balanceModel models.BalanceModel
	var response CustomHttp.Response

	data := balanceModel.GetAll()
	h := gin.H{
		"balances": data,
	}
	response.Data = h
	response.Status = http.StatusOK
	return response

}

func (receiver BalanceController) SaveBalance(c *gin.Context) CustomHttp.Response {

	var balanceApi models.BalanceApi
	var response CustomHttp.Response
	err := c.BindJSON(&balanceApi)
	if err != nil {
		panic(err)
	}
	db := mysql.InitDb()
	db.Model(balanceApi).Create(&balanceApi)
	response.Status = http.StatusOK
	h := gin.H{
		"balance": balanceApi,
	}
	response.Data = h
	response.ReferenceId = balanceApi.ID
	return response

}
