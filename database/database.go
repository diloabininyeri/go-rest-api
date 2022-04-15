package database

import "super_pay/database/mysql"

type Book struct {
}

func main() {

	var book Book
	db := mysql.InitDb()
	db.Create(&book)

}
