package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
	"errors"
	"strings"
)

type LocationService interface {
	Create(request request.LocationCreateInput) (model.Location, error)
	Update(inputID request.LocationFindById, request request.LocationUpdateInput) (model.Location, error)
	FindById(locationId int) (model.Location, error)
	FindAll(page int, limit int, q string) ([]model.Location, error)
	Delete(locationId int) (model.Location, error)
}

type locationService struct {
	repository repository.LocationRepository
}

func NewLocationService(repository repository.LocationRepository) LocationService {
	return &locationService{
		repository,
	}
}

func (s *locationService) Create(request request.LocationCreateInput) (model.Location, error) {
	location := model.Location{}
	location.Name = request.Name
	location.Code = request.Code
	location.Address = request.Address

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(location.Name))
	if err != nil {
		return location, err
	}

	if len(checkIfExist) > 0 {
		return location, errors.New("Location with same name already exist")
	}

	newLocation, err := s.repository.Create(location)
	if err != nil {
		return newLocation, err
	}

	return newLocation, nil
}

func (s *locationService) Update(inputID request.LocationFindById, request request.LocationUpdateInput) (model.Location, error) {
	location, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return location, err
	}

	location.Name = request.Name
	location.Code = request.Code
	location.Address = request.Address

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(location.Name))
	if err != nil {
		return location, err
	}

	if len(checkIfExist) > 0 {
		return location, errors.New("Location with same name already exist")
	}

	updatedLocation, err := s.repository.Update(location)
	if err != nil {
		return updatedLocation, err
	}

	return updatedLocation, nil

}

func (s *locationService) FindById(locationId int) (model.Location, error) {
	location, err := s.repository.FindById(locationId)
	if err != nil {
		return location, err
	}

	if location.ID == 0 {
		return location, errors.New("Location not found")
	}

	return location, nil
}

func (s *locationService) FindAll(page int, limit int, q string) ([]model.Location, error) {
	locations, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return locations, err
	}
	return locations, nil
}

func (s *locationService) Delete(locationId int) (model.Location, error) {
	location, err := s.repository.Delete(locationId)
	if err != nil {
		return location, err
	}

	return location, nil
}
