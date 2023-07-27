package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Image struct {
	ID        uint      `json:"id"`
	Url       string    `json:"url"`
	IsPrimary *bool     `json:"isPrimary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatImage(image model.Image) Image {
	var dataImage Image

	dataImage.ID = image.ID
	dataImage.Url = image.Url
	dataImage.IsPrimary = &image.IsPrimary
	dataImage.CreatedAt = image.CreatedAt
	dataImage.UpdatedAt = image.UpdatedAt

	return dataImage
}

func FormatImages(image []model.Image) []Image {
	var dataImages []Image

	for _, value := range image {
		singleDataImage := FormatImage(value)

		dataImages = append(dataImages, singleDataImage)
	}

	return dataImages
}
