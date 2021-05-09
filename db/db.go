package db

import (
	"github.com/zoshigayan/rss_reader/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(sqlite.Open("rss_reader.db"), &gorm.Config{})
	if err != nil {
		panic("DB接続に失敗したで")
	}

	db.AutoMigrate(&models.Feed{})
	db.AutoMigrate(&models.Entry{})
	log.Println("接続したで")
}

func DbManager() *gorm.DB {
	return db
}
