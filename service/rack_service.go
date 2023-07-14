package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
	"strings"
)

type RackService interface {
	Create(request request.RackCreateInput) (model.Rack, error)
	Update(inputID request.RackFindById, request request.RackUpdateInput) (model.Rack, error)
	FindById(rackId int) (model.Rack, error)
	FindAll(page int, limit int, q string) ([]model.Rack, error)
	Delete(rackId int) (model.Rack, error)
}

type rackService struct {
	repository repository.RackRepository
}

func NewRackService(repository repository.RackRepository) RackService {
	return &rackService{
		repository,
	}
}

func (s *rackService) Create(request request.RackCreateInput) (model.Rack, error) {
	rack := model.Rack{}
	rack.Title = request.Title

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(rack.Title))
	if err != nil {
		return rack, err
	}

	if len(checkIfExist) > 0 {
		return rack, errors.New("Rack with same name already exist")
	}

	newRack, err := s.repository.Create(rack)
	if err != nil {
		return newRack, err
	}

	return newRack, nil
}

func (s *rackService) Update(inputID request.RackFindById, request request.RackUpdateInput) (model.Rack, error) {
	rack, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return rack, err
	}

	rack.Title = request.Title

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(rack.Title))
	if err != nil {
		return rack, err
	}

	if len(checkIfExist) > 0 {
		return rack, errors.New("Rack with same name already exist")
	}

	updatedRack, err := s.repository.Update(rack)
	if err != nil {
		return updatedRack, err
	}

	return updatedRack, nil

}

func (s *rackService) FindById(rackId int) (model.Rack, error) {
	rack, err := s.repository.FindById(rackId)
	if err != nil {
		return rack, err
	}

	if rack.ID == 0 {
		return rack, errors.New("Rack not found")
	}

	return rack, nil
}

func (s *rackService) FindAll(page int, limit int, q string) ([]model.Rack, error) {
	racks, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return racks, err
	}
	return racks, nil
}

func (s *rackService) Delete(rackId int) (model.Rack, error) {
	rack, err := s.repository.Delete(rackId)
	if err != nil {
		return rack, err
	}

	return rack, nil
}
