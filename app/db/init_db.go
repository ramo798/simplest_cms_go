package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	USER := "root"
	PASS := "root"
	HOST := "mysql_123:3306"
	DBNAME := "mydatabase"
	db, err := gorm.Open("mysql", USER+":"+PASS+"@tcp("+HOST+")/"+DBNAME+"?charset=utf8&parseTime=True&loc=Asia%2FTokyo")
	if err != nil {
		panic(err.Error())
	}
	return db
}
