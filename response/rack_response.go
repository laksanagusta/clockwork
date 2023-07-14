package response

import (
	"clockwork-server/model"
	"time"
)

type Rack struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Code      string    `json:"code"`
	MaxQty    int       `json:"maxQty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatRack(rack model.Rack) Rack {
	var dataRack Rack

	dataRack.ID = rack.ID
	dataRack.Title = rack.Title
	dataRack.Code = rack.Code
	dataRack.MaxQty = rack.MaxQty
	dataRack.CreatedAt = rack.CreatedAt
	dataRack.UpdatedAt = rack.UpdatedAt

	return dataRack
}

func FormatRacks(order []model.Rack) []Rack {
	var dataRacks []Rack

	for _, value := range order {
		singleDataRack := FormatRack(value)
		dataRacks = append(dataRacks, singleDataRack)
	}

	return dataRacks
}
