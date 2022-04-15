package postgress

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"super_pay/helpers"
)

func GetConnection() *sql.DB {

	driver := helpers.ReadEnv("DATABASE_DRIVER")
	db, err := sql.Open(driver, getConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}

func getConnectionString() string {

	host := helpers.ReadEnv("DATABASE_HOST")
	user := helpers.ReadEnv("DATABASE_USER")
	password := helpers.ReadEnv("DATABASE_PASSWORD")
	dbname := helpers.ReadEnv("DATABASE_NAME")
	port := helpers.ReadEnv("DATABASE_PORT")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname = %s sslmode=disable", host, port, user, password, dbname)

}
