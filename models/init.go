package models

import (
	"database/sql"

	mgo "gopkg.in/mgo.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// InitModels 连接数据库
func InitModels(rawDB *sql.DB, driverName string) (err error) {
	db = sqlx.NewDb(rawDB, driverName)
	err = db.Ping()
	return
}

// GetDB 获取*sqlx.DB对象
func GetDB() *sqlx.DB {
	return db
}

var mongoSession *mgo.Session

func InitMongodb(sess *mgo.Session) {
	mongoSession = sess
}

func GetMongo() *mgo.Session {
	return mongoSession
}
