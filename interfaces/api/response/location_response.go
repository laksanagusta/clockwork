package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Location struct {
	ID        uint      `json:"id"`
	Name      string    `json:"title"`
	Code      string    `json:"code"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatLocation(organization model.Location) Location {
	var dataLocation Location

	dataLocation.ID = organization.ID
	dataLocation.Name = organization.Name
	dataLocation.Address = organization.Address
	dataLocation.CreatedAt = organization.CreatedAt
	dataLocation.UpdatedAt = organization.UpdatedAt

	return dataLocation
}

func FormatLocations(order []model.Location) []Location {
	var dataLocations []Location

	for _, value := range order {
		singleDataLocation := FormatLocation(value)
		dataLocations = append(dataLocations, singleDataLocation)
	}

	return dataLocations
}
