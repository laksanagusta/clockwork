package repository

import (
	"clockwork-server/domain/model"

	"gorm.io/gorm"
)

type ProductAttributeRepository interface {
	DeleteByProductId(productId int) error
	CreateMany(productAttributes []model.ProductAttribute) error
}

type productAttributeRepository struct {
	db *gorm.DB
}

func NewProductAttributeRepository(db *gorm.DB) ProductAttributeRepository {
	return &productAttributeRepository{db}
}

func (pr *productAttributeRepository) DeleteByProductId(productId int) error {
	err := pr.db.Where("product_id = ?", productId).Delete(nil).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *productAttributeRepository) CreateMany(productAttributes []model.ProductAttribute) error {
	err := pr.db.Create(productAttributes).Error
	if err != nil {
		return err
	}

	return nil
}
