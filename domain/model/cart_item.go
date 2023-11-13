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

type CartItems []CartItem

func (ci CartItems) PopulateProductIDs() []int8 {
	var ids []int8
	for _, v := range ci {
		ids = append(ids, int8(v.Product.ID))
	}

	return ids
}
