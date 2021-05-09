package models

import (
	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	ID          int
	ListId      int
	List        List
	Title       string
	Description string
	Link        string
}
