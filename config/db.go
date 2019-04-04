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
	// structで指定したPublishedDateの型([]uint8)と，SQLのpublished_dateの型(*time.Time)が違がうことを許容する．
	OPTION := "parseTime=true"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}
