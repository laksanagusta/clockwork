package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint64 `gorm:"size:64;not null;uniqueIndex;primary"`
	Name         string `gorm:"size:50;not null;" json:"name"`
	PhoneNumber  string `gorm:"size:12;not null;" json:"phoneNumber"`
	Address      string `gorm:"size:255;" json:"address"`
	Email        string `gorm:"size:40;" json:"email"`
	Occupation   string `gorm:"size:40;" json:"occupation"`
	PasswordHash string `gorm:"size:255;" json:"password_hash"`
	Username     string `gorm:"size:20;not null;unique" json:"username"`
	Role         string `gorm:"size:6;not null;unique" json:"role"`
}
