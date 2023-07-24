package model

import "time"

type CartItemAttributeItem struct {
	ID               uint `gorm:"primarykey"`
	CartItemID       uint `gorm:"not null;index" json:"orderItemId"`
	CartItem         CartItem
	AttributeItemID  uint `gorm:"not null;index" json:"attributeItemId"`
	AttributeItem    AttributeItem
	AdditionalCharge int `gorm:"not null;" json:"AddditionalCharge"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
