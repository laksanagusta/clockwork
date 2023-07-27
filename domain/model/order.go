package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	BaseAmount             int    `gorm:"size:20;not null;" json:"baseAmount"`
	AdditionalChargeAmount int    `gorm:"size:20;not null;" json:"additionalChargeAmount"`
	DiscountAmount         int    `gorm:"size:20;" json:"discountAmount"`
	TaxAmount              int    `gorm:"size:20;" json:"taxAmount"`
	GrandTotal             int    `gorm:"size:20;" json:"grandTotal"`
	Status                 string `gorm:"size:20;not null;" json:"string"`
	TransactionNumber      string `gorm:"size:20;not null;" json:"transactionNumber"`
	SnapUrl                string `gorm:"size:100;" json:"snapUrl"`
	PaymentID              uint   `gorm:"size:100;" json:"paymentId"`
	Payment                Payment
	CartID                 uint `gorm:"size:20;" json:"cartId"`
	Cart                   Cart
	CustomerID             uint `gorm:"size:20;" json:"customerId"`
	Customer               Customer
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
