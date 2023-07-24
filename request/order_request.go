package request

type OrderCreateRequest struct {
	CartID        int    `json:"cartId" validate:"required" binding:"required,numeric"`
	PaymentMethod string `json:"paymentMethod" validate:"required" binding:"required"`
}

type OrderUpdateRequest struct {
	CartID        int    `json:"cartId" validate:"required" binding:"required,numeric"`
	PaymentMethod string `json:"paymentMethod" validate:"required" binding:"required"`
}

type OrderFindById struct {
	ID int `uri:"id" binding:"required"`
}

type OrderFindByCode struct {
	Code string `uri:"code" binding:"required"`
}
