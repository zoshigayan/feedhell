package services

import (
	"github.com/mmcdole/gofeed"
	"github.com/zoshigayan/feedhell/db"
	"github.com/zoshigayan/feedhell/models"
	"log"
)

func Sync() error {
	db := db.DbManager()

	feeds := []models.Feed{}
	db.Find(&feeds)

	fp := gofeed.NewParser()
	for _, feed := range feeds {
		resp, err := fp.ParseURL(feed.Link)
		if err != nil {
			return err
		}

		entries := []models.Entry{}
		for _, item := range resp.Items {
			entries = append(
				entries,
				models.Entry{
					Feed:        feed,
					Title:       item.Title,
					Description: item.Description,
					Link:        item.Link,
				})
			log.Println(item.Title)
		}

		db.Create(entries)
	}

	return nil
}
