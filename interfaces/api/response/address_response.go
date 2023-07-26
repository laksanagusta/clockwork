package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Address struct {
	ID         uint      `json:"id"`
	Street     string    `json:"street"`
	Title      string    `json:"title"`
	Note       string    `json:"note"`
	City       string    `json:"city"`
	CustomerId uint      `json:"customerId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func FormatAddress(address model.Address) Address {
	var dataAddress Address

	dataAddress.ID = address.ID
	dataAddress.Title = address.Title
	dataAddress.Street = address.Street
	dataAddress.Note = address.Note
	dataAddress.City = address.City
	dataAddress.CustomerId = address.CustomerID
	dataAddress.CreatedAt = address.CreatedAt
	dataAddress.UpdatedAt = address.UpdatedAt

	return dataAddress
}

func FormatAddresss(address []model.Address) []Address {
	var dataAddresss []Address

	for _, value := range address {
		singleDataAddress := FormatAddress(value)

		dataAddresss = append(dataAddresss, singleDataAddress)
	}

	return dataAddresss
}
