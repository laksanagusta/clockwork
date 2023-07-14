package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
)

type OrderItemService interface {
	Create(orderItemReq request.OrderItemCreateInput, customerId int) (model.Order, error)
	Update(inputID request.OrderItemFindById, request request.OrderItemUpdateInput) (model.Order, error)
	FindById(orderItemId int) (model.OrderItem, error)
	FindByCode(code string) (model.OrderItem, error)
	FindAll(page int, page_size int, q string) ([]model.OrderItem, error)
	Delete(orderItemId int) (model.OrderItem, error)
}

type orderItemService struct {
	repository          repository.OrderItemRepository
	orderRepository     repository.OrderRepository
	inventoryRepository repository.InventoryRepository
	midtransService     MidtransService
}

func NewOrderItemService(orderRepository repository.OrderRepository, repository repository.OrderItemRepository, inventoryRepository repository.InventoryRepository, midtransService MidtransService) OrderItemService {
	return &orderItemService{
		repository,
		orderRepository,
		inventoryRepository,
		midtransService,
	}
}

func (s *orderItemService) Create(orderItemReq request.OrderItemCreateInput, customerId int) (model.Order, error) {
	orderItem := model.OrderItem{}

	orderItem.Qty = orderItemReq.Qty
	orderItem.UnitPrice = orderItemReq.UnitPrice
	orderItem.SubTotal = orderItemReq.Qty * orderItemReq.UnitPrice
	orderItem.ProductID = orderItemReq.ProductID

	inventory, err := s.inventoryRepository.FindByProductId(int(orderItemReq.ProductID))
	if err != nil {
		return model.Order{}, nil
	}

	if inventory.SalableQty < orderItemReq.Qty {
		return model.Order{}, errors.New("Qty insufficient")
	}

	order, err := s.orderRepository.FindOngoingOrder(int(orderItemReq.OrderID))
	if err != nil {
		return order, err
	}

	if order.ID == 0 {
		orderItem.OrderID = order.ID

		order := model.Order{}

		order.GrandTotal = orderItemReq.SubTotal
		order.Status = "created"

		order, err := s.orderRepository.Create(order)
		if err != nil {
			return order, err
		}

		_, err = s.repository.Create(orderItem)
		if err != nil {
			return order, err
		}

		updatedOrder, err := s.RecalculateOrderGrandtotal(order)
		if err != nil {
			return updatedOrder, err
		}

		return updatedOrder, nil
	} else {
		updatedOrder, err := s.RecalculateOrderGrandtotal(order)
		if err != nil {
			return updatedOrder, err
		}

		return updatedOrder, nil
	}
}

func (s *orderItemService) Update(inputID request.OrderItemFindById, orderItemReq request.OrderItemUpdateInput) (model.Order, error) {
	var order model.Order

	orderItem, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return order, err
	}

	order = orderItem.Order

	orderItem.Qty = orderItemReq.Qty
	orderItem.UnitPrice = orderItemReq.UnitPrice
	orderItem.SubTotal = orderItemReq.Qty * orderItemReq.UnitPrice

	_, err = s.repository.Update(orderItem)
	if err != nil {
		return order, err
	}

	updatedOrder, err := s.RecalculateOrderGrandtotal(order)
	if err != nil {
		return updatedOrder, err
	}

	return updatedOrder, nil
}

func (s *orderItemService) RecalculateOrderGrandtotal(order model.Order) (model.Order, error) {
	orderItems, err := s.repository.FindByOrderId(int(order.ID))
	if err != nil {
		return order, err
	}

	var grandTotal int
	for _, v := range orderItems {
		grandTotal += v.SubTotal
	}

	order.GrandTotal = grandTotal

	updatedOrder, err := s.orderRepository.Update(order)
	if err != nil {
		return updatedOrder, err
	}

	return updatedOrder, nil
}

func (s *orderItemService) FindById(orderItemId int) (model.OrderItem, error) {
	orderItem, err := s.repository.FindById(orderItemId)
	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}

func (s *orderItemService) FindByCode(code string) (model.OrderItem, error) {
	orderItem, err := s.repository.FindByCode(code)
	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}

func (s *orderItemService) FindAll(page int, pageSize int, q string) ([]model.OrderItem, error) {
	orderItems, err := s.repository.FindAll(page, pageSize, q)
	if err != nil {
		return orderItems, err
	}
	return orderItems, nil
}

func (s *orderItemService) Delete(orderItemId int) (model.OrderItem, error) {
	orderItem, err := s.repository.Delete(orderItemId)
	if err != nil {
		return orderItem, err
	}

	return orderItem, nil
}
