package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"todo/app/http/controllers"
)

func main() {
	router := gin.Default()

	router.POST("/article/store", articles.Store)
	router.GET("/article/edit/:id", articles.Edit)
	err := router.Run()
	if err != nil {
		return
	}
}
