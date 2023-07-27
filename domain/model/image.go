package model

import (
	"time"
)

type Image struct {
	ID        uint   `gorm:"primarykey"`
	Url       string `gorm:"size:60;not null;" json:"url"`
	ProductID uint
	IsPrimary bool `gorm:"not null;" json:"isPrimary"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
