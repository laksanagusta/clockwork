package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
	"strings"
)

type AttributeItemService interface {
	Create(request request.AttributeItemCreateRequest) (model.AttributeItem, error)
	Update(inputID request.AttributeItemFindById, request request.AttributeItemUpdateRequest) (model.AttributeItem, error)
	FindById(attributeItemId int) (model.AttributeItem, error)
	FindAll(page int, limit int, q string) ([]model.AttributeItem, error)
	Delete(attributeItemId int) (model.AttributeItem, error)
}

type attributeItemService struct {
	repository repository.AttributeItemRepository
}

func NewAttributeItemService(repository repository.AttributeItemRepository) AttributeItemService {
	return &attributeItemService{
		repository,
	}
}

func (s *attributeItemService) Create(request request.AttributeItemCreateRequest) (model.AttributeItem, error) {
	attributeItem := model.AttributeItem{}
	attributeItem.Title = request.Title
	attributeItem.AdditionalCharge = request.AdditionalCharge
	attributeItem.AttributeID = request.AttributeID

	newAttributeItem, err := s.repository.Create(attributeItem)
	if err != nil {
		return newAttributeItem, err
	}

	return newAttributeItem, nil
}

func (s *attributeItemService) Update(inputID request.AttributeItemFindById, request request.AttributeItemUpdateRequest) (model.AttributeItem, error) {
	attributeItem, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return attributeItem, err
	}

	attributeItem.Title = request.Title
	attributeItem.AdditionalCharge = request.AdditionalCharge
	attributeItem.AttributeID = request.AttributeID

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(attributeItem.Title))
	if err != nil {
		return attributeItem, err
	}

	if len(checkIfExist) > 0 {
		return attributeItem, errors.New("AttributeItem with same name already exist")
	}

	updatedAttributeItem, err := s.repository.Update(attributeItem)
	if err != nil {
		return updatedAttributeItem, err
	}

	return updatedAttributeItem, nil

}

func (s *attributeItemService) FindById(attributeItemId int) (model.AttributeItem, error) {
	attributeItem, err := s.repository.FindById(attributeItemId)
	if err != nil {
		return attributeItem, err
	}

	if attributeItem.ID == 0 {
		return attributeItem, errors.New("AttributeItem not found")
	}

	return attributeItem, nil
}

func (s *attributeItemService) FindAll(page int, limit int, q string) ([]model.AttributeItem, error) {
	attributeItems, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return attributeItems, err
	}
	return attributeItems, nil
}

func (s *attributeItemService) Delete(attributeItemId int) (model.AttributeItem, error) {
	attributeItem, err := s.repository.Delete(attributeItemId)
	if err != nil {
		return attributeItem, err
	}

	return attributeItem, nil
}
