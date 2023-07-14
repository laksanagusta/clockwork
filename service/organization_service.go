package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
	"strings"
)

type OrganizationService interface {
	Create(request request.OrganizationCreateInput) (model.Organization, error)
	Update(inputID request.OrganizationFindById, request request.OrganizationUpdateInput) (model.Organization, error)
	FindById(organizationId int) (model.Organization, error)
	FindAll(page int, limit int, q string) ([]model.Organization, error)
	Delete(organizationId int) (model.Organization, error)
}

type organizationService struct {
	repository repository.OrganizationRepository
}

func NewOrganizationService(repository repository.OrganizationRepository) OrganizationService {
	return &organizationService{
		repository,
	}
}

func (s *organizationService) Create(request request.OrganizationCreateInput) (model.Organization, error) {
	organization := model.Organization{}
	organization.Name = request.Name
	organization.Address = request.Address

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(organization.Name))
	if err != nil {
		return organization, err
	}

	if len(checkIfExist) > 0 {
		return organization, errors.New("Organization with same name already exist")
	}

	newOrganization, err := s.repository.Create(organization)
	if err != nil {
		return newOrganization, err
	}

	return newOrganization, nil
}

func (s *organizationService) Update(inputID request.OrganizationFindById, request request.OrganizationUpdateInput) (model.Organization, error) {
	organization, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return organization, err
	}

	organization.Name = request.Name
	organization.Address = request.Address

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(organization.Name))
	if err != nil {
		return organization, err
	}

	if len(checkIfExist) > 0 {
		return organization, errors.New("Organization with same name already exist")
	}

	updatedOrganization, err := s.repository.Update(organization)
	if err != nil {
		return updatedOrganization, err
	}

	return updatedOrganization, nil

}

func (s *organizationService) FindById(organizationId int) (model.Organization, error) {
	organization, err := s.repository.FindById(organizationId)
	if err != nil {
		return organization, err
	}

	if organization.ID == 0 {
		return organization, errors.New("Organization not found")
	}

	return organization, nil
}

func (s *organizationService) FindAll(page int, limit int, q string) ([]model.Organization, error) {
	organizations, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return organizations, err
	}
	return organizations, nil
}

func (s *organizationService) Delete(organizationId int) (model.Organization, error) {
	organization, err := s.repository.Delete(organizationId)
	if err != nil {
		return organization, err
	}

	return organization, nil
}
