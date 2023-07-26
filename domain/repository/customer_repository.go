package repository

import (
	"clockwork-server/domain/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(customer model.Customer) (model.Customer, error)
	FindByEmail(email string) (model.Customer, error)
	FindById(id uint64) (model.Customer, error)
	FindByPhoneNumber(phoneNumber string) (model.Customer, error)
	FindAll() ([]model.Customer, error)
	Update(customer model.Customer) (model.Customer, error)
	Delete(customerID uint64) (string, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) Save(customer model.Customer) (model.Customer, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) FindByEmail(email string) (model.Customer, error) {
	var customer model.Customer
	err := r.db.Where("email = ?", email).Find(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) FindByPhoneNumber(phoneNumber string) (model.Customer, error) {
	var customer model.Customer
	err := r.db.Where("phone_number = ?", phoneNumber).Find(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) FindById(id uint64) (model.Customer, error) {
	var customer model.Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) FindAll() ([]model.Customer, error) {
	var customer []model.Customer
	err := r.db.Find(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) Update(customer model.Customer) (model.Customer, error) {
	err := r.db.Save(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *customerRepository) Delete(customerID uint64) (string, error) {
	err := r.db.Delete(&model.Customer{}, customerID).Error
	if err != nil {
		return "error delete customer", err
	}

	return "success delete customer", nil
}
