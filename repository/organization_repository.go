package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
	Create(organization model.Organization) (model.Organization, error)
	Update(organization model.Organization) (model.Organization, error)
	FindById(organizationId int) (model.Organization, error)
	FindAll(page int, page_size int, q string) ([]model.Organization, error)
	Delete(organizationId int) (model.Organization, error)
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db}
}

func (pr *organizationRepository) Create(organization model.Organization) (model.Organization, error) {
	err := pr.db.Create(&organization).Error
	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (pr *organizationRepository) Update(organization model.Organization) (model.Organization, error) {
	err := pr.db.Save(&organization).Error
	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (pr *organizationRepository) FindById(organizationId int) (model.Organization, error) {
	organization := model.Organization{}

	err := pr.db.First(&organization, organizationId).Error
	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (pr *organizationRepository) FindAll(page int, limit int, q string) ([]model.Organization, error) {
	var organization []model.Organization

	querydb := pr.db

	if limit > 0 {
		querydb = querydb.Limit(limit)
	} else {
		querydb = querydb.Limit(helper.QUERY_LIMITATION)
	}

	if page > 0 {
		querydb = querydb.Offset(page - 1)
	}

	if q != "" {
		querydb = querydb.Where("lower(name) LIKE ?", "%"+q+"%")
	}

	err := querydb.Find(&organization).Error
	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (pr *organizationRepository) Delete(organizationId int) (model.Organization, error) {
	var organization model.Organization
	err := pr.db.Where("id = ?", organizationId).Delete(&organization).Error

	if err != nil {
		return organization, err
	}

	return organization, nil
}
