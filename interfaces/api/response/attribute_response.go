package response

import (
	"clockwork-server/domain/model"
)

type Attribute struct {
	ID             uint            `json:"id"`
	Title          string          `json:"title"`
	IsMultiple     *bool           `json:"isMultiple"`
	IsRequired     *bool           `json:"isRequred"`
	AttributeItems []AttributeItem `json:"attributeItems"`
	// CreatedAt      time.Time       `json:"createdAt"`
	// UpdatedAt      time.Time       `json:"updatedAt"`
}

func FormatAttribute(attribute model.Attribute) Attribute {
	attributeRes := Attribute{}
	attributeRes.ID = attribute.ID
	attributeRes.Title = attribute.Title
	attributeRes.IsMultiple = attribute.IsMultiple
	attributeRes.IsRequired = attribute.IsRequired
	attributeRes.AttributeItems = FormatAttributeItems(attribute.AttributeItem)
	// attributeRes.CreatedAt = attribute.CreatedAt
	// attributeRes.UpdatedAt = attribute.UpdatedAt

	return attributeRes
}

func FormatAttributes(attributes []model.Attribute) []Attribute {
	attributesRes := []Attribute{}

	for _, value := range attributes {
		attribute := FormatAttribute(value)
		attributesRes = append(attributesRes, attribute)
	}

	return attributesRes
}
