package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
)

type CartService interface {
	Create(customerId int) (model.Cart, error)
	FindById(cartId int) (model.Cart, error)
	FindAll(page int, limit int, q string) ([]model.Cart, error)
	Delete(cartId int) (model.Cart, error)
	RecalculateCart(cart model.Cart) model.Cart
	CheckActiveCart(customerId int) (model.Cart, error)
	Update(inputID request.CartFindById, req request.CartUpdateRequest) (model.Cart, error)
}

type cartService struct {
	repository repository.CartRepository
}

func NewCartService(repository repository.CartRepository) CartService {
	return &cartService{
		repository,
	}
}

func (s *cartService) Create(customerId int) (model.Cart, error) {
	cartActive, _ := s.repository.FindOneByCustomerAndStatus(customerId, "active")
	if cartActive.ID != 0 {
		return cartActive, nil
	}

	cart := model.Cart{
		BaseAmount: 0,
		TotalItem:  0,
		Status:     "active",
		CustomerID: customerId,
	}

	newCart, err := s.repository.Create(cart)
	if err != nil {
		return newCart, err
	}

	return newCart, nil
}

func (s *cartService) Update(inputID request.CartFindById, req request.CartUpdateRequest) (model.Cart, error) {
	cart, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return cart, err
	}

	cart.BaseAmount = req.BaseAmount
	cart.TotalItem = req.TotalItem
	cart.Status = req.Status

	updateCart, err := s.repository.Update(cart)
	if err != nil {
		return updateCart, err
	}

	return updateCart, nil
}

func (s *cartService) RecalculateCart(cart model.Cart) model.Cart {
	for _, v := range cart.CartItems {
		cart.BaseAmount += v.SubTotal
		cart.TotalItem += 1
	}

	return cart
}

func (s *cartService) FindById(cartId int) (model.Cart, error) {
	cart, err := s.repository.FindById(cartId)
	if err != nil {
		return cart, err
	}

	if cart.ID == 0 {
		return cart, errors.New("Cart not found")
	}

	return cart, nil
}

func (s *cartService) FindAll(page int, limit int, q string) ([]model.Cart, error) {
	carts, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return carts, err
	}
	return carts, nil
}

func (s *cartService) CheckActiveCart(customerId int) (model.Cart, error) {
	cart, err := s.repository.FindOneByCustomerAndStatus(customerId, "active")
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (s *cartService) Delete(cartId int) (model.Cart, error) {
	cart, err := s.repository.Delete(cartId)
	if err != nil {
		return cart, err
	}

	return cart, nil
}
