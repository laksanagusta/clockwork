package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID           uint64    `gorm:"size:64;not null;uniqueIndex;primary"`
	Name         string    `gorm:"size:20;not null;" json:"name"`
	PhoneNumber  string    `gorm:"size:12;not null;" json:"phoneNumber"`
	Email        string    `gorm:"size:40;not null;" json:"email"`
	PasswordHash string    `gorm:"size:255;" json:"password_hash"`
	Address      []Address `json:"address"`
}
