package repository

import (
	"clockwork-server/helper"
	"clockwork-server/model"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment model.Payment) (model.Payment, error)
	Update(payment model.Payment) (model.Payment, error)
	FindById(paymentId int) (model.Payment, error)
	FindAll(page int, page_size int, q string) ([]model.Payment, error)
	Delete(paymentId int) (model.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (pr *paymentRepository) Create(payment model.Payment) (model.Payment, error) {
	err := pr.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (pr *paymentRepository) Update(payment model.Payment) (model.Payment, error) {
	err := pr.db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (pr *paymentRepository) FindById(paymentId int) (model.Payment, error) {
	payment := model.Payment{}

	err := pr.db.First(&payment, paymentId).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (pr *paymentRepository) FindAll(page int, limit int, q string) ([]model.Payment, error) {
	var payment []model.Payment

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
		querydb = querydb.Where("lower(title) LIKE ?", "%"+q+"%")
	}

	err := querydb.Find(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (pr *paymentRepository) Delete(paymentId int) (model.Payment, error) {
	var payment model.Payment
	err := pr.db.Where("id = ?", paymentId).Delete(&payment).Error

	if err != nil {
		return payment, err
	}

	return payment, nil
}
