package model

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Title      string `gorm:"size:60;not null;" json:"title"`
	Code       string `gorm:"size:40;not null;" json:"code"`
	StartTime  string `gotm:"size:12;not null;" json:"startTime"`
	EndTime    string `gotm:"size:12;not null;" json:"endTime"`
	DiscAmount int    `gorm:"size:8;not null;" json:"discAmount"`
	IsActive   *bool  `gorm:"not null" json:"isActive"`
}
