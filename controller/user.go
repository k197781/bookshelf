package controller

import (
	"github.com/gin-gonic/gin"
	"bookshelf/models"
)

func UserCreate(c *gin.Context) {
	var data models.User
	data.Name = c.Query("name")

	// データを保存する
	result := db.Create(&data)
	if result.Error != nil {
		panic(result.Error)
	}

	// 元の形に戻す
	result.Scan(&data)
	c.JSON(201, data)
}
