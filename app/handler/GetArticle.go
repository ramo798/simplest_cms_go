package handler

import (
	"app/db"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticleAll is 全記事取得
func GetArticleAll(c *gin.Context) {
	posts := db.FeachAllArticle()
	log.Println(posts)
	c.JSON(200, posts)
}

// GetArticleOnNumber is IDを指定して記事を取得
func GetArticleOnNumber(c *gin.Context) {
	articlenum := c.Param("no")
	i, _ := strconv.Atoi(articlenum)
	post := db.FeachOneArticle(i)
	c.JSON(200, post)
}
