package services

import (
	"github.com/mmcdole/gofeed"
)

func GetFeedInfo(url string) (string, string, error) {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(url)
	if err != nil {
		return "", "", err
	}

	return feed.Title, feed.Description, nil
}
