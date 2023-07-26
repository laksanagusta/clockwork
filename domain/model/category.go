package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title string `gorm:"size:20;not null;" json:"title"`
}
