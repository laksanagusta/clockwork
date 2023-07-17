package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"
	"strings"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product model.Product) (model.Product, error)
	Update(product model.Product) (model.Product, error)
	FindById(productId int) (model.Product, error)
	FindBySerialNumberAndTitle(serialNumber string, title string) (model.Product, error)
	FindAll(page int, page_size int, q string) ([]model.Product, error)
	Delete(productId int) (model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) Create(product model.Product) (model.Product, error) {
	err := pr.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepository) Update(product model.Product) (model.Product, error) {
	err := pr.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepository) FindById(productId int) (model.Product, error) {
	product := model.Product{}
	err := pr.db.Preload("Inventory").Preload("User").Preload("Category").Preload("Images").First(&product, productId).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepository) FindBySerialNumberAndTitle(serialNumber string, title string) (model.Product, error) {
	product := model.Product{}

	err := pr.db.Where("LOWER(serial_number) = ?", strings.ToLower(serialNumber)).Or("LOWER(title) = ?", strings.ToLower(title)).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepository) FindAll(page int, limit int, q string) ([]model.Product, error) {
	var product []model.Product

	querydb := pr.db

	if limit > 0 {
		querydb = querydb.Limit(limit)
	} else {
		querydb = querydb.Limit(helper.QUERY_LIMITATION)
	}

	if page > 0 {
		querydb = querydb.Offset(page - 1)
	}

	if q != "" {
		querydb = querydb.Where("lower(title) LIKE ?", "%"+q+"%")
	}

	err := querydb.Preload("Inventory").Preload("User").Preload("Category").Preload("Images").Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepository) Delete(productId int) (model.Product, error) {
	var product model.Product
	err := pr.db.Model(&product).Where("id = ?", productId).Update("is_deleted", 1).Error

	if err != nil {
		return product, err
	}

	return product, nil
}
