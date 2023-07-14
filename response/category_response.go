package response

import (
	"clockwork-server/model"
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatCategory(category model.Category) Category {
	var dataCategory Category

	dataCategory.ID = category.ID
	dataCategory.Title = category.Title
	dataCategory.CreatedAt = category.CreatedAt
	dataCategory.UpdatedAt = category.UpdatedAt

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
