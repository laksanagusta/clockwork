package model

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Qty                   int    `gorm:"size:20;not null;" json:"qty"`
	UnitPrice             int    `gorm:"size:20;not null;" json:"unitPrice"`
	SubTotal              int    `gorm:"size:20;not null;" json:"subTotal"`
	Note                  string `gorm:"size:100;not null;" json:"note"`
	AttributeItemSorted   string `gorm:"size:20;not null" json:"attributeItemSorted"`
	ProductID             uint   `gorm:"size:20;not null;index" json:"productId"`
	Product               Product
	CartID                uint `gorm:"size:20;not null;index" json:"cartId"`
	Cart                  Cart
	CartItemAttributeItem []CartItemAttributeItem
}
