package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Order struct {
	BaseAmount             int       `json:"baseAmount"`
	AdditionalChargeAmount int       `json:"additionalChargeAmount"`
	DiscountAmount         int       `json:"discountAmount"`
	TaxAmount              int       `json:"taxAmount"`
	GrandTotal             int       `json:"grandTotal"`
	Status                 string    `json:"string"`
	TransactionNumber      string    `json:"transactionNumber"`
	SnapUrl                string    `json:"snapUrl"`
	PaymentID              uint      `json:"paymentId"`
	Payment                Payment   `json:"payment"`
	CartID                 uint      `json:"cartId"`
	Cart                   Cart      `json:"cart"`
	CustomerID             uint      `json:"customerId"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}

func FormatOrder(order model.Order) Order {
	var dataOrder Order

	dataOrder.GrandTotal = order.GrandTotal
	dataOrder.BaseAmount = order.BaseAmount
	dataOrder.DiscountAmount = order.DiscountAmount

	dataOrder.TaxAmount = order.TaxAmount
	dataOrder.Status = order.Status
	dataOrder.TransactionNumber = order.TransactionNumber
	dataOrder.SnapUrl = order.SnapUrl

	dataOrder.CustomerID = order.CustomerID

	dataOrder.PaymentID = order.PaymentID
	dataOrder.Payment = FormatPayment(order.Payment)

	dataOrder.CartID = order.CartID
	dataOrder.Cart = FormatCart(order.Cart)

	dataOrder.CreatedAt = order.CreatedAt
	dataOrder.UpdatedAt = order.UpdatedAt

	return dataOrder
}

func FormatOrders(order []model.Order) []Order {
	var dataOrders []Order

	for _, value := range order {
		singleDataOrder := FormatOrder(value)
		dataOrders = append(dataOrders, singleDataOrder)
	}

	return dataOrders
}
