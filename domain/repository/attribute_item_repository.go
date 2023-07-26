package repository

import (
	"clockwork-server/domain/model"
	"clockwork-server/helper"

	"gorm.io/gorm"
)

type AttributeItemRepository interface {
	Create(attributeItem model.AttributeItem) (model.AttributeItem, error)
	Update(attributeItem model.AttributeItem) (model.AttributeItem, error)
	FindById(attributeItemId int) (model.AttributeItem, error)
	FindAll(page int, page_size int, q string) ([]model.AttributeItem, error)
	Delete(attributeItemId int) (model.AttributeItem, error)
}

type attributeItemRepository struct {
	db *gorm.DB
}

func NewAttributeItemRepository(db *gorm.DB) AttributeItemRepository {
	return &attributeItemRepository{db}
}

func (pr *attributeItemRepository) Create(attributeItem model.AttributeItem) (model.AttributeItem, error) {
	err := pr.db.Create(&attributeItem).Error
	if err != nil {
		return attributeItem, err
	}

	return attributeItem, nil
}

func (pr *attributeItemRepository) Update(attributeItem model.AttributeItem) (model.AttributeItem, error) {
	err := pr.db.Save(&attributeItem).Error
	if err != nil {
		return attributeItem, err
	}

	return attributeItem, nil
}

func (pr *attributeItemRepository) FindById(attributeItemId int) (model.AttributeItem, error) {
	attributeItem := model.AttributeItem{}

	err := pr.db.First(&attributeItem, attributeItemId).Error
	if err != nil {
		return attributeItem, err
	}

	return attributeItem, nil
}

func (pr *attributeItemRepository) FindAll(page int, limit int, q string) ([]model.AttributeItem, error) {
	var attributeItem []model.AttributeItem

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

	err := querydb.Find(&attributeItem).Error
	if err != nil {
		return attributeItem, err
	}

	return attributeItem, nil
}

func (pr *attributeItemRepository) Delete(attributeItemId int) (model.AttributeItem, error) {
	var attributeItem model.AttributeItem
	err := pr.db.Where("id = ?", attributeItemId).Delete(&attributeItem).Error

	if err != nil {
		return attributeItem, err
	}

	return attributeItem, nil
}
