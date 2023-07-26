package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Order struct {
	GrandTotal int       `json:"grandTotal"`
	Cart       Cart      `json:"cart"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func FormatOrder(order model.Order) Order {
	var dataOrder Order

	dataOrder.GrandTotal = order.GrandTotal
	dataOrder.CreatedAt = order.CreatedAt
	dataOrder.UpdatedAt = order.UpdatedAt
	dataOrder.Cart = FormatCart(order.Cart)

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
