package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
)

type OrderService interface {
	Create(request request.OrderCreateRequest) (model.Order, error)
	Update(inputID request.OrderFindById, request request.OrderUpdateRequest) (model.Order, error)
	FindById(orderId int) (model.Order, error)
	FindByCode(code string) (model.Order, error)
	FindAll(page int, page_size int, q string) ([]model.Order, error)
	Delete(orderId int) (model.Order, error)
	PlaceOrder(request request.OrderCreateRequest, orderId request.OrderFindById) (model.Order, error)
}

type orderService struct {
	repository      repository.OrderRepository
	midtransService MidtransService
}

func NewOrderService(repository repository.OrderRepository, midtransService MidtransService) OrderService {
	return &orderService{
		repository,
		midtransService,
	}
}

func (s *orderService) PlaceOrder(orderReq request.OrderCreateRequest, orderId request.OrderFindById) (model.Order, error) {
	order, err := s.repository.FindById(orderId.ID)
	if err != nil {
		return order, err
	}

	order.Status = "waiting_for_payment"

	createOrder, err := s.repository.Create(order)
	if err != nil {
		return order, err
	}

	snapUrl, err := s.midtransService.GenerateSnapUrl(createOrder)
	if err != nil {
		return createOrder, err
	}

	createOrder.SnapUrl = snapUrl

	assignSnapUrl, err := s.repository.Update(createOrder)
	if err != nil {
		return createOrder, err
	}

	return assignSnapUrl, nil
}

func (s *orderService) Create(request request.OrderCreateRequest) (model.Order, error) {
	order := model.Order{}

	newOrder, err := s.repository.Create(order)
	if err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func (s *orderService) Update(inputID request.OrderFindById, request request.OrderUpdateRequest) (model.Order, error) {
	order, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return order, err
	}

	order.GrandTotal = request.GrandTotal
	order.TransactionNumber = request.TransactionNumber

	updatedOrder, err := s.repository.Update(order)
	if err != nil {
		return updatedOrder, err
	}

	return updatedOrder, nil
}

func (s *orderService) FindById(orderId int) (model.Order, error) {
	order, err := s.repository.FindById(orderId)
	if err != nil {
		return order, err
	}

	if order.ID == 0 {
		return order, errors.New("Order not found")
	}

	return order, nil
}

func (s *orderService) FindByCode(code string) (model.Order, error) {
	order, err := s.repository.FindByCode(code)
	if err != nil {
		return order, err
	}

	if order.ID == 0 {
		return order, errors.New("Order not found")
	}

	return order, nil
}

func (s *orderService) FindAll(page int, pageSize int, q string) ([]model.Order, error) {
	orders, err := s.repository.FindAll(page, pageSize, q)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *orderService) Delete(orderId int) (model.Order, error) {
	order, err := s.repository.Delete(orderId)
	if err != nil {
		return order, err
	}

	return order, nil
}
