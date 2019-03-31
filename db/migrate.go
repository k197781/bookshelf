package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"bookshelf/models"
	"bookshelf/config"
)

func main(){
	db := database.GetConnection()
	db.CreateTable(&models.Book{})
}
