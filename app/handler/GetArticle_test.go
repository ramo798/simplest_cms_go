package handler

import (
	"app/db"
	"log"
	"testing"
)

func TestGetArticleAll(t *testing.T) {
	posts := db.FeachAllArticle()
	log.Println(posts)

}
