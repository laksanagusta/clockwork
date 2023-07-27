package response

import (
	"clockwork-server/domain/model"
	"time"
)

type Voucher struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Code      string    `json:"code"`
	IsActive  *bool     `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatVoucher(voucher model.Voucher) Voucher {
	var dataVoucher Voucher

	dataVoucher.ID = voucher.ID
	dataVoucher.Title = voucher.Title
	dataVoucher.Code = voucher.Code
	dataVoucher.IsActive = voucher.IsActive
	dataVoucher.CreatedAt = voucher.CreatedAt
	dataVoucher.UpdatedAt = voucher.UpdatedAt

	return dataVoucher
}

func FormatVouchers(order []model.Voucher) []Voucher {
	var dataVouchers []Voucher

	for _, value := range order {
		singleDataVoucher := FormatVoucher(value)
		dataVouchers = append(dataVouchers, singleDataVoucher)
	}

	return dataVouchers
}
