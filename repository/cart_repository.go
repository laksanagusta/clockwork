package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	Create(cart model.Cart) (model.Cart, error)
	Update(cart model.Cart) (model.Cart, error)
	FindById(cartId int) (model.Cart, error)
	FindAll(page int, page_size int, q string) ([]model.Cart, error)
	FindOneByCustomerAndStatus(customerId int, status string) (model.Cart, error)
	Delete(cartId int) (model.Cart, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) Create(cart model.Cart) (model.Cart, error) {
	err := r.db.Create(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartRepository) Update(cart model.Cart) (model.Cart, error) {
	err := r.db.Save(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartRepository) FindById(cartId int) (model.Cart, error) {
	cart := model.Cart{}

	err := r.db.Preload("CartItems").First(&cart, cartId).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartRepository) FindOneByCustomerAndStatus(customerId int, status string) (model.Cart, error) {
	cart := model.Cart{}

	err := r.db.Where("customer_id = ?", customerId).Where("status = ?", status).Preload("CartItems").First(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartRepository) FindAll(page int, limit int, q string) ([]model.Cart, error) {
	var cart []model.Cart

	querydb := r.db

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

	err := querydb.Find(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartRepository) Delete(cartId int) (model.Cart, error) {
	var cart model.Cart
	err := r.db.Where("id = ?", cartId).Delete(&cart).Error

	if err != nil {
		return cart, err
	}

	return cart, nil
}
