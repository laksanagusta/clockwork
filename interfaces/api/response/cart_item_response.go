package response

import (
	"clockwork-server/domain/model"
	"time"
)

type CartItem struct {
	ID                    uint                    `json:"id"`
	Qty                   int                     `json:"qty"`
	UnitPrice             int                     `json:"unitPrice"`
	SubTotal              int                     `json:"subTotal"`
	ProductID             uint                    `json:"productId"`
	Note                  string                  `json:"note"`
	CartID                uint                    `json:"cartId"`
	Product               Product                 `json:"product"`
	CartItemAttributeItem []CartItemAttributeItem `json:"cartItemAttributeItem"`
	CreatedAt             time.Time               `json:"createdAt"`
	UpdatedAt             time.Time               `json:"updatedAt"`
}

func FormatCartItem(cartItem model.CartItem) CartItem {
	var newCartItem CartItem

	newCartItem.ID = cartItem.ID
	newCartItem.Qty = cartItem.Qty
	newCartItem.UnitPrice = cartItem.UnitPrice
	newCartItem.SubTotal = cartItem.SubTotal
	newCartItem.ProductID = cartItem.ProductID
	newCartItem.CartID = cartItem.CartID
	newCartItem.Note = cartItem.Note
	newCartItem.CartItemAttributeItem = FormatCartItemAttributeItems(cartItem.CartItemAttributeItem)
	newCartItem.CreatedAt = cartItem.CreatedAt
	newCartItem.UpdatedAt = cartItem.UpdatedAt

	return newCartItem
}

func FormatCartItems(cartItem []model.CartItem) []CartItem {
	var dataCartItems []CartItem

	for _, value := range cartItem {
		singleDataCartItem := FormatCartItem(value)
		singleDataCartItem.Product = FormatProduct(value.Product)
		dataCartItems = append(dataCartItems, singleDataCartItem)
	}

	return dataCartItems
}
