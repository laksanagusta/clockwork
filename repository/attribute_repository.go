package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"

	"gorm.io/gorm"
)

type AttributeRepository interface {
	Create(attribute model.Attribute) (model.Attribute, error)
	Update(attribute model.Attribute) (model.Attribute, error)
	FindById(attributeId int) (model.Attribute, error)
	FindAll(page int, page_size int, q string) ([]model.Attribute, error)
	Delete(attributeId int) (model.Attribute, error)
	FindAttributeWithAttributeItemByProductId(productId int) ([]model.Attribute, error)
}

type attributeRepository struct {
	db *gorm.DB
}

func NewAttributeRepository(db *gorm.DB) AttributeRepository {
	return &attributeRepository{db}
}

func (pr *attributeRepository) Create(attribute model.Attribute) (model.Attribute, error) {
	err := pr.db.Create(&attribute).Error
	if err != nil {
		return attribute, err
	}

	return attribute, nil
}

func (pr *attributeRepository) Update(attribute model.Attribute) (model.Attribute, error) {
	err := pr.db.Save(&attribute).Error
	if err != nil {
		return attribute, err
	}

	return attribute, nil
}

func (pr *attributeRepository) FindById(attributeId int) (model.Attribute, error) {
	attribute := model.Attribute{}

	err := pr.db.First(&attribute, attributeId).Error
	if err != nil {
		return attribute, err
	}

	return attribute, nil
}

func (pr *attributeRepository) FindAll(page int, limit int, q string) ([]model.Attribute, error) {
	var attribute []model.Attribute

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

	err := querydb.Find(&attribute).Error
	if err != nil {
		return attribute, err
	}

	return attribute, nil
}

func (pr *attributeRepository) Delete(attributeId int) (model.Attribute, error) {
	var attribute model.Attribute
	err := pr.db.Where("id = ?", attributeId).Delete(&attribute).Error

	if err != nil {
		return attribute, err
	}

	return attribute, nil
}

func (pr *attributeRepository) FindAttributeWithAttributeItemByProductId(productId int) ([]model.Attribute, error) {
	attributes := []model.Attribute{}

	err := pr.db.Table("attributes").
		Select("attributes.*").
		Joins("JOIN product_attributes on attributes.id = product_attributes.attribute_id").
		Where("product_attributes.product_id = ?", productId).
		Find(&attributes).Error

	if err != nil {
		return attributes, err
	}

	for k, v := range attributes {
		attributeItems := []model.AttributeItem{}
		err := pr.db.Where("attribute_id = ?", v.ID).Find(&attributeItems).Error
		if err != nil {
			return attributes, err
		}

		attributes[k].AttributeItem = attributeItems
	}

	return attributes, nil
}
