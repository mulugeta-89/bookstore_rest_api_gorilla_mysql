package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db *gorm.DB
)

func connect() {
	b, err := gorm.Open("mysql", "root:Mulu2835../library?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db =  b
}

func getDb() *gorm.DB {
	return db
}