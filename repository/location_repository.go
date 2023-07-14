package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"

	"gorm.io/gorm"
)

type LocationRepository interface {
	Create(location model.Location) (model.Location, error)
	Update(location model.Location) (model.Location, error)
	FindById(locationId int) (model.Location, error)
	FindAll(page int, page_size int, q string) ([]model.Location, error)
	Delete(locationId int) (model.Location, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db}
}

func (pr *locationRepository) Create(location model.Location) (model.Location, error) {
	err := pr.db.Create(&location).Error
	if err != nil {
		return location, err
	}

	return location, nil
}

func (pr *locationRepository) Update(location model.Location) (model.Location, error) {
	err := pr.db.Save(&location).Error
	if err != nil {
		return location, err
	}

	return location, nil
}

func (pr *locationRepository) FindById(locationId int) (model.Location, error) {
	location := model.Location{}

	err := pr.db.First(&location, locationId).Error
	if err != nil {
		return location, err
	}

	return location, nil
}

func (pr *locationRepository) FindAll(page int, limit int, q string) ([]model.Location, error) {
	var location []model.Location

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
		querydb = querydb.Where("lower(name) LIKE ?", "%"+q+"%")
	}

	err := querydb.Find(&location).Error
	if err != nil {
		return location, err
	}

	return location, nil
}

func (pr *locationRepository) Delete(locationId int) (model.Location, error) {
	var location model.Location
	err := pr.db.Where("id = ?", locationId).Delete(&location).Error

	if err != nil {
		return location, err
	}

	return location, nil
}
