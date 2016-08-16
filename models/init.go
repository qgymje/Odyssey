package models

import (
	"database/sql"

	"github.com/diegogub/aranGO"
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

var session *aranGO.Session

func InitArango(s *aranGO.Session) error {
	session = s

	return initCollection(s)
}

func GetSession() *aranGO.Session {
	return session
}

const (
	DB_NAME      = "odyssey"
	DOC_SMSCodes = "smscodes"
)

func initCollection(s *aranGO.Session) (err error) {
	if err = createCollection(DB_NAME, DOC_SMSCodes); err != nil {
		return err
	}
	return
}

func createCollection(db, col string) error {
	if !GetSession().DB(db).ColExist(col) {
		collection := aranGO.NewCollectionOptions(col, false)
		err := GetSession().DB(db).CreateCollection(collection)
		if err != nil {
			return err
		}
	}
	return nil
}
