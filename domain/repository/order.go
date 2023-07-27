package repository

import (
	"clockwork-server/domain/model"
	"clockwork-server/helper"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order model.Order) (model.Order, error)
	Update(order model.Order) (model.Order, error)
	FindById(orderId int) (model.Order, error)
	FindByCode(code string) (model.Order, error)
	FindAll(page int, limit int, customerID string) ([]model.Order, error)
	FindOngoingOrder(userId int) (model.Order, error)
	Delete(orderId int) (model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (pr *orderRepository) Create(order model.Order) (model.Order, error) {
	err := pr.db.Create(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (pr *orderRepository) Update(order model.Order) (model.Order, error) {
	err := pr.db.Save(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (pr *orderRepository) FindById(orderId int) (model.Order, error) {
	order := model.Order{}
	err := pr.db.Preload("Payment").Preload("Cart").Preload("Customer").Find(&order, orderId).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (pr *orderRepository) FindOngoingOrder(orderId int) (model.Order, error) {
	order := model.Order{}
	err := pr.db.Where("status = ?", "created").First(&order, orderId).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (pr *orderRepository) FindByCode(code string) (model.Order, error) {
	order := model.Order{}
	err := pr.db.Where("code = ?", code).Find(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (pr *orderRepository) FindAll(page int, limit int, customerID string) ([]model.Order, error) {
	var order []model.Order

	querydb := pr.db

	if limit > 0 {
		querydb = querydb.Limit(limit)
	} else {
		querydb = querydb.Limit(helper.QUERY_LIMITATION)
	}

	if page > 0 {
		querydb = querydb.Offset(page - 1)
	}

	if customerID != "" {
		querydb.Where("customer_id = ?", customerID)
	}

	err := querydb.Preload("Cart").Preload("Payment").Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (pr *orderRepository) Delete(orderId int) (model.Order, error) {
	var order model.Order
	err := pr.db.Model(&order).Where("id = ?", orderId).Update("is_deleted", 1).Error

	if err != nil {
		return order, err
	}

	return order, nil
}
