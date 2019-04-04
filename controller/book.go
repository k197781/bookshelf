package controller

import (
	"fmt"
	"time"
	"strconv"
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

func BookUpdate(c *gin.Context) {
	var data models.Book
	var book models.Book

	if c.Query("Title") != "" {
		data.Title = c.Query("Title")
	}
	if c.Query("Author") != "" {
		data.Author = c.Query("Author")
	}
	if c.Query("Publisher") != "" {
		data.Publisher = c.Query("Publisher")
	}
	if c.Query("PublishedDate") != "" {
		layout := "2006-01-02"
		t,err := time.Parse(layout, c.Query("PublishedDate"))
		if err != nil{
			panic(err)
		}
		data.PublishedDate = t
	}
	now := time.Now()
	data.UpdatedAt = now

	book.Id, _ = strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	// Modelで変更するインスタンスを指定し，dataにある空(0や'')ではないメンバだけを更新する
	result := db.Model(&book).Updates(&data)
	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(204, nil)
}

func BookDelete(c *gin.Context) {
	var book models.Book
	book.Id, _ = strconv.ParseInt(c.Params.ByName("id"), 10, 64)

	result := db.Delete(&book)
	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(204, nil)
}
