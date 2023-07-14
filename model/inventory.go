package model

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	StockQty    int  `gorm:"size:20;not null;" json:"stockQty"`
	IsInStock   bool `gorm:"size:20;not null;" json:"isInStock"`
	ReservedQty int  `gorm:"size:20;not null;" json:"reservedQty"`
	SalableQty  int  `gorm:"size:20;not null;" json:"salableQty"`
	ProductID   uint `gorm:"size:20;not null;" json:"productId"`
}
