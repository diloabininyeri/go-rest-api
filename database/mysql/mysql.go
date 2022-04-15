package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"super_pay/helpers"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error

	db, err := gorm.Open(mysql.Open(getConnectionString()), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		panic(err)
		return nil
	}

	return db
}

func getConnectionString() string {

	user := helpers.ReadEnv("DATABASE_USER")
	host := helpers.ReadEnv("DATABASE_HOST")
	port := helpers.ReadEnv("DATABASE_PORT")
	password := helpers.ReadEnv("DATABASE_PASSWORD")
	dbName := helpers.ReadEnv("DATABASE_NAME")
	dsn := user + ":" + password + "@tcp" + "(" + host + ":" + port + ")/" + dbName + "?" + "parseTime=true&loc=Local"
	return dsn
}
