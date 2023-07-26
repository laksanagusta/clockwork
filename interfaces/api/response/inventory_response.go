package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Inventory struct {
	ID          uint `json:"id"`
	StockQty    int  `json:"stockQty"`
	IsInStock   bool `json:"isInStock"`
	SalableQty  int  `json:"salableQty"`
	ReservedQty int  `json:"reservedQty"`
	// ProductId int       `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatInventory(inventory model.Inventory) Inventory {
	var dataInventory Inventory

	dataInventory.ID = inventory.ID
	dataInventory.StockQty = inventory.StockQty
	dataInventory.IsInStock = inventory.IsInStock
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
