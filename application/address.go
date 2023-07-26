package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
	"errors"
)

type AddressService interface {
	Create(request request.AddressCreateRequest) (model.Address, error)
	Update(id request.AddressFindById, request request.AddressUpdateRequest) (model.Address, error)
	FindById(id int) (model.Address, error)
	FindAll(page int, limit int, q string) ([]model.Address, error)
	Delete(id int) (model.Address, error)
}

type addressService struct {
	addressRepo  repository.AddressRepository
	customerRepo repository.CustomerRepository
}

func NewAddressService(addressRepo repository.AddressRepository, customerRepo repository.CustomerRepository) AddressService {
	return &addressService{addressRepo, customerRepo}
}

func (service *addressService) Create(request request.AddressCreateRequest) (model.Address, error) {
	address := model.Address{}

	address.Title = request.Title
	address.Street = request.Street
	address.Note = request.Note
	address.City = request.City
	address.CustomerID = request.CustomerID

	isCustomerExist, e := service.customerRepo.FindById(uint64(address.CustomerID))
	if e != nil {
		return address, e
	}

	if isCustomerExist.ID == 0 {
		return address, errors.New("Customer not found")
	}

	newAddress, err := service.addressRepo.Create(address)
	if err != nil {
		return newAddress, err
	}

	return newAddress, nil
}

func (service *addressService) Update(id request.AddressFindById, request request.AddressUpdateRequest) (model.Address, error) {
	address, e := service.addressRepo.FindById(id.ID)
	if e != nil {
		return address, e
	}

	address.Title = request.Title
	address.Street = request.Street
	address.Note = request.Note
	address.City = request.City

	newAddress, err := service.addressRepo.Update(address)
	if err != nil {
		return newAddress, err
	}

	return newAddress, nil
}

func (service *addressService) FindById(id int) (model.Address, error) {
	address, e := service.addressRepo.FindById(id)
	if e != nil {
		return address, e
	}

	return address, nil
}

func (service *addressService) FindAll(page int, limit int, q string) ([]model.Address, error) {
	address, e := service.addressRepo.FindAll(page, limit, q)
	if e != nil {
		return address, e
	}

	return address, nil
}

func (service *addressService) Delete(id int) (model.Address, error) {
	address, e := service.addressRepo.Delete(id)
	if e != nil {
		return address, e
	}

	return address, nil
}
