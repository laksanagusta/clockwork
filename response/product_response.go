package response

import (
	"clockwork-server/model"
	"time"
)

type Product struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	SerialNumber string    `json:"serialNumber"`
	UnitPrice    int       `json:"unitPrice"`
	Inventory    Inventory `json:"inventory"`
	User         User      `json:"user"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
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

	if product.Inventory.ID != 0 {
		dataProduct.Inventory = FormatInventory(product.Inventory)
	}

	if product.User.ID != 0 {
		dataProduct.User = FormatUser(product.User, "")
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
