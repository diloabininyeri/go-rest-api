package models

import (
	"super_pay/database/mysql"
	"time"
)

var tableName = "balances"

type BalanceModel struct {
}

type BalanceStructure struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b BalanceStructure) TableName() string {
	return tableName
}

func GetAll() []BalanceStructure {
	var balances []BalanceStructure
	mysql.InitDb() //for singleton connection you can direct set db variable
	db := mysql.Db
	//db.Find(&balances, "user_id=? and id=?", 2, 1) with where example
	db.Find(&balances)
	print(balances)
	return balances
}
