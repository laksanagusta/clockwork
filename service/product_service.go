package service

import (
	"clockwork-server/helper"
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
	repository      repository.ProductRepository
	categoryRepo    repository.CategoryRepository
	inventoryRepo   repository.InventoryRepository
	productAttrRepo repository.ProductAttributeRepository
}

func NewProductService(
	repository repository.ProductRepository,
	categoryRepo repository.CategoryRepository,
	inventoryRepo repository.InventoryRepository,
	productAttrRepo repository.ProductAttributeRepository,
) ProductService {
	return &productService{
		repository,
		categoryRepo,
		inventoryRepo,
		productAttrRepo,
	}
}

const PRODUCT_ALREADY_EXIST_NOTIF = "create product failed, product with sama name or title already exist"
const PRODUCT_NOT_FOUND = "Product not found"

func (s *productService) Create(request request.ProductCreateInput) (model.Product, error) {
	product := model.Product{}
	product.Title = request.Title
	product.Description = request.Description
	product.SerialNumber = request.SerialNumber
	product.UnitPrice = request.UnitPrice
	product.UserID = request.User.ID
	product.CategoryID = request.CategoryID

	_, err := s.categoryRepo.FindById(int(product.CategoryID))
	if err != nil {
		return product, errors.New(helper.CATEGORY_NOT_FOUND_MESSAGE)
	}

	checkSameProduct, err := s.repository.FindBySerialNumberAndTitle(product.SerialNumber, product.Title)
	if err != nil {
		return product, err
	}

	if checkSameProduct.ID != 0 {
		return product, errors.New(helper.PRODUCT_EXIST_MESSAGE)
	}

	inventory := model.Inventory{
		StockQty:    0,
		IsInStock:   false,
		ReservedQty: 0,
		SalableQty:  0,
	}

	newInventory, err := s.inventoryRepo.Create(inventory)
	if err != nil {
		return product, nil
	}

	product.InventoryID = newInventory.ID

	newProduct, err := s.repository.Create(product)
	if err != nil {
		return newProduct, err
	}

	productAttributes := []model.ProductAttribute{}

	for _, value := range request.Attributes {
		productAttribute := model.ProductAttribute{
			ProductID:   int(newProduct.ID),
			AttributeID: value,
		}

		productAttributes = append(productAttributes, productAttribute)
	}

	err = s.productAttrRepo.CreateMany(productAttributes)
	if err != nil {
		return newProduct, err
	}

	getNewProduct, err := s.repository.FindById(int(newProduct.ID))
	if err != nil {
		return getNewProduct, err
	}

	return getNewProduct, nil
}

func (s *productService) Update(inputID request.ProductFindById, request request.ProductUpdateInput) (model.Product, error) {
	product, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return product, errors.New(helper.PRODUCT_NOT_FOUND_MESSAGE)
	}

	product.Title = request.Title
	product.Description = request.Description
	product.SerialNumber = request.SerialNumber
	product.UnitPrice = request.UnitPrice
	product.CategoryID = request.CategoryID

	_, err = s.categoryRepo.FindById(int(product.CategoryID))
	if err != nil {
		return product, errors.New(helper.CATEGORY_NOT_FOUND_MESSAGE)
	}

	checkSameProduct, err := s.repository.FindBySerialNumberAndTitle(product.SerialNumber, product.Title)
	if err != nil {
		return product, err
	}

	if checkSameProduct.ID != 0 {
		return product, errors.New(helper.PRODUCT_EXIST_MESSAGE)
	}

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	s.productAttrRepo.DeleteByProductId(int(updatedProduct.ID))

	productAttributes := []model.ProductAttribute{}

	for _, value := range request.Attributes {
		productAttribute := model.ProductAttribute{
			ProductID:   inputID.ID,
			AttributeID: value,
		}

		productAttributes = append(productAttributes, productAttribute)
	}

	err = s.productAttrRepo.CreateMany(productAttributes)
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
		return product, errors.New(PRODUCT_NOT_FOUND)
	}

	return product, nil
}

func (s *productService) FindBySerialNumber(serialNumber string) (model.Product, error) {
	product, err := s.repository.FindBySerialNumberAndTitle(serialNumber, "")
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New(PRODUCT_NOT_FOUND)
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
