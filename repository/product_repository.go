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

type productQueryResult struct {
	ProductTitle string
}

func (pr *productRepository) FindById(productId int) (model.Product, error) {
	product := model.Product{}
	attributes := []model.Attribute{}
	// err := pr.db.Table("products").
	// 	Select("products.title as product_title, products.unit_price as products_unit_price, attributes.title as attributes_title, attribute_items.title as attribute_items_title").
	// 	Joins("LEFT JOIN product_attributes on product_attributes.product_id = products.id").
	// 	Joins("LEFT JOIN attributes on attributes.id = product_attributes.attribute_id").
	// 	Joins("LEFT JOIN attribute_items on attribute_items.attribute_id = attributes.id").
	// 	Where("products.id = ?", productId).
	// 	Find(&result).Error

	err := pr.db.Preload("Category").Preload("User").Preload("Images").Find(&product, productId).Error
	if err != nil {
		return product, err
	}

	err = pr.db.Table("attributes").
		Select("attributes.*").
		Joins("JOIN product_attributes on attributes.id = product_attributes.attribute_id").
		Where("product_attributes.product_id = ?", productId).
		Find(&attributes).Error

	for k, v := range attributes {
		attributeItems := []model.AttributeItem{}
		err := pr.db.Where("attribute_id = ?", v.ID).Find(&attributeItems).Error
		if err != nil {
			return product, err
		}

		attributes[k].AttributeItem = attributeItems
	}

	product.Attributes = attributes

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
