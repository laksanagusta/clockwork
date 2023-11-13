package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	BaseAmount int    `gorm:"size:20;not null" json:"baseAmount"`
	TotalItem  int    `gorm:"size:20;not null" json:"totalItem"`
	Status     string `gorm:"size:10;not null" json:"status"`
	UserID     int    `gorm:"size:10;not null" json:"userId"`
	User       User
	VoucherID  uint `gorm:"size:8;" json:"voucherId"`
	Voucher    Voucher
	CartItems  []CartItem
}
