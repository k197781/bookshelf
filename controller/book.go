package controller

import (
	"fmt"
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
