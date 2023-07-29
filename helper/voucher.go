package helper

import (
	"errors"
	"time"
)

type VoucherHelper interface {
	ValidateTimeValid(st string, et string) error
}

type voucherHelper struct {
	globalHelper GlobalHelper
}

func NewVoucherHelper() VoucherHelper {
	globalHelper := NewGlobalHelper()

	return &voucherHelper{globalHelper}
}

func (h *voucherHelper) ValidateTimeValid(st string, et string) error {
	parsedStartTime, err := time.Parse("2006-01-02", st)
	if err != nil {
		return errors.New("Invalid start time value")
	}

	parsedEndTime, err := time.Parse("2006-01-02", et)
	if err != nil {
		return errors.New("Invalid end time value")
	}

	err = h.globalHelper.StartTimeEndTimeValidation(parsedStartTime, parsedEndTime)
	if err != nil {
		return err
	}

	return nil
}
