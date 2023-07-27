package request

type VoucherCreateInput struct {
	Title    string `json:"title" validate:"required" binding:"required"`
	Code     string `json:"code" validate:"required" binding:"required"`
	IsActive *bool  `json:"isActive" validate:"required" binding:"required,boolean"`
}

type VoucherUpdateInput struct {
	Title    string `json:"title" validate:"required" binding:"required"`
	Code     string `json:"code" validate:"required" binding:"required"`
	IsActive *bool  `json:"isActive" validate:"required" binding:"required,boolean"`
}

type VoucherFindById struct {
	ID int `uri:"id" binding:"required"`
}
