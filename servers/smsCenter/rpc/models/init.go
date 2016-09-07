// Package models 表示数据操作层, 不涉及任何具体的业务逻辑, 只做数据操作, 保持功能唯一
package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var ErrNotFound = sql.ErrNoRows

const DBName = "sms"

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
