package models

import (
	"Odyssey/utils"
	"log"
	"os"
	"strconv"

	pg "gopkg.in/pg.v4"
)

var db *pg.DB

//const driverName = "postgres"

// InitModels 连接数据库
func InitModels() (err error) {
	c := utils.GetConf().GetStringMapString("database")

	//dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c["username"], c["password"], c["host"], c["port"], c["dbname"], c["sslmode"])
	//dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", c["host"], c["username"], c["dbname"], c["sslmode"], c["password"])

	var sslmode bool
	if c["sslmode"] != "disable" {
		sslmode = true
	}
	poolsize, _ := strconv.Atoi(c["poolsize"])

	db = pg.Connect(&pg.Options{
		Addr:     c["host"],
		User:     c["username"],
		Password: c["password"],
		Database: c["dbname"],
		SSL:      sslmode,
		PoolSize: poolsize,
	})

	pg.SetQueryLogger(log.New(os.Stdout, "", log.LstdFlags))

	return
}

// GetDB 获取*pg.DB对象
func GetDB() *pg.DB {
	return db
}
