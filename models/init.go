package models

import (
	"Odyssey/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const driverName = "postgres"

// 接收一个参数用于启动db
func InitModels() error {
	c := utils.GetConf().GetStringMapString("database")
	utils.GetLog().Debug("database config = %v", c)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c["username"], c["password"], c["host"], c["port"], c["dbname"], c["sslmode"])

	var err error
	db, err = sql.Open(driverName, dsn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("connect to ", driverName, " success!")
	return nil
}
