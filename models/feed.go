package models

import (
	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	Link        string
}
