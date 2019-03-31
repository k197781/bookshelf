package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "bookshelf"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	return DBMS, CONNECT
}
