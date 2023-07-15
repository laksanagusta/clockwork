package request

type OrderItemCreateRequest struct {
	Qty       int  `json:"qty" validate:"required" binding:"required"`
	UnitPrice int  `json:"unitPrice" validate:"required" binding:"required"`
	SubTotal  int  `json:"subTotal" validate:"required" binding:"required"`
	ProductID uint `json:"productId" validate:"required" binding:"required"`
	OrderID   uint `json:"orderId" validate:"required" binding:"required"`
}

type OrderItemUpdateRequest struct {
	Qty       int `json:"qty" validate:"required" binding:"required"`
	UnitPrice int `json:"unitPrice" validate:"required" binding:"required"`
	SubTotal  int `json:"subTotal" validate:"required" binding:"required"`
}

type OrderItemFindById struct {
	ID int `uri:"id" binding:"required"`
}

type OrderItemFindByCode struct {
	Code string `uri:"code" binding:"required"`
}
