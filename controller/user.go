package controller

import (
	"strconv"
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

func UserUpdate(c *gin.Context) {
	var user models.User
	var data models.User
	data.Name = c.Query("name")

	user.Id, _ = strconv.Atoi(c.Params.ByName("id"))

	// データを更新する
	result := db.Model(&user).Updates(&data)
	if result.Error != nil {
		panic(result.Error)
	}

	// 元の形に戻す
	result.Scan(&data)
	c.JSON(204, nil)
}
