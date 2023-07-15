package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Qty       int  `gorm:"size:20;not null;" json:"qty"`
	UnitPrice int  `gorm:"size:20;not null;" json:"unitPrice"`
	SubTotal  int  `gorm:"size:20;not null;" json:"subTotal"`
	ProductID uint `gorm:"size:20;not null;" json:"productId"`
	Product   Product
	OrderID   uint `gorm:"size:20;not null;" json:"orderId"`
	Order     Order
}
