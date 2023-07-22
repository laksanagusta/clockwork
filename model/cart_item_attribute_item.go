package model

import "gorm.io/gorm"

type CartItemAttributeItem struct {
	gorm.Model
	CartItemID       uint `gorm:"not null;index" json:"orderItemId"`
	CartItem         CartItem
	AttributeItemID  uint `gorm:"not null;index" json:"attributeItemId"`
	AttributeItem    AttributeItem
	AdditionalCharge int `gorm:"not null;" json:"AddditionalCharge"`
}
