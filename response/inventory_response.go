package response

import (
	"clockwork-server/model"
	"time"
)

type Inventory struct {
	StockQty  int       `json:"stockQty"`
	IsInStock bool      `json:"isInStock"`
	ProductId int       `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatInventory(inventory model.Inventory) Inventory {
	var dataInventory Inventory

	dataInventory.StockQty = inventory.StockQty
	dataInventory.IsInStock = inventory.IsInStock
	dataInventory.ProductId = int(inventory.ProductID)
	dataInventory.CreatedAt = inventory.CreatedAt
	dataInventory.UpdatedAt = inventory.UpdatedAt

	return dataInventory
}

func FormatInventories(inventory []model.Inventory) []Inventory {
	var dataInventories []Inventory

	for _, value := range inventory {
		singleDataInventory := FormatInventory(value)

		dataInventories = append(dataInventories, singleDataInventory)
	}

	return dataInventories
}
