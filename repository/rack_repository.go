package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"

	"gorm.io/gorm"
)

type RackRepository interface {
	Create(rack model.Rack) (model.Rack, error)
	Update(rack model.Rack) (model.Rack, error)
	FindById(rackId int) (model.Rack, error)
	FindAll(page int, page_size int, q string) ([]model.Rack, error)
	Delete(rackId int) (model.Rack, error)
}

type rackRepository struct {
	db *gorm.DB
}

func NewRackRepository(db *gorm.DB) RackRepository {
	return &rackRepository{db}
}

func (pr *rackRepository) Create(rack model.Rack) (model.Rack, error) {
	err := pr.db.Create(&rack).Error
	if err != nil {
		return rack, err
	}

	return rack, nil
}

func (pr *rackRepository) Update(rack model.Rack) (model.Rack, error) {
	err := pr.db.Save(&rack).Error
	if err != nil {
		return rack, err
	}

	return rack, nil
}

func (pr *rackRepository) FindById(rackId int) (model.Rack, error) {
	rack := model.Rack{}

	err := pr.db.First(&rack, rackId).Error
	if err != nil {
		return rack, err
	}

	return rack, nil
}

func (pr *rackRepository) FindAll(page int, limit int, q string) ([]model.Rack, error) {
	var rack []model.Rack

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

	err := querydb.Find(&rack).Error
	if err != nil {
		return rack, err
	}

	return rack, nil
}

func (pr *rackRepository) Delete(rackId int) (model.Rack, error) {
	var rack model.Rack
	err := pr.db.Where("id = ?", rackId).Delete(&rack).Error

	if err != nil {
		return rack, err
	}

	return rack, nil
}
