package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name    string `gorm:"size:40;not null;" json:"name"`
	Address string `gorm:"size:20;not null;" json:"address"`
}
