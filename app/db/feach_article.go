package db

import (
	"app/model"
)

func FeachAllArticle() []model.Article {
	db := InitDB()
	defer db.Close()

	articles := []model.Article{}
	db.Find(&articles)

	// fmt.Println(articles)

	return articles
}
