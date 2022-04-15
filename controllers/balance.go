package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"super_pay/database/mysql"
)

type Balance struct {
}

type balances struct {
	//gorm.Model
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	Amount int `json:"amount"`
}

func (receiver Balance) GetAll(c *gin.Context) gin.H {

	var data []balances
	mysql.InitDb() //for singleton connection you can direct set db variable
	db := mysql.Db
	//db.Find(&data, "user_id=? and id=?", 2, 1) with where example
	db.Find(&data)

	h := gin.H{
		"success": true,
		"data":    data,
	}
	return h
}
