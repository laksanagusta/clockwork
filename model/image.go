package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Url       string `gorm:"size:60;not null;" json:"url"`
	ProductID uint
}
