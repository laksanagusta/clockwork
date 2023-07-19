package model

import "gorm.io/gorm"

type Attribute struct {
	gorm.Model
	Title         string `gorm:"size:20;not null" json:"title"`
	IsMultiple    *bool  `gorm:"not null" json:"isMultiple"`
	IsRequired    *bool  `gorm:"not null" json:"isRequired"`
	AttributeItem []AttributeItem
}
