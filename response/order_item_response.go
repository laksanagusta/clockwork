package response

import (
	"clockwork-server/model"
	"time"
)

type OrderItem struct {
	ID        uint      `json:"id"`
	Qty       int       `json:"grandTotal"`
	UnitPrice int       `json:"unitPrice"`
	SubTotal  int       `json:"subTotal"`
	ProductID uint      `json:"productId"`
	OrderID   uint      `json:"orderId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatOrderItem(orderItem model.OrderItem) OrderItem {
	var newOrderItem OrderItem

	newOrderItem.Qty = orderItem.Qty
	newOrderItem.UnitPrice = orderItem.UnitPrice
	newOrderItem.SubTotal = orderItem.SubTotal
	newOrderItem.ProductID = orderItem.ProductID
	newOrderItem.OrderID = orderItem.OrderID
	newOrderItem.CreatedAt = orderItem.CreatedAt
	newOrderItem.UpdatedAt = orderItem.UpdatedAt

	return newOrderItem
}

func FormatOrderItems(orderItem []model.OrderItem) []OrderItem {
	var dataOrderItems []OrderItem

	for _, value := range orderItem {
		singleDataOrderItem := FormatOrderItem(value)
		dataOrderItems = append(dataOrderItems, singleDataOrderItem)
	}

	return dataOrderItems
}
