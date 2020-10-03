package model

// Article is 記事の構造体
type Article struct {
	ID         int    `json:"post_id"`
	UsersID    string `json:"users_id"`
	Title      string `json:"title"`
	MainText   string `json:"maintext"`
	Tag        string `json:"tag"`
	ReleseDate string `json:"relesedate"`
}
