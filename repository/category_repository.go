package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category model.Category) (model.Category, error)
	Update(category model.Category) (model.Category, error)
	FindById(categoryId int) (model.Category, error)
	FindAll(page int, page_size int, q string) ([]model.Category, error)
	Delete(categoryId int) (model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (pr *categoryRepository) Create(category model.Category) (model.Category, error) {
	err := pr.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (pr *categoryRepository) Update(category model.Category) (model.Category, error) {
	err := pr.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (pr *categoryRepository) FindById(categoryId int) (model.Category, error) {
	category := model.Category{}

	err := pr.db.First(&category, categoryId).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (pr *categoryRepository) FindAll(page int, limit int, q string) ([]model.Category, error) {
	var category []model.Category

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

	err := querydb.Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (pr *categoryRepository) Delete(categoryId int) (model.Category, error) {
	var category model.Category
	err := pr.db.Where("id = ?", categoryId).Delete(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
