package model

import "gorm.io/gorm"

type AttributeItem struct {
	gorm.Model
	Title            string `gorm:"size:20;not null" json:"title"`
	AdditionalCharge int    `gorm:"size:20;not null" json:"additionalCharge"`
	AttributeID      uint
	Attribute        Attribute
}
