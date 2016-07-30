package models

import (
	"Odyssey/utils"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const driverName = "mysql"

// InitModels 连接数据库
func InitModels() (err error) {
	c := utils.GetConf().GetStringMapString("database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset-utf8&parseTime=True&loc=Local", c["username"], c["password"], c["host"], c["port"], c["dbname"])

	db, err = sqlx.Connect(driverName, dsn)
	if err != nil {
		panic("db connections failed.")
	}

	db.DB.SetMaxIdleConns(5)
	db.DB.SetMaxOpenConns(10)

	return
}

// GetDB 获取*sqlx.DB对象
func GetDB() *sqlx.DB {
	return db
}
