package repository

import (
	"clockwork-server/model"

	"gorm.io/gorm"
)

type CartItemAttributeItemRepository interface {
	Create(cartItemAttributeItem model.CartItemAttributeItem) (model.CartItemAttributeItem, error)
	Update(cartItemAttributeItem model.CartItemAttributeItem) (model.CartItemAttributeItem, error)
	FindById(cartItemAttributeItemId int) (model.CartItemAttributeItem, error)
	Delete(cartItemAttributeItemId int) (model.CartItemAttributeItem, error)
	CreateMany(cartItemAttributeItem []model.CartItemAttributeItem) error
	DeleteByCartItemId(cartItemId uint) error
}

type cartItemAttributeItemRepository struct {
	db *gorm.DB
}

func NewCartItemAttributeItemRepository(db *gorm.DB) CartItemAttributeItemRepository {
	return &cartItemAttributeItemRepository{db}
}

func (pr *cartItemAttributeItemRepository) Create(cartItemAttributeItem model.CartItemAttributeItem) (model.CartItemAttributeItem, error) {
	err := pr.db.Create(&cartItemAttributeItem).Error
	if err != nil {
		return cartItemAttributeItem, err
	}

	return cartItemAttributeItem, nil
}

func (pr *cartItemAttributeItemRepository) CreateMany(cartItemAttributeItem []model.CartItemAttributeItem) error {
	err := pr.db.Create(&cartItemAttributeItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *cartItemAttributeItemRepository) Update(cartItemAttributeItem model.CartItemAttributeItem) (model.CartItemAttributeItem, error) {
	err := pr.db.Save(&cartItemAttributeItem).Error
	if err != nil {
		return cartItemAttributeItem, err
	}

	return cartItemAttributeItem, nil
}

func (pr *cartItemAttributeItemRepository) FindById(cartItemAttributeItemId int) (model.CartItemAttributeItem, error) {
	cartItemAttributeItem := model.CartItemAttributeItem{}

	err := pr.db.First(&cartItemAttributeItem, cartItemAttributeItemId).Error
	if err != nil {
		return cartItemAttributeItem, err
	}

	return cartItemAttributeItem, nil
}

func (pr *cartItemAttributeItemRepository) DeleteByCartItemId(cartItemId uint) error {
	err := pr.db.Where("order_item_id = ?", cartItemId).Delete(nil).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *cartItemAttributeItemRepository) Delete(cartItemAttributeItemId int) (model.CartItemAttributeItem, error) {
	var cartItemAttributeItem model.CartItemAttributeItem
	err := pr.db.Where("id = ?", cartItemAttributeItemId).Delete(&cartItemAttributeItem).Error

	if err != nil {
		return cartItemAttributeItem, err
	}

	return cartItemAttributeItem, nil
}
