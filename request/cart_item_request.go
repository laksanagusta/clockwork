package request

type CartItemCreateRequest struct {
	Qty           int `json:"qty" validate:"required" binding:"required"`
	UnitPrice     int `json:"unitPrice" validate:"required" binding:"required,numeric"`
	SubTotal      int `json:"subTotal" validate:"required" binding:"required,numeric"`
	Note          string
	AttributeItem []AttributeItem
	ProductID     uint `json:"productId" validate:"required" binding:"required"`
	OrderID       uint `json:"orderId" validate:"required" binding:"required"`
}

type CartItemUpdateRequest struct {
	Qty           int    `json:"qty" validate:"required" binding:"required"`
	UnitPrice     int    `json:"unitPrice" validate:"required" binding:"required,numeric"`
	SubTotal      int    `json:"subTotal" validate:"required" binding:"required,numeric"`
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
