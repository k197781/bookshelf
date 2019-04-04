package main

import (
	"bookshelf/models"
	"bookshelf/config"
)

func main(){
	db := database.GetConnection()
	db.CreateTable(&models.Book{})
}
