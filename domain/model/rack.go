package model

import "gorm.io/gorm"

type Rack struct {
	gorm.Model
	Title  string `gorm:"size:40;not null;" json:"title"`
	Code   string `gorm:"size:20;not null;uniqueIndex" json:"code"`
	MaxQty int    `gorm:"size:32;not null;" json:"maxQty"`
}
