package controller

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"bookshelf/models"
	"bookshelf/config"
)

var db = database.GetConnection()

func BookIndex(c *gin.Context) {
	Books := [] models.Book{} // 複数件取得する場合、構造体を配列にする
	db.Find(&Books)
	fmt.Println(Books)

	c.JSON(200, Books)
}

func BookCreate(c *gin.Context) {
	var data models.Book
	data.Title = c.Query("Title")
	data.Author = c.Query("Author")
	data.Publisher = c.Query("Publisher")
	layout := "2006-01-02"
	now := time.Now()
	t,err := time.Parse(layout, c.Query("PublishedDate"))
	if err != nil{
		panic(err)
	}
	data.PublishedDate = t
	data.CreatedAt = now
	data.UpdatedAt = now

	// データを保存する
	result := db.Create(&data)
	if result.Error != nil {
		panic(result.Error)
	}

	// 元の形に戻す
	result.Scan(&data)
	c.JSON(201, data)
}
