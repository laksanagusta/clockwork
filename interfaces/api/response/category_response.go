package response

import (
	"clockwork-server/domain/model"
)

type Category struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func FormatCategory(category model.Category) Category {
	var dataCategory Category

	dataCategory.ID = category.ID
	dataCategory.Title = category.Title

	return dataCategory
}

func FormatCategories(order []model.Category) []Category {
	var dataCategories []Category

	for _, value := range order {
		singleDataCategory := FormatCategory(value)
		dataCategories = append(dataCategories, singleDataCategory)
	}

	return dataCategories
}
