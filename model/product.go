package model

import (
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title        string `gorm:"size:60;not null;" json:"title"`
	Description  string `gorm:"size:60;not null;" json:"description"`
	SerialNumber string `gorm:"size:60;not null;" json:"serialNumber"`
	UnitPrice    int    `gorm:"size:60;not null;" json:"unitPrice"`
	UserID       uint64
	User         User
	Inventory    Inventory
	Images       []Image
}

func (c *Product) BeforeCreate(trx *gorm.DB) error {
	c.SerialNumber = strings.ToUpper(c.SerialNumber)
	return nil
}
