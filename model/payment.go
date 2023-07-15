package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentMethod   string `gorm:"size:20;not null;" json:"paymentMethod"`
	PaymentResponse string `gorm:"size:20;not null;" json:"string"`
	PaymentStatus   string `gorm:"size:20;not null;" json:"transactionNumber"`
}
