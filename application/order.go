package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/helper"
	"clockwork-server/interfaces/api/request"
	"fmt"
)

type OrderService interface {
	Create(request request.OrderCreateRequest) (model.Order, error)
	Update(inputID request.OrderFindById, request request.OrderUpdateRequest) (model.Order, error)
	FindById(orderId int) (model.Order, error)
	FindByCode(code string) (model.Order, error)
	FindAll(page int, page_size int, q string) ([]model.Order, error)
	Delete(orderId int) (model.Order, error)
	PlaceOrder(request request.PlaceOrderRequest) (model.Order, error)
}

type orderService struct {
	repository      repository.OrderRepository
	midtransService MidtransService
	cartRepo        repository.CartRepository
	paymentRepo     repository.PaymentRepository
	orderHelper     helper.OrderHelper
}

func NewOrderService(repository repository.OrderRepository,
	midtransService MidtransService,
	cartRepo repository.CartRepository,
	paymentRepo repository.PaymentRepository,
	orderHelper helper.OrderHelper) OrderService {
	return &orderService{
		repository,
		midtransService,
		cartRepo,
		paymentRepo,
		orderHelper,
	}
}

func (s *orderService) PlaceOrder(orderReq request.PlaceOrderRequest) (model.Order, error) {
	order := model.Order{}

	cart, err := s.cartRepo.FindById(orderReq.CartID)
	if err != nil {
		return order, err
	}

	payment, err := s.paymentRepo.Create(model.Payment{
		PaymentMethod:   orderReq.PaymentMethod,
		PaymentStatus:   "pending",
		PaymentResponse: "",
	})

	if err != nil {
		return order, err
	}

	order.Status = "waiting_for_payment"
	order.BaseAmount = cart.BaseAmount
	order.AdditionalChargeAmount = s.orderHelper.CalculateTotalAdditionalCharge(cart.CartItems)

	totalAmount := order.BaseAmount + order.AdditionalChargeAmount

	order.TaxAmount = cart.BaseAmount * 11 / 100
	order.GrandTotal = totalAmount + order.TaxAmount
	order.CartID = cart.ID
	order.DiscountAmount = 0
	order.SnapUrl = ""
	order.PaymentID = payment.ID

	padZeros, _ := fmt.Printf("%06d", cart.ID)
	order.TransactionNumber = fmt.Sprint(padZeros)

	createOrder, err := s.repository.Create(order)
	if err != nil {
		return order, err
	}

	snapUrl, err := s.midtransService.GenerateSnapUrl(createOrder)
	if err != nil {
		return createOrder, err
	}

	createOrder.SnapUrl = snapUrl

	updateOrder, err := s.repository.Update(createOrder)
	if err != nil {
		return createOrder, err
	}

	return updateOrder, nil
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

	return order, nil
}

func (s *orderService) FindByCode(code string) (model.Order, error) {
	order, err := s.repository.FindByCode(code)
	if err != nil {
		return order, err
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
