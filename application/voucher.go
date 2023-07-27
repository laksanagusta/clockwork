package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
	"errors"
	"strings"
)

type VoucherService interface {
	Create(request request.VoucherCreateInput) (model.Voucher, error)
	Update(inputID request.VoucherFindById, request request.VoucherUpdateInput) (model.Voucher, error)
	FindById(voucherId int) (model.Voucher, error)
	FindAll(page int, limit int, q string) ([]model.Voucher, error)
	Delete(voucherId int) (model.Voucher, error)
}

type voucherService struct {
	repository repository.VoucherRepository
}

func NewVoucherService(repository repository.VoucherRepository) VoucherService {
	return &voucherService{
		repository,
	}
}

func (s *voucherService) Create(request request.VoucherCreateInput) (model.Voucher, error) {
	voucher := model.Voucher{}
	voucher.Title = request.Title

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(voucher.Title))
	if err != nil {
		return voucher, err
	}

	if len(checkIfExist) > 0 {
		return voucher, errors.New("Voucher with same name already exist")
	}

	newVoucher, err := s.repository.Create(voucher)
	if err != nil {
		return newVoucher, err
	}

	return newVoucher, nil
}

func (s *voucherService) Update(inputID request.VoucherFindById, request request.VoucherUpdateInput) (model.Voucher, error) {
	voucher, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return voucher, err
	}

	voucher.Title = request.Title

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(voucher.Title))
	if err != nil {
		return voucher, err
	}

	if len(checkIfExist) > 0 {
		return voucher, errors.New("Voucher with same name already exist")
	}

	updatedVoucher, err := s.repository.Update(voucher)
	if err != nil {
		return updatedVoucher, err
	}

	return updatedVoucher, nil

}

func (s *voucherService) FindById(voucherId int) (model.Voucher, error) {
	voucher, err := s.repository.FindById(voucherId)
	if err != nil {
		return voucher, err
	}

	if voucher.ID == 0 {
		return voucher, errors.New("Voucher not found")
	}

	return voucher, nil
}

func (s *voucherService) FindAll(page int, limit int, q string) ([]model.Voucher, error) {
	vouchers, err := s.repository.FindAll(page, limit, q)
	if err != nil {
		return vouchers, err
	}
	return vouchers, nil
}

func (s *voucherService) Delete(voucherId int) (model.Voucher, error) {
	voucher, err := s.repository.Delete(voucherId)
	if err != nil {
		return voucher, err
	}

	return voucher, nil
}
