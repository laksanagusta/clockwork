package request

type CartItemCreateRequest struct {
	Qty           int    `json:"qty" validate:"required" binding:"required"`
	Note          string `json:"note" validate:"required" binding:"required"`
	AttributeItem []AttributeItem
	ProductID     uint `json:"productId" validate:"required" binding:"required,numeric"`
	CartID        uint `json:"cartId" validate:"required" binding:"required,numeric"`
}

type CartItemUpdateRequest struct {
	Qty           int    `json:"qty" validate:"required" binding:"required"`
	Note          string `json:"note" validate:"required" binding:"required"`
	AttributeItem []AttributeItem
}

type AttributeItem struct {
	ID               uint `json:"id" validate:"required" binding:"required,numeric"`
	AdditionalCharge int  `json:"additionalCharge" validate:"required" binding:"required,numeric"`
}

type CartItemFindById struct {
	ID int `uri:"id" binding:"required"`
}

type CartItemFindByCode struct {
	Code string `uri:"code" binding:"required"`
}
