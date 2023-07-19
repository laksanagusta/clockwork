package response

import (
	"clockwork-server/model"
	"time"
)

type AttributeItem struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	AdditionalCharge int       `json:"additionalCharge"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func FormatAttributeItem(attribute model.AttributeItem) AttributeItem {
	attributeRes := AttributeItem{}

	attributeRes.ID = attribute.ID
	attributeRes.Title = attribute.Title
	attributeRes.AdditionalCharge = attribute.AdditionalCharge
	attributeRes.CreatedAt = attribute.CreatedAt
	attributeRes.UpdatedAt = attribute.UpdatedAt

	return attributeRes
}

func FormatAttributeItems(attributes []model.AttributeItem) []AttributeItem {
	attributesRes := []AttributeItem{}

	for _, value := range attributes {
		attribute := FormatAttributeItem(value)
		attributesRes = append(attributesRes, attribute)
	}

	return attributesRes
}
