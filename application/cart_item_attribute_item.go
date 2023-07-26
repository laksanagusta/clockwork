package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
)

type CartItemAttributeItemService interface {
	CreateMany(cartItemReq []request.AttributeItem, cartItemId uint) ([]model.CartItemAttributeItem, error)
}

type cartItemAttributeItemService struct {
	repository repository.CartItemAttributeItemRepository
}

func NewCartItemAttributeItemService(repository repository.CartItemAttributeItemRepository) CartItemAttributeItemService {
	return &cartItemAttributeItemService{
		repository,
	}
}

func (s *cartItemAttributeItemService) CreateMany(attributeItem []request.AttributeItem, cartItemId uint) ([]model.CartItemAttributeItem, error) {
	attributeItems := []model.CartItemAttributeItem{}
	for _, value := range attributeItem {
		attributeItem := model.CartItemAttributeItem{
			CartItemID:       cartItemId,
			AttributeItemID:  value.ID,
			AdditionalCharge: value.AdditionalCharge,
		}

		attributeItems = append(attributeItems, attributeItem)
	}

	err := s.repository.CreateMany(attributeItems)
	if err != nil {
		return attributeItems, err
	}

	return attributeItems, nil
}
