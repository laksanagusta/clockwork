package response

import (
	"clockwork-server/model"
	"time"
)

type Order struct {
	GrandTotal int         `json:"grandTotal"`
	OrderItem  []OrderItem `json:"orderItem"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}

func FormatOrder(order model.Order) Order {
	var dataOrder Order

	dataOrder.GrandTotal = order.GrandTotal
	dataOrder.CreatedAt = order.CreatedAt
	dataOrder.UpdatedAt = order.UpdatedAt

	for _, valueOrderItem := range order.OrderItem {
		newOrderItem := FormatOrderItem(valueOrderItem)
		dataOrder.OrderItem = append(dataOrder.OrderItem, newOrderItem)
	}

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
