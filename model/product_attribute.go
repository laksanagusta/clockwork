package model

type ProductAttribute struct {
	ProductID   int `gorm:"not null" json:"productId"`
	AttributeID int `gorm:"not null" json:"attributeId"`
}
