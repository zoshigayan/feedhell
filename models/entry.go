package models

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	FeedId      int
	Feed        Feed
	Title       string
	Description string
	Link        string
}
