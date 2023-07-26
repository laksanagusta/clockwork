package response

import "clockwork-server/domain/model"

type CartItemAttributeItem struct {
	AdditionalCharge int `json:"additionalCharge"`
	AttributeItem    AttributeItem
}

func FormatCartItemAttributeItem(cartItemAttributeItem model.CartItemAttributeItem) CartItemAttributeItem {
	data := CartItemAttributeItem{
		AdditionalCharge: cartItemAttributeItem.AdditionalCharge,
		AttributeItem:    FormatAttributeItem(cartItemAttributeItem.AttributeItem),
	}

	return data
}

func FormatCartItemAttributeItems(data []model.CartItemAttributeItem) []CartItemAttributeItem {
	datas := []CartItemAttributeItem{}

	for _, v := range data {
		singleData := FormatCartItemAttributeItem(v)
		datas = append(datas, singleData)
	}

	return datas
}
