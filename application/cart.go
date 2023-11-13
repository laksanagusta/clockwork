package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
	"errors"
	"fmt"
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
	imageRepo  repository.ImageRepository
}

func NewCartService(repository repository.CartRepository, imageRepo repository.ImageRepository) CartService {
	return &cartService{
		repository,
		imageRepo,
	}
}

func (s *cartService) Create(userId int) (model.Cart, error) {
	cartActive, _ := s.repository.FindOneByCustomerAndStatus(userId, "active")
	if cartActive.ID != 0 {
		return cartActive, nil
	}

	cart := model.Cart{
		BaseAmount: 0,
		TotalItem:  0,
		Status:     "active",
		UserID:     userId,
		//TODO : Need to check this
		VoucherID: 1,
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
	cart.BaseAmount = 0
	cart.TotalItem = 0
	for _, v := range cart.CartItems {
		fmt.Println(v.SubTotal)
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

func (s *cartService) CheckActiveCart(userId int) (model.Cart, error) {
	cart, err := s.repository.FindOneByCustomerAndStatus(userId, "active")
	if err != nil {
		return cart, err
	}

	var cartItems model.CartItems
	cartItems = cart.CartItems

	images, err := s.imageRepo.FindByProductIDs(cartItems.PopulateProductIDs())
	if err != nil {
		return cart, err
	}

	productImagesMapped := model.MappingImagesToProduct(images)

	for k, cartItem := range cartItems {
		val, ok := productImagesMapped[cartItem.Product.ID]
		if ok {
			cartItems[k].Product.Images = val
		}
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
