package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
)

type ProductService interface {
	Create(request request.ProductCreateInput) (model.Product, error)
	Update(inputID request.ProductFindById, request request.ProductUpdateInput) (model.Product, error)
	FindById(productId int) (model.Product, error)
	FindBySerialNumber(serialNumber string) (model.Product, error)
	FindAll(page int, page_size int, q string) ([]model.Product, error)
	Delete(productId int) (model.Product, error)
}

type productService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{
		repository,
	}
}

const PRODUCT_ALREADY_EXIST_NOTIF = "create product failed, product with sama name or title already exist"

func (s *productService) Create(request request.ProductCreateInput) (model.Product, error) {
	product := model.Product{}
	product.Title = request.Title
	product.Description = request.Description
	product.SerialNumber = request.SerialNumber
	product.UnitPrice = request.UnitPrice
	product.UserID = request.User.ID

	checkProduct, err := s.repository.FindBySerialNumberAndTitle(product.SerialNumber, product.Title)
	if err != nil {
		return product, err
	}

	if checkProduct.ID != 0 {
		return product, errors.New(PRODUCT_ALREADY_EXIST_NOTIF)
	}

	newProduct, err := s.repository.Create(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *productService) Update(inputID request.ProductFindById, request request.ProductUpdateInput) (model.Product, error) {
	product, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return product, err
	}

	product.Title = request.Title
	product.Description = request.Description
	product.UnitPrice = request.UnitPrice
	product.SerialNumber = request.SerialNumber
	product.UnitPrice = request.UnitPrice

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil
}

func (s *productService) FindById(productId int) (model.Product, error) {
	product, err := s.repository.FindById(productId)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("Product not found")
	}

	return product, nil
}

func (s *productService) FindBySerialNumber(serialNumber string) (model.Product, error) {
	product, err := s.repository.FindBySerialNumberAndTitle(serialNumber, "")
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("Product not found")
	}

	return product, nil
}

func (s *productService) FindAll(page int, pageSize int, q string) ([]model.Product, error) {
	products, err := s.repository.FindAll(page, pageSize, q)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *productService) Delete(productId int) (model.Product, error) {
	product, err := s.repository.Delete(productId)
	if err != nil {
		return product, err
	}

	return product, nil
}
