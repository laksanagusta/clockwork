package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
	"errors"
	"strings"
)

type CategoryService interface {
	Create(request request.CategoryCreateInput) (model.Category, error)
	Update(inputID request.CategoryFindById, request request.CategoryUpdateInput) (model.Category, error)
	FindById(categoryId int) (model.Category, error)
	FindAll(page int, limit int, q string) ([]model.Category, error)
	Delete(categoryId int) (model.Category, error)
}

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &categoryService{
		repository,
	}
}

func (s *categoryService) Create(request request.CategoryCreateInput) (model.Category, error) {
	category := model.Category{}
	category.Title = request.Title

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(category.Title))
	if err != nil {
		return category, err
	}

	if len(checkIfExist) > 0 {
		return category, errors.New("Category with same name already exist")
	}

	newCategory, err := s.repository.Create(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *categoryService) Update(inputID request.CategoryFindById, request request.CategoryUpdateInput) (model.Category, error) {
	category, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return category, err
	}

	category.Title = request.Title

	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil

}

func (s *categoryService) FindById(categoryId int) (model.Category, error) {
	category, err := s.repository.FindById(categoryId)
	if err != nil {
		return category, err
	}

	if category.ID == 0 {
		return category, errors.New("Category not found")
	}

	return category, nil
}

func (s *categoryService) FindAll(page int, limit int, q string) ([]model.Category, error) {
	categorys, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return categorys, err
	}
	return categorys, nil
}

func (s *categoryService) Delete(categoryId int) (model.Category, error) {
	category, err := s.repository.FindById(categoryId)
	if err != nil {
		return category, errors.New("category data not found")
	}

	_, err = s.repository.Delete(categoryId)
	if err != nil {
		return category, err
	}

	return category, nil
}
