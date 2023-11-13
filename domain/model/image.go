package model

import (
	"time"
)

type Image struct {
	ID        uint   `gorm:"primarykey"`
	Url       string `gorm:"size:60;not null;" json:"url"`
	ProductID uint
	IsPrimary bool `gorm:"not null;" json:"isPrimary"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MappingImagesToProduct(i []Image) map[uint][]Image {
	res := make(map[uint][]Image)

	for _, v := range i {
		_, ok := res[v.ProductID]
		if !ok {
			res[v.ProductID] = []Image{}
		}
		res[v.ProductID] = append(res[v.ProductID], v)
	}

	return res
}
