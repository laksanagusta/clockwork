package helper

import (
	"errors"
	"strconv"
	"time"
)

type GlobalHelper interface {
	ConvertArrayOfStringToInt(strings []string) ([]int, error)
	StartTimeEndTimeValidation(st time.Time, et time.Time) error
	NewTrue() *bool
}

type globalHelper struct {
}

func NewGlobalHelper() GlobalHelper {
	return &globalHelper{}
}

func (h *globalHelper) ConvertArrayOfStringToInt(strings []string) ([]int, error) {
	ints := []int{}

	for _, v := range strings {
		convertString, err := strconv.Atoi(v)
		if err != nil {
			return ints, errors.New("Error converting string to number, please check your input")
		}

		ints = append(ints, convertString)
	}

	return ints, nil
}

func (h *globalHelper) StartTimeEndTimeValidation(st time.Time, et time.Time) error {
	if !st.Before(et) {
		return errors.New("Start time must be lower than end time")
	}

	return nil
}

func (h *globalHelper) NewTrue() *bool {
	b := true
	return &b
}
