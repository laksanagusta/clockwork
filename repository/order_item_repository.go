package repository

import (
	"clockwork-server/model"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(orderItem model.OrderItem) (model.OrderItem, error)
	Update(orderItem model.OrderItem) (model.OrderItem, error)
	FindById(orderItemId int) (model.OrderItem, error)
	FindByCode(code string) (model.OrderItem, error)
	FindAll(page int, page_size int, q string) ([]model.OrderItem, error)
	Delete(orderItemId int) (model.OrderItem, error)
	FindByOrderId(orderId int) ([]model.OrderItem, error)
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{db}
}

func (pr *orderItemRepository) Create(orderItem model.OrderItem) (model.OrderItem, error) {
	err := pr.db.Create(&orderItem).Error

	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}

func (pr *orderItemRepository) Update(orderItem model.OrderItem) (model.OrderItem, error) {
	err := pr.db.Save(&orderItem).Error

	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}

func (pr *orderItemRepository) FindById(orderItemId int) (model.OrderItem, error) {
	orderItem := model.OrderItem{}
	err := pr.db.Preload("order").First(&orderItem, orderItemId).Error

	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}

func (pr *orderItemRepository) FindByCode(code string) (model.OrderItem, error) {
	orderItem := model.OrderItem{}
	err := pr.db.Where("code = ?", code).Find(&orderItem).Error

	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}

func (pr *orderItemRepository) FindByOrderId(orderId int) ([]model.OrderItem, error) {
	var orderItem []model.OrderItem

	err := pr.db.Where("orderId = ?", orderId).Find(&orderItem).Error

	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}

func (pr *orderItemRepository) FindAll(page int, pageSize int, q string) ([]model.OrderItem, error) {
	var orderItem []model.OrderItem

	querydb := pr.db.Offset(page).Limit(pageSize).Where("is_deleted = ?", 0)

	if q != "" {
		querydb.Where("title LIKE ?", "%"+q+"%")
	}

	err := querydb.Find(&orderItem).Error
	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}

func (pr *orderItemRepository) Delete(orderItemId int) (model.OrderItem, error) {
	var orderItem model.OrderItem
	err := pr.db.Model(&orderItem).Where("id = ?", orderItemId).Update("is_deleted", 1).Error

	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}
