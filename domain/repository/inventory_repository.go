package repository

import (
	"clockwork-server/domain/model"
	"clockwork-server/helper"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	Create(inventory model.Inventory) (model.Inventory, error)
	Update(inventory model.Inventory) (model.Inventory, error)
	FindById(inventoryId int) (model.Inventory, error)
	FindByProductId(productId int) (model.Inventory, error)
	FindAll(page int, page_size int, q string) ([]model.Inventory, error)
	Delete(inventoryId int) (model.Inventory, error)
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db}
}

func (pr *inventoryRepository) Create(inventory model.Inventory) (model.Inventory, error) {
	err := pr.db.Create(&inventory).Error

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func (pr *inventoryRepository) Update(inventory model.Inventory) (model.Inventory, error) {
	err := pr.db.Save(&inventory).Error

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func (pr *inventoryRepository) FindById(inventoryId int) (model.Inventory, error) {
	inventory := model.Inventory{}
	err := pr.db.First(&inventory, inventoryId).Error

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func (pr *inventoryRepository) FindByProductId(productId int) (model.Inventory, error) {
	inventory := model.Inventory{}
	err := pr.db.Where("product_id = ?", productId).Find(&inventory).Error

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func (pr *inventoryRepository) FindAll(page int, limit int, q string) ([]model.Inventory, error) {
	var inventory []model.Inventory

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

	err := querydb.Find(&inventory).Error
	if err != nil {
		return inventory, err
	}

	return inventory, nil
}

func (pr *inventoryRepository) Delete(inventoryId int) (model.Inventory, error) {
	var inventory model.Inventory
	err := pr.db.Model(&inventory).Where("id = ?", inventoryId).Update("is_deleted", 1).Error

	if err != nil {
		return inventory, err
	}

	return inventory, nil
}
