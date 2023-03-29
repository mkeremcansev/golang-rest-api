package articles

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/app/http/requests"
	"todo/config/database"
)

func Store(c *gin.Context) {

	var article requests.ArticleStoreRequest
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Connection()
	insert, err := db.Exec("INSERT INTO articles(title, content) VALUES(?, ?)", article.Title, article.Content)
	if err != nil {
		panic(err.Error())
	}

	id, err := insert.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Article created successfully",
		"id":      id,
	})
}
