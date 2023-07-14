package request

type OrderCreateRequest struct {
	GrandTotal        int    `json:"grandTotal" validate:"required" binding:"required"`
	TransactionNumber string `json:"transactionNumber" validate:"required" binding:"required"`
}

type OrderUpdateRequest struct {
	GrandTotal        int    `json:"grandTotal" validate:"required" binding:"required"`
	TransactionNumber string `json:"transactionNumber" validate:"required" binding:"required"`
}

type OrderFindById struct {
	ID int `uri:"id" binding:"required"`
}

type OrderFindByCode struct {
	Code string `uri:"code" binding:"required"`
}
