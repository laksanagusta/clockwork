package request

type CartUpdateRequest struct {
	BaseAmount int    `json:"baseAmount" validate:"required" binding:"required,numberic"`
	TotalItem  int    `json:"totalItem" validate:"required" binding:"required,numeric"`
	Status     string `json:"status" validate:"required" binding:"required"`
}

type CartFindById struct {
	ID int `uri:"id" binding:"required"`
}
