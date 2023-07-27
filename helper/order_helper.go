package helper

import "clockwork-server/domain/model"

type OrderHelper interface {
	CalculateTotalAdditionalCharge([]model.CartItem) int
}

type orderHelper struct {
}

func NewOrderHelper() OrderHelper {
	return &orderHelper{}
}

func (h *orderHelper) CalculateTotalAdditionalCharge(data []model.CartItem) int {
	result := 0

	for _, v := range data {
		for _, x := range v.CartItemAttributeItem {
			result += x.AdditionalCharge
		}
	}

	return result
}
