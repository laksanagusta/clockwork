package helper

import (
	"clockwork-server/request"
	"sort"
	"strconv"
	"strings"
)

type CartItemHelper interface {
	SortAttributeItemId(attributeItems []request.AttributeItem) string
}

type cartItemHelper struct {
}

func NewCartItemHelper() CartItemHelper {
	return &cartItemHelper{}
}

func (h *cartItemHelper) SortAttributeItemId(attributeItems []request.AttributeItem) string {
	attributeItemId := []string{}

	for _, v := range attributeItems {
		attributeItemId = append(attributeItemId, strconv.FormatUint(uint64(v.ID), 10))
	}

	sort.Strings(attributeItemId)
	implodeString := strings.Join(attributeItemId, "_")

	return implodeString
}
