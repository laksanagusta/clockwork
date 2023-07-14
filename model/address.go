package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Title      string `gorm:"size:40;not null;" json:"title"`
	Street     string `gorm:"size:225;not null;" json:"street"`
	Note       string `gorm:"size:225;not null;" json:"note"`
	City       string `gorm:"size:20;not null;" json:"city"`
	CustomerID uint   `gorm:"size:20;not null;" json:"customerId"`
}
