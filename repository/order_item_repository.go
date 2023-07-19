package repository

import (
	"clockwork-server/model"

	"gorm.io/gorm"
)

type CartItemRepository interface {
	Create(cartItem model.CartItem) (model.CartItem, error)
	Update(cartItem model.CartItem) (model.CartItem, error)
	FindById(cartItemId int) (model.CartItem, error)
	FindByCode(code string) (model.CartItem, error)
	FindAll(page int, page_size int, q string) ([]model.CartItem, error)
	Delete(cartItemId int) (model.CartItem, error)
	FindByOrderId(orderId int) ([]model.CartItem, error)
	FindByProductId(productId int) ([]model.CartItem, error)
}

type cartItemRepository struct {
	db *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) CartItemRepository {
	return &cartItemRepository{db}
}

func (pr *cartItemRepository) Create(cartItem model.CartItem) (model.CartItem, error) {
	err := pr.db.Create(&cartItem).Error

	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (pr *cartItemRepository) Update(cartItem model.CartItem) (model.CartItem, error) {
	err := pr.db.Save(&cartItem).Error

	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (pr *cartItemRepository) FindById(cartItemId int) (model.CartItem, error) {
	cartItem := model.CartItem{}
	err := pr.db.Preload("order").First(&cartItem, cartItemId).Error

	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (pr *cartItemRepository) FindByCode(code string) (model.CartItem, error) {
	cartItem := model.CartItem{}
	err := pr.db.Where("code = ?", code).Find(&cartItem).Error

	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (pr *cartItemRepository) FindByOrderId(orderId int) ([]model.CartItem, error) {
	var cartItem []model.CartItem

	err := pr.db.Where("orderId = ?", orderId).Find(&cartItem).Error

	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (pr *cartItemRepository) FindByProductId(productId int) ([]model.CartItem, error) {
	var cartItem []model.CartItem

	err := pr.db.Where("productId = ?", productId).Find(&cartItem).Error

	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (pr *cartItemRepository) FindAll(page int, pageSize int, q string) ([]model.CartItem, error) {
	var cartItem []model.CartItem

	querydb := pr.db.Offset(page).Limit(pageSize).Where("is_deleted = ?", 0)

	if q != "" {
		querydb.Where("title LIKE ?", "%"+q+"%")
	}

	err := querydb.Find(&cartItem).Error
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (pr *cartItemRepository) Delete(cartItemId int) (model.CartItem, error) {
	var cartItem model.CartItem
	err := pr.db.Model(&cartItem).Where("id = ?", cartItemId).Update("is_deleted", 1).Error

	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}
