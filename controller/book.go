package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"bookshelf/models"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db  *gorm.DB

func BookIndex(c *gin.Context) {
	// 初期化時にDBと接続
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "bookshelf"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}
	Books := [] models.Book{} // 複数件取得する場合、構造体を配列にする
	db.Find(&Books)
	fmt.Println(Books)

	c.JSON(200, Books)
}
