package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	BaseAmount int    `gorm:"size:20;not null" json:"baseAmount"`
	TotalItem  int    `gorm:"size:20;not null" json:"totalItem"`
	Status     string `gorm:"size:10;not null" json:"status"`
	CustomerID int    `gorm:"size:10;not null" json:"customerId"`
	Customer   Customer
	CartItems  []CartItem
}
