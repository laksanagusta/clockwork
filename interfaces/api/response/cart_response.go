package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Cart struct {
	ID         uint       `json:"id"`
	BaseAmount int        `json:"baseAmount"`
	TotalItem  int        `json:"TotalItem"`
	Status     string     `json:"status"`
	CartItems  []CartItem `json:"cartItems"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

func FormatCart(cart model.Cart) Cart {
	var dataCart Cart

	dataCart.ID = cart.ID
	dataCart.TotalItem = cart.TotalItem
	dataCart.Status = cart.Status
	dataCart.BaseAmount = cart.BaseAmount
	dataCart.CreatedAt = cart.CreatedAt
	dataCart.UpdatedAt = cart.UpdatedAt
	dataCart.CartItems = FormatCartItems(cart.CartItems)

	return dataCart
}

func FormatCarts(order []model.Cart) []Cart {
	var dataCarts []Cart

	for _, value := range order {
		singleDataCart := FormatCart(value)
		dataCarts = append(dataCarts, singleDataCart)
	}

	return dataCarts
}
