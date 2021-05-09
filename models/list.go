package models

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	ID    int
	Title string
}
