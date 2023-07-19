package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
)

type CartItemService interface {
	Create(cartItemReq request.CartItemCreateRequest, customerId int) (model.Cart, error)
	Update(inputID request.CartItemFindById, request request.CartItemUpdateRequest, customerId int) (model.Cart, error)
	FindById(cartItemId int) (model.CartItem, error)
	FindByCode(code string) (model.CartItem, error)
	FindAll(page int, page_size int, q string) ([]model.CartItem, error)
	Delete(cartItemId int) (model.CartItem, error)
}

type cartItemService struct {
	repository                      repository.CartItemRepository
	cartRepository                  repository.CartRepository
	inventoryRepository             repository.InventoryRepository
	inventoryService                InventoryService
	cartItemAttributeItemRepository repository.CartItemAttributeItemRepository
	cartItemAttributeItemService    CartItemAttributeItemService
	cartService                     CartService
}

func NewCartItemService(inventoryService InventoryService,
	cartRepository repository.CartRepository,
	repository repository.CartItemRepository,
	inventoryRepository repository.InventoryRepository,
	cartItemAttributeItemRepository repository.CartItemAttributeItemRepository,
	cartItemAttributeItemService CartItemAttributeItemService,
	cartService CartService,
) CartItemService {
	return &cartItemService{
		repository,
		cartRepository,
		inventoryRepository,
		inventoryService,
		cartItemAttributeItemRepository,
		cartItemAttributeItemService,
		cartService,
	}
}

func (s *cartItemService) Create(cartItemReq request.CartItemCreateRequest, customerId int) (model.Cart, error) {
	cartItem := model.CartItem{}

	cartItem.Qty = cartItemReq.Qty
	cartItem.UnitPrice = cartItemReq.UnitPrice
	cartItem.SubTotal = cartItemReq.Qty * cartItemReq.UnitPrice
	cartItem.ProductID = cartItemReq.ProductID

	// inventory, err := s.inventoryRepository.FindByProductId(int(cartItemReq.ProductID))
	// if err != nil {
	// 	return model.Order{}, err
	// }

	// if inventory.SalableQty < cartItemReq.Qty {
	// 	return model.Order{}, errors.New("Qty insufficient")
	// }

	cart, err := s.cartRepository.FindOneByCustomerAndStatus(customerId, "active")
	if err != nil {
		return cart, errors.New("Theres no active cart found, Please create cart first !")
	}

	cartItem, err = s.repository.Create(cartItem)
	if err != nil {
		return cart, err
	}

	if len(cartItemReq.AttributeItem) > 0 {
		_, err = s.cartItemAttributeItemService.CreateMany(cartItemReq.AttributeItem, cartItem.ID)
		if err != nil {
			return cart, err
		}
	}

	cart, err = s.cartRepository.FindById(int(cart.ID))
	if err != nil {
		return cart, err
	}

	// inventory.ReservedQty += cartItem.Qty
	// inventory.SalableQty -= cartItem.Qty

	// _, err = s.inventoryRepository.Update(inventory)
	// if err != nil {
	// 	return model.Order{}, err
	// }

	return cart, nil
}

func (s *cartItemService) Update(inputID request.CartItemFindById, cartItemReq request.CartItemUpdateRequest, customerId int) (model.Cart, error) {
	var cart model.Cart

	cartItem, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return model.Cart{}, err
	}

	cart = cartItem.Cart

	cartItem.Qty = cartItemReq.Qty
	cartItem.UnitPrice = cartItemReq.UnitPrice
	cartItem.SubTotal = cartItemReq.Qty * cartItemReq.UnitPrice

	_, err = s.repository.Update(cartItem)
	if err != nil {
		return model.Cart{}, err
	}

	// _, err = s.inventoryService.syncInventoryAfterUpdateCartItem(int(cartItem.ProductID), cartItem, cartItemReq)
	// if err != nil {
	// 	return cart, err
	// }

	err = s.cartItemAttributeItemRepository.DeleteByCartItemId(cartItem.ID)
	if err != nil {
		return cart, err
	}

	if len(cartItemReq.AttributeItem) > 0 {
		_, err = s.cartItemAttributeItemService.CreateMany(cartItemReq.AttributeItem, cartItem.ID)
		if err != nil {
			return cart, err
		}
	}

	existCart, err := s.cartRepository.FindById(int(cartItem.CartID))
	if err != nil {
		return cart, err
	}

	updateCart, err := s.cartRepository.Update(s.cartService.RecalculateCart(existCart))
	if err != nil {
		return updateCart, err
	}

	cartAfterUpdate, err := s.cartRepository.FindById(int(updateCart.ID))
	if err != nil {
		return cartAfterUpdate, err
	}

	return cartAfterUpdate, nil
}

func (s *cartItemService) FindById(cartItemId int) (model.CartItem, error) {
	cartItem, err := s.repository.FindById(cartItemId)
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (s *cartItemService) FindByCode(code string) (model.CartItem, error) {
	cartItem, err := s.repository.FindByCode(code)
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

func (s *cartItemService) FindAll(page int, pageSize int, q string) ([]model.CartItem, error) {
	cartItems, err := s.repository.FindAll(page, pageSize, q)
	if err != nil {
		return cartItems, err
	}
	return cartItems, nil
}

func (s *cartItemService) Delete(cartItemId int) (model.CartItem, error) {
	cartItem, err := s.repository.Delete(cartItemId)
	if err != nil {
		return cartItem, err
	}

	_, err = s.inventoryService.syncInventoryAfterDeleteItemFromCart(cartItem)
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}

// func (s *cartItemService) RecalculateOrderGrandtotal(order model.Order) (model.Order, error) {
// 	cartItems, err := s.repository.FindByOrderId(int(order.ID))
// 	if err != nil {
// 		return model.Order{}, err
// 	}

// 	for _, v := range cartItems {
// 		order.BaseAmount += v.SubTotal
// 	}

// 	updatedOrder, err := s.orderRepository.Update(order)
// 	if err != nil {
// 		return updatedOrder, err
// 	}

// 	return updatedOrder, nil
// }
