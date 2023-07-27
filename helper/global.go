package helper

import (
	"errors"
	"strconv"
)

type GlobalHelper interface {
	ConvertArrayOfStringToInt(strings []string) ([]int, error)
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
