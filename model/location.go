package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name    string `gorm:"size:40;not null;" json:"name"`
	Code    string `gorm:"size:20;not null;uniqueIndex" json:"code"`
	Address string `gorm:"size:20;not null;" json:"address"`
}
