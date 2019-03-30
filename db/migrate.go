package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"bookshelf/models"
)

func main(){
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "bookshelf"
	// OPTION := "charset=utf8&parseTime=True&loc=Local"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}
	db.CreateTable(&models.Book{})
}
