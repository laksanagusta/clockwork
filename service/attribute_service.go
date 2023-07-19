package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
	"strings"
)

type AttributeService interface {
	Create(request request.AttributeCreateRequest) (model.Attribute, error)
	Update(inputID request.AttributeFindById, request request.AttributeUpdateRequest) (model.Attribute, error)
	FindById(attributeId int) (model.Attribute, error)
	FindAll(page int, limit int, q string) ([]model.Attribute, error)
	Delete(attributeId int) (model.Attribute, error)
}

type attributeService struct {
	repository repository.AttributeRepository
}

func NewAttributeService(repository repository.AttributeRepository) AttributeService {
	return &attributeService{
		repository,
	}
}

func (s *attributeService) Create(request request.AttributeCreateRequest) (model.Attribute, error) {
	attribute := model.Attribute{}
	attribute.Title = request.Title
	attribute.IsMultiple = request.IsMultiple
	attribute.IsRequired = request.IsRequired

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(attribute.Title))
	if err != nil {
		return attribute, err
	}

	if len(checkIfExist) > 0 {
		return attribute, errors.New("Attribute with same name already exist")
	}

	newAttribute, err := s.repository.Create(attribute)
	if err != nil {
		return newAttribute, err
	}

	return newAttribute, nil
}

func (s *attributeService) Update(inputID request.AttributeFindById, request request.AttributeUpdateRequest) (model.Attribute, error) {
	attribute, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return attribute, err
	}

	attribute.Title = request.Title
	attribute.IsMultiple = request.IsMultiple
	attribute.IsRequired = request.IsRequired

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(attribute.Title))
	if err != nil {
		return attribute, err
	}

	if len(checkIfExist) > 0 {
		return attribute, errors.New("Attribute with same name already exist")
	}

	updatedAttribute, err := s.repository.Update(attribute)
	if err != nil {
		return updatedAttribute, err
	}

	return updatedAttribute, nil

}

func (s *attributeService) FindById(attributeId int) (model.Attribute, error) {
	attribute, err := s.repository.FindById(attributeId)
	if err != nil {
		return attribute, err
	}

	if attribute.ID == 0 {
		return attribute, errors.New("Attribute not found")
	}

	return attribute, nil
}

func (s *attributeService) FindAll(page int, limit int, q string) ([]model.Attribute, error) {
	attributes, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return attributes, err
	}
	return attributes, nil
}

func (s *attributeService) Delete(attributeId int) (model.Attribute, error) {
	attribute, err := s.repository.Delete(attributeId)
	if err != nil {
		return attribute, err
	}

	return attribute, nil
}
