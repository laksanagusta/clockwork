package response

import (
	"clockwork-server/domain/model"
	"time"
)

type User struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Occupation  string    `json:"occupation"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
	Token       string    `json:"token"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FormatUser(user model.User, token string) User {
	var dataUser User

	dataUser.ID = user.ID
	dataUser.Name = user.Name
	dataUser.Address = user.Address
	dataUser.PhoneNumber = user.PhoneNumber
	dataUser.Occupation = user.Occupation
	dataUser.Token = token
	dataUser.CreatedAt = user.CreatedAt
	dataUser.UpdatedAt = user.UpdatedAt

	return dataUser
}

func FormatUsers(user []model.User) []User {
	var dataUsers []User

	for _, value := range user {
		singleDataUser := FormatUser(value, "")
		dataUsers = append(dataUsers, singleDataUser)
	}

	return dataUsers
}
