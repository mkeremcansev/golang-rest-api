package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"todo/app/http/controllers"
)

func main() {
	r := gin.Default()
	r.POST("/store", articles.Store)
	err := r.Run()
	if err != nil {
		return
	}
}
