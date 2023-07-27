package repository

import (
	"clockwork-server/domain/model"
	"clockwork-server/helper"

	"gorm.io/gorm"
)

type VoucherRepository interface {
	Create(voucher model.Voucher) (model.Voucher, error)
	Update(voucher model.Voucher) (model.Voucher, error)
	FindById(voucherId int) (model.Voucher, error)
	FindAll(page int, page_size int, q string) ([]model.Voucher, error)
	Delete(voucherId int) (model.Voucher, error)
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db}
}

func (pr *voucherRepository) Create(voucher model.Voucher) (model.Voucher, error) {
	err := pr.db.Create(&voucher).Error
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}

func (pr *voucherRepository) Update(voucher model.Voucher) (model.Voucher, error) {
	err := pr.db.Save(&voucher).Error
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}

func (pr *voucherRepository) FindById(voucherId int) (model.Voucher, error) {
	voucher := model.Voucher{}

	err := pr.db.First(&voucher, voucherId).Error
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}

func (pr *voucherRepository) FindAll(page int, limit int, q string) ([]model.Voucher, error) {
	var voucher []model.Voucher

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

	err := querydb.Find(&voucher).Error
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}

func (pr *voucherRepository) Delete(voucherId int) (model.Voucher, error) {
	var voucher model.Voucher
	err := pr.db.Where("id = ?", voucherId).Delete(&voucher).Error

	if err != nil {
		return voucher, err
	}

	return voucher, nil
}
