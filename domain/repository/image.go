package repository

import (
	"clockwork-server/domain/model"

	"gorm.io/gorm"
)

type ImageRepository interface {
	Create(model.Image) (model.Image, error)
	Remove(id int8) (int8, error)
	GetImagesByProductId(productId int) ([]model.Image, error)
	UpdateIsPrimaryFalse(productId int) error
	FindById(id int8) (model.Image, error)
	Update(image model.Image) (model.Image, error)
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{db}
}

func (r *imageRepository) Create(image model.Image) (model.Image, error) {
	err := r.db.Create(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *imageRepository) Update(image model.Image) (model.Image, error) {
	err := r.db.Save(&image).Error
	if err != nil {
		return image, err
	}

	return image, nil
}

func (r *imageRepository) GetImagesByProductId(productId int) ([]model.Image, error) {
	images := []model.Image{}

	err := r.db.Where("product_id = ?", productId).Find(&images).Error
	if err != nil {
		return images, err
	}

	return images, nil
}

func (r *imageRepository) UpdateIsPrimaryFalse(productId int) error {
	images := []model.Image{}
	err := r.db.Raw("UPDATE images SET is_primary = false WHERE product_id = ?", productId).Scan(&images).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *imageRepository) Remove(id int8) (int8, error) {
	image := model.Image{}

	err := r.db.Delete(&image, id).Error
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *imageRepository) FindById(id int8) (model.Image, error) {
	var image model.Image

	err := r.db.First(&image, id).Error
	if err != nil {
		return image, err
	}

	return image, nil
}
