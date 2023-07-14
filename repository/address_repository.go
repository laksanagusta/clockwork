package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"

	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address model.Address) (model.Address, error)
	Update(address model.Address) (model.Address, error)
	FindById(addressId int) (model.Address, error)
	FindAll(page int, page_size int, q string) ([]model.Address, error)
	Delete(addressId int) (model.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (pr *addressRepository) Create(address model.Address) (model.Address, error) {
	err := pr.db.Create(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (pr *addressRepository) Update(address model.Address) (model.Address, error) {
	err := pr.db.Save(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (pr *addressRepository) FindById(addressId int) (model.Address, error) {
	address := model.Address{}

	err := pr.db.First(&address, addressId).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (pr *addressRepository) FindAll(page int, limit int, q string) ([]model.Address, error) {
	var address []model.Address

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

	err := querydb.Find(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}

func (pr *addressRepository) Delete(addressId int) (model.Address, error) {
	var address model.Address
	err := pr.db.Where("id = ?", addressId).Delete(&address).Error

	if err != nil {
		return address, err
	}

	return address, nil
}
