package db

import (
	"app/model"
	"fmt"
)

// FeachAllArticle is 記事を全て取得する
func FeachAllArticle() []model.Article {
	db := InitDB()
	defer db.Close()

	articles := []model.Article{}
	db.Find(&articles)

	// fmt.Println(articles)

	return articles
}

// FeachOneArticle is 指定のIDの記事を取得する
func FeachOneArticle(num int) model.Article {
	db := InitDB()
	defer db.Close()

	article := model.Article{}
	article.ID = num

	db.First(&article)

	fmt.Println(article)

	return article
}
