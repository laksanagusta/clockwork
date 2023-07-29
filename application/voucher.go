package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/helper"
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
	ApplyVoucher(req request.VoucherApply) (model.Cart, error)
}

type voucherService struct {
	repository   repository.VoucherRepository
	cartRepo     repository.CartRepository
	globalHelper helper.GlobalHelper
}

func NewVoucherService(repository repository.VoucherRepository, cartRepo repository.CartRepository, globalHelper helper.GlobalHelper) VoucherService {

	return &voucherService{
		repository,
		cartRepo,
		globalHelper,
	}
}

func (s *voucherService) Create(request request.VoucherCreateInput) (model.Voucher, error) {
	voucher := model.Voucher{
		Title:      request.Title,
		Code:       request.Code,
		IsActive:   request.IsActive,
		DiscAmount: request.DiscAmount,
		EndTime:    request.EndTime,
		StartTime:  request.StartTime,
	}

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
	voucher.Code = request.Code
	voucher.IsActive = request.IsActive
	voucher.DiscAmount = request.DiscAmount
	voucher.StartTime = request.StartTime
	voucher.EndTime = request.EndTime

	checkIfExist, err := s.repository.FindAll(0, 1, strings.ToLower(voucher.Title))
	if err != nil {
		return voucher, err
	}

	if len(checkIfExist) > 0 {
		if checkIfExist[0].ID != uint(inputID.ID) {
			return voucher, errors.New("Voucher with same name already exist")
		}
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

func (s *voucherService) ApplyVoucher(req request.VoucherApply) (model.Cart, error) {
	cart, _ := s.cartRepo.FindById(int(req.CartID))

	if cart.ID == 0 || cart.Status == "inactive" {
		return cart, errors.New("cart not valid")
	}

	voucher, _ := s.repository.FindById(int(req.CartID))

	if voucher.ID == 0 || voucher.IsActive == s.globalHelper.NewTrue() {
		return cart, errors.New("voucher is not valid")
	}

	updateCart, err := s.repository.ApplyVoucher(int(req.CartID), int(req.VoucherID))
	if err != nil {
		return updateCart, err
	}

	latestCart, err := s.cartRepo.FindById(int(req.CartID))
	if err != nil {
		return latestCart, err
	}

	return latestCart, nil
}
