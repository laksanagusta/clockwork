package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Product struct {
	ID           uint        `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	SerialNumber string      `json:"serialNumber"`
	UnitPrice    int         `json:"unitPrice"`
	Category     Category    `json:"category"`
	Attributes   []Attribute `json:"attributes"`
	Images       []Image     `jsom:"images"`
	User         User        `json:"user"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

func FormatProduct(product model.Product) Product {
	var dataProduct Product

	dataProduct.ID = product.ID
	dataProduct.Title = product.Title
	dataProduct.Description = product.Description
	dataProduct.SerialNumber = product.SerialNumber
	dataProduct.UnitPrice = product.UnitPrice
	dataProduct.CreatedAt = product.CreatedAt
	dataProduct.UpdatedAt = product.UpdatedAt

	dataProduct.Attributes = FormatAttributes(product.Attributes)

	if product.User.ID != 0 {
		dataProduct.User = FormatUser(product.User, "")
	}

	if product.Category.ID != 0 {
		dataProduct.Category = FormatCategory(product.Category)
	}

	if len(product.Images) > 0 {
		dataProduct.Images = FormatImages(product.Images)
	}

	return dataProduct
}

func FormatProducts(product []model.Product) []Product {
	var dataProducts []Product

	for _, value := range product {
		singleDataProduct := FormatProduct(value)
		dataProducts = append(dataProducts, singleDataProduct)
	}

	return dataProducts
}
