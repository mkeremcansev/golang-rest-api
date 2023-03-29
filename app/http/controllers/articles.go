package articles

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/app/http/requests"
	"todo/app/models"
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
		"data": models.Article{
			ID:      uint(id),
			Title:   article.Title,
			Content: article.Content,
		},
	})
}

func Edit(c *gin.Context) {
	id := c.Param("id")

	db := database.Connection()

	var article models.Article
	err := db.QueryRow("SELECT id, title, content FROM articles WHERE id = ?", id).Scan(&article.ID, &article.Title, &article.Content)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": article,
		})
	}
}
