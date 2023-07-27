package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Payment struct {
	ID              uint      `json:"id"`
	PaymentMethod   string    `json:"paymentMethod"`
	PaymentResponse string    `json:"paymentResponse"`
	PaymentStatus   string    `json:"paymentStatus"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func FormatPayment(payment model.Payment) Payment {
	var dataPayment Payment

	dataPayment.ID = payment.ID
	dataPayment.PaymentMethod = payment.PaymentMethod
	dataPayment.PaymentResponse = payment.PaymentResponse
	dataPayment.PaymentStatus = payment.PaymentStatus
	dataPayment.CreatedAt = payment.CreatedAt
	dataPayment.UpdatedAt = payment.UpdatedAt

	return dataPayment
}

func FormatPayments(order []model.Payment) []Payment {
	var dataPayments []Payment

	for _, value := range order {
		singleDataPayment := FormatPayment(value)
		dataPayments = append(dataPayments, singleDataPayment)
	}

	return dataPayments
}
