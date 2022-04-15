package models

import (
	"super_pay/database/mysql"
	"time"
)

var tableName = "balances"

type BalanceModel struct {
}

type BalanceApi struct {
	//ID uint `json:"id" gorm:"primary_key"`
	ID int `json:"id" gorm:"primary_key"`
	//Id     string `json:"id"`
	UserId int `json:"user_id"`
	Amount int `json:"amount"`
}

func (m BalanceModel) GetAll() []BalanceStructure {
	return getAll()
}

type BalanceStructure struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (receiver BalanceApi) TableName() string {
	return tableName
}
func (b BalanceStructure) TableName() string {
	return tableName
}

func getAll() []BalanceStructure {
	var balances []BalanceStructure
	mysql.InitDb() //for singleton connection you can direct set db variable
	db := mysql.Db
	//db.Find(&balances, "user_id=? and id=?", 2, 1) with where example
	db.Find(&balances)
	print(balances)
	return balances
}
