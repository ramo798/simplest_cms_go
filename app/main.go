package main

import (
	"app/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello, API")
	})

	router.GET("/article/all", handler.GetArticleAll)

	router.GET("/article/id/:no", handler.GetArticleOnNumber)

	router.Run(":8080")

}
