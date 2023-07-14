package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	GrandTotal        int    `gorm:"size:20;not null;" json:"grandTotal"`
	Status            string `gorm:"size:20;not null;" json:"string"`
	TransactionNumber string `gorm:"size:20;not null;" json:"transactionNumber"`
	SnapUrl           string `gorm:"size:100;" json:"snapUrl"`
	OrderItem         []OrderItem
}

type OrderStatus struct {
	Created           string
	WaitingForPayment string
	PaymentCompleted  string
	PaymentFailed     string
	Shipped           string
	Delivered         string
	Completed         string
}

func GetOrderStatus() *OrderStatus {
	return &OrderStatus{
		Created:           "created",
		WaitingForPayment: "waiting_for_payment",
		PaymentCompleted:  "payment_completed",
		PaymentFailed:     "payment_failed",
		Shipped:           "shipped",
		Delivered:         "delivered",
		Completed:         "completed",
	}
}
