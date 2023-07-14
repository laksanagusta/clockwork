package repository

import (
	"clockwork-server/model"

	"gorm.io/gorm"
)

type ImageRepository interface {
	Create(model.Image) (model.Image, error)
	Remove(id int8) (int8, error)
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{db}
}

func (imageRepository *imageRepository) Create(image model.Image) (model.Image, error) {
	err := imageRepository.db.Create(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (imageRepository *imageRepository) Remove(id int8) (int8, error) {
	err := imageRepository.db.Delete(id).Error
	if err != nil {
		return id, err
	}

	return id, nil
}
