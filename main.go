package main

import (
	// $GOPATH/srcからのパスで指定する
	"bookshelf/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// controllerパッケージから読み込む
	router.GET("/", controller.IndexGET)
	router.GET("/books", controller.BookIndex)
	router.POST("/books", controller.BookCreate)
	router.PUT("/books/:id", controller.BookUpdate)
	router.DELETE("/books/:id", controller.BookDelete)
	router.POST("/users", controller.UserCreate)
	router.Run(":8080")
}
