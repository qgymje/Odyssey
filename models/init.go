package models

import (
	"Odyssey/utils"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

const driverName = "mysql"

// InitModels 连接数据库
func InitModels() (err error) {
	c := utils.GetConf().GetStringMapString("database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset-utf8&parseTime=True&loc=Local", c["username"], c["password"], c["host"], c["port"], c["dbname"])

	db, err = gorm.Open(driverName, dsn)

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)

	db.LogMode(true)

	if err = createTables(); err != nil {
		panic("create tables error")
	}

	return
}

func createTables() (err error) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&SMSCode{})
	return
}

// GetDB 获取*pg.DB对象
func GetDB() *gorm.DB {
	return db
}
