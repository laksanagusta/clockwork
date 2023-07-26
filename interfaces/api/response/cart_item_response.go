package response

import (
	"clockwork-server/domain/model"
	"time"
)

type CartItem struct {
	ID        uint      `json:"id"`
	Qty       int       `json:"grandTotal"`
	UnitPrice int       `json:"unitPrice"`
	SubTotal  int       `json:"subTotal"`
	ProductID uint      `json:"productId"`
	CartID    uint      `json:"cartId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatCartItem(cartItem model.CartItem) CartItem {
	var newCartItem CartItem

	newCartItem.Qty = cartItem.Qty
	newCartItem.UnitPrice = cartItem.UnitPrice
	newCartItem.SubTotal = cartItem.SubTotal
	newCartItem.ProductID = cartItem.ProductID
	newCartItem.CartID = cartItem.CartID
	newCartItem.CreatedAt = cartItem.CreatedAt
	newCartItem.UpdatedAt = cartItem.UpdatedAt

	return newCartItem
}

func FormatCartItems(cartItem []model.CartItem) []CartItem {
	var dataCartItems []CartItem

	for _, value := range cartItem {
		singleDataCartItem := FormatCartItem(value)
		dataCartItems = append(dataCartItems, singleDataCartItem)
	}

	return dataCartItems
}
