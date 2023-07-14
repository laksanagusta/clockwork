package response

import (
	"clockwork-server/model"
	"time"
)

type Customer struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	Token       string    `json:"token"`
	Email       string    `json:"email"`
	Address     []Address `json:"address"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FormatCustomer(customer model.Customer, token string) Customer {
	var dataCustomer Customer

	dataCustomer.Name = customer.Name
	dataCustomer.Email = customer.Email
	dataCustomer.PhoneNumber = customer.PhoneNumber
	dataCustomer.Token = token
	dataCustomer.Address = FormatAddresss(customer.Address)
	dataCustomer.CreatedAt = customer.CreatedAt
	dataCustomer.UpdatedAt = customer.UpdatedAt

	return dataCustomer
}

func FormatCustomers(customer []model.Customer) []Customer {
	var dataCustomers []Customer

	for _, value := range customer {
		singleDataCustomer := FormatCustomer(value, "")

		dataCustomers = append(dataCustomers, singleDataCustomer)
	}

	return dataCustomers
}
