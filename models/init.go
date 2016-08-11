package models

import (
	"Odyssey/utils"
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"

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

var mongoSession *mgo.Session

func dialMongo() (err error) {
	m := utils.GetConf().GetStringMapString("mongodb")
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{m["host"]},
		Timeout:  60 * time.Second,
		Database: m["dbname"],
		Username: m["username"],
		Password: m["password"],
	}
	mongoSession, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic("mongodb connections failed")
	}
	mongoSession.SetMode(mgo.Monotonic, true)
}
