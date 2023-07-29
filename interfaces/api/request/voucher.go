package request

type VoucherCreateInput struct {
	Title      string `json:"title" validate:"required" binding:"required"`
	Code       string `json:"code" validate:"required" binding:"required"`
	StartTime  string `json:"startTime" binding:"required"`
	EndTime    string `json:"endTime" binding:"required"`
	IsActive   *bool  `json:"isActive" validate:"required" binding:"required,boolean"`
	DiscAmount int    `json:"discAmount" validate:"required" binding:"required,number"`
}

type VoucherUpdateInput struct {
	Title      string `json:"title" validate:"required" binding:"required"`
	Code       string `json:"code" validate:"required" binding:"required"`
	StartTime  string `json:"startTime" binding:"required,ltefield=EndTime" time_format:"2006-01-02"`
	EndTime    string `json:"endTime" binding:"required" time_format:"2006-01-02"`
	IsActive   *bool  `json:"isActive" validate:"required" binding:"required,boolean"`
	DiscAmount int    `json:"discAmount" validate:"required" binding:"required,number"`
}

type VoucherApply struct {
	CartID    uint `json:"cartId" validate:"required" binding:"required"`
	VoucherID uint `json:"voucherId" validate:"required" binding:"required"`
}

type VoucherFindById struct {
	ID int `uri:"id" binding:"required"`
}
