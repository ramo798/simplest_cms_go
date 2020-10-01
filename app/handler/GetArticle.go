package handler

import (
	"app/db"
	"log"

	"github.com/gin-gonic/gin"
)

func GetArticleAll(c *gin.Context) {
	posts := db.FeachAllArticle()
	log.Println(posts)
	c.JSON(200, posts)
}
