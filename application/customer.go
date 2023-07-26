package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type CustomerService interface {
	RegisterCustomer(input request.RegisterCustomerRequest) (model.Customer, error)
	Login(input request.LoginRequest) (model.Customer, error)
	CustomerDetails(id uint64) (model.Customer, error)
	CustomerFindAll() ([]model.Customer, error)
	UpdateCustomer(inputID request.CustomerFindById, inputData request.UpdateCustomerRequest) (model.Customer, error)
	DeleteCustomer(inputID request.CustomerFindById) (string, error)
	CreateCustomerBulk(workerIndex int, counter int, jobs []interface{}) (model.Customer, error)
}

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{repository}
}

func (s *customerService) RegisterCustomer(input request.RegisterCustomerRequest) (model.Customer, error) {
	customer := model.Customer{}
	customer.Name = input.Name
	customer.Email = strings.ToLower(input.Email)
	customer.PhoneNumber = input.PhoneNumber

	isEmailExist, err := s.repository.FindByEmail(customer.Email)
	if err != nil {
		return customer, err
	}

	if isEmailExist.ID != 0 {
		return customer, errors.New("Register failed, Email already been taken")
	}

	isPhoneNumberExist, err := s.repository.FindByPhoneNumber(customer.PhoneNumber)
	if err != nil {
		return customer, err
	}

	if isPhoneNumberExist.ID != 0 {
		return customer, errors.New("Register failed, phone number already been taken")
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return customer, err
	}

	customer.PasswordHash = string(PasswordHash)

	newCustomer, err := s.repository.Save(customer)
	if err != nil {
		return newCustomer, err
	}

	return newCustomer, nil
}

func (s *customerService) Login(input request.LoginRequest) (model.Customer, error) {
	email := input.Email
	password := input.Password

	customer, err := s.repository.FindByEmail(email)
	if err != nil {
		return customer, err
	}

	if customer.ID == 0 {
		return customer, errors.New("Customer not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(password))
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (s *customerService) CustomerDetails(id uint64) (model.Customer, error) {
	customer_id := id
	customer, err := s.repository.FindById(customer_id)
	if customer.ID == 0 {
		return customer, errors.New("Customer not found")
	}
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (s *customerService) CustomerFindAll() ([]model.Customer, error) {
	customers, err := s.repository.FindAll()
	if err != nil {
		return customers, err
	}
	return customers, nil
}

func (s *customerService) UpdateCustomer(inputID request.CustomerFindById, input request.UpdateCustomerRequest) (model.Customer, error) {
	customer, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return customer, err
	}

	customer.Name = input.Name
	customer.Email = strings.ToLower(input.Email)
	customer.PhoneNumber = input.PhoneNumber

	updatedCustomer, err := s.repository.Update(customer)
	if err != nil {
		return updatedCustomer, err
	}

	return updatedCustomer, nil

}

func (s *customerService) DeleteCustomer(inputID request.CustomerFindById) (string, error) {
	updatedCustomer, err := s.repository.Delete(inputID.ID)
	if err != nil {
		return updatedCustomer, err
	}

	return updatedCustomer, nil
}

func (s *customerService) CreateCustomerBulk(workerIndex int, counter int, jobs []interface{}) (model.Customer, error) {
	customer := model.Customer{}
	customer.Name = fmt.Sprintf("%v", jobs[0])

	newCustomers, err := s.repository.Save(customer)
	if err != nil {
		return newCustomers, err
	}

	return newCustomers, nil
}
