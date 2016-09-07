package utils

import (
	"database/sql"
	"fmt"
)

func InitMysql() *sql.DB {
	const driverName = "mysql"

	c := GetConf().GetStringMapString("database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset-utf8&parseTime=True&loc=Local", c["username"], c["password"], c["host"], c["port"], c["dbname"])

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		panic("connect db failed.")
	}
	//defer db.Close()

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	return db
}
